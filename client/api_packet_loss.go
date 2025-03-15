/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
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


// PacketLossApiService PacketLossApi service
type PacketLossApiService service

type PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest struct {
	ctx context.Context
	ApiService *PacketLossApiService
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
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) NetworkIds(networkIds []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by SSID number.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Ssids(ssids []int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.ssids = &ssids
	return r
}

// Filter results by band. Valid bands are: 2.4, 5, and 6.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Bands(bands []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.bands = &bands
	return r
}

// Filter results by client mac address(es).
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Macs(macs []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.macs = &macs
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) PerPage(perPage int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) StartingAfter(startingAfter string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) EndingBefore(endingBefore string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) T0(t0 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) T1(t1 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Timespan(timespan float32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	r.timespan = &timespan
	return r
}

func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) Execute() ([]InlineResponse200338, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesPacketLossByClientExecute(r)
}

/*
GetOrganizationWirelessDevicesPacketLossByClient Get average packet loss for the given timespan for all clients in the organization.

Get average packet loss for the given timespan for all clients in the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest
*/
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByClient(ctx context.Context, organizationId string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest {
	return PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200338
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByClientExecute(r PacketLossApiGetOrganizationWirelessDevicesPacketLossByClientRequest) ([]InlineResponse200338, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200338
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PacketLossApiService.GetOrganizationWirelessDevicesPacketLossByClient")
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

type PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest struct {
	ctx context.Context
	ApiService *PacketLossApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	ssids *[]int32
	bands *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
}

// Filter results by network.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) NetworkIds(networkIds []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) Serials(serials []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.serials = &serials
	return r
}

// Filter results by SSID number.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) Ssids(ssids []int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.ssids = &ssids
	return r
}

// Filter results by band. Valid bands are: 2.4, 5, and 6.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) Bands(bands []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.bands = &bands
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) PerPage(perPage int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) StartingAfter(startingAfter string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) EndingBefore(endingBefore string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) T0(t0 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) T1(t1 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) Timespan(timespan float32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	r.timespan = &timespan
	return r
}

func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) Execute() ([]InlineResponse200339, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesPacketLossByDeviceExecute(r)
}

/*
GetOrganizationWirelessDevicesPacketLossByDevice Get average packet loss for the given timespan for all devices in the organization

Get average packet loss for the given timespan for all devices in the organization. Does not include device's own traffic.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest
*/
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByDevice(ctx context.Context, organizationId string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest {
	return PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200339
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByDeviceExecute(r PacketLossApiGetOrganizationWirelessDevicesPacketLossByDeviceRequest) ([]InlineResponse200339, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200339
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PacketLossApiService.GetOrganizationWirelessDevicesPacketLossByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/packetLoss/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.ssids != nil {
		localVarQueryParams.Add("ssids", parameterToString(*r.ssids, "csv"))
	}
	if r.bands != nil {
		localVarQueryParams.Add("bands", parameterToString(*r.bands, "csv"))
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

type PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest struct {
	ctx context.Context
	ApiService *PacketLossApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	ssids *[]int32
	bands *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
}

// Filter results by network.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) NetworkIds(networkIds []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) Serials(serials []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.serials = &serials
	return r
}

// Filter results by SSID number.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) Ssids(ssids []int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.ssids = &ssids
	return r
}

// Filter results by band. Valid bands are: 2.4, 5, and 6.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) Bands(bands []string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.bands = &bands
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) PerPage(perPage int32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) StartingAfter(startingAfter string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) EndingBefore(endingBefore string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) T0(t0 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) T1(t1 string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) Timespan(timespan float32) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	r.timespan = &timespan
	return r
}

func (r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) Execute() ([]InlineResponse200340, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesPacketLossByNetworkExecute(r)
}

/*
GetOrganizationWirelessDevicesPacketLossByNetwork Get average packet loss for the given timespan for all networks in the organization.

Get average packet loss for the given timespan for all networks in the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest
*/
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByNetwork(ctx context.Context, organizationId string) PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest {
	return PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200340
func (a *PacketLossApiService) GetOrganizationWirelessDevicesPacketLossByNetworkExecute(r PacketLossApiGetOrganizationWirelessDevicesPacketLossByNetworkRequest) ([]InlineResponse200340, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200340
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PacketLossApiService.GetOrganizationWirelessDevicesPacketLossByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/packetLoss/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.ssids != nil {
		localVarQueryParams.Add("ssids", parameterToString(*r.ssids, "csv"))
	}
	if r.bands != nil {
		localVarQueryParams.Add("bands", parameterToString(*r.bands, "csv"))
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
