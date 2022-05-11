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

// L3FirewallRulesApiService L3FirewallRulesApi service
type L3FirewallRulesApiService service

type L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest struct {
	ctx _context.Context
	ApiService *L3FirewallRulesApiService
	networkId string
}


func (r L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallL3FirewallRulesExecute(r)
}

/*
GetNetworkApplianceFirewallL3FirewallRules Return the L3 firewall rules for an MX network

Return the L3 firewall rules for an MX network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest
*/
func (a *L3FirewallRulesApiService) GetNetworkApplianceFirewallL3FirewallRules(ctx _context.Context, networkId string) L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest {
	return L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L3FirewallRulesApiService) GetNetworkApplianceFirewallL3FirewallRulesExecute(r L3FirewallRulesApiApiGetNetworkApplianceFirewallL3FirewallRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L3FirewallRulesApiService.GetNetworkApplianceFirewallL3FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/l3FirewallRules"
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

type L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest struct {
	ctx _context.Context
	ApiService *L3FirewallRulesApiService
	networkId string
	number string
}


func (r L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidFirewallL3FirewallRulesExecute(r)
}

/*
GetNetworkWirelessSsidFirewallL3FirewallRules Return the L3 firewall rules for an SSID on an MR network

Return the L3 firewall rules for an SSID on an MR network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest
*/
func (a *L3FirewallRulesApiService) GetNetworkWirelessSsidFirewallL3FirewallRules(ctx _context.Context, networkId string, number string) L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	return L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L3FirewallRulesApiService) GetNetworkWirelessSsidFirewallL3FirewallRulesExecute(r L3FirewallRulesApiApiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L3FirewallRulesApiService.GetNetworkWirelessSsidFirewallL3FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules"
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

type L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest struct {
	ctx _context.Context
	ApiService *L3FirewallRulesApiService
	networkId string
	updateNetworkApplianceFirewallL3FirewallRules *InlineObject30
}

func (r L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest) UpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules InlineObject30) L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest {
	r.updateNetworkApplianceFirewallL3FirewallRules = &updateNetworkApplianceFirewallL3FirewallRules
	return r
}

func (r L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkApplianceFirewallL3FirewallRulesExecute(r)
}

/*
UpdateNetworkApplianceFirewallL3FirewallRules Update the L3 firewall rules of an MX network

Update the L3 firewall rules of an MX network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest
*/
func (a *L3FirewallRulesApiService) UpdateNetworkApplianceFirewallL3FirewallRules(ctx _context.Context, networkId string) L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest {
	return L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L3FirewallRulesApiService) UpdateNetworkApplianceFirewallL3FirewallRulesExecute(r L3FirewallRulesApiApiUpdateNetworkApplianceFirewallL3FirewallRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L3FirewallRulesApiService.UpdateNetworkApplianceFirewallL3FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/l3FirewallRules"
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
	localVarPostBody = r.updateNetworkApplianceFirewallL3FirewallRules
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

type L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct {
	ctx _context.Context
	ApiService *L3FirewallRulesApiService
	networkId string
	number string
	updateNetworkWirelessSsidFirewallL3FirewallRules *InlineObject137
}

func (r L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) UpdateNetworkWirelessSsidFirewallL3FirewallRules(updateNetworkWirelessSsidFirewallL3FirewallRules InlineObject137) L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	r.updateNetworkWirelessSsidFirewallL3FirewallRules = &updateNetworkWirelessSsidFirewallL3FirewallRules
	return r
}

func (r L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidFirewallL3FirewallRulesExecute(r)
}

/*
UpdateNetworkWirelessSsidFirewallL3FirewallRules Update the L3 firewall rules of an SSID on an MR network

Update the L3 firewall rules of an SSID on an MR network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param number
 @return L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest
*/
func (a *L3FirewallRulesApiService) UpdateNetworkWirelessSsidFirewallL3FirewallRules(ctx _context.Context, networkId string, number string) L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	return L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L3FirewallRulesApiService) UpdateNetworkWirelessSsidFirewallL3FirewallRulesExecute(r L3FirewallRulesApiApiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L3FirewallRulesApiService.UpdateNetworkWirelessSsidFirewallL3FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules"
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
	localVarPostBody = r.updateNetworkWirelessSsidFirewallL3FirewallRules
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
