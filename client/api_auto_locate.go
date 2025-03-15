/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
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


// AutoLocateApiService AutoLocateApi service
type AutoLocateApiService service

type AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	networkId string
	batchNetworkFloorPlansAutoLocateJobs *InlineObject103
}

func (r AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest) BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs InlineObject103) AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest {
	r.batchNetworkFloorPlansAutoLocateJobs = &batchNetworkFloorPlansAutoLocateJobs
	return r
}

func (r AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest) Execute() (*InlineResponse200101, *http.Response, error) {
	return r.ApiService.BatchNetworkFloorPlansAutoLocateJobsExecute(r)
}

/*
BatchNetworkFloorPlansAutoLocateJobs Schedule auto locate jobs for one or more floor plans in a network

Schedule auto locate jobs for one or more floor plans in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest
*/
func (a *AutoLocateApiService) BatchNetworkFloorPlansAutoLocateJobs(ctx context.Context, networkId string) AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest {
	return AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200101
func (a *AutoLocateApiService) BatchNetworkFloorPlansAutoLocateJobsExecute(r AutoLocateApiBatchNetworkFloorPlansAutoLocateJobsRequest) (*InlineResponse200101, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.BatchNetworkFloorPlansAutoLocateJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/autoLocate/jobs/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchNetworkFloorPlansAutoLocateJobs == nil {
		return localVarReturnValue, nil, reportError("batchNetworkFloorPlansAutoLocateJobs is required and must be specified")
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
	localVarPostBody = r.batchNetworkFloorPlansAutoLocateJobs
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	networkId string
	jobId string
}

func (r AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelNetworkFloorPlansAutoLocateJobExecute(r)
}

/*
CancelNetworkFloorPlansAutoLocateJob Cancel a scheduled or running auto locate job

Cancel a scheduled or running auto locate job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param jobId Job ID
 @return AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest
*/
func (a *AutoLocateApiService) CancelNetworkFloorPlansAutoLocateJob(ctx context.Context, networkId string, jobId string) AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest {
	return AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		jobId: jobId,
	}
}

// Execute executes the request
func (a *AutoLocateApiService) CancelNetworkFloorPlansAutoLocateJobExecute(r AutoLocateApiCancelNetworkFloorPlansAutoLocateJobRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.CancelNetworkFloorPlansAutoLocateJob")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	floorPlanIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) PerPage(perPage int32) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) StartingAfter(startingAfter string) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) EndingBefore(endingBefore string) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter devices by one or more network IDs
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) NetworkIds(networkIds []string) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter devices by one or more floorplan IDs
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) FloorPlanIds(floorPlanIds []string) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	r.floorPlanIds = &floorPlanIds
	return r
}

func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) Execute() ([]InlineResponse200277, *http.Response, error) {
	return r.ApiService.GetOrganizationFloorPlansAutoLocateDevicesExecute(r)
}

/*
GetOrganizationFloorPlansAutoLocateDevices List auto locate details for each device in your organization

List auto locate details for each device in your organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest
*/
func (a *AutoLocateApiService) GetOrganizationFloorPlansAutoLocateDevices(ctx context.Context, organizationId string) AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest {
	return AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200277
func (a *AutoLocateApiService) GetOrganizationFloorPlansAutoLocateDevicesExecute(r AutoLocateApiGetOrganizationFloorPlansAutoLocateDevicesRequest) ([]InlineResponse200277, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200277
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.GetOrganizationFloorPlansAutoLocateDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/floorPlans/autoLocate/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.floorPlanIds != nil {
		localVarQueryParams.Add("floorPlanIds", parameterToString(*r.floorPlanIds, "csv"))
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
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	floorPlanIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) PerPage(perPage int32) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) StartingAfter(startingAfter string) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) EndingBefore(endingBefore string) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter floorplans by one or more network IDs
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) NetworkIds(networkIds []string) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter floorplans by one or more floorplan IDs
func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) FloorPlanIds(floorPlanIds []string) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	r.floorPlanIds = &floorPlanIds
	return r
}

func (r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) Execute() ([]InlineResponse200278, *http.Response, error) {
	return r.ApiService.GetOrganizationFloorPlansAutoLocateStatusesExecute(r)
}

/*
GetOrganizationFloorPlansAutoLocateStatuses List the status of auto locate for each floorplan in your organization

List the status of auto locate for each floorplan in your organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest
*/
func (a *AutoLocateApiService) GetOrganizationFloorPlansAutoLocateStatuses(ctx context.Context, organizationId string) AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest {
	return AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200278
func (a *AutoLocateApiService) GetOrganizationFloorPlansAutoLocateStatusesExecute(r AutoLocateApiGetOrganizationFloorPlansAutoLocateStatusesRequest) ([]InlineResponse200278, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200278
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.GetOrganizationFloorPlansAutoLocateStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/floorPlans/autoLocate/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.floorPlanIds != nil {
		localVarQueryParams.Add("floorPlanIds", parameterToString(*r.floorPlanIds, "csv"))
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
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	networkId string
	jobId string
	publishNetworkFloorPlansAutoLocateJob *InlineObject104
}

func (r AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest) PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob InlineObject104) AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest {
	r.publishNetworkFloorPlansAutoLocateJob = &publishNetworkFloorPlansAutoLocateJob
	return r
}

func (r AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest) Execute() (*InlineResponse200102, *http.Response, error) {
	return r.ApiService.PublishNetworkFloorPlansAutoLocateJobExecute(r)
}

/*
PublishNetworkFloorPlansAutoLocateJob Update the status of a finished auto locate job to be published, and update device locations

Update the status of a finished auto locate job to be published, and update device locations

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param jobId Job ID
 @return AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest
*/
func (a *AutoLocateApiService) PublishNetworkFloorPlansAutoLocateJob(ctx context.Context, networkId string, jobId string) AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest {
	return AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return InlineResponse200102
func (a *AutoLocateApiService) PublishNetworkFloorPlansAutoLocateJobExecute(r AutoLocateApiPublishNetworkFloorPlansAutoLocateJobRequest) (*InlineResponse200102, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200102
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.PublishNetworkFloorPlansAutoLocateJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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
	localVarPostBody = r.publishNetworkFloorPlansAutoLocateJob
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest struct {
	ctx context.Context
	ApiService *AutoLocateApiService
	networkId string
	jobId string
	recalculateNetworkFloorPlansAutoLocateJob *InlineObject105
}

func (r AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest) RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob InlineObject105) AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest {
	r.recalculateNetworkFloorPlansAutoLocateJob = &recalculateNetworkFloorPlansAutoLocateJob
	return r
}

func (r AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest) Execute() (*InlineResponse200103, *http.Response, error) {
	return r.ApiService.RecalculateNetworkFloorPlansAutoLocateJobExecute(r)
}

/*
RecalculateNetworkFloorPlansAutoLocateJob Trigger auto locate recalculation for a job, and optionally set anchors

Trigger auto locate recalculation for a job, and optionally set anchors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param jobId Job ID
 @return AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest
*/
func (a *AutoLocateApiService) RecalculateNetworkFloorPlansAutoLocateJob(ctx context.Context, networkId string, jobId string) AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest {
	return AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return InlineResponse200103
func (a *AutoLocateApiService) RecalculateNetworkFloorPlansAutoLocateJobExecute(r AutoLocateApiRecalculateNetworkFloorPlansAutoLocateJobRequest) (*InlineResponse200103, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200103
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoLocateApiService.RecalculateNetworkFloorPlansAutoLocateJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/recalculate"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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
	localVarPostBody = r.recalculateNetworkFloorPlansAutoLocateJob
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
