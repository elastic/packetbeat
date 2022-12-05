// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package kubernetes

import (
	"testing"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes/metadata"

	"github.com/elastic/beats/v7/libbeat/common"

	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
)

func TestGenerateNodeData(t *testing.T) {
	uid := "005f3b90-4b9d-12f8-acf0-31020a840133"
	node := &kubernetes.Node{
		ObjectMeta: kubernetes.ObjectMeta{
			Name: "testnode",
			UID:  types.UID(uid),
			Labels: map[string]string{
				"foo":        "bar",
				"with-dash":  "dash-value",
				"with/slash": "some/path",
			},
			Annotations: map[string]string{
				"baz": "ban",
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		Status: v1.NodeStatus{
			Conditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionTrue}},
			Addresses:  []v1.NodeAddress{{Type: v1.NodeHostName, Address: "node1"}},
		},
	}

	data := generateNodeData(node, &Config{}, &nodeMeta{})

	mapping := map[string]interface{}{
		"node": common.MapStr{
			"uid":  string(node.GetUID()),
			"name": node.GetName(),
			"ip":   "node1",
		},
		"annotations": common.MapStr{
			"baz": "ban",
		},
		"labels": common.MapStr{
			"foo":        "bar",
			"with-dash":  "dash-value",
			"with/slash": "some/path",
		},
	}

	processors := map[string]interface{}{
		"orchestrator": common.MapStr{
			"cluster": common.MapStr{
				"name": "devcluster",
				"url":  "8.8.8.8:9090"},
		}, "kubernetes": common.MapStr{
			"labels": common.MapStr{
				"foo":        "bar",
				"with-dash":  "dash-value",
				"with/slash": "some/path",
			},
			"annotations": common.MapStr{"baz": "ban"},
			"node": common.MapStr{
				"ip":   "node1",
				"name": "testnode",
				"uid":  "005f3b90-4b9d-12f8-acf0-31020a840133"},
		},
	}
	assert.Equal(t, node, data.node)
	assert.Equal(t, mapping, data.mapping)
	for _, v := range data.processors {
		k := v["add_fields"].(map[string]interface{})
		target := k["target"].(string)
		fields := k["fields"]
		assert.Equal(t, processors[target], fields)
	}
}

type nodeMeta struct{}

// Generate generates node metadata from a resource object
// Metadata map is in the following form:
// {
// 	  "kubernetes": {},
//    "some.ecs.field": "asdf"
// }
// All Kubernetes fields that need to be stored under kubernetes. prefix are populated by
// GenerateK8s method while fields that are part of ECS are generated by GenerateECS method
func (n *nodeMeta) Generate(obj kubernetes.Resource, opts ...metadata.FieldOptions) common.MapStr {
	ecsFields := n.GenerateECS(obj)
	meta := common.MapStr{
		"kubernetes": n.GenerateK8s(obj, opts...),
	}
	meta.DeepUpdate(ecsFields)
	return meta
}

// GenerateECS generates node ECS metadata from a resource object
func (n *nodeMeta) GenerateECS(obj kubernetes.Resource) common.MapStr {
	return common.MapStr{
		"orchestrator": common.MapStr{
			"cluster": common.MapStr{
				"name": "devcluster",
				"url":  "8.8.8.8:9090",
			},
		},
	}
}

// GenerateK8s generates node metadata from a resource object
func (n *nodeMeta) GenerateK8s(obj kubernetes.Resource, opts ...metadata.FieldOptions) common.MapStr {
	k8sNode := obj.(*kubernetes.Node)
	return common.MapStr{
		"node": common.MapStr{
			"uid":  string(k8sNode.GetUID()),
			"name": k8sNode.GetName(),
			"ip":   "node1",
		},
		"labels": common.MapStr{
			"foo":        "bar",
			"with-dash":  "dash-value",
			"with/slash": "some/path",
		},
		"annotations": common.MapStr{
			"baz": "ban",
		},
	}
}

// GenerateFromName generates node metadata from a node name
func (n *nodeMeta) GenerateFromName(name string, opts ...metadata.FieldOptions) common.MapStr {
	return nil
}
