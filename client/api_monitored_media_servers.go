/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
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


// MonitoredMediaServersApiService MonitoredMediaServersApi service
type MonitoredMediaServersApiService service

type MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	createOrganizationInsightMonitoredMediaServer *InlineObject266
}

func (r MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest) CreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer InlineObject266) MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest {
	r.createOrganizationInsightMonitoredMediaServer = &createOrganizationInsightMonitoredMediaServer
	return r
}

func (r MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest) Execute() (*InlineResponse200291, *http.Response, error) {
	return r.ApiService.CreateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
CreateOrganizationInsightMonitoredMediaServer Add a media server to be monitored for this organization

Add a media server to be monitored for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) CreateOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string) MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200291
func (a *MonitoredMediaServersApiService) CreateOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiCreateOrganizationInsightMonitoredMediaServerRequest) (*InlineResponse200291, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200291
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.CreateOrganizationInsightMonitoredMediaServer")
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

type MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
}

func (r MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
DeleteOrganizationInsightMonitoredMediaServer Delete a monitored media server from this organization

Delete a monitored media server from this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param monitoredMediaServerId Monitored media server ID
 @return MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) DeleteOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
func (a *MonitoredMediaServersApiService) DeleteOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiDeleteOrganizationInsightMonitoredMediaServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.DeleteOrganizationInsightMonitoredMediaServer")
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

type MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
}

func (r MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest) Execute() (*InlineResponse200291, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServer Return a monitored media server for this organization

Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param monitoredMediaServerId Monitored media server ID
 @return MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return InlineResponse200291
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServerRequest) (*InlineResponse200291, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200291
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.GetOrganizationInsightMonitoredMediaServer")
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

type MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest struct {
	ctx context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
}

func (r MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest) Execute() ([]InlineResponse200291, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServersExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServers List the monitored media servers for this organization

List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest
*/
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServers(ctx context.Context, organizationId string) MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest {
	return MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200291
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServersExecute(r MonitoredMediaServersApiGetOrganizationInsightMonitoredMediaServersRequest) ([]InlineResponse200291, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200291
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.GetOrganizationInsightMonitoredMediaServers")
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

type MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
	updateOrganizationInsightMonitoredMediaServer *InlineObject267
}

func (r MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest) UpdateOrganizationInsightMonitoredMediaServer(updateOrganizationInsightMonitoredMediaServer InlineObject267) MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	r.updateOrganizationInsightMonitoredMediaServer = &updateOrganizationInsightMonitoredMediaServer
	return r
}

func (r MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest) Execute() (*InlineResponse200291, *http.Response, error) {
	return r.ApiService.UpdateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
UpdateOrganizationInsightMonitoredMediaServer Update a monitored media server for this organization

Update a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param monitoredMediaServerId Monitored media server ID
 @return MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) UpdateOrganizationInsightMonitoredMediaServer(ctx context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return InlineResponse200291
func (a *MonitoredMediaServersApiService) UpdateOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiUpdateOrganizationInsightMonitoredMediaServerRequest) (*InlineResponse200291, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200291
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.UpdateOrganizationInsightMonitoredMediaServer")
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
