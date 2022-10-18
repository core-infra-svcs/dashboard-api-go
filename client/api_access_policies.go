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

// AccessPoliciesApiService AccessPoliciesApi service
type AccessPoliciesApiService service

type AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest struct {
	ctx _context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	createNetworkSwitchAccessPolicy *InlineObject110
}

func (r AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest) CreateNetworkSwitchAccessPolicy(createNetworkSwitchAccessPolicy InlineObject110) AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest {
	r.createNetworkSwitchAccessPolicy = &createNetworkSwitchAccessPolicy
	return r
}

func (r AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest) Execute() (InlineResponse20046, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkSwitchAccessPolicyExecute(r)
}

/*
CreateNetworkSwitchAccessPolicy Create an access policy for a switch network

Create an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) CreateNetworkSwitchAccessPolicy(ctx _context.Context, networkId string) AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20046
func (a *AccessPoliciesApiService) CreateNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiApiCreateNetworkSwitchAccessPolicyRequest) (InlineResponse20046, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20046
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.CreateNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createNetworkSwitchAccessPolicy == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchAccessPolicy is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchAccessPolicy
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

type AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest struct {
	ctx _context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
}


func (r AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkSwitchAccessPolicyExecute(r)
}

/*
DeleteNetworkSwitchAccessPolicy Delete an access policy for a switch network

Delete an access policy for a switch network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) DeleteNetworkSwitchAccessPolicy(ctx _context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
func (a *AccessPoliciesApiService) DeleteNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiApiDeleteNetworkSwitchAccessPolicyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.DeleteNetworkSwitchAccessPolicy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", _neturl.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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

type AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest struct {
	ctx _context.Context
	ApiService *AccessPoliciesApiService
	networkId string
}


func (r AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest) Execute() ([]InlineResponse20046, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchAccessPoliciesExecute(r)
}

/*
GetNetworkSwitchAccessPolicies List the access policies for a switch network

List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest
*/
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicies(ctx _context.Context, networkId string) AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest {
	return AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20046
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPoliciesExecute(r AccessPoliciesApiApiGetNetworkSwitchAccessPoliciesRequest) ([]InlineResponse20046, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20046
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.GetNetworkSwitchAccessPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies"
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

type AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest struct {
	ctx _context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
}


func (r AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest) Execute() (InlineResponse20046, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchAccessPolicyExecute(r)
}

/*
GetNetworkSwitchAccessPolicy Return a specific access policy for a switch network

Return a specific access policy for a switch network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicy(ctx _context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
//  @return InlineResponse20046
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiApiGetNetworkSwitchAccessPolicyRequest) (InlineResponse20046, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20046
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.GetNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", _neturl.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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

type AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest struct {
	ctx _context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
	updateNetworkSwitchAccessPolicy *InlineObject111
}

func (r AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest) UpdateNetworkSwitchAccessPolicy(updateNetworkSwitchAccessPolicy InlineObject111) AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest {
	r.updateNetworkSwitchAccessPolicy = &updateNetworkSwitchAccessPolicy
	return r
}

func (r AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest) Execute() (InlineResponse20046, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkSwitchAccessPolicyExecute(r)
}

/*
UpdateNetworkSwitchAccessPolicy Update an access policy for a switch network

Update an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) UpdateNetworkSwitchAccessPolicy(ctx _context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
//  @return InlineResponse20046
func (a *AccessPoliciesApiService) UpdateNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiApiUpdateNetworkSwitchAccessPolicyRequest) (InlineResponse20046, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20046
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.UpdateNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", _neturl.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchAccessPolicy
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
