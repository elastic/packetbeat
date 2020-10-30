#!/usr/bin/env groovy

@Library('apm@current') _

import groovy.transform.Field

/**
 This is required to store the test suites we will use to trigger the E2E tests.
*/
@Field def e2eTestSuites = []
@Field def lock = false

pipeline {
  agent none
  environment {
    REPO = 'beats'
    BASE_DIR = "src/github.com/elastic/${env.REPO}"
    JOB_GCS_BUCKET = 'beats-ci-artifacts'
    JOB_GCS_BUCKET_STASH = 'beats-ci-temp'
    JOB_GCS_CREDENTIALS = 'beats-ci-gcs-plugin'
    DOCKERELASTIC_SECRET = 'secret/observability-team/ci/docker-registry/prod'
    DOCKER_REGISTRY = 'docker.elastic.co'
    GITHUB_CHECK_E2E_TESTS_NAME = 'E2E Tests'
    SNAPSHOT = "true"
    PIPELINE_LOG_LEVEL = "INFO"
  }
  options {
    timeout(time: 3, unit: 'HOURS')
    buildDiscarder(logRotator(numToKeepStr: '20', artifactNumToKeepStr: '20', daysToKeepStr: '30'))
    timestamps()
    ansiColor('xterm')
    disableResume()
    durabilityHint('PERFORMANCE_OPTIMIZED')
    disableConcurrentBuilds()
  }
  triggers {
    issueCommentTrigger('(?i)^\\/packag[ing|e]$')
    // disable upstream trigger on a PR basis
    upstream("Beats/beats/${ env.JOB_BASE_NAME.startsWith('PR-') ? 'none' : env.JOB_BASE_NAME }")
  }
  parameters {
    booleanParam(name: 'linux', defaultValue: true, description: 'Build linux binaries.')
    booleanParam(name: 'macos', defaultValue: true, description: 'Build macOS binaries.')
    booleanParam(name: 'windows', defaultValue: true, description: 'Build windows binaries.')

    booleanParam(name: 'i386_amd64', defaultValue: true, description: 'Build i386/amd64 binaries.')
    booleanParam(name: 'arm', defaultValue: true, description: 'Build ARM binaries.')
    booleanParam(name: 'ppc_mips_x390s', defaultValue: true, description: 'Build PPC/MIPS/x390s binaries.')

    booleanParam(name: 'auditbeat', defaultValue: true, description: 'Build auditbeat binaries.')
    booleanParam(name: 'filebeat', defaultValue: true, description: 'Build filebeat binaries.')
    booleanParam(name: 'heartbeat', defaultValue: true, description: 'Build heartbeat binaries.')
    booleanParam(name: 'journalbeat', defaultValue: true, description: 'Build journalbeat binaries.')
    booleanParam(name: 'metricbeat', defaultValue: true, description: 'Build metricbeat binaries.')
    booleanParam(name: 'packetbeat', defaultValue: true, description: 'Build packetbeat binaries.')
    booleanParam(name: 'winlogbeat', defaultValue: true, description: 'Build winlogbeat binaries.')
    booleanParam(name: 'elastic_agent', defaultValue: true, description: 'Build elastic-agent binaries.')
    booleanParam(name: 'dockerlogbeat', defaultValue: true, description: 'Build dockerlogbeat binaries.')
    booleanParam(name: 'functionbeat', defaultValue: true, description: 'Build functionbeat binaries.')

    booleanParam(name: 'sign_on_macos', defaultValue: false, description: 'Sign macOS binaries.')
    booleanParam(name: 'archive_on_gcp', defaultValue: true, description: 'Enable/Disable the archive of binaries on GCP.')
    booleanParam(name: 'push_docker', defaultValue: true, description: 'Enable/Disable the push of the Docker images.')
    booleanParam(name: 'dry_run', defaultValue: false, description: 'Execute the pipeline without generate packages.')
    booleanParam(name: 'force', defaultValue: false, description: 'Force to Execute the pipeline.')
  }
  stages {
    stage('Filter build') {
      agent { label 'ubuntu-18 && immutable' }
      when {
        beforeAgent true
        anyOf {
          triggeredBy cause: "IssueCommentCause"
          expression {
            def ret = isUserTrigger() || isUpstreamTrigger() || params.dry_run || params.force
            if(!ret){
              currentBuild.result = 'NOT_BUILT'
              currentBuild.description = "The build has been skipped"
              currentBuild.displayName = "#${BUILD_NUMBER}-(Skipped)"
              echo("the build has been skipped due the trigger is a branch scan and the allow ones are manual, GitHub comment, and upstream job")
            }
            return ret
          }
        }
      }
      stages {
        stage('Checkout') {
          options { skipDefaultCheckout() }
          steps {
            deleteDir()
            checkoutComitId()
            setEnvVar("GO_VERSION", readFile("${BASE_DIR}/.go-version").trim())
            withMageEnv(){
              dir("${BASE_DIR}"){
                setEnvVar('BEAT_VERSION', sh(label: 'Get beat version', script: 'make get-version', returnStdout: true)?.trim())
              }
            }
            stashV2(name: 'source', bucket: "${JOB_GCS_BUCKET_STASH}", credentialsId: "${JOB_GCS_CREDENTIALS}")
          }
        }
        stage('Build Packages'){
          options { skipDefaultCheckout() }
          steps {
            packageTask()
          }
        }
        // stage('Run E2E Tests for Packages'){
        //   agent { label 'ubuntu-18 && immutable' }
        //   options { skipDefaultCheckout() }
        //   steps {
        //     runE2ETests()
        //   }
        // }
      }
    }
  }
}

def checkoutComitId(){
  if(isUpstreamTrigger()) {
    try {
      gitCheckout(basedir: "${BASE_DIR}", branch: getCommitId())
    } catch(err) {
      // Fallback to the head of the branch as used to be.
      gitCheckout(basedir: "${BASE_DIR}")
    }
  } else {
    gitCheckout(basedir: "${BASE_DIR}")
  }
}

def getCommitId(){
  copyArtifacts(filter: 'packaging.properties',
                flatten: true,
                projectName: "Beats/beats/${env.JOB_BASE_NAME}",
                selector: upstream(fallbackToLastSuccessful: true))
  def props = readProperties(file: 'packaging.properties')
  return props.COMMIT
}

def beats(){
  beatsList = []
  if(params.auditbeat){
    beatsList << 'auditbeat' << 'x-pack/auditbeat'
  }
  if(params.filebeat){
    beatsList << 'filebeat' << 'x-pack/filebeat'
  }
  if(params.heartbeat){
    beatsList << 'heartbeat' << 'x-pack/heartbeat'
  }
  if(params.journalbeat){
    beatsList << 'journalbeat'
  }
  if(params.metricbeat){
    beatsList << 'metricbeat' << 'x-pack/metricbeat'
  }
  if(params.packetbeat){
    beatsList << 'packetbeat' << 'x-pack/packetbeat'
  }
  if(params.elastic_agent){
    beatsList << 'x-pack/elastic-agent'
  }
  if(params.dockerlogbeat){
    beatsList << 'x-pack/dockerlogbeat'
  }
  if(params.functionbeat){
    beatsList << 'x-pack/functionbeat'
  }
  return beatsList
}

def platform(){
  platformsList = []
  if(params.linux){
    if(params.i386_amd64) {
      platformsList << 'linux/amd64 linux/386'
    }
    if(params.arm) {
      platformsList << 'linux/arm64 linux/armv7'
    }
    if(params.ppc_mips_x390s) {
      platformsList << 'linux/ppc64le linux/mips64 linux/s390x'
    }
  }
  if(params.windows){
    platformsList << 'windows/amd64 windows/386'
  }
  if(params.macos){
    platformsList << 'darwin/amd64'
  }
  return platformsList
}

def excludes(){
  excludeList = [:]
  excludeList['x-pack/elastic-agent:linux/ppc64le linux/mips64 linux/s390x'] = true
  return excludeList
}

def signOnMacOS(){
  return params.macos && params.sign_on_macos && env.PLATFORM == 'darwin/amd64'
}

def packageTask(){
  def tasks = [:]
  def excludes = excludes()
  beats().each{ beat ->
    platform().each{ platform ->
      def key = "${beat}:${platform}"
      if(excludes[key] == null){
        tasks["Packaging ${key}"] = {
          packageMacOS(beat, platform)
          packageLinux(beat, platform)
          prepareE2ETestForPackage("${beat}")
        }
      }
    }
  }
  parallel tasks
}

def packageMacOS(beat, platform){
  if(signOnMacOS()){
    node('macosx-10.12'){
      withEnv([
        "HOME=${env.WORKSPACE}",
        "PLATFORMS=+all ${platform}",
        "BEATS_FOLDER=${beat}"
      ]){
        withGithubNotify(context: "Packaging MacOS ${BEATS_FOLDER}") {
          deleteDir()
          withMacOSEnv(){
            release()
            pushCIDockerImages()
            publishPackages("${env.BEATS_FOLDER}")
          }
        }
      }
    }
  }
}

def packageLinux(beat, platform){
  if(!signOnMacOS()){
    node('ubuntu-18 && immutable'){
      withEnv([
        "HOME=${env.WORKSPACE}",
        "PLATFORMS=+all ${platform}",
        "BEATS_FOLDER=${beat}"
      ]){
        withGithubNotify(context: "Packaging Linux ${BEATS_FOLDER}") {
          deleteDir()
          release()
          pushCIDockerImages()
          publishPackages("${env.BEATS_FOLDER}")
        }
      }
      post {
        success {
          writeFile(file: 'beats-tester.properties',
                    text: """\
                    ## To be consumed by the beats-tester pipeline
                    COMMIT=${env.GIT_BASE_COMMIT}
                    BEATS_URL_BASE=https://storage.googleapis.com/${env.JOB_GCS_BUCKET}/commits/${env.GIT_BASE_COMMIT}
                    VERSION=${env.BEAT_VERSION}-SNAPSHOT""".stripIndent()) // stripIdent() requires '''/
          archiveArtifacts artifacts: 'beats-tester.properties'
        }
      }
    }
  }
}

def pushCIDockerImages(){
  catchError(buildResult: 'UNSTABLE', message: 'Unable to push Docker images', stageResult: 'FAILURE') {
    if (env?.BEATS_FOLDER?.endsWith('auditbeat')) {
      tagAndPush('auditbeat')
    } else if (env?.BEATS_FOLDER?.endsWith('filebeat')) {
      tagAndPush('filebeat')
    } else if (env?.BEATS_FOLDER?.endsWith('heartbeat')) {
      tagAndPush('heartbeat')
    } else if ("${env.BEATS_FOLDER}" == "journalbeat"){
      tagAndPush('journalbeat')
    } else if (env?.BEATS_FOLDER?.endsWith('metricbeat')) {
      tagAndPush('metricbeat')
    } else if ("${env.BEATS_FOLDER}" == "packetbeat"){
      tagAndPush('packetbeat')
    } else if ("${env.BEATS_FOLDER}" == "x-pack/elastic-agent") {
      tagAndPush('elastic-agent')
    }
  }
}

def tagAndPush(beatName){
  def libbetaVer = env.BEAT_VERSION
  def aliasVersion = ""
  if("${env.SNAPSHOT}" == "true"){
    aliasVersion = libbetaVer.substring(0, libbetaVer.lastIndexOf(".")) // remove third number in version

    libbetaVer += "-SNAPSHOT"
    aliasVersion += "-SNAPSHOT"
  }

  def tagName = "${libbetaVer}"
  if (isPR()) {
    tagName = "pr-${env.CHANGE_ID}"
  }

  dockerLogin(secret: "${DOCKERELASTIC_SECRET}", registry: "${DOCKER_REGISTRY}")

  // supported image flavours
  def variants = ["", "-oss", "-ubi8"]
  variants.each { variant ->
    doTagAndPush(beatName, variant, libbetaVer, tagName)
    doTagAndPush(beatName, variant, libbetaVer, "${env.GIT_BASE_COMMIT}")

    if (!isPR() && aliasVersion != "") {
      doTagAndPush(beatName, variant, libbetaVer, aliasVersion)
    }
  }
}

/**
* @param beatName name of the Beat
* @param variant name of the variant used to build the docker image name
* @param sourceTag tag to be used as source for the docker tag command, usually under the 'beats' namespace
* @param targetTag tag to be used as target for the docker tag command, usually under the 'observability-ci' namespace
*/
def doTagAndPush(beatName, variant, sourceTag, targetTag) {
  def sourceName = "${DOCKER_REGISTRY}/beats/${beatName}${variant}:${sourceTag}"
  def targetName = "${DOCKER_REGISTRY}/observability-ci/${beatName}${variant}:${targetTag}"

  if(params.dry_run || !params.push_docker){
    return
  }

  def exists = sh(label: "Check if ${targetName} Docker image exist", script: "docker inspect ${sourceName}", returnStatus: true)
  if(exists == 0){
    def iterations = 0
    retryWithSleep(retries: 3, seconds: 5, backoff: true) {
      iterations++
      def status = sh(label: "Change tag and push ${targetName}", script: """
        docker tag ${sourceName} ${targetName}
        docker push ${targetName}
      """, returnStatus: true)

      if ( status > 0 && iterations < 3) {
        error("tag and push failed for ${beatName}, retry")
      } else if ( status > 0 ) {
        log(level: 'WARN', text: "${beatName} doesn't have ${variant} docker images. See https://github.com/elastic/beats/pull/21621")
      }
    }
  }
}

def prepareE2ETestForPackage(String beat){
  if ("${beat}" == "filebeat" || "${beat}" == "x-pack/filebeat") {
    e2eTestSuites.push('fleet')
    e2eTestSuites.push('helm')
  } else if ("${beat}" == "metricbeat" || "${beat}" == "x-pack/metricbeat") {
    e2eTestSuites.push('ALL')
    echo("${beat} adds all test suites to the E2E tests job.")
  } else if ("${beat}" == "x-pack/elastic-agent") {
    e2eTestSuites.push('fleet')
  } else {
    echo("${beat} does not add any test suite to the E2E tests job.")
    return
  }
}

def release(){
  withBeatsEnv(){
    withEnv([
      "DEV=true"
    ]) {
      dir("${env.BEATS_FOLDER}") {
        sh(label: "Release ${env.BEATS_FOLDER} ${env.PLATFORMS}", script: 'mage package')
      }
    }
  }
}

def runE2ETests(){
  if (e2eTestSuites.size() == 0) {
    echo("Not triggering E2E tests for PR-${env.CHANGE_ID} because the changes does not affect the E2E.")
    return
  }

  def suites = '' // empty value represents all suites in the E2E tests

  catchError(buildResult: 'UNSTABLE', message: 'Unable to run e2e tests', stageResult: 'FAILURE') {
    def suitesSet = e2eTestSuites.toSet()

    if (!suitesSet.contains('ALL')) {
      suitesSet.each { suite ->
        suites += "${suite},"
      };
    }

    triggerE2ETests(suites)
  }
}

def triggerE2ETests(String suite) {
  echo("Triggering E2E tests for PR-${env.CHANGE_ID}. Test suites: ${suite}.")

  def branchName = isPR() ? "${env.CHANGE_TARGET}" : "${env.JOB_BASE_NAME}"
  def e2eTestsPipeline = "e2e-tests/e2e-testing-mbp/${branchName}"

  def parameters = [
    booleanParam(name: 'forceSkipGitChecks', value: true),
    booleanParam(name: 'forceSkipPresubmit', value: true),
    booleanParam(name: 'notifyOnGreenBuilds', value: !isPR()),
    string(name: 'runTestsSuites', value: suite),
    string(name: 'GITHUB_CHECK_NAME', value: env.GITHUB_CHECK_E2E_TESTS_NAME),
    string(name: 'GITHUB_CHECK_REPO', value: env.REPO),
    string(name: 'GITHUB_CHECK_SHA1', value: env.GIT_BASE_COMMIT),
  ]
  if (isPR()) {
    def version = "pr-${env.CHANGE_ID}"
    parameters.push(booleanParam(name: 'USE_CI_SNAPSHOTS', value: true))
    parameters.push(string(name: 'ELASTIC_AGENT_VERSION', value: "${version}"))
    parameters.push(string(name: 'METRICBEAT_VERSION', value: "${version}"))
  }

  if(params.dry_run){
    return
  }

  build(job: "${e2eTestsPipeline}",
    parameters: parameters,
    propagate: false,
    wait: false
  )

  def notifyContext = "${env.GITHUB_CHECK_E2E_TESTS_NAME}"
  githubNotify(context: "${notifyContext}", description: "${notifyContext} ...", status: 'PENDING', targetUrl: "${env.JENKINS_URL}search/?q=${e2eTestsPipeline.replaceAll('/','+')}")
}

def withMacOSEnv(Closure body){
  withEnvMask( vars: [
      [var: "KEYCHAIN_PASS", password: getVaultSecret(secret: "secret/jenkins-ci/macos-codesign-keychain").data.password],
      [var: "KEYCHAIN", password: "/var/lib/jenkins/Library/Keychains/Elastic.keychain-db"],
      [var: "APPLE_SIGNING_ENABLED", password: "true"],
  ]){
    if(params.dry_run){
      return
    }
    body()
  }
}

def publishPackages(baseDir){
  dir("${env.BASE_DIR}"){
    def bucketUri = "gs://${JOB_GCS_BUCKET}/snapshots"
    if (isPR()) {
      bucketUri = "gs://${JOB_GCS_BUCKET}/pull-requests/pr-${env.CHANGE_ID}"
    }
    def beatsFolderName = getBeatsName(baseDir)
    uploadPackages("${bucketUri}/${beatsFolderName}", baseDir)

    // Copy those files to another location with the sha commit to test them
    // aftewords.
    bucketUri = "gs://${JOB_GCS_BUCKET}/commits/${env.GIT_BASE_COMMIT}"
    uploadPackages("${bucketUri}/${beatsFolderName}", baseDir)
  }
}

@NonCPS
def setLock(){
  synchronized(lock) {
    lock = true
  }
}

@NonCPS
def setUnlock(){
  synchronized(lock) {
    lock = false
  }
}

@NonCPS
def isLock(){
  synchronized(lock) {
    return lock
  }
}

def waitForUnlock(){
  while(isLock()){
    sleep 10
  }
}

def uploadPackages(bucketUri, baseDir){
  if(params.dry_run || !params.archive_on_gcp){
    return
  }
  waitForUnlock()
  setLock()
  googleStorageUpload(bucket: bucketUri,
    credentialsId: "${JOB_GCS_CREDENTIALS}",
    pathPrefix: "${baseDir}/build/distributions/",
    pattern: "${baseDir}/build/distributions/**/*",
    sharedPublicly: true,
    showInline: true
  )
  setUnlock()
}

/**
* There is a specific folder structure in https://staging.elastic.co/ and https://artifacts.elastic.co/downloads/
* therefore the storage bucket in GCP should follow the same folder structure.
* This is required by https://github.com/elastic/beats-tester
* e.g.
* baseDir=name -> return name
* baseDir=name1/name2/name3-> return name2
*/
def getBeatsName(baseDir) {
  return baseDir.replace('x-pack/', '')
}

def withBeatsEnv(Closure body) {
  withMageEnv(){
    withEnv([
      "PYTHON_ENV=${WORKSPACE}/python-env"
    ]) {
      unstashV2(name: 'source', bucket: "${JOB_GCS_BUCKET_STASH}", credentialsId: "${JOB_GCS_CREDENTIALS}")
      if(params.dry_run){
        return
      }
      dir("${env.BASE_DIR}"){
        body()
      }
    }
  }
}
