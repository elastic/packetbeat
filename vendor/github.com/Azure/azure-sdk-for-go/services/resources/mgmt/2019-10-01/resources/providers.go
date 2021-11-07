package resources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProvidersClient is the provides operations for working with resources and resource groups.
type ProvidersClient struct {
	BaseClient
}

// NewProvidersClient creates an instance of the ProvidersClient client.
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return NewProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProvidersClientWithBaseURI creates an instance of the ProvidersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return ProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified resource provider.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// expand - the $expand query parameter. For example, to include property aliases in response, use
// $expand=resourceTypes/aliases.
func (client ProvidersClient) Get(ctx context.Context, resourceProviderNamespace string, expand string) (result Provider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceProviderNamespace, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProvidersClient) GetPreparer(ctx context.Context, resourceProviderNamespace string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProvidersClient) GetResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtTenantScope gets the specified resource provider at the tenant level.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// expand - the $expand query parameter. For example, to include property aliases in response, use
// $expand=resourceTypes/aliases.
func (client ProvidersClient) GetAtTenantScope(ctx context.Context, resourceProviderNamespace string, expand string) (result Provider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.GetAtTenantScope")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAtTenantScopePreparer(ctx, resourceProviderNamespace, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "GetAtTenantScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtTenantScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "GetAtTenantScope", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtTenantScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "GetAtTenantScope", resp, "Failure responding to request")
		return
	}

	return
}

// GetAtTenantScopePreparer prepares the GetAtTenantScope request.
func (client ProvidersClient) GetAtTenantScopePreparer(ctx context.Context, resourceProviderNamespace string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{resourceProviderNamespace}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAtTenantScopeSender sends the GetAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) GetAtTenantScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtTenantScopeResponder handles the response to the GetAtTenantScope request. The method always
// closes the http.Response Body.
func (client ProvidersClient) GetAtTenantScopeResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all resource providers for a subscription.
// Parameters:
// top - the number of results to return. If null is passed returns all deployments.
// expand - the properties to include in the results. For example, use &$expand=metadata in the query string to
// retrieve resource provider metadata. To include property aliases in response, use
// $expand=resourceTypes/aliases.
func (client ProvidersClient) List(ctx context.Context, top *int32, expand string) (result ProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.List")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, top, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ProvidersClient) ListPreparer(ctx context.Context, top *int32, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProvidersClient) ListResponder(resp *http.Response) (result ProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProvidersClient) listNextResults(ctx context.Context, lastResults ProviderListResult) (result ProviderListResult, err error) {
	req, err := lastResults.providerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProvidersClient) ListComplete(ctx context.Context, top *int32, expand string) (result ProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, top, expand)
	return
}

// ListAtTenantScope gets all resource providers for the tenant.
// Parameters:
// top - the number of results to return. If null is passed returns all providers.
// expand - the properties to include in the results. For example, use &$expand=metadata in the query string to
// retrieve resource provider metadata. To include property aliases in response, use
// $expand=resourceTypes/aliases.
func (client ProvidersClient) ListAtTenantScope(ctx context.Context, top *int32, expand string) (result ProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.ListAtTenantScope")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAtTenantScopeNextResults
	req, err := client.ListAtTenantScopePreparer(ctx, top, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "ListAtTenantScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAtTenantScopeSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "ListAtTenantScope", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListAtTenantScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "ListAtTenantScope", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAtTenantScopePreparer prepares the ListAtTenantScope request.
func (client ProvidersClient) ListAtTenantScopePreparer(ctx context.Context, top *int32, expand string) (*http.Request, error) {
	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAtTenantScopeSender sends the ListAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) ListAtTenantScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAtTenantScopeResponder handles the response to the ListAtTenantScope request. The method always
// closes the http.Response Body.
func (client ProvidersClient) ListAtTenantScopeResponder(resp *http.Response) (result ProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAtTenantScopeNextResults retrieves the next set of results, if any.
func (client ProvidersClient) listAtTenantScopeNextResults(ctx context.Context, lastResults ProviderListResult) (result ProviderListResult, err error) {
	req, err := lastResults.providerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "listAtTenantScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAtTenantScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "listAtTenantScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAtTenantScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "listAtTenantScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAtTenantScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProvidersClient) ListAtTenantScopeComplete(ctx context.Context, top *int32, expand string) (result ProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.ListAtTenantScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAtTenantScope(ctx, top, expand)
	return
}

// Register registers a subscription with a resource provider.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider to register.
func (client ProvidersClient) Register(ctx context.Context, resourceProviderNamespace string) (result Provider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.Register")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegisterPreparer(ctx, resourceProviderNamespace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", resp, "Failure sending request")
		return
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", resp, "Failure responding to request")
		return
	}

	return
}

// RegisterPreparer prepares the Register request.
func (client ProvidersClient) RegisterPreparer(ctx context.Context, resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/register", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) RegisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client ProvidersClient) RegisterResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Unregister unregisters a subscription from a resource provider.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider to unregister.
func (client ProvidersClient) Unregister(ctx context.Context, resourceProviderNamespace string) (result Provider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProvidersClient.Unregister")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UnregisterPreparer(ctx, resourceProviderNamespace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", nil, "Failure preparing request")
		return
	}

	resp, err := client.UnregisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", resp, "Failure sending request")
		return
	}

	result, err = client.UnregisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", resp, "Failure responding to request")
		return
	}

	return
}

// UnregisterPreparer prepares the Unregister request.
func (client ProvidersClient) UnregisterPreparer(ctx context.Context, resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/unregister", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnregisterSender sends the Unregister request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) UnregisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UnregisterResponder handles the response to the Unregister request. The method always
// closes the http.Response Body.
func (client ProvidersClient) UnregisterResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
