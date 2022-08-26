// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !aix
// +build !aix

package gcs

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	cursor "github.com/elastic/beats/v7/filebeat/input/v2/input-cursor"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/gcs/job"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/gcs/pool"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/gcs/state"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/gcs/types"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/go-concert/timed"
)

type scheduler interface {
	createJobs(pager *azblob.ContainerListBlobFlatPager) ([]job.Job, error)
	Schedule(ctx context.Context) error
}

type gcsInputScheduler struct {
	publisher cursor.Publisher
	bucket    *storage.BucketHandle
	src       *types.Source
	cfg       *config
	state     *state.State
	log       *logp.Logger
}

// NewGcsInputScheduler , returns a new scheduler instance
func NewGcsInputScheduler(publisher cursor.Publisher, bucket *storage.BucketHandle, src *types.Source, cfg *config,
	state *state.State, log *logp.Logger,
) scheduler {
	return &gcsInputScheduler{
		publisher: publisher,
		bucket:    bucket,
		src:       src,
		cfg:       cfg,
		state:     state,
		log:       log,
	}
}

// Schedule , is responsible for fetching & scheduling jobs using the workerpool model
func (gcs *gcsInputScheduler) Schedule(ctx context.Context) error {
	var pager *azblob.ContainerListBlobFlatPager
	var availableWorkers int32

	workerPool := pool.NewWorkerPool(ctx, gcs.src.MaxWorkers, gcs.log)
	workerPool.Start()

	if !gcs.src.Poll {
		availableWorkers = workerPool.AvailableWorkers()
		pager = gcs.fetchBlobPager(availableWorkers)
		return gcs.scheduleOnce(ctx, pager, workerPool)
	}

	for {
		availableWorkers = workerPool.AvailableWorkers()
		if availableWorkers == 0 {
			continue
		}

		// availableWorkers is used as the batch size for a blob page so that
		// work distribution remains efficient
		pager = gcs.fetchBlobPager(availableWorkers)
		err := gcs.scheduleOnce(ctx, pager, workerPool)
		if err != nil {
			return err
		}

		err = timed.Wait(ctx, gcs.src.PollInterval)
		if err != nil {
			return err
		}
	}
}

func (gcs *gcsInputScheduler) scheduleOnce(ctx context.Context, pager *azblob.ContainerListBlobFlatPager, workerPool pool.Pool) error {
	for pager.NextPage(ctx) {
		jobs, err := gcs.createJobs(pager)
		if err != nil {
			gcs.log.Errorf("Job creation failed for container %s with error %v", gcs.src.BucketName, err)
			return err
		}

		// If previous checkpoint was saved then look up starting point for new jobs
		if gcs.state.Checkpoint().LatestEntryTime != nil {
			jobs = gcs.moveToLastSeenJob(jobs)
		}

		// Submits job to worker pool for further processing
		for _, job := range jobs {
			workerPool.Submit(job)
		}
	}
	return nil
}

func (gcs *gcsInputScheduler) createJobs(pager *azblob.ContainerListBlobFlatPager) ([]job.Job, error) {
	var jobs []job.Job
	pageMarker := pager.PageResponse().Marker

	for _, v := range pager.PageResponse().Segment.BlobItems {
		blobURL := gcs.serviceURL + gcs.src.ContainerName + "/" + *v.Name
		blobCreds := &types.BlobCredentials{
			ServiceCreds:  gcs.credential,
			BlobName:      *v.Name,
			ContainerName: gcs.src.ContainerName,
		}

		blobClient, err := fetchBlobClient(blobURL, blobCreds, gcs.log)
		if err != nil {
			return nil, err
		}

		job := job.NewAzureInputJob(blobClient, v, blobURL, pageMarker, gcs.state, gcs.src, gcs.publisher)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

// fetchBlobPager fetches the current blob page object given a batch size & a page marker.
// The page marker has been disabled since it was found that it operates on the basis of
// lexicographical order , and not on the basis of the latest file uploaded, meaning if a blob with a name
// of lesser lexicographic value is uploaded after a blob with a name of higher value , the latest
// marker stored in the checkpoint will not retrieve that new blob , this distorts the polling logic
// hence disabling it for now , until more feedback is given. Disabling this how ever makes the sheduler loop
// through all the blobs on every poll action to arrive at the latest checkpoint.
// [NOTE] : There are no api's / sdk functions that list blobs via timestamp/latest entry , it's always lexicographical order
func (gcs *gcsInputScheduler) fetchBlobPager(batchSize int32) *azblob.ContainerListBlobFlatPager {
	pager := gcs.client.ListBlobsFlat(&azblob.ContainerListBlobsFlatOptions{
		Include: []azblob.ListBlobsIncludeItem{
			azblob.ListBlobsIncludeItemMetadata,
			azblob.ListBlobsIncludeItemTags,
		},
		// Marker:     gcs.state.Checkpoint().Marker,
		MaxResults: &batchSize,
	})

	return pager
}

// moveToLastSeenJob , moves to the latest job position past the last seen job
// Jobs are stored in lexicographical order always , hence the latest position can be found either on the basis of job name or timestamp
func (gcs *gcsInputScheduler) moveToLastSeenJob(jobs []job.Job) []job.Job {
	var latestJobs []job.Job
	jobsToReturn := make([]job.Job, 0)
	counter := 0
	flag := false
	ignore := false

	for _, job := range jobs {
		if job.Timestamp().After(*gcs.state.Checkpoint().LatestEntryTime) {
			latestJobs = append(latestJobs, job)
		} else if job.Name() == gcs.state.Checkpoint().BlobName {
			flag = true
			break
		} else if job.Name() > gcs.state.Checkpoint().BlobName {
			flag = true
			counter--
			break
		} else if job.Name() <= gcs.state.Checkpoint().BlobName && !ignore {
			ignore = true
		}
		counter++
	}

	if flag && (counter < len(jobs)-1) {
		jobsToReturn = jobs[counter+1:]
	} else if !flag && !ignore {
		jobsToReturn = jobs
	}

	// in a senario where there are some jobs which have a later time stamp
	// but lesser alphanumeric order and some jobs have greater alphanumeric order
	// than the current checkpoint
	if len(jobsToReturn) != len(jobs) && len(latestJobs) > 0 {
		jobsToReturn = append(latestJobs, jobsToReturn...)
	}

	return jobsToReturn
}
