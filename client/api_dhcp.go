/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
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

// Linger please
var (
	_ context.Context
)

// DhcpApiService DhcpApi service
type DhcpApiService service

type DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
}


func (r DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceApplianceDhcpSubnetsExecute(r)
}

/*
GetDeviceApplianceDhcpSubnets Return the DHCP subnet information for an appliance

Return the DHCP subnet information for an appliance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest
*/
func (a *DhcpApiService) GetDeviceApplianceDhcpSubnets(ctx context.Context, serial string) DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest {
	return DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DhcpApiService) GetDeviceApplianceDhcpSubnetsExecute(r DhcpApiApiGetDeviceApplianceDhcpSubnetsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetDeviceApplianceDhcpSubnets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/dhcp/subnets"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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

type DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
	interfaceId string
}


func (r DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceSwitchRoutingInterfaceDhcpExecute(r)
}

/*
GetDeviceSwitchRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch

Return a layer 3 interface DHCP configuration for a switch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @param interfaceId
 @return DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) GetDeviceSwitchRoutingInterfaceDhcp(ctx context.Context, serial string, interfaceId string) DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest {
	return DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) GetDeviceSwitchRoutingInterfaceDhcpExecute(r DhcpApiApiGetDeviceSwitchRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetDeviceSwitchRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterToString(r.interfaceId, "")), -1)

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

type DhcpApiApiGetNetworkCellularGatewayDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
}


func (r DhcpApiApiGetNetworkCellularGatewayDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkCellularGatewayDhcpExecute(r)
}

/*
GetNetworkCellularGatewayDhcp List common DHCP settings of MGs

List common DHCP settings of MGs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpApiApiGetNetworkCellularGatewayDhcpRequest
*/
func (a *DhcpApiService) GetNetworkCellularGatewayDhcp(ctx context.Context, networkId string) DhcpApiApiGetNetworkCellularGatewayDhcpRequest {
	return DhcpApiApiGetNetworkCellularGatewayDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) GetNetworkCellularGatewayDhcpExecute(r DhcpApiApiGetNetworkCellularGatewayDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetNetworkCellularGatewayDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/dhcp"
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

type DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	switchStackId string
	interfaceId string
}


func (r DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchStackRoutingInterfaceDhcpExecute(r)
}

/*
GetNetworkSwitchStackRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch stack

Return a layer 3 interface DHCP configuration for a switch stack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param switchStackId
 @param interfaceId
 @return DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) GetNetworkSwitchStackRoutingInterfaceDhcp(ctx context.Context, networkId string, switchStackId string, interfaceId string) DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		switchStackId: switchStackId,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) GetNetworkSwitchStackRoutingInterfaceDhcpExecute(r DhcpApiApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetNetworkSwitchStackRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"switchStackId"+"}", url.PathEscape(parameterToString(r.switchStackId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterToString(r.interfaceId, "")), -1)

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

type DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
	interfaceId string
	updateDeviceSwitchRoutingInterfaceDhcp *InlineObject17
}

func (r DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) UpdateDeviceSwitchRoutingInterfaceDhcp(updateDeviceSwitchRoutingInterfaceDhcp InlineObject17) DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest {
	r.updateDeviceSwitchRoutingInterfaceDhcp = &updateDeviceSwitchRoutingInterfaceDhcp
	return r
}

func (r DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateDeviceSwitchRoutingInterfaceDhcpExecute(r)
}

/*
UpdateDeviceSwitchRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch

Update a layer 3 interface DHCP configuration for a switch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @param interfaceId
 @return DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) UpdateDeviceSwitchRoutingInterfaceDhcp(ctx context.Context, serial string, interfaceId string) DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest {
	return DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) UpdateDeviceSwitchRoutingInterfaceDhcpExecute(r DhcpApiApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateDeviceSwitchRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterToString(r.interfaceId, "")), -1)

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
	localVarPostBody = r.updateDeviceSwitchRoutingInterfaceDhcp
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

type DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	updateNetworkCellularGatewayDhcp *InlineObject59
}

func (r DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest) UpdateNetworkCellularGatewayDhcp(updateNetworkCellularGatewayDhcp InlineObject59) DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest {
	r.updateNetworkCellularGatewayDhcp = &updateNetworkCellularGatewayDhcp
	return r
}

func (r DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkCellularGatewayDhcpExecute(r)
}

/*
UpdateNetworkCellularGatewayDhcp Update common DHCP settings of MGs

Update common DHCP settings of MGs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest
*/
func (a *DhcpApiService) UpdateNetworkCellularGatewayDhcp(ctx context.Context, networkId string) DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest {
	return DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) UpdateNetworkCellularGatewayDhcpExecute(r DhcpApiApiUpdateNetworkCellularGatewayDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateNetworkCellularGatewayDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/dhcp"
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
	localVarPostBody = r.updateNetworkCellularGatewayDhcp
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

type DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	switchStackId string
	interfaceId string
	updateNetworkSwitchStackRoutingInterfaceDhcp *InlineObject115
}

func (r DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) UpdateNetworkSwitchStackRoutingInterfaceDhcp(updateNetworkSwitchStackRoutingInterfaceDhcp InlineObject115) DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	r.updateNetworkSwitchStackRoutingInterfaceDhcp = &updateNetworkSwitchStackRoutingInterfaceDhcp
	return r
}

func (r DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchStackRoutingInterfaceDhcpExecute(r)
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch stack

Update a layer 3 interface DHCP configuration for a switch stack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param switchStackId
 @param interfaceId
 @return DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx context.Context, networkId string, switchStackId string, interfaceId string) DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		switchStackId: switchStackId,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) UpdateNetworkSwitchStackRoutingInterfaceDhcpExecute(r DhcpApiApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateNetworkSwitchStackRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"switchStackId"+"}", url.PathEscape(parameterToString(r.switchStackId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterToString(r.interfaceId, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchStackRoutingInterfaceDhcp
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
