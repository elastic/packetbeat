from heartbeat import BaseTest
import urllib2
import json
import nose.tools


class Test(BaseTest):
    def __init__(self, *args):
        self.proc = None
        super(Test, self).__init__(*args)

    def test_telemetry(self):
        """
        Test that telemetry metrics are correctly registered and increment / decrement
        """
        server = self.start_server("hello world", 200)
        try:
            self.setup_dynamic(["-E", "http.enabled=true"])

            cfg_file = "test.yml"

            self.write_dyn_config(
                cfg_file, self.http_cfg("http://localhost:8185")
            )

            self.wait_until(lambda: self.output_has(lines=1))

            self.assert_state({
                "http": {
                    "monitor_starts": 1,
                    "monitor_stops": 0,
                    "endpoint_starts": 1,
                    "endpoint_stops": 0,
                }
            })

            tcp_hosts = ["localhost:8185", "localhost:12345"]

            self.write_dyn_config(
                cfg_file, self.tcp_cfg(*tcp_hosts)
            )

            for tcp_host in tcp_hosts:
                self.wait_until(lambda: self.log_contains(
                    "Start job 'tcp-tcp@{}".format(tcp_host)))

            init_lines = self.output_lines()
            self.wait_until(lambda: self.output_has(lines=init_lines+2))

            self.assert_state({
                "http": {
                    "monitor_starts": 1,
                    "monitor_stops": 1,
                    "endpoint_starts": 1,
                    "endpoint_stops": 1,
                },
                "tcp": {
                    "monitor_starts": 1,
                    "monitor_stops": 0,
                    "endpoint_starts": 2,
                    "endpoint_stops": 0,
                }
            })
        finally:
            self.proc.check_kill_and_wait()
            server.shutdown()

    @staticmethod
    def assert_state(expected={}):
        stats = json.loads(urllib2.urlopen(
            "http://localhost:5066/state").read())

        for proto in ("http", "tcp", "icmp"):
            proto_expected = expected.get(proto, {})
            nose.tools.assert_dict_equal(stats['heartbeat'][proto], {
                'monitor_starts': proto_expected.get("monitor_starts", 0),
                'monitor_stops': proto_expected.get("monitor_stops", 0),
                'endpoint_starts': proto_expected.get("endpoint_starts", 0),
                'endpoint_stops': proto_expected.get("endpoint_stops", 0),
            })
