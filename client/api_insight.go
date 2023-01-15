/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
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


// InsightApiService InsightApi service
type InsightApiService service

type InsightApiCreateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
	createOrganizationInsightMonitoredMediaServer *InlineObject195
}

func (r InsightApiCreateOrganizationInsightMonitoredMediaServerRequest) CreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer InlineObject195) InsightApiCreateOrganizationInsightMonitoredMediaServerRequest {
	r.createOrganizationInsightMonitoredMediaServer = &createOrganizationInsightMonitoredMediaServer
	return r
}

func (r InsightApiCreateOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
CreateOrganizationInsightMonitoredMediaServer Add a media server to be monitored for this organization

Add a media server to be monitored for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InsightApiCreateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *InsightApiService) CreateOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string) InsightApiCreateOrganizationInsightMonitoredMediaServerRequest {
	return InsightApiCreateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InsightApiService) CreateOrganizationInsightMonitoredMediaServerExecute(r InsightApiCreateOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.CreateOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInsightMonitoredMediaServer == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInsightMonitoredMediaServer is required and must be specified")
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
	localVarPostBody = r.createOrganizationInsightMonitoredMediaServer
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

type InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
	monitoredMediaServerId string
}

func (r InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
DeleteOrganizationInsightMonitoredMediaServer Delete a monitored media server from this organization

Delete a monitored media server from this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest
*/
func (a *InsightApiService) DeleteOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest {
	return InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
func (a *InsightApiService) DeleteOrganizationInsightMonitoredMediaServerExecute(r InsightApiDeleteOrganizationInsightMonitoredMediaServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.DeleteOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", url.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

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

type InsightApiGetNetworkInsightApplicationHealthByTimeRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	networkId string
	applicationId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
func (r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) T0(t0 string) InsightApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) T1(t1 string) InsightApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
func (r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) Timespan(timespan float32) InsightApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300.
func (r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) Resolution(resolution int32) InsightApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.resolution = &resolution
	return r
}

func (r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) Execute() ([]InlineResponse20027, *http.Response, error) {
	return r.ApiService.GetNetworkInsightApplicationHealthByTimeExecute(r)
}

/*
GetNetworkInsightApplicationHealthByTime Get application health by time

Get application health by time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param applicationId
 @return InsightApiGetNetworkInsightApplicationHealthByTimeRequest
*/
func (a *InsightApiService) GetNetworkInsightApplicationHealthByTime(ctx context.Context, networkId string, applicationId string) InsightApiGetNetworkInsightApplicationHealthByTimeRequest {
	return InsightApiGetNetworkInsightApplicationHealthByTimeRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20027
func (a *InsightApiService) GetNetworkInsightApplicationHealthByTimeExecute(r InsightApiGetNetworkInsightApplicationHealthByTimeRequest) ([]InlineResponse20027, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20027
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.GetNetworkInsightApplicationHealthByTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/insight/applications/{applicationId}/healthByTime"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
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

type InsightApiGetOrganizationInsightApplicationsRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
}

func (r InsightApiGetOrganizationInsightApplicationsRequest) Execute() ([]InlineResponse20092, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightApplicationsExecute(r)
}

/*
GetOrganizationInsightApplications List all Insight tracked applications

List all Insight tracked applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InsightApiGetOrganizationInsightApplicationsRequest
*/
func (a *InsightApiService) GetOrganizationInsightApplications(ctx context.Context, organizationId string) InsightApiGetOrganizationInsightApplicationsRequest {
	return InsightApiGetOrganizationInsightApplicationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20092
func (a *InsightApiService) GetOrganizationInsightApplicationsExecute(r InsightApiGetOrganizationInsightApplicationsRequest) ([]InlineResponse20092, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20092
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.GetOrganizationInsightApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type InsightApiGetOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
	monitoredMediaServerId string
}

func (r InsightApiGetOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServer Return a monitored media server for this organization

Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return InsightApiGetOrganizationInsightMonitoredMediaServerRequest
*/
func (a *InsightApiService) GetOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) InsightApiGetOrganizationInsightMonitoredMediaServerRequest {
	return InsightApiGetOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InsightApiService) GetOrganizationInsightMonitoredMediaServerExecute(r InsightApiGetOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.GetOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", url.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

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

type InsightApiGetOrganizationInsightMonitoredMediaServersRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
}

func (r InsightApiGetOrganizationInsightMonitoredMediaServersRequest) Execute() ([]InlineResponse20093, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServersExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServers List the monitored media servers for this organization

List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InsightApiGetOrganizationInsightMonitoredMediaServersRequest
*/
func (a *InsightApiService) GetOrganizationInsightMonitoredMediaServers(ctx context.Context, organizationId string) InsightApiGetOrganizationInsightMonitoredMediaServersRequest {
	return InsightApiGetOrganizationInsightMonitoredMediaServersRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20093
func (a *InsightApiService) GetOrganizationInsightMonitoredMediaServersExecute(r InsightApiGetOrganizationInsightMonitoredMediaServersRequest) ([]InlineResponse20093, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20093
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.GetOrganizationInsightMonitoredMediaServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *InsightApiService
	organizationId string
	monitoredMediaServerId string
	updateOrganizationInsightMonitoredMediaServer *InlineObject196
}

func (r InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest) UpdateOrganizationInsightMonitoredMediaServer(updateOrganizationInsightMonitoredMediaServer InlineObject196) InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	r.updateOrganizationInsightMonitoredMediaServer = &updateOrganizationInsightMonitoredMediaServer
	return r
}

func (r InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
UpdateOrganizationInsightMonitoredMediaServer Update a monitored media server for this organization

Update a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *InsightApiService) UpdateOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	return InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InsightApiService) UpdateOrganizationInsightMonitoredMediaServerExecute(r InsightApiUpdateOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightApiService.UpdateOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", url.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

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
	localVarPostBody = r.updateOrganizationInsightMonitoredMediaServer
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
