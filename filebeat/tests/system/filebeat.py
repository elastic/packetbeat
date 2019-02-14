import json
import os
import stat
import sys

curdir = os.path.dirname(__file__)
sys.path.append(os.path.join(curdir, '../../../libbeat/tests/system'))

from beat.beat import TestCase, TimeoutError, REGEXP_TYPE

default_registry_file = 'registry/filebeat/data.json'


class BaseTest(TestCase):

    @classmethod
    def setUpClass(self):
        if not hasattr(self, "beat_name"):
            self.beat_name = "filebeat"
        if not hasattr(self, "beat_path"):
            self.beat_path = os.path.abspath(os.path.join(curdir, "../../"))

        super(BaseTest, self).setUpClass()

    def registry(self, name=None, data_path=None):
        data_path = data_path if data_path else self.working_dir
        return Registry(data_path, name)

    def log_access(self, file=None):
        file = file if file else self.beat_name + ".log"
        return LogState(os.path.join(self.working_dir, file))

    def input_logs(self):
        return InputLogs(os.path.join(self.working_dir, "log"))

    def has_registry(self, name=None, data_path=None):
        return self.registry(name, data_path).exists()

    def get_registry(self, name=None, data_path=None, filter=None):
        reg = self.registry(name, data_path)
        self.wait_until(reg.exists)
        return reg.load(filter=filter)

    def get_registry_entry_by_path(self, path):
        """
        Fetches the registry file and checks if an entry for the given path exists
        If the path exists, the state for the given path is returned
        If a path exists multiple times (which is possible because of file rotation)
        the most recent version is returned
        """

        def hasPath(entry):
            return entry["source"] == path

        entries = self.get_registry(filter=hasPath)
        entries.sort(key=lambda x: x["timestamp"])

        # return entry with largest timestamp
        return None if len(entries) == 0 else entries[-1]

    def file_permissions(self, path):
        full_path = os.path.join(self.working_dir, path)
        return oct(stat.S_IMODE(os.lstat(full_path).st_mode))

class InputLogs:
    def __init__(self, home):
        self.home = home
        if not os.path.isdir(self.home):
            os.mkdir(self.home)

    def write(self, name, contents):
        self._write_to(name, 'w', contents)

    def append(self, name, contents):
        self._write_to(name, 'a', contents)

    def size(self, name):
        return os.path.getsize(os.path.join(self.home, name))

    def _write_to(self, name, mode, contents):
        path = os.path.join(self.home, name)
        with open(path, mode) as f:
            f.write(contents)

    def remove(self, name):
        path = os.path.join(self.home, name)
        os.remove(path)


class Registry:
    def __init__(self, home, name=None):
        if not name:
            name = default_registry_file
        self.path = os.path.join(home, name)

    def exists(self):
        return os.path.isfile(self.path)

    def load(self, filter=None):
        with open(self.path) as f:
            entries = json.load(f)

        if filter:
            entries = [x for x in entries if filter(x)]
        return entries

    def count(self, filter=None):
        if not self.exists():
            return 0
        return len(self.load(filter=filter))


class LogState:
    def __init__(self, path):
        self.path = path
        self.off = 0

    def checkpoint(self):
        self.off = os.path.getsize(self.path)

    def lines(self, filter=None):
        if not filter:
            filter = lambda x: True
        with open(self.path, "r") as f:
            f.seek(self.off)
            return [l for l in f if filter(l)]

    def contains(self, msg, ignore_case=False, count=1):
        if ignore_case:
            msg = msg.lower()

        if type(msg) == REGEXP_TYPE:
            match = lambda x: msg.search(x) is not None
        else:
            match = lambda x: x.find(msg)

        pred = match
        if ignore_case:
            pred = lambda x: match(x.lower())

        return len(self.lines(filter=pred)) >= count

    def next(self, msg, ignore_case=False, count=1):
        ok = self.contains(msg, ignore_case, count)
        if ok:
            self.checkpoint()
        return ok

    def nextCheck(self, msg, ignore_case=False, count=1):
        return lambda: self.next(msg, ignore_case, count)
