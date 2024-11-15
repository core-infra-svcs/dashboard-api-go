/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
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


// WirelessControllersApiService WirelessControllersApi service
type WirelessControllersApiService service

type WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest struct {
	ctx context.Context
	ApiService *WirelessControllersApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	controllerSerials *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// Optional parameter to filter access points by network ID. This filter uses multiple exact matches.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) NetworkIds(networkIds []string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) Serials(serials []string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) ControllerSerials(controllerSerials []string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.controllerSerials = &controllerSerials
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) PerPage(perPage int32) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) StartingAfter(startingAfter string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) EndingBefore(endingBefore string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) Execute() (*InlineResponse200321, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesWirelessControllersByDeviceExecute(r)
}

/*
GetOrganizationWirelessDevicesWirelessControllersByDevice List of Catalyst access points information

List of Catalyst access points information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest
*/
func (a *WirelessControllersApiService) GetOrganizationWirelessDevicesWirelessControllersByDevice(ctx context.Context, organizationId string) WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest {
	return WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200321
func (a *WirelessControllersApiService) GetOrganizationWirelessDevicesWirelessControllersByDeviceExecute(r WirelessControllersApiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest) (*InlineResponse200321, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200321
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WirelessControllersApiService.GetOrganizationWirelessDevicesWirelessControllersByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/wirelessControllers/byDevice"
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
	if r.controllerSerials != nil {
		localVarQueryParams.Add("controllerSerials", parameterToString(*r.controllerSerials, "csv"))
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
