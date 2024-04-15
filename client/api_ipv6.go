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


// Ipv6ApiService Ipv6Api service
type Ipv6ApiService service

type Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct {
	ctx context.Context
	ApiService *Ipv6ApiService
	serial string
	updateDeviceWirelessAlternateManagementInterfaceIpv6 *InlineObject33
}

func (r Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) UpdateDeviceWirelessAlternateManagementInterfaceIpv6(updateDeviceWirelessAlternateManagementInterfaceIpv6 InlineObject33) Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	r.updateDeviceWirelessAlternateManagementInterfaceIpv6 = &updateDeviceWirelessAlternateManagementInterfaceIpv6
	return r
}

func (r Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Execute() (*InlineResponse20038, *http.Response, error) {
	return r.ApiService.UpdateDeviceWirelessAlternateManagementInterfaceIpv6Execute(r)
}

/*
UpdateDeviceWirelessAlternateManagementInterfaceIpv6 Update alternate management interface IPv6 address

Update alternate management interface IPv6 address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request
*/
func (a *Ipv6ApiService) UpdateDeviceWirelessAlternateManagementInterfaceIpv6(ctx context.Context, serial string) Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	return Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return InlineResponse20038
func (a *Ipv6ApiService) UpdateDeviceWirelessAlternateManagementInterfaceIpv6Execute(r Ipv6ApiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) (*InlineResponse20038, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20038
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Ipv6ApiService.UpdateDeviceWirelessAlternateManagementInterfaceIpv6")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/wireless/alternateManagementInterface/ipv6"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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
	localVarPostBody = r.updateDeviceWirelessAlternateManagementInterfaceIpv6
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
