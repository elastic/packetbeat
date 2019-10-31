// Copyright 2013-2017 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

import (
	"strings"
)

// RemoveTask is used to poll for UDF registration completion.
type RemoveTask struct {
	*baseTask

	packageName string
}

// NewRemoveTask initializes a RemoveTask with fields needed to query server nodes.
func NewRemoveTask(cluster *Cluster, packageName string) *RemoveTask {
	return &RemoveTask{
		baseTask:    newTask(cluster, false),
		packageName: packageName,
	}
}

// IsDone will query all nodes for task completion status.
func (tskr *RemoveTask) IsDone() (bool, error) {
	command := "udf-list"
	nodes := tskr.cluster.GetNodes()
	done := false

	for _, node := range nodes {
		responseMap, err := node.RequestInfo(command)
		if err != nil {
			return false, err
		}

		for _, response := range responseMap {
			find := "filename=" + tskr.packageName
			index := strings.Index(response, find)

			if index >= 0 {
				return false, nil
			}
			done = true
		}
	}
	return done, nil
}

// OnComplete returns a channel that will be closed as soon as the task is finished.
// If an error is encountered during operation, an error will be sent on the channel.
func (tskr *RemoveTask) OnComplete() chan error {
	return tskr.onComplete(tskr)
}
