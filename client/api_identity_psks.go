/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
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

// IdentityPsksApiService IdentityPsksApi service
type IdentityPsksApiService service

type IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest struct {
	ctx _context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	createNetworkWirelessSsidIdentityPsk *InlineObject160
}

func (r IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest) CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk InlineObject160) IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest {
	r.createNetworkWirelessSsidIdentityPsk = &createNetworkWirelessSsidIdentityPsk
	return r
}

func (r IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkWirelessSsidIdentityPskExecute(r)
}

/*
CreateNetworkWirelessSsidIdentityPsk Create an Identity PSK

Create an Identity PSK

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) CreateNetworkWirelessSsidIdentityPsk(ctx _context.Context, networkId string, number string) IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IdentityPsksApiService) CreateNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiApiCreateNetworkWirelessSsidIdentityPskRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.CreateNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createNetworkWirelessSsidIdentityPsk == nil {
		return localVarReturnValue, nil, reportError("createNetworkWirelessSsidIdentityPsk is required and must be specified")
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
	localVarPostBody = r.createNetworkWirelessSsidIdentityPsk
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

type IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest struct {
	ctx _context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
}


func (r IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkWirelessSsidIdentityPskExecute(r)
}

/*
DeleteNetworkWirelessSsidIdentityPsk Delete an Identity PSK

Delete an Identity PSK

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @param identityPskId
 @return IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) DeleteNetworkWirelessSsidIdentityPsk(ctx _context.Context, networkId string, number string, identityPskId string) IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
func (a *IdentityPsksApiService) DeleteNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiApiDeleteNetworkWirelessSsidIdentityPskRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.DeleteNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", _neturl.PathEscape(parameterToString(r.identityPskId, "")), -1)

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

type IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest struct {
	ctx _context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
}


func (r IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest) Execute() (InlineResponse20061, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidIdentityPskExecute(r)
}

/*
GetNetworkWirelessSsidIdentityPsk Return an Identity PSK

Return an Identity PSK

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @param identityPskId
 @return IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsk(ctx _context.Context, networkId string, number string, identityPskId string) IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
//  @return InlineResponse20061
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiApiGetNetworkWirelessSsidIdentityPskRequest) (InlineResponse20061, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20061
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.GetNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", _neturl.PathEscape(parameterToString(r.identityPskId, "")), -1)

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

type IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest struct {
	ctx _context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
}


func (r IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest) Execute() ([]InlineResponse20061, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidIdentityPsksExecute(r)
}

/*
GetNetworkWirelessSsidIdentityPsks List all Identity PSKs in a wireless network

List all Identity PSKs in a wireless network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest
*/
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsks(ctx _context.Context, networkId string, number string) IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest {
	return IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return []InlineResponse20061
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsksExecute(r IdentityPsksApiApiGetNetworkWirelessSsidIdentityPsksRequest) ([]InlineResponse20061, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20061
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.GetNetworkWirelessSsidIdentityPsks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)

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

type IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest struct {
	ctx _context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
	updateNetworkWirelessSsidIdentityPsk *InlineObject161
}

func (r IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest) UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk InlineObject161) IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest {
	r.updateNetworkWirelessSsidIdentityPsk = &updateNetworkWirelessSsidIdentityPsk
	return r
}

func (r IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidIdentityPskExecute(r)
}

/*
UpdateNetworkWirelessSsidIdentityPsk Update an Identity PSK

Update an Identity PSK

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @param identityPskId
 @return IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) UpdateNetworkWirelessSsidIdentityPsk(ctx _context.Context, networkId string, number string, identityPskId string) IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IdentityPsksApiService) UpdateNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiApiUpdateNetworkWirelessSsidIdentityPskRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.UpdateNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", _neturl.PathEscape(parameterToString(r.identityPskId, "")), -1)

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
	localVarPostBody = r.updateNetworkWirelessSsidIdentityPsk
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
