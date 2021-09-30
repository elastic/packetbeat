import os
import unittest
from filebeat import BaseTest
from beat.beat import INTEGRATION_TESTS
from beat import common_tests


class Test(BaseTest, common_tests.TestExportsMixin, common_tests.TestDashboardMixin):

    def test_base(self):
        """
        Test if the basic fields exist.
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/test.log",
        )

        with open(self.working_dir + "/test.log", "w") as f:
            f.write("test message\n")

        filebeat = self.start_beat()
        self.wait_until(lambda: self.output_has(lines=1))
        filebeat.check_kill_and_wait()

        output = self.read_output()[0]
        assert "@timestamp" in output
        assert "input.type" in output

    @unittest.skipUnless(INTEGRATION_TESTS, "integration test")
    def test_index_management(self):
        """
        Test that the template can be loaded with `setup --index-management`
        """
        es_url = self.get_elasticsearch_url()
        es = self.get_elasticsearch_instance(url=es_url)
        self.render_config_template(
            elasticsearch={"host": es_url},
        )
        exit_code = self.run_beat(extra_args=["setup", "--index-management"])

        assert exit_code == 0
        assert self.log_contains('Loaded index template')
        assert len(es.cat.templates(name='filebeat-*', h='name')) > 0

    @unittest.skipUnless(INTEGRATION_TESTS, "integration test")
    def test_template_migration(self):
        """
        Test that the template can be loaded with `setup --template`
        """
        es_url = self.get_elasticsearch_url()
        es = self.get_elasticsearch_instance(url=es_url)
        self.render_config_template(
            elasticsearch={"host": es_url},
        )
        exit_code = self.run_beat(extra_args=["setup", "--template",
                                              "-E", "setup.template.overwrite=true", "-E", "migration.6_to_7.enabled=true"])

        assert exit_code == 0
        assert self.log_contains('Loaded index template')
        assert len(es.cat.templates(name='filebeat-*', h='name')) > 0
