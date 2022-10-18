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

// RulesApiService RulesApi service
type RulesApiService service

type RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest struct {
	ctx _context.Context
	ApiService *RulesApiService
	networkId string
}


func (r RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkApplianceTrafficShapingRulesExecute(r)
}

/*
GetNetworkApplianceTrafficShapingRules Display the traffic shaping settings rules for an MX network

Display the traffic shaping settings rules for an MX network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest
*/
func (a *RulesApiService) GetNetworkApplianceTrafficShapingRules(ctx _context.Context, networkId string) RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest {
	return RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RulesApiService) GetNetworkApplianceTrafficShapingRulesExecute(r RulesApiApiGetNetworkApplianceTrafficShapingRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesApiService.GetNetworkApplianceTrafficShapingRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/rules"
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

type RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest struct {
	ctx _context.Context
	ApiService *RulesApiService
	networkId string
	number string
}


func (r RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidTrafficShapingRulesExecute(r)
}

/*
GetNetworkWirelessSsidTrafficShapingRules Display the traffic shaping settings for a SSID on an MR network

Display the traffic shaping settings for a SSID on an MR network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest
*/
func (a *RulesApiService) GetNetworkWirelessSsidTrafficShapingRules(ctx _context.Context, networkId string, number string) RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest {
	return RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RulesApiService) GetNetworkWirelessSsidTrafficShapingRulesExecute(r RulesApiApiGetNetworkWirelessSsidTrafficShapingRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesApiService.GetNetworkWirelessSsidTrafficShapingRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules"
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

type RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest struct {
	ctx _context.Context
	ApiService *RulesApiService
	networkId string
	updateNetworkApplianceTrafficShapingRules *InlineObject51
}

func (r RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest) UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules InlineObject51) RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest {
	r.updateNetworkApplianceTrafficShapingRules = &updateNetworkApplianceTrafficShapingRules
	return r
}

func (r RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkApplianceTrafficShapingRulesExecute(r)
}

/*
UpdateNetworkApplianceTrafficShapingRules Update the traffic shaping settings rules for an MX network

Update the traffic shaping settings rules for an MX network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest
*/
func (a *RulesApiService) UpdateNetworkApplianceTrafficShapingRules(ctx _context.Context, networkId string) RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest {
	return RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RulesApiService) UpdateNetworkApplianceTrafficShapingRulesExecute(r RulesApiApiUpdateNetworkApplianceTrafficShapingRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesApiService.UpdateNetworkApplianceTrafficShapingRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/rules"
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
	localVarPostBody = r.updateNetworkApplianceTrafficShapingRules
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

type RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct {
	ctx _context.Context
	ApiService *RulesApiService
	networkId string
	number string
	updateNetworkWirelessSsidTrafficShapingRules *InlineObject164
}

func (r RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest) UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules InlineObject164) RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	r.updateNetworkWirelessSsidTrafficShapingRules = &updateNetworkWirelessSsidTrafficShapingRules
	return r
}

func (r RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidTrafficShapingRulesExecute(r)
}

/*
UpdateNetworkWirelessSsidTrafficShapingRules Update the traffic shaping settings for an SSID on an MR network

Update the traffic shaping settings for an SSID on an MR network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest
*/
func (a *RulesApiService) UpdateNetworkWirelessSsidTrafficShapingRules(ctx _context.Context, networkId string, number string) RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	return RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RulesApiService) UpdateNetworkWirelessSsidTrafficShapingRulesExecute(r RulesApiApiUpdateNetworkWirelessSsidTrafficShapingRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesApiService.UpdateNetworkWirelessSsidTrafficShapingRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", _neturl.PathEscape(parameterToString(r.number, "")), -1)

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
	localVarPostBody = r.updateNetworkWirelessSsidTrafficShapingRules
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
