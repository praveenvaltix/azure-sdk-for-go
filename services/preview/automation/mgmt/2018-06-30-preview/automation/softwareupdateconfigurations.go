package automation

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SoftwareUpdateConfigurationsClient is the automation Client
type SoftwareUpdateConfigurationsClient struct {
	BaseClient
}

// NewSoftwareUpdateConfigurationsClient creates an instance of the SoftwareUpdateConfigurationsClient client.
func NewSoftwareUpdateConfigurationsClient(subscriptionID string, countType1 CountType) SoftwareUpdateConfigurationsClient {
	return NewSoftwareUpdateConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID, countType1)
}

// NewSoftwareUpdateConfigurationsClientWithBaseURI creates an instance of the SoftwareUpdateConfigurationsClient
// client.
func NewSoftwareUpdateConfigurationsClientWithBaseURI(baseURI string, subscriptionID string, countType1 CountType) SoftwareUpdateConfigurationsClient {
	return SoftwareUpdateConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID, countType1)}
}

// Create create a new software update configuration with the name given in the URI.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// softwareUpdateConfigurationName - the name of the software update configuration to be created.
// parameters - request body.
// clientRequestID - identifies this specific client request.
func (client SoftwareUpdateConfigurationsClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, clientRequestID string) (result SoftwareUpdateConfiguration, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SoftwareUpdateConfigurationProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SoftwareUpdateConfigurationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, parameters, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SoftwareUpdateConfigurationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":           autorest.Encode("path", automationAccountName),
		"resourceGroupName":               autorest.Encode("path", resourceGroupName),
		"softwareUpdateConfigurationName": autorest.Encode("path", softwareUpdateConfigurationName),
		"subscriptionId":                  autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationsClient) CreateResponder(resp *http.Response) (result SoftwareUpdateConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a specific software update configuration.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// softwareUpdateConfigurationName - the name of the software update configuration to be created.
// clientRequestID - identifies this specific client request.
func (client SoftwareUpdateConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, clientRequestID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SoftwareUpdateConfigurationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SoftwareUpdateConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":           autorest.Encode("path", automationAccountName),
		"resourceGroupName":               autorest.Encode("path", resourceGroupName),
		"softwareUpdateConfigurationName": autorest.Encode("path", softwareUpdateConfigurationName),
		"subscriptionId":                  autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetByName get a single software update configuration by name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// softwareUpdateConfigurationName - the name of the software update configuration to be created.
// clientRequestID - identifies this specific client request.
func (client SoftwareUpdateConfigurationsClient) GetByName(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, clientRequestID string) (result SoftwareUpdateConfiguration, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SoftwareUpdateConfigurationsClient", "GetByName", err.Error())
	}

	req, err := client.GetByNamePreparer(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "GetByName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "GetByName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "GetByName", resp, "Failure responding to request")
	}

	return
}

// GetByNamePreparer prepares the GetByName request.
func (client SoftwareUpdateConfigurationsClient) GetByNamePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":           autorest.Encode("path", automationAccountName),
		"resourceGroupName":               autorest.Encode("path", resourceGroupName),
		"softwareUpdateConfigurationName": autorest.Encode("path", softwareUpdateConfigurationName),
		"subscriptionId":                  autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByNameSender sends the GetByName request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationsClient) GetByNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByNameResponder handles the response to the GetByName request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationsClient) GetByNameResponder(resp *http.Response) (result SoftwareUpdateConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all software update configurations for the account.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// clientRequestID - identifies this specific client request.
// filter - the filter to apply on the operation.
func (client SoftwareUpdateConfigurationsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, clientRequestID string, filter string) (result SoftwareUpdateConfigurationListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SoftwareUpdateConfigurationsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, automationAccountName, clientRequestID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SoftwareUpdateConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, clientRequestID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationsClient) ListResponder(resp *http.Response) (result SoftwareUpdateConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
