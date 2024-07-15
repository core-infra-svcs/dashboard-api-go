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


// WakeOnLanApiService WakeOnLanApi service
type WakeOnLanApiService service

type WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest struct {
	ctx context.Context
	ApiService *WakeOnLanApiService
	serial string
	createDeviceLiveToolsWakeOnLan *InlineObject21
}

func (r WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest) CreateDeviceLiveToolsWakeOnLan(createDeviceLiveToolsWakeOnLan InlineObject21) WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest {
	r.createDeviceLiveToolsWakeOnLan = &createDeviceLiveToolsWakeOnLan
	return r
}

func (r WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest) Execute() (*InlineResponse2015, *http.Response, error) {
	return r.ApiService.CreateDeviceLiveToolsWakeOnLanExecute(r)
}

/*
CreateDeviceLiveToolsWakeOnLan Enqueue a job to send a Wake-on-LAN packet from the device

Enqueue a job to send a Wake-on-LAN packet from the device. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest
*/
func (a *WakeOnLanApiService) CreateDeviceLiveToolsWakeOnLan(ctx context.Context, serial string) WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest {
	return WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return InlineResponse2015
func (a *WakeOnLanApiService) CreateDeviceLiveToolsWakeOnLanExecute(r WakeOnLanApiCreateDeviceLiveToolsWakeOnLanRequest) (*InlineResponse2015, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2015
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WakeOnLanApiService.CreateDeviceLiveToolsWakeOnLan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/wakeOnLan"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDeviceLiveToolsWakeOnLan == nil {
		return localVarReturnValue, nil, reportError("createDeviceLiveToolsWakeOnLan is required and must be specified")
	}

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
	localVarPostBody = r.createDeviceLiveToolsWakeOnLan
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

type WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest struct {
	ctx context.Context
	ApiService *WakeOnLanApiService
	serial string
	wakeOnLanId string
}

func (r WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest) Execute() (*InlineResponse20024, *http.Response, error) {
	return r.ApiService.GetDeviceLiveToolsWakeOnLanExecute(r)
}

/*
GetDeviceLiveToolsWakeOnLan Return a Wake-on-LAN job

Return a Wake-on-LAN job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param wakeOnLanId Wake on lan ID
 @return WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest
*/
func (a *WakeOnLanApiService) GetDeviceLiveToolsWakeOnLan(ctx context.Context, serial string, wakeOnLanId string) WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest {
	return WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		wakeOnLanId: wakeOnLanId,
	}
}

// Execute executes the request
//  @return InlineResponse20024
func (a *WakeOnLanApiService) GetDeviceLiveToolsWakeOnLanExecute(r WakeOnLanApiGetDeviceLiveToolsWakeOnLanRequest) (*InlineResponse20024, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20024
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WakeOnLanApiService.GetDeviceLiveToolsWakeOnLan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wakeOnLanId"+"}", url.PathEscape(parameterToString(r.wakeOnLanId, "")), -1)

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
