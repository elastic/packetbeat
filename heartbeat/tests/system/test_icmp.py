import logging
import os
import platform
import socket
import subprocess
import sys
import time
import unittest
import re
from beat.beat import INTEGRATION_TESTS
from elasticsearch import Elasticsearch
from heartbeat import BaseTest


class Test(BaseTest):
    def test_base(self):
        """
        Basic test with icmp root non privilege ICMP test.

        """

        config = {
            "monitors": [
                {
                    "type": "icmp",
                    "schedule": "*/5 * * * * * *",
                    "hosts": ["127.0.0.1"],
                }
            ]
        }

        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*",
            **config
        )

        proc = self.start_beat()

        def has_started_message(): return self.log_contains("ICMP loop successfully initialized")

        def has_failed_message(): return self.log_contains("Failed to initialize ICMP loop")

        # We don't know if the system tests are running is configured to support or not support ping, but we can at least check that the ICMP loop
        # was initiated. In the future we should start up VMs with the correct perms configured and be more specific. In addition to that
        # we should run pings on those machines and make sure they work.
        self.wait_until(lambda: has_started_message() or has_failed_message(), 30)

        self.wait_until(lambda: self.output_has(lines=1))
        output = self.read_output()
        monitor_status = output[0]["monitor.status"]
        monitor_error = output[0]["error.message"]
        if has_failed_message():
            assert monitor_status == "down"
            self.assertRegex(monitor_error, ".*Insufficient privileges to perform ICMP ping.*")
        else:
            assert monitor_status == "up"
