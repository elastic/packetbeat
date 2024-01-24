// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package kprobes

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func Test_InotifyWatcher(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping on non-linux")
	}

	tmpDir, err := os.MkdirTemp("", "kprobe_unit_test")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	watcher, err := newInotifyWatcher()
	require.NoError(t, err)

	added, err := watcher.Add(1, 1, tmpDir)
	require.NoError(t, err)
	require.True(t, added)

	added, err = watcher.Add(1, 1, filepath.Join(tmpDir, "test"))
	require.NoError(t, err)
	require.False(t, added)

	added, err = watcher.Add(2, 2, tmpDir)
	require.NoError(t, err)
	require.False(t, added)

	tmpDir2, err := os.MkdirTemp("", "kprobe_unit_test")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(tmpDir2)
	}()
	added, err = watcher.Add(2, 2, tmpDir2)
	require.NoError(t, err)
	require.True(t, added)

	require.NoError(t, watcher.Close())

	_, err = watcher.Add(1, 1, tmpDir)
	require.Error(t, err)
}

func Test_InotifyWatcher_Add_Err(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping on non-linux")
	}

	watcher, err := newInotifyWatcher()
	require.NoError(t, err)

	inotifyAddWatch = func(fd int, pathname string, mask uint32) (int, error) {
		return -1, os.ErrInvalid
	}
	defer func() {
		inotifyAddWatch = unix.InotifyAddWatch
	}()

	_, err = watcher.Add(1, 1, "non_existent")
	require.ErrorIs(t, err, os.ErrInvalid)

	require.NoError(t, watcher.Close())
}

func Test_InotifyWatcher_Close_Err(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping on non-linux")
	}

	tmpDir, err := os.MkdirTemp("", "kprobe_unit_test")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	watcher, err := newInotifyWatcher()
	require.NoError(t, err)

	added, err := watcher.Add(1, 1, tmpDir)
	require.NoError(t, err)
	require.True(t, added)

	err = os.RemoveAll(tmpDir)
	require.NoError(t, err)

	require.Error(t, watcher.Close())
}
