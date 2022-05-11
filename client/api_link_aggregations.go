/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
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

// LinkAggregationsApiService LinkAggregationsApi service
type LinkAggregationsApiService service

type LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest struct {
	ctx _context.Context
	ApiService *LinkAggregationsApiService
	networkId string
	createNetworkSwitchLinkAggregation *InlineObject97
}

func (r LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest) CreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation InlineObject97) LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest {
	r.createNetworkSwitchLinkAggregation = &createNetworkSwitchLinkAggregation
	return r
}

func (r LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkSwitchLinkAggregationExecute(r)
}

/*
CreateNetworkSwitchLinkAggregation Create a link aggregation group

Create a link aggregation group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest
*/
func (a *LinkAggregationsApiService) CreateNetworkSwitchLinkAggregation(ctx _context.Context, networkId string) LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest {
	return LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LinkAggregationsApiService) CreateNetworkSwitchLinkAggregationExecute(r LinkAggregationsApiApiCreateNetworkSwitchLinkAggregationRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkAggregationsApiService.CreateNetworkSwitchLinkAggregation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/linkAggregations"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

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
	localVarPostBody = r.createNetworkSwitchLinkAggregation
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

type LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest struct {
	ctx _context.Context
	ApiService *LinkAggregationsApiService
	networkId string
	linkAggregationId string
}


func (r LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkSwitchLinkAggregationExecute(r)
}

/*
DeleteNetworkSwitchLinkAggregation Split a link aggregation group into separate ports

Split a link aggregation group into separate ports

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param linkAggregationId
 @return LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest
*/
func (a *LinkAggregationsApiService) DeleteNetworkSwitchLinkAggregation(ctx _context.Context, networkId string, linkAggregationId string) LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest {
	return LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		linkAggregationId: linkAggregationId,
	}
}

// Execute executes the request
func (a *LinkAggregationsApiService) DeleteNetworkSwitchLinkAggregationExecute(r LinkAggregationsApiApiDeleteNetworkSwitchLinkAggregationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkAggregationsApiService.DeleteNetworkSwitchLinkAggregation")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkAggregationId"+"}", _neturl.PathEscape(parameterToString(r.linkAggregationId, "")), -1)

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

type LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest struct {
	ctx _context.Context
	ApiService *LinkAggregationsApiService
	networkId string
}


func (r LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchLinkAggregationsExecute(r)
}

/*
GetNetworkSwitchLinkAggregations List link aggregation groups

List link aggregation groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest
*/
func (a *LinkAggregationsApiService) GetNetworkSwitchLinkAggregations(ctx _context.Context, networkId string) LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest {
	return LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LinkAggregationsApiService) GetNetworkSwitchLinkAggregationsExecute(r LinkAggregationsApiApiGetNetworkSwitchLinkAggregationsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkAggregationsApiService.GetNetworkSwitchLinkAggregations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/linkAggregations"
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

type LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest struct {
	ctx _context.Context
	ApiService *LinkAggregationsApiService
	networkId string
	linkAggregationId string
	updateNetworkSwitchLinkAggregation *InlineObject98
}

func (r LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest) UpdateNetworkSwitchLinkAggregation(updateNetworkSwitchLinkAggregation InlineObject98) LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest {
	r.updateNetworkSwitchLinkAggregation = &updateNetworkSwitchLinkAggregation
	return r
}

func (r LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkSwitchLinkAggregationExecute(r)
}

/*
UpdateNetworkSwitchLinkAggregation Update a link aggregation group

Update a link aggregation group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param linkAggregationId
 @return LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest
*/
func (a *LinkAggregationsApiService) UpdateNetworkSwitchLinkAggregation(ctx _context.Context, networkId string, linkAggregationId string) LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest {
	return LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		linkAggregationId: linkAggregationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LinkAggregationsApiService) UpdateNetworkSwitchLinkAggregationExecute(r LinkAggregationsApiApiUpdateNetworkSwitchLinkAggregationRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkAggregationsApiService.UpdateNetworkSwitchLinkAggregation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkAggregationId"+"}", _neturl.PathEscape(parameterToString(r.linkAggregationId, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchLinkAggregation
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
