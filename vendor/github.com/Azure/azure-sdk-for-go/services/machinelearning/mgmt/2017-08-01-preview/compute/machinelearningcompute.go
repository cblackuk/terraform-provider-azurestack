package compute

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// MachineLearningComputeClient is the these APIs allow end users to operate on Azure Machine Learning Compute
// resources. They support the following operations:<ul><li>Create or update a cluster</li><li>Get a
// cluster</li><li>Patch a cluster</li><li>Delete a cluster</li><li>Get keys for a cluster</li><li>Check if updates are
// available for system services in a cluster</li><li>Update system services in a cluster</li><li>Get all clusters in a
// resource group</li><li>Get all clusters in a subscription</li></ul>
type MachineLearningComputeClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// NewMachineLearningComputeClient creates an instance of the MachineLearningComputeClient client.
func NewMachineLearningComputeClient(subscriptionID string) MachineLearningComputeClient {
	return NewMachineLearningComputeClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// NewMachineLearningComputeClientWithBaseURI creates an instance of the MachineLearningComputeClient client.
func NewMachineLearningComputeClientWithBaseURI(baseURI string, subscriptionID string) MachineLearningComputeClient {
	return MachineLearningComputeClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// ListAvailableOperations gets all available operations.
func (client MachineLearningComputeClient) ListAvailableOperations(ctx context.Context) (result AvailableOperations, err error) {
	req, err := client.ListAvailableOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.MachineLearningComputeClient", "ListAvailableOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.MachineLearningComputeClient", "ListAvailableOperations", resp, "Failure sending request")
		return
	}

	result, err = client.ListAvailableOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.MachineLearningComputeClient", "ListAvailableOperations", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// ListAvailableOperationsPreparer prepares the ListAvailableOperations request.
func (client MachineLearningComputeClient) ListAvailableOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.MachineLearningCompute/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// ListAvailableOperationsSender sends the ListAvailableOperations request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) ListAvailableOperationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute instead.
// ListAvailableOperationsResponder handles the response to the ListAvailableOperations request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) ListAvailableOperationsResponder(resp *http.Response) (result AvailableOperations, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
