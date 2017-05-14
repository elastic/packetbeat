import sys
sys.path.append('../../vendor/github.com/elastic/beats/libbeat/tests/system')
from beat.beat import TestCase

class BaseTest(TestCase):

    @classmethod
    def setUpClass(self):
        self.beat_name = "{{cookiecutter.beat}}"
        self.build_path = "../../build/system-tests/"
        self.beat_path = "../../{{cookiecutter.beat}}.test"
