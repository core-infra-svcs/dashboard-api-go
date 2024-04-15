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


// SmOwnersForKeyApiService SmOwnersForKeyApi service
type SmOwnersForKeyApiService service

type SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest struct {
	ctx context.Context
	ApiService *SmOwnersForKeyApiService
	networkId string
	username *string
	email *string
	mac *string
	serial *string
	imei *string
	bluetoothMac *string
}

// The username of a Systems Manager user
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Username(username string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.username = &username
	return r
}

// The email of a network user account or a Systems Manager device
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Email(email string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.email = &email
	return r
}

// The MAC of a network client device or a Systems Manager device
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Mac(mac string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.mac = &mac
	return r
}

// The serial of a Systems Manager device
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Serial(serial string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.serial = &serial
	return r
}

// The IMEI of a Systems Manager device
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Imei(imei string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.imei = &imei
	return r
}

// The MAC of a Bluetooth client
func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) BluetoothMac(bluetoothMac string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	r.bluetoothMac = &bluetoothMac
	return r
}

func (r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) Execute() (*map[string][]string, *http.Response, error) {
	return r.ApiService.GetNetworkPiiSmOwnersForKeyExecute(r)
}

/*
GetNetworkPiiSmOwnersForKey Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier

Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/smOwnersForKey
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest
*/
func (a *SmOwnersForKeyApiService) GetNetworkPiiSmOwnersForKey(ctx context.Context, networkId string) SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest {
	return SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string][]string
func (a *SmOwnersForKeyApiService) GetNetworkPiiSmOwnersForKeyExecute(r SmOwnersForKeyApiGetNetworkPiiSmOwnersForKeyRequest) (*map[string][]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmOwnersForKeyApiService.GetNetworkPiiSmOwnersForKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/smOwnersForKey"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.username != nil {
		localVarQueryParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.email != nil {
		localVarQueryParams.Add("email", parameterToString(*r.email, ""))
	}
	if r.mac != nil {
		localVarQueryParams.Add("mac", parameterToString(*r.mac, ""))
	}
	if r.serial != nil {
		localVarQueryParams.Add("serial", parameterToString(*r.serial, ""))
	}
	if r.imei != nil {
		localVarQueryParams.Add("imei", parameterToString(*r.imei, ""))
	}
	if r.bluetoothMac != nil {
		localVarQueryParams.Add("bluetoothMac", parameterToString(*r.bluetoothMac, ""))
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
