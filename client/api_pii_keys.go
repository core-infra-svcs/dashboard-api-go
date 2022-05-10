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

// PiiKeysApiService PiiKeysApi service
type PiiKeysApiService service

type PiiKeysApiApiGetNetworkPiiPiiKeysRequest struct {
	ctx context.Context
	ApiService *PiiKeysApiService
	networkId string
	username *string
	email *string
	mac *string
	serial *string
	imei *string
	bluetoothMac *string
}

// The username of a Systems Manager user
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Username(username string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.username = &username
	return r
}
// The email of a network user account or a Systems Manager device
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Email(email string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.email = &email
	return r
}
// The MAC of a network client device or a Systems Manager device
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Mac(mac string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.mac = &mac
	return r
}
// The serial of a Systems Manager device
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Serial(serial string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.serial = &serial
	return r
}
// The IMEI of a Systems Manager device
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Imei(imei string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.imei = &imei
	return r
}
// The MAC of a Bluetooth client
func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) BluetoothMac(bluetoothMac string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	r.bluetoothMac = &bluetoothMac
	return r
}

func (r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkPiiPiiKeysExecute(r)
}

/*
GetNetworkPiiPiiKeys List the keys required to access Personally Identifiable Information (PII) for a given identifier

List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/piiKeys
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return PiiKeysApiApiGetNetworkPiiPiiKeysRequest
*/
func (a *PiiKeysApiService) GetNetworkPiiPiiKeys(ctx context.Context, networkId string) PiiKeysApiApiGetNetworkPiiPiiKeysRequest {
	return PiiKeysApiApiGetNetworkPiiPiiKeysRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PiiKeysApiService) GetNetworkPiiPiiKeysExecute(r PiiKeysApiApiGetNetworkPiiPiiKeysRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiKeysApiService.GetNetworkPiiPiiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/piiKeys"
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
