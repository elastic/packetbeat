from filebeat import BaseTest
from beat.beat import INTEGRATION_TESTS
import os
import unittest
import glob
import shutil
import subprocess
from elasticsearch import Elasticsearch
import json
import logging
from parameterized import parameterized

EXTENDED_COMPARE = ["elasticsearch"]


def load_fileset_test_cases():
    """
    Creates a list of all modules, filesets and testfiles inside for testing.
    To execute tests for only 1 module, set the env variable TESTING_FILEBEAT_MODULES
    to the specific module name or a , separated lists of modules.
    """
    current_dir = os.path.dirname(os.path.abspath(__file__))
    modules_dir = os.path.join(current_dir, "..", "..", "module")

    modules = os.getenv("TESTING_FILEBEAT_MODULES")
    if modules:
        modules = modules.split(",")
    else:
        modules = os.listdir(modules_dir)

    test_cases = []

    for module in modules:
        path = os.path.join(modules_dir, module)

        if not os.path.isdir(path):
            continue

        for fileset in os.listdir(path):
            if not os.path.isdir(os.path.join(path, fileset)):
                continue

            if not os.path.isfile(os.path.join(path, fileset, "manifest.yml")):
                continue

            test_files = glob.glob(os.path.join(modules_dir, module,
                                                fileset, "test", "*.log"))
            for test_file in test_files:
                test_cases.append([module, fileset, test_file])

    return test_cases


class Test(BaseTest):

    def init(self):
        self.elasticsearch_url = self.get_elasticsearch_url()
        print("Using elasticsearch: {}".format(self.elasticsearch_url))
        self.es = Elasticsearch([self.elasticsearch_url])
        logging.getLogger("urllib3").setLevel(logging.WARNING)
        logging.getLogger("elasticsearch").setLevel(logging.ERROR)

        self.modules_path = os.path.abspath(self.working_dir +
                                            "/../../../../module")

        self.filebeat = os.path.abspath(self.working_dir +
                                        "/../../../../filebeat.test")

        self.index_name = "test-filebeat-modules"

    @parameterized.expand(load_fileset_test_cases)
    @unittest.skipIf(not INTEGRATION_TESTS,
                     "integration tests are disabled, run with INTEGRATION_TESTS=1 to enable them.")
    @unittest.skipIf(os.getenv("TESTING_ENVIRONMENT") == "2x",
                     "integration test not available on 2.x")
    def test_fileset_file(self, module, fileset, test_file):
        self.init()

        # generate a minimal configuration
        cfgfile = os.path.join(self.working_dir, "filebeat.yml")
        self.render_config_template(
            template_name="filebeat_modules",
            output=cfgfile,
            index_name=self.index_name,
            elasticsearch_url=self.elasticsearch_url
        )

        self.run_on_file(
            module=module,
            fileset=fileset,
            test_file=test_file,
            cfgfile=cfgfile)

    def run_on_file(self, module, fileset, test_file, cfgfile):
        print("Testing {}/{} on {}".format(module, fileset, test_file))

        try:
            self.es.indices.delete(index=self.index_name)
        except:
            pass
        self.wait_until(lambda: not self.es.indices.exists(self.index_name))

        cmd = [
            self.filebeat, "-systemTest",
            "-e", "-d", "*", "-once",
            "-c", cfgfile,
            "-modules={}".format(module),
            "-M", "{module}.*.enabled=false".format(module=module),
            "-M", "{module}.{fileset}.enabled=true".format(
                module=module, fileset=fileset),
            "-M", "{module}.{fileset}.var.paths=[{test_file}]".format(
                module=module, fileset=fileset, test_file=test_file),
            "-M", "*.*.input.close_eof=true",
        ]

        output_path = os.path.join(self.working_dir)
        output = open(os.path.join(output_path, "output.log"), "ab")
        output.write(" ".join(cmd) + "\n")
        subprocess.Popen(cmd,
                         stdin=None,
                         stdout=output,
                         stderr=subprocess.STDOUT,
                         bufsize=0).wait()

        # Make sure index exists
        self.wait_until(lambda: self.es.indices.exists(self.index_name))

        self.es.indices.refresh(index=self.index_name)
        # Loads the first 100 events to be checked
        res = self.es.search(index=self.index_name,
                             body={"query": {"match_all": {}}, "size": 100})
        objects = [o["_source"] for o in res["hits"]["hits"]]
        assert len(objects) > 0
        for obj in objects:
            assert obj["fileset"]["module"] == module, "expected fileset.module={} but got {}".format(
                module, obj["fileset"]["module"])

            assert "error" not in obj, "not error expected but got: {}".format(
                obj)

            if (module == "auditd" and fileset == "log") \
                    or (module == "osquery" and fileset == "result"):
                # There are dynamic fields that are not documented.
                pass
            else:
                self.assert_fields_are_documented(obj)

        if os.path.exists(test_file + "-expected.json"):
            self._test_expected_events(module, test_file, objects)

    def _test_expected_events(self, module, test_file, objects):
        with open(test_file + "-expected.json", "r") as f:
            expected = json.load(f)

        assert len(expected) == len(objects), "expected {} events to compare but got {}".format(
            len(expected), len(objects))

        for ev in expected:
            found = False
            for obj in objects:

                # Skip read timestamp as always different
                del obj["read_timestamp"]

                if "event" in obj and "created" in obj["event"]:
                    del obj["event"]["created"]

                del ev["_source"]["event"]["created"]
                del ev["_source"]["host"]["name"]
                del obj["host"]["name"]
                del ev["_source"]["beat"]
                del obj["source"]
                del ev["_source"]["source"]
                del obj["beat"]
                print self.flatten_object(obj, {}, "")
                print self.flatten_object(ev["_source"], {}, "")

                # All modules in EXTENDED_COMPARE are checked in more detail
                if module in EXTENDED_COMPARE:
                    if ev["_source"] == obj:
                        found = True
                        break
                else:
                    if ev["_source"][module] == obj[module]:
                        found = True
                        break

            assert found, "The following expected object was not found:\n {}\nSearched in: \n{}".format(
                pretty_json(ev["_source"]), pretty_json(objects))


def pretty_json(obj):
    return json.dumps(obj, indent=2, separators=(',', ': '))
