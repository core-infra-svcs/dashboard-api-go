/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// MonitoredMediaServersApiService MonitoredMediaServersApi service
type MonitoredMediaServersApiService service

type MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx _context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	createOrganizationInsightMonitoredMediaServer *InlineObject176
}

func (r MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest) CreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer InlineObject176) MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest {
	r.createOrganizationInsightMonitoredMediaServer = &createOrganizationInsightMonitoredMediaServer
	return r
}

func (r MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
CreateOrganizationInsightMonitoredMediaServer Add a media server to be monitored for this organization

Add a media server to be monitored for this organization. Only valid for organizations with Meraki Insight.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) CreateOrganizationInsightMonitoredMediaServer(ctx _context.Context, organizationId string) MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MonitoredMediaServersApiService) CreateOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiApiCreateOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.CreateOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest struct {
	ctx _context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
}


func (r MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
DeleteOrganizationInsightMonitoredMediaServer Delete a monitored media server from this organization

Delete a monitored media server from this organization. Only valid for organizations with Meraki Insight.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) DeleteOrganizationInsightMonitoredMediaServer(ctx _context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
func (a *MonitoredMediaServersApiService) DeleteOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiApiDeleteOrganizationInsightMonitoredMediaServerRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.DeleteOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", _neturl.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest struct {
	ctx _context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
}


func (r MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServer Return a monitored media server for this organization

Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServer(ctx _context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.GetOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", _neturl.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest struct {
	ctx _context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
}


func (r MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationInsightMonitoredMediaServersExecute(r)
}

/*
GetOrganizationInsightMonitoredMediaServers List the monitored media servers for this organization

List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest
*/
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServers(ctx _context.Context, organizationId string) MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest {
	return MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *MonitoredMediaServersApiService) GetOrganizationInsightMonitoredMediaServersExecute(r MonitoredMediaServersApiApiGetOrganizationInsightMonitoredMediaServersRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.GetOrganizationInsightMonitoredMediaServers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest struct {
	ctx _context.Context
	ApiService *MonitoredMediaServersApiService
	organizationId string
	monitoredMediaServerId string
	updateOrganizationInsightMonitoredMediaServer *InlineObject177
}

func (r MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest) UpdateOrganizationInsightMonitoredMediaServer(updateOrganizationInsightMonitoredMediaServer InlineObject177) MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	r.updateOrganizationInsightMonitoredMediaServer = &updateOrganizationInsightMonitoredMediaServer
	return r
}

func (r MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateOrganizationInsightMonitoredMediaServerExecute(r)
}

/*
UpdateOrganizationInsightMonitoredMediaServer Update a monitored media server for this organization

Update a monitored media server for this organization. Only valid for organizations with Meraki Insight.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param monitoredMediaServerId
 @return MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest
*/
func (a *MonitoredMediaServersApiService) UpdateOrganizationInsightMonitoredMediaServer(ctx _context.Context, organizationId string, monitoredMediaServerId string) MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest {
	return MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		monitoredMediaServerId: monitoredMediaServerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MonitoredMediaServersApiService) UpdateOrganizationInsightMonitoredMediaServerExecute(r MonitoredMediaServersApiApiUpdateOrganizationInsightMonitoredMediaServerRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredMediaServersApiService.UpdateOrganizationInsightMonitoredMediaServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"monitoredMediaServerId"+"}", _neturl.PathEscape(parameterToString(r.monitoredMediaServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
