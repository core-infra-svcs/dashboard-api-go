/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
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


// PiiApiService PiiApi service
type PiiApiService service

type PiiApiCreateNetworkPiiRequestRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	createNetworkPiiRequest *InlineObject118
}

func (r PiiApiCreateNetworkPiiRequestRequest) CreateNetworkPiiRequest(createNetworkPiiRequest InlineObject118) PiiApiCreateNetworkPiiRequestRequest {
	r.createNetworkPiiRequest = &createNetworkPiiRequest
	return r
}

func (r PiiApiCreateNetworkPiiRequestRequest) Execute() (*InlineResponse200114, *http.Response, error) {
	return r.ApiService.CreateNetworkPiiRequestExecute(r)
}

/*
CreateNetworkPiiRequest Submit a new delete or restrict processing PII request

Submit a new delete or restrict processing PII request

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/requests
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return PiiApiCreateNetworkPiiRequestRequest
*/
func (a *PiiApiService) CreateNetworkPiiRequest(ctx context.Context, networkId string) PiiApiCreateNetworkPiiRequestRequest {
	return PiiApiCreateNetworkPiiRequestRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200114
func (a *PiiApiService) CreateNetworkPiiRequestExecute(r PiiApiCreateNetworkPiiRequestRequest) (*InlineResponse200114, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200114
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.CreateNetworkPiiRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/requests"
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
	localVarPostBody = r.createNetworkPiiRequest
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

type PiiApiDeleteNetworkPiiRequestRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	requestId string
}

func (r PiiApiDeleteNetworkPiiRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkPiiRequestExecute(r)
}

/*
DeleteNetworkPiiRequest Delete a restrict processing PII request

Delete a restrict processing PII request

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/requests/{requestId}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param requestId Request ID
 @return PiiApiDeleteNetworkPiiRequestRequest
*/
func (a *PiiApiService) DeleteNetworkPiiRequest(ctx context.Context, networkId string, requestId string) PiiApiDeleteNetworkPiiRequestRequest {
	return PiiApiDeleteNetworkPiiRequestRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		requestId: requestId,
	}
}

// Execute executes the request
func (a *PiiApiService) DeleteNetworkPiiRequestExecute(r PiiApiDeleteNetworkPiiRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.DeleteNetworkPiiRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterToString(r.requestId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PiiApiGetNetworkPiiPiiKeysRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	username *string
	email *string
	mac *string
	serial *string
	imei *string
	bluetoothMac *string
}

// The username of a Systems Manager user
func (r PiiApiGetNetworkPiiPiiKeysRequest) Username(username string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.username = &username
	return r
}

// The email of a network user account or a Systems Manager device
func (r PiiApiGetNetworkPiiPiiKeysRequest) Email(email string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.email = &email
	return r
}

// The MAC of a network client device or a Systems Manager device
func (r PiiApiGetNetworkPiiPiiKeysRequest) Mac(mac string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.mac = &mac
	return r
}

// The serial of a Systems Manager device
func (r PiiApiGetNetworkPiiPiiKeysRequest) Serial(serial string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.serial = &serial
	return r
}

// The IMEI of a Systems Manager device
func (r PiiApiGetNetworkPiiPiiKeysRequest) Imei(imei string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.imei = &imei
	return r
}

// The MAC of a Bluetooth client
func (r PiiApiGetNetworkPiiPiiKeysRequest) BluetoothMac(bluetoothMac string) PiiApiGetNetworkPiiPiiKeysRequest {
	r.bluetoothMac = &bluetoothMac
	return r
}

func (r PiiApiGetNetworkPiiPiiKeysRequest) Execute() (*map[string]InlineResponse200113, *http.Response, error) {
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
 @param networkId Network ID
 @return PiiApiGetNetworkPiiPiiKeysRequest
*/
func (a *PiiApiService) GetNetworkPiiPiiKeys(ctx context.Context, networkId string) PiiApiGetNetworkPiiPiiKeysRequest {
	return PiiApiGetNetworkPiiPiiKeysRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]InlineResponse200113
func (a *PiiApiService) GetNetworkPiiPiiKeysExecute(r PiiApiGetNetworkPiiPiiKeysRequest) (*map[string]InlineResponse200113, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]InlineResponse200113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.GetNetworkPiiPiiKeys")
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

type PiiApiGetNetworkPiiRequestRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	requestId string
}

func (r PiiApiGetNetworkPiiRequestRequest) Execute() (*InlineResponse200114, *http.Response, error) {
	return r.ApiService.GetNetworkPiiRequestExecute(r)
}

/*
GetNetworkPiiRequest Return a PII request

Return a PII request

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/requests/{requestId}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param requestId Request ID
 @return PiiApiGetNetworkPiiRequestRequest
*/
func (a *PiiApiService) GetNetworkPiiRequest(ctx context.Context, networkId string, requestId string) PiiApiGetNetworkPiiRequestRequest {
	return PiiApiGetNetworkPiiRequestRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		requestId: requestId,
	}
}

// Execute executes the request
//  @return InlineResponse200114
func (a *PiiApiService) GetNetworkPiiRequestExecute(r PiiApiGetNetworkPiiRequestRequest) (*InlineResponse200114, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200114
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.GetNetworkPiiRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterToString(r.requestId, "")), -1)

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

type PiiApiGetNetworkPiiRequestsRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
}

func (r PiiApiGetNetworkPiiRequestsRequest) Execute() ([]InlineResponse200114, *http.Response, error) {
	return r.ApiService.GetNetworkPiiRequestsExecute(r)
}

/*
GetNetworkPiiRequests List the PII requests for this network or organization

List the PII requests for this network or organization

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/requests
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return PiiApiGetNetworkPiiRequestsRequest
*/
func (a *PiiApiService) GetNetworkPiiRequests(ctx context.Context, networkId string) PiiApiGetNetworkPiiRequestsRequest {
	return PiiApiGetNetworkPiiRequestsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200114
func (a *PiiApiService) GetNetworkPiiRequestsExecute(r PiiApiGetNetworkPiiRequestsRequest) ([]InlineResponse200114, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200114
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.GetNetworkPiiRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/requests"
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

type PiiApiGetNetworkPiiSmDevicesForKeyRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	username *string
	email *string
	mac *string
	serial *string
	imei *string
	bluetoothMac *string
}

// The username of a Systems Manager user
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Username(username string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.username = &username
	return r
}

// The email of a network user account or a Systems Manager device
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Email(email string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.email = &email
	return r
}

// The MAC of a network client device or a Systems Manager device
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Mac(mac string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.mac = &mac
	return r
}

// The serial of a Systems Manager device
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Serial(serial string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.serial = &serial
	return r
}

// The IMEI of a Systems Manager device
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Imei(imei string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.imei = &imei
	return r
}

// The MAC of a Bluetooth client
func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) BluetoothMac(bluetoothMac string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	r.bluetoothMac = &bluetoothMac
	return r
}

func (r PiiApiGetNetworkPiiSmDevicesForKeyRequest) Execute() (*map[string][]string, *http.Response, error) {
	return r.ApiService.GetNetworkPiiSmDevicesForKeyExecute(r)
}

/*
GetNetworkPiiSmDevicesForKey Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier

Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.

## ALTERNATE PATH

```
/organizations/{organizationId}/pii/smDevicesForKey
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return PiiApiGetNetworkPiiSmDevicesForKeyRequest
*/
func (a *PiiApiService) GetNetworkPiiSmDevicesForKey(ctx context.Context, networkId string) PiiApiGetNetworkPiiSmDevicesForKeyRequest {
	return PiiApiGetNetworkPiiSmDevicesForKeyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string][]string
func (a *PiiApiService) GetNetworkPiiSmDevicesForKeyExecute(r PiiApiGetNetworkPiiSmDevicesForKeyRequest) (*map[string][]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.GetNetworkPiiSmDevicesForKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/pii/smDevicesForKey"
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

type PiiApiGetNetworkPiiSmOwnersForKeyRequest struct {
	ctx context.Context
	ApiService *PiiApiService
	networkId string
	username *string
	email *string
	mac *string
	serial *string
	imei *string
	bluetoothMac *string
}

// The username of a Systems Manager user
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Username(username string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.username = &username
	return r
}

// The email of a network user account or a Systems Manager device
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Email(email string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.email = &email
	return r
}

// The MAC of a network client device or a Systems Manager device
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Mac(mac string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.mac = &mac
	return r
}

// The serial of a Systems Manager device
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Serial(serial string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.serial = &serial
	return r
}

// The IMEI of a Systems Manager device
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Imei(imei string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.imei = &imei
	return r
}

// The MAC of a Bluetooth client
func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) BluetoothMac(bluetoothMac string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	r.bluetoothMac = &bluetoothMac
	return r
}

func (r PiiApiGetNetworkPiiSmOwnersForKeyRequest) Execute() (*map[string][]string, *http.Response, error) {
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
 @return PiiApiGetNetworkPiiSmOwnersForKeyRequest
*/
func (a *PiiApiService) GetNetworkPiiSmOwnersForKey(ctx context.Context, networkId string) PiiApiGetNetworkPiiSmOwnersForKeyRequest {
	return PiiApiGetNetworkPiiSmOwnersForKeyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string][]string
func (a *PiiApiService) GetNetworkPiiSmOwnersForKeyExecute(r PiiApiGetNetworkPiiSmOwnersForKeyRequest) (*map[string][]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PiiApiService.GetNetworkPiiSmOwnersForKey")
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
