package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// UsagesClient is the the Azure SQL Database management API provides a RESTful set of web services that interact with
// Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and delete
// databases.
type UsagesClient struct {
	BaseClient
}

// NewUsagesClient creates an instance of the UsagesClient client.
func NewUsagesClient(subscriptionID string) UsagesClient {
	return NewUsagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsagesClientWithBaseURI creates an instance of the UsagesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return UsagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByInstancePool gets all instance pool usage metrics
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// instancePoolName - the name of the instance pool to be retrieved.
// expandChildren - optional request parameter to include managed instance usages within the instance pool.
func (client UsagesClient) ListByInstancePool(ctx context.Context, resourceGroupName string, instancePoolName string, expandChildren *bool) (result UsageListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesClient.ListByInstancePool")
		defer func() {
			sc := -1
			if result.ulr.Response.Response != nil {
				sc = result.ulr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInstancePoolNextResults
	req, err := client.ListByInstancePoolPreparer(ctx, resourceGroupName, instancePoolName, expandChildren)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.UsagesClient", "ListByInstancePool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstancePoolSender(req)
	if err != nil {
		result.ulr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.UsagesClient", "ListByInstancePool", resp, "Failure sending request")
		return
	}

	result.ulr, err = client.ListByInstancePoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.UsagesClient", "ListByInstancePool", resp, "Failure responding to request")
		return
	}
	if result.ulr.hasNextLink() && result.ulr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByInstancePoolPreparer prepares the ListByInstancePool request.
func (client UsagesClient) ListByInstancePoolPreparer(ctx context.Context, resourceGroupName string, instancePoolName string, expandChildren *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instancePoolName":  autorest.Encode("path", instancePoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if expandChildren != nil {
		queryParameters["expandChildren"] = autorest.Encode("query", *expandChildren)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstancePoolSender sends the ListByInstancePool request. The method will close the
// http.Response Body if it receives an error.
func (client UsagesClient) ListByInstancePoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByInstancePoolResponder handles the response to the ListByInstancePool request. The method always
// closes the http.Response Body.
func (client UsagesClient) ListByInstancePoolResponder(resp *http.Response) (result UsageListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstancePoolNextResults retrieves the next set of results, if any.
func (client UsagesClient) listByInstancePoolNextResults(ctx context.Context, lastResults UsageListResult) (result UsageListResult, err error) {
	req, err := lastResults.usageListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.UsagesClient", "listByInstancePoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstancePoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.UsagesClient", "listByInstancePoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstancePoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.UsagesClient", "listByInstancePoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstancePoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsagesClient) ListByInstancePoolComplete(ctx context.Context, resourceGroupName string, instancePoolName string, expandChildren *bool) (result UsageListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesClient.ListByInstancePool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInstancePool(ctx, resourceGroupName, instancePoolName, expandChildren)
	return
}
