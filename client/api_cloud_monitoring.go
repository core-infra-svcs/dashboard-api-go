/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CloudMonitoringApiService CloudMonitoringApi service
type CloudMonitoringApiService service

type CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringImport *InlineObject198
}

func (r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport InlineObject198) CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringImport = &createOrganizationInventoryOnboardingCloudMonitoringImport
	return r
}

func (r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Execute() ([]InlineResponse2015, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
*/
func (a *CloudMonitoringApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx context.Context, organizationId string) CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2015
func (a *CloudMonitoringApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ([]InlineResponse2015, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2015
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringImport == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringImport is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringImport
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringPrepare *InlineObject199
}

func (r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare InlineObject199) CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringPrepare = &createOrganizationInventoryOnboardingCloudMonitoringPrepare
	return r
}

func (r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Execute() ([]InlineResponse2016, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringPrepare Initiates or updates an import session

Initiates or updates an import session. An import ID will be generated and used when you are ready to commit the import.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
*/
func (a *CloudMonitoringApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx context.Context, organizationId string) CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2016
func (a *CloudMonitoringApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r CloudMonitoringApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ([]InlineResponse2016, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringPrepare == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringPrepare is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringPrepare
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringApiService
	organizationId string
	importIds *[]string
}

// import ids from an imports
func (r CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ImportIds(importIds []string) CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	r.importIds = &importIds
	return r
}

func (r CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) Execute() ([]InlineResponse20092, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation

Check the status of a committed Import operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest
*/
func (a *CloudMonitoringApiService) GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx context.Context, organizationId string) CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	return CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20092
func (a *CloudMonitoringApiService) GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r CloudMonitoringApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ([]InlineResponse20092, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20092
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringApiService.GetOrganizationInventoryOnboardingCloudMonitoringImports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importIds == nil {
		return localVarReturnValue, nil, reportError("importIds is required and must be specified")
	}

	localVarQueryParams.Add("importIds", parameterToString(*r.importIds, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
