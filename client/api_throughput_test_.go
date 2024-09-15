/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
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


// ThroughputTestApiService ThroughputTestApi service
type ThroughputTestApiService service

type ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest struct {
	ctx context.Context
	ApiService *ThroughputTestApiService
	serial string
	createDeviceLiveToolsThroughputTest *InlineObject20
}

func (r ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest) CreateDeviceLiveToolsThroughputTest(createDeviceLiveToolsThroughputTest InlineObject20) ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest {
	r.createDeviceLiveToolsThroughputTest = &createDeviceLiveToolsThroughputTest
	return r
}

func (r ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest) Execute() (*InlineResponse2014, *http.Response, error) {
	return r.ApiService.CreateDeviceLiveToolsThroughputTestExecute(r)
}

/*
CreateDeviceLiveToolsThroughputTest Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput. This endpoint has a rate limit of one request every five seconds per device.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest
*/
func (a *ThroughputTestApiService) CreateDeviceLiveToolsThroughputTest(ctx context.Context, serial string) ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest {
	return ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return InlineResponse2014
func (a *ThroughputTestApiService) CreateDeviceLiveToolsThroughputTestExecute(r ThroughputTestApiCreateDeviceLiveToolsThroughputTestRequest) (*InlineResponse2014, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2014
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThroughputTestApiService.CreateDeviceLiveToolsThroughputTest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/throughputTest"
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
	localVarPostBody = r.createDeviceLiveToolsThroughputTest
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

type ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest struct {
	ctx context.Context
	ApiService *ThroughputTestApiService
	serial string
	throughputTestId string
}

func (r ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest) Execute() (*InlineResponse20024, *http.Response, error) {
	return r.ApiService.GetDeviceLiveToolsThroughputTestExecute(r)
}

/*
GetDeviceLiveToolsThroughputTest Return a throughput test job

Return a throughput test job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param throughputTestId Throughput test ID
 @return ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest
*/
func (a *ThroughputTestApiService) GetDeviceLiveToolsThroughputTest(ctx context.Context, serial string, throughputTestId string) ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest {
	return ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		throughputTestId: throughputTestId,
	}
}

// Execute executes the request
//  @return InlineResponse20024
func (a *ThroughputTestApiService) GetDeviceLiveToolsThroughputTestExecute(r ThroughputTestApiGetDeviceLiveToolsThroughputTestRequest) (*InlineResponse20024, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20024
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThroughputTestApiService.GetDeviceLiveToolsThroughputTest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/throughputTest/{throughputTestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"throughputTestId"+"}", url.PathEscape(parameterToString(r.throughputTestId, "")), -1)

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
