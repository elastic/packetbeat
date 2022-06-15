package consumption

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// UsageDetailsClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type UsageDetailsClient struct {
	BaseClient
}

// NewUsageDetailsClient creates an instance of the UsageDetailsClient client.
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return NewUsageDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageDetailsClientWithBaseURI creates an instance of the UsageDetailsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return UsageDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or
// later. For more information on using this API, including how to specify a date range, please see:
// https://docs.microsoft.com/en-us/azure/cost-management-billing/costs/manage-automation
// Parameters:
// scope - the scope associated with usage details operations. This includes '/subscriptions/{subscriptionId}/'
// for subscription scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing
// Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope and
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope. For
// subscription, billing account, department, enrollment account and management group, you can also add billing
// period to the scope using '/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'. For e.g. to
// specify billing period at department scope use
// '/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'.
// Also, Modern Commerce Account scopes are '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
// for billingAccount scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners.
// expand - may be used to expand the properties/additionalInfo or properties/meterDetails within a list of
// usage details. By default, these fields are not included when listing usage details.
// filter - may be used to filter usageDetails by properties/resourceGroup, properties/instanceName,
// properties/resourceId, properties/chargeType, properties/reservationId, properties/publisherType or tags.
// The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:). PublisherType
// Filter accepts two values azure and marketplace and it is currently supported for Web Direct Offer Type
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
// metric - allows to select different type of cost/usage records.
func (client UsageDetailsClient) List(ctx context.Context, scope string, expand string, filter string, skiptoken string, top *int32, metric Metrictype, startDate string, endDate string) (result UsageDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.List")
		defer func() {
			sc := -1
			if result.udlr.Response.Response != nil {
				sc = result.udlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.UsageDetailsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, expand, filter, skiptoken, top, metric, startDate, endDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.udlr.hasNextLink() && result.udlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageDetailsClient) ListPreparer(ctx context.Context, scope string, expand string, filter string, skiptoken string, top *int32, metric Metrictype, startDate string, endDate string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(string(metric)) > 0 {
		queryParameters["metric"] = autorest.Encode("query", metric)
	}
	if len(startDate) > 0 {
		queryParameters["startDate"] = autorest.Encode("query", startDate)
	}
	if len(endDate) > 0 {
		queryParameters["endDate"] = autorest.Encode("query", endDate)
	}

	fmt.Println(queryParameters)

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageDetailsClient) ListResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UsageDetailsClient) listNextResults(ctx context.Context, lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.usageDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// // ListComplete enumerates all values, automatically crossing page boundaries as required.
// func (client UsageDetailsClient) ListComplete(ctx context.Context, scope string, expand string, filter string, skiptoken string, top *int32, metric Metrictype, startDate string endDate string) (result UsageDetailsListResultIterator, err error) {
// 	if tracing.IsEnabled() {
// 		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.List")
// 		defer func() {
// 			sc := -1
// 			if result.Response().Response.Response != nil {
// 				sc = result.page.Response().Response.Response.StatusCode
// 			}
// 			tracing.EndSpan(ctx, sc, err)
// 		}()
// 	}
// 	result.page, err = client.List(ctx, scope, expand, filter, skiptoken, top, metric, sta)
// 	return
// }
