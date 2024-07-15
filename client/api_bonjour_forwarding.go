/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
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


// BonjourForwardingApiService BonjourForwardingApi service
type BonjourForwardingApiService service

type BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest struct {
	ctx context.Context
	ApiService *BonjourForwardingApiService
	networkId string
	number string
}

func (r BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest) Execute() (*InlineResponse200187, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidBonjourForwardingExecute(r)
}

/*
GetNetworkWirelessSsidBonjourForwarding List the Bonjour forwarding setting and rules for the SSID

List the Bonjour forwarding setting and rules for the SSID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest
*/
func (a *BonjourForwardingApiService) GetNetworkWirelessSsidBonjourForwarding(ctx context.Context, networkId string, number string) BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest {
	return BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return InlineResponse200187
func (a *BonjourForwardingApiService) GetNetworkWirelessSsidBonjourForwardingExecute(r BonjourForwardingApiGetNetworkWirelessSsidBonjourForwardingRequest) (*InlineResponse200187, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200187
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BonjourForwardingApiService.GetNetworkWirelessSsidBonjourForwarding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/bonjourForwarding"
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

type BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest struct {
	ctx context.Context
	ApiService *BonjourForwardingApiService
	networkId string
	number string
	updateNetworkWirelessSsidBonjourForwarding *InlineObject185
}

func (r BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest) UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding InlineObject185) BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest {
	r.updateNetworkWirelessSsidBonjourForwarding = &updateNetworkWirelessSsidBonjourForwarding
	return r
}

func (r BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest) Execute() (*InlineResponse200187, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidBonjourForwardingExecute(r)
}

/*
UpdateNetworkWirelessSsidBonjourForwarding Update the bonjour forwarding setting and rules for the SSID

Update the bonjour forwarding setting and rules for the SSID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest
*/
func (a *BonjourForwardingApiService) UpdateNetworkWirelessSsidBonjourForwarding(ctx context.Context, networkId string, number string) BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest {
	return BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return InlineResponse200187
func (a *BonjourForwardingApiService) UpdateNetworkWirelessSsidBonjourForwardingExecute(r BonjourForwardingApiUpdateNetworkWirelessSsidBonjourForwardingRequest) (*InlineResponse200187, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200187
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BonjourForwardingApiService.UpdateNetworkWirelessSsidBonjourForwarding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/bonjourForwarding"
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
	localVarPostBody = r.updateNetworkWirelessSsidBonjourForwarding
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
