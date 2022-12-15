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


// OnboardingApiService OnboardingApi service
type OnboardingApiService service

type OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	ctx context.Context
	ApiService *OnboardingApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringImport *InlineObject198
}

func (r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport InlineObject198) OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringImport = &createOrganizationInventoryOnboardingCloudMonitoringImport
	return r
}

func (r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Execute() ([]InlineResponse2015, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
*/
func (a *OnboardingApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx context.Context, organizationId string) OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2015
func (a *OnboardingApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ([]InlineResponse2015, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2015
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnboardingApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImport")
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

type OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	ctx context.Context
	ApiService *OnboardingApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringPrepare *InlineObject199
}

func (r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare InlineObject199) OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringPrepare = &createOrganizationInventoryOnboardingCloudMonitoringPrepare
	return r
}

func (r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Execute() ([]InlineResponse2016, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringPrepare Initiates or updates an import session

Initiates or updates an import session. An import ID will be generated and used when you are ready to commit the import.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
*/
func (a *OnboardingApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx context.Context, organizationId string) OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2016
func (a *OnboardingApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r OnboardingApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ([]InlineResponse2016, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnboardingApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare")
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

type OnboardingApiGetOrganizationCameraOnboardingStatusesRequest struct {
	ctx context.Context
	ApiService *OnboardingApiService
	organizationId string
	serials *[]string
	networkIds *[]string
}

// A list of serial numbers. The returned cameras will be filtered to only include these serials.
func (r OnboardingApiGetOrganizationCameraOnboardingStatusesRequest) Serials(serials []string) OnboardingApiGetOrganizationCameraOnboardingStatusesRequest {
	r.serials = &serials
	return r
}

// A list of network IDs. The returned cameras will be filtered to only include these networks.
func (r OnboardingApiGetOrganizationCameraOnboardingStatusesRequest) NetworkIds(networkIds []string) OnboardingApiGetOrganizationCameraOnboardingStatusesRequest {
	r.networkIds = &networkIds
	return r
}

func (r OnboardingApiGetOrganizationCameraOnboardingStatusesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraOnboardingStatusesExecute(r)
}

/*
GetOrganizationCameraOnboardingStatuses Fetch onboarding status of cameras

Fetch onboarding status of cameras

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OnboardingApiGetOrganizationCameraOnboardingStatusesRequest
*/
func (a *OnboardingApiService) GetOrganizationCameraOnboardingStatuses(ctx context.Context, organizationId string) OnboardingApiGetOrganizationCameraOnboardingStatusesRequest {
	return OnboardingApiGetOrganizationCameraOnboardingStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *OnboardingApiService) GetOrganizationCameraOnboardingStatusesExecute(r OnboardingApiGetOrganizationCameraOnboardingStatusesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnboardingApiService.GetOrganizationCameraOnboardingStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/onboarding/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
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

type OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct {
	ctx context.Context
	ApiService *OnboardingApiService
	organizationId string
	importIds *[]string
}

// import ids from an imports
func (r OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ImportIds(importIds []string) OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	r.importIds = &importIds
	return r
}

func (r OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) Execute() ([]InlineResponse20092, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation

Check the status of a committed Import operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest
*/
func (a *OnboardingApiService) GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx context.Context, organizationId string) OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	return OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20092
func (a *OnboardingApiService) GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r OnboardingApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ([]InlineResponse20092, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20092
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnboardingApiService.GetOrganizationInventoryOnboardingCloudMonitoringImports")
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

type OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest struct {
	ctx context.Context
	ApiService *OnboardingApiService
	organizationId string
	updateOrganizationCameraOnboardingStatuses *InlineObject187
}

func (r OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest) UpdateOrganizationCameraOnboardingStatuses(updateOrganizationCameraOnboardingStatuses InlineObject187) OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest {
	r.updateOrganizationCameraOnboardingStatuses = &updateOrganizationCameraOnboardingStatuses
	return r
}

func (r OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateOrganizationCameraOnboardingStatusesExecute(r)
}

/*
UpdateOrganizationCameraOnboardingStatuses Notify that credential handoff to camera has completed

Notify that credential handoff to camera has completed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest
*/
func (a *OnboardingApiService) UpdateOrganizationCameraOnboardingStatuses(ctx context.Context, organizationId string) OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest {
	return OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *OnboardingApiService) UpdateOrganizationCameraOnboardingStatusesExecute(r OnboardingApiUpdateOrganizationCameraOnboardingStatusesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnboardingApiService.UpdateOrganizationCameraOnboardingStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/onboarding/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateOrganizationCameraOnboardingStatuses
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
