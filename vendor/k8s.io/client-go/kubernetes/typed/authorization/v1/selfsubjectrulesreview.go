/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// SelfSubjectRulesReviewsGetter has a method to return a SelfSubjectRulesReviewInterface.
// A group's client should implement this interface.
type SelfSubjectRulesReviewsGetter interface {
	SelfSubjectRulesReviews() SelfSubjectRulesReviewInterface
}

// SelfSubjectRulesReviewInterface has methods to work with SelfSubjectRulesReview resources.
type SelfSubjectRulesReviewInterface interface {
	Create(ctx context.Context, selfSubjectRulesReview *v1.SelfSubjectRulesReview, opts metav1.CreateOptions) (*v1.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewExpansion
}

// selfSubjectRulesReviews implements SelfSubjectRulesReviewInterface
type selfSubjectRulesReviews struct {
	client rest.Interface
}

// newSelfSubjectRulesReviews returns a SelfSubjectRulesReviews
func newSelfSubjectRulesReviews(c *AuthorizationV1Client) *selfSubjectRulesReviews {
	return &selfSubjectRulesReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a selfSubjectRulesReview and creates it.  Returns the server's representation of the selfSubjectRulesReview, and an error, if there is any.
func (c *selfSubjectRulesReviews) Create(ctx context.Context, selfSubjectRulesReview *v1.SelfSubjectRulesReview, opts metav1.CreateOptions) (result *v1.SelfSubjectRulesReview, err error) {
	result = &v1.SelfSubjectRulesReview{}
	err = c.client.Post().
		Resource("selfsubjectrulesreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectRulesReview).
		Do(ctx).
		Into(result)
	return
}
