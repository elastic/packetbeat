from filebeat import BaseTest
import os
import socket
import tempfile


class Test(BaseTest):
    """
    Test filebeat UNIX input
    """

    def test_unix_with_newline_delimiter(self):
        """
        Test UNIX input with a new line delimiter
        """
        self.send_events_with_delimiter("\n")

    def test_unix_with_custom_char_delimiter(self):
        """
        Test UNIX input with a custom single char delimiter
        """
        self.send_events_with_delimiter(";")

    def test_unix_with_custom_word_delimiter(self):
        """
        Test UNIX input with a custom single char delimiter
        """
        self.send_events_with_delimiter("<END>")

    def send_events_with_delimiter(self, delimiter):
        # we create the socket in a temporary directory because
        # go will fail to create a unix socket if the path length
        # is longer than 108 characters. See https://github.com/golang/go/issues/6895
        with tempfile.TemporaryDirectory() as tempdir:
            path = os.path.join(tempdir, "filebeat.sock")
            input_raw = """
- type: unix
  path: "{}"
  enabled: true
"""

            # Use default of \n and stripping \r
            if delimiter != "":
                input_raw += "\n  line_delimiter: {}".format(delimiter)

            input_raw = input_raw.format(path)

            self.render_config_template(
                input_raw=input_raw,
                inputs=False,
            )

            filebeat = self.start_beat()

            self.wait_until(lambda: self.log_contains("Started listening for UNIX connection"))

            sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)  # UNIX
            sock.connect(path)

            for n in range(0, 2):
                sock.send(bytes("Hello World: " + str(n) + delimiter, "utf-8"))

            self.wait_until(lambda: self.output_count(lambda x: x >= 2))

            filebeat.check_kill_and_wait()

            output = self.read_output()

            assert len(output) == 2
            assert output[0]["input.type"] == "unix"

            sock.close()
