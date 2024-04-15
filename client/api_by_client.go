/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
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


// ByClientApiService ByClientApi service
type ByClientApiService service

type ByClientApiGetNetworkPoliciesByClientRequest struct {
	ctx context.Context
	ApiService *ByClientApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	timespan *float32
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r ByClientApiGetNetworkPoliciesByClientRequest) PerPage(perPage int32) ByClientApiGetNetworkPoliciesByClientRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByClientApiGetNetworkPoliciesByClientRequest) StartingAfter(startingAfter string) ByClientApiGetNetworkPoliciesByClientRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByClientApiGetNetworkPoliciesByClientRequest) EndingBefore(endingBefore string) ByClientApiGetNetworkPoliciesByClientRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ByClientApiGetNetworkPoliciesByClientRequest) T0(t0 string) ByClientApiGetNetworkPoliciesByClientRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ByClientApiGetNetworkPoliciesByClientRequest) Timespan(timespan float32) ByClientApiGetNetworkPoliciesByClientRequest {
	r.timespan = &timespan
	return r
}

func (r ByClientApiGetNetworkPoliciesByClientRequest) Execute() ([]InlineResponse200102, *http.Response, error) {
	return r.ApiService.GetNetworkPoliciesByClientExecute(r)
}

/*
GetNetworkPoliciesByClient Get policies for all clients with policies

Get policies for all clients with policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ByClientApiGetNetworkPoliciesByClientRequest
*/
func (a *ByClientApiService) GetNetworkPoliciesByClient(ctx context.Context, networkId string) ByClientApiGetNetworkPoliciesByClientRequest {
	return ByClientApiGetNetworkPoliciesByClientRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200102
func (a *ByClientApiService) GetNetworkPoliciesByClientExecute(r ByClientApiGetNetworkPoliciesByClientRequest) ([]InlineResponse200102, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200102
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByClientApiService.GetNetworkPoliciesByClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/policies/byClient"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
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

type ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest struct {
	ctx context.Context
	ApiService *ByClientApiService
	organizationId string
	networkIds *[]string
	ssids *[]int32
	bands *[]string
	macs *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
}

// Filter results by network.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) NetworkIds(networkIds []string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by SSID number.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Ssids(ssids []int32) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.ssids = &ssids
	return r
}

// Filter results by band. Valid bands are: 2.4, 5, and 6.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Bands(bands []string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.bands = &bands
	return r
}

// Filter results by client mac address(es).
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Macs(macs []string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.macs = &macs
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) PerPage(perPage int32) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) StartingAfter(startingAfter string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) EndingBefore(endingBefore string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) T0(t0 string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) T1(t1 string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Timespan(timespan float32) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.timespan = &timespan
	return r
}

func (r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Execute() ([]InlineResponse200286, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesPacketLossByClientExecute(r)
}

/*
GetOrganizationWirelessDevicesPacketLossByClient Get average packet loss for the given timespan for all clients in the organization.

Get average packet loss for the given timespan for all clients in the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest
*/
func (a *ByClientApiService) GetOrganizationWirelessDevicesPacketLossByClient(ctx context.Context, organizationId string) ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	return ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200286
func (a *ByClientApiService) GetOrganizationWirelessDevicesPacketLossByClientExecute(r ByClientApiGetOrganizationWirelessDevicesPacketLossByClientRequest) ([]InlineResponse200286, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200286
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByClientApiService.GetOrganizationWirelessDevicesPacketLossByClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/packetLoss/byClient"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.ssids != nil {
		localVarQueryParams.Add("ssids", parameterToString(*r.ssids, "csv"))
	}
	if r.bands != nil {
		localVarQueryParams.Add("bands", parameterToString(*r.bands, "csv"))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
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
