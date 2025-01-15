/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
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


// L7FirewallRulesApiService L7FirewallRulesApi service
type L7FirewallRulesApiService service

type L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest struct {
	ctx context.Context
	ApiService *L7FirewallRulesApiService
	networkId string
}

func (r L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallL7FirewallRulesExecute(r)
}

/*
GetNetworkApplianceFirewallL7FirewallRules List the MX L7 firewall rules for an MX network

List the MX L7 firewall rules for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest
*/
func (a *L7FirewallRulesApiService) GetNetworkApplianceFirewallL7FirewallRules(ctx context.Context, networkId string) L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest {
	return L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L7FirewallRulesApiService) GetNetworkApplianceFirewallL7FirewallRulesExecute(r L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L7FirewallRulesApiService.GetNetworkApplianceFirewallL7FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/l7FirewallRules"
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

type L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest struct {
	ctx context.Context
	ApiService *L7FirewallRulesApiService
	networkId string
}

func (r L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest) Execute() (*InlineResponse20052, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesExecute(r)
}

/*
GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories Return the L7 firewall application categories and their associated applications for an MX network

Return the L7 firewall application categories and their associated applications for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest
*/
func (a *L7FirewallRulesApiService) GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(ctx context.Context, networkId string) L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest {
	return L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20052
func (a *L7FirewallRulesApiService) GetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesExecute(r L7FirewallRulesApiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest) (*InlineResponse20052, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20052
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L7FirewallRulesApiService.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories"
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

type L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest struct {
	ctx context.Context
	ApiService *L7FirewallRulesApiService
	networkId string
	number string
}

func (r L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest) Execute() (*InlineResponse200197, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidFirewallL7FirewallRulesExecute(r)
}

/*
GetNetworkWirelessSsidFirewallL7FirewallRules Return the L7 firewall rules for an SSID on an MR network

Return the L7 firewall rules for an SSID on an MR network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest
*/
func (a *L7FirewallRulesApiService) GetNetworkWirelessSsidFirewallL7FirewallRules(ctx context.Context, networkId string, number string) L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	return L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return InlineResponse200197
func (a *L7FirewallRulesApiService) GetNetworkWirelessSsidFirewallL7FirewallRulesExecute(r L7FirewallRulesApiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest) (*InlineResponse200197, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200197
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L7FirewallRulesApiService.GetNetworkWirelessSsidFirewallL7FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterToString(r.number, "")), -1)

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

type L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest struct {
	ctx context.Context
	ApiService *L7FirewallRulesApiService
	networkId string
	updateNetworkApplianceFirewallL7FirewallRules *InlineObject47
}

func (r L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest) UpdateNetworkApplianceFirewallL7FirewallRules(updateNetworkApplianceFirewallL7FirewallRules InlineObject47) L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	r.updateNetworkApplianceFirewallL7FirewallRules = &updateNetworkApplianceFirewallL7FirewallRules
	return r
}

func (r L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceFirewallL7FirewallRulesExecute(r)
}

/*
UpdateNetworkApplianceFirewallL7FirewallRules Update the MX L7 firewall rules for an MX network

Update the MX L7 firewall rules for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest
*/
func (a *L7FirewallRulesApiService) UpdateNetworkApplianceFirewallL7FirewallRules(ctx context.Context, networkId string) L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	return L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *L7FirewallRulesApiService) UpdateNetworkApplianceFirewallL7FirewallRulesExecute(r L7FirewallRulesApiUpdateNetworkApplianceFirewallL7FirewallRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L7FirewallRulesApiService.UpdateNetworkApplianceFirewallL7FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/l7FirewallRules"
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
	localVarPostBody = r.updateNetworkApplianceFirewallL7FirewallRules
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

type L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct {
	ctx context.Context
	ApiService *L7FirewallRulesApiService
	networkId string
	number string
	updateNetworkWirelessSsidFirewallL7FirewallRules *InlineObject194
}

func (r L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) UpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules InlineObject194) L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	r.updateNetworkWirelessSsidFirewallL7FirewallRules = &updateNetworkWirelessSsidFirewallL7FirewallRules
	return r
}

func (r L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) Execute() (*InlineResponse200197, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidFirewallL7FirewallRulesExecute(r)
}

/*
UpdateNetworkWirelessSsidFirewallL7FirewallRules Update the L7 firewall rules of an SSID on an MR network

Update the L7 firewall rules of an SSID on an MR network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest
*/
func (a *L7FirewallRulesApiService) UpdateNetworkWirelessSsidFirewallL7FirewallRules(ctx context.Context, networkId string, number string) L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	return L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return InlineResponse200197
func (a *L7FirewallRulesApiService) UpdateNetworkWirelessSsidFirewallL7FirewallRulesExecute(r L7FirewallRulesApiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) (*InlineResponse200197, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200197
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L7FirewallRulesApiService.UpdateNetworkWirelessSsidFirewallL7FirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterToString(r.number, "")), -1)

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
	localVarPostBody = r.updateNetworkWirelessSsidFirewallL7FirewallRules
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
