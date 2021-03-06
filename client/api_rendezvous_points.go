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

// RendezvousPointsApiService RendezvousPointsApi service
type RendezvousPointsApiService service

type RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx _context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	createNetworkSwitchRoutingMulticastRendezvousPoint *InlineObject108
}

func (r RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) CreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint InlineObject108) RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.createNetworkSwitchRoutingMulticastRendezvousPoint = &createNetworkSwitchRoutingMulticastRendezvousPoint
	return r
}

func (r RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
CreateNetworkSwitchRoutingMulticastRendezvousPoint Create a multicast rendezvous point

Create a multicast rendezvous point

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx _context.Context, networkId string) RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.CreateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createNetworkSwitchRoutingMulticastRendezvousPoint == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchRoutingMulticastRendezvousPoint is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchRoutingMulticastRendezvousPoint
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

type RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx _context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
}


func (r RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
DeleteNetworkSwitchRoutingMulticastRendezvousPoint Delete a multicast rendezvous point

Delete a multicast rendezvous point

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx _context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
func (a *RendezvousPointsApiService) DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.DeleteNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", _neturl.PathEscape(parameterToString(r.rendezvousPointId, "")), -1)

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

type RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx _context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
}


func (r RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoint Return a multicast rendezvous point

Return a multicast rendezvous point

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx _context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.GetNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", _neturl.PathEscape(parameterToString(r.rendezvousPointId, "")), -1)

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

type RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct {
	ctx _context.Context
	ApiService *RendezvousPointsApiService
	networkId string
}


func (r RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) Execute() ([][]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoints List multicast rendezvous points

List multicast rendezvous points

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest
*/
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx _context.Context, networkId string) RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest {
	return RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return [][]map[string]interface{}
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r RendezvousPointsApiApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) ([][]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  [][]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.GetNetworkSwitchRoutingMulticastRendezvousPoints")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

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

type RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx _context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
	updateNetworkSwitchRoutingMulticastRendezvousPoint *InlineObject109
}

func (r RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint InlineObject109) RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.updateNetworkSwitchRoutingMulticastRendezvousPoint = &updateNetworkSwitchRoutingMulticastRendezvousPoint
	return r
}

func (r RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
UpdateNetworkSwitchRoutingMulticastRendezvousPoint Update a multicast rendezvous point

Update a multicast rendezvous point

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx _context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.UpdateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", _neturl.PathEscape(parameterToString(r.rendezvousPointId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateNetworkSwitchRoutingMulticastRendezvousPoint == nil {
		return localVarReturnValue, nil, reportError("updateNetworkSwitchRoutingMulticastRendezvousPoint is required and must be specified")
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
	localVarPostBody = r.updateNetworkSwitchRoutingMulticastRendezvousPoint
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
