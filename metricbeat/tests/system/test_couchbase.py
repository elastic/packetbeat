import os
import metricbeat
import unittest
from parameterized import parameterized


@unittest.skip("See https://github.com/elastic/beats/issues/14660")
class Test(metricbeat.BaseTest):

    COMPOSE_SERVICES = ['couchbase']
    FIELDS = ['couchbase']

    @parameterized.expand([
        ("bucket"),
        ("cluster"),
        ("node"),
    ])
    @unittest.skipUnless(metricbeat.INTEGRATION_TESTS, "integration test")
    def test_couchbase(self, metricset):
        """
        couchbase metricsets tests
        """
        self.check_metricset("couchbase", metricset, self.get_hosts(), self.FIELDS)
