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

// SplashAuthorizationStatusApiService SplashAuthorizationStatusApi service
type SplashAuthorizationStatusApiService service

type SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest struct {
	ctx _context.Context
	ApiService *SplashAuthorizationStatusApiService
	networkId string
	clientId string
}


func (r SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkClientSplashAuthorizationStatusExecute(r)
}

/*
GetNetworkClientSplashAuthorizationStatus Return the splash authorization for a client, for each SSID they've associated with through splash

Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param clientId
 @return SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest
*/
func (a *SplashAuthorizationStatusApiService) GetNetworkClientSplashAuthorizationStatus(ctx _context.Context, networkId string, clientId string) SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest {
	return SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SplashAuthorizationStatusApiService) GetNetworkClientSplashAuthorizationStatusExecute(r SplashAuthorizationStatusApiApiGetNetworkClientSplashAuthorizationStatusRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplashAuthorizationStatusApiService.GetNetworkClientSplashAuthorizationStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/clients/{clientId}/splashAuthorizationStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.PathEscape(parameterToString(r.clientId, "")), -1)

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

type SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest struct {
	ctx _context.Context
	ApiService *SplashAuthorizationStatusApiService
	networkId string
	clientId string
	updateNetworkClientSplashAuthorizationStatus *InlineObject64
}

func (r SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest) UpdateNetworkClientSplashAuthorizationStatus(updateNetworkClientSplashAuthorizationStatus InlineObject64) SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest {
	r.updateNetworkClientSplashAuthorizationStatus = &updateNetworkClientSplashAuthorizationStatus
	return r
}

func (r SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkClientSplashAuthorizationStatusExecute(r)
}

/*
UpdateNetworkClientSplashAuthorizationStatus Update a client's splash authorization

Update a client's splash authorization. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param clientId
 @return SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest
*/
func (a *SplashAuthorizationStatusApiService) UpdateNetworkClientSplashAuthorizationStatus(ctx _context.Context, networkId string, clientId string) SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest {
	return SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SplashAuthorizationStatusApiService) UpdateNetworkClientSplashAuthorizationStatusExecute(r SplashAuthorizationStatusApiApiUpdateNetworkClientSplashAuthorizationStatusRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplashAuthorizationStatusApiService.UpdateNetworkClientSplashAuthorizationStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/clients/{clientId}/splashAuthorizationStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.PathEscape(parameterToString(r.clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateNetworkClientSplashAuthorizationStatus == nil {
		return localVarReturnValue, nil, reportError("updateNetworkClientSplashAuthorizationStatus is required and must be specified")
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
	localVarPostBody = r.updateNetworkClientSplashAuthorizationStatus
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
