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


// AlternateManagementInterfaceApiService AlternateManagementInterfaceApi service
type AlternateManagementInterfaceApiService service

type AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest struct {
	ctx context.Context
	ApiService *AlternateManagementInterfaceApiService
	networkId string
}

func (r AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchAlternateManagementInterfaceExecute(r)
}

/*
GetNetworkSwitchAlternateManagementInterface Return the switch alternate management interface for the network

Return the switch alternate management interface for the network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest
*/
func (a *AlternateManagementInterfaceApiService) GetNetworkSwitchAlternateManagementInterface(ctx context.Context, networkId string) AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest {
	return AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlternateManagementInterfaceApiService) GetNetworkSwitchAlternateManagementInterfaceExecute(r AlternateManagementInterfaceApiGetNetworkSwitchAlternateManagementInterfaceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternateManagementInterfaceApiService.GetNetworkSwitchAlternateManagementInterface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/alternateManagementInterface"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest struct {
	ctx context.Context
	ApiService *AlternateManagementInterfaceApiService
	networkId string
}

func (r AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessAlternateManagementInterfaceExecute(r)
}

/*
GetNetworkWirelessAlternateManagementInterface Return alternate management interface and devices with IP assigned

Return alternate management interface and devices with IP assigned

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest
*/
func (a *AlternateManagementInterfaceApiService) GetNetworkWirelessAlternateManagementInterface(ctx context.Context, networkId string) AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest {
	return AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlternateManagementInterfaceApiService) GetNetworkWirelessAlternateManagementInterfaceExecute(r AlternateManagementInterfaceApiGetNetworkWirelessAlternateManagementInterfaceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternateManagementInterfaceApiService.GetNetworkWirelessAlternateManagementInterface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/alternateManagementInterface"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest struct {
	ctx context.Context
	ApiService *AlternateManagementInterfaceApiService
	networkId string
	updateNetworkSwitchAlternateManagementInterface *InlineObject110
}

func (r AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest) UpdateNetworkSwitchAlternateManagementInterface(updateNetworkSwitchAlternateManagementInterface InlineObject110) AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest {
	r.updateNetworkSwitchAlternateManagementInterface = &updateNetworkSwitchAlternateManagementInterface
	return r
}

func (r AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchAlternateManagementInterfaceExecute(r)
}

/*
UpdateNetworkSwitchAlternateManagementInterface Update the switch alternate management interface for the network

Update the switch alternate management interface for the network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest
*/
func (a *AlternateManagementInterfaceApiService) UpdateNetworkSwitchAlternateManagementInterface(ctx context.Context, networkId string) AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest {
	return AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlternateManagementInterfaceApiService) UpdateNetworkSwitchAlternateManagementInterfaceExecute(r AlternateManagementInterfaceApiUpdateNetworkSwitchAlternateManagementInterfaceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternateManagementInterfaceApiService.UpdateNetworkSwitchAlternateManagementInterface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/alternateManagementInterface"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchAlternateManagementInterface
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

type AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest struct {
	ctx context.Context
	ApiService *AlternateManagementInterfaceApiService
	networkId string
	updateNetworkWirelessAlternateManagementInterface *InlineObject146
}

func (r AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest) UpdateNetworkWirelessAlternateManagementInterface(updateNetworkWirelessAlternateManagementInterface InlineObject146) AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest {
	r.updateNetworkWirelessAlternateManagementInterface = &updateNetworkWirelessAlternateManagementInterface
	return r
}

func (r AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessAlternateManagementInterfaceExecute(r)
}

/*
UpdateNetworkWirelessAlternateManagementInterface Update alternate management interface and device static IP

Update alternate management interface and device static IP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest
*/
func (a *AlternateManagementInterfaceApiService) UpdateNetworkWirelessAlternateManagementInterface(ctx context.Context, networkId string) AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest {
	return AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlternateManagementInterfaceApiService) UpdateNetworkWirelessAlternateManagementInterfaceExecute(r AlternateManagementInterfaceApiUpdateNetworkWirelessAlternateManagementInterfaceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternateManagementInterfaceApiService.UpdateNetworkWirelessAlternateManagementInterface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/alternateManagementInterface"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	localVarPostBody = r.updateNetworkWirelessAlternateManagementInterface
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
