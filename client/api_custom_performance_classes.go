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


// CustomPerformanceClassesApiService CustomPerformanceClassesApi service
type CustomPerformanceClassesApiService service

type CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	ctx context.Context
	ApiService *CustomPerformanceClassesApiService
	networkId string
	createNetworkApplianceTrafficShapingCustomPerformanceClass *InlineObject68
}

func (r CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) CreateNetworkApplianceTrafficShapingCustomPerformanceClass(createNetworkApplianceTrafficShapingCustomPerformanceClass InlineObject68) CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	r.createNetworkApplianceTrafficShapingCustomPerformanceClass = &createNetworkApplianceTrafficShapingCustomPerformanceClass
	return r
}

func (r CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Execute() (*InlineResponse20068, *http.Response, error) {
	return r.ApiService.CreateNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r)
}

/*
CreateNetworkApplianceTrafficShapingCustomPerformanceClass Add a custom performance class for an MX network

Add a custom performance class for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
*/
func (a *CustomPerformanceClassesApiService) CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx context.Context, networkId string) CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20068
func (a *CustomPerformanceClassesApiService) CreateNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r CustomPerformanceClassesApiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) (*InlineResponse20068, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20068
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPerformanceClassesApiService.CreateNetworkApplianceTrafficShapingCustomPerformanceClass")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkApplianceTrafficShapingCustomPerformanceClass == nil {
		return localVarReturnValue, nil, reportError("createNetworkApplianceTrafficShapingCustomPerformanceClass is required and must be specified")
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
	localVarPostBody = r.createNetworkApplianceTrafficShapingCustomPerformanceClass
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

type CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	ctx context.Context
	ApiService *CustomPerformanceClassesApiService
	networkId string
	customPerformanceClassId string
}

func (r CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r)
}

/*
DeleteNetworkApplianceTrafficShapingCustomPerformanceClass Delete a custom performance class from an MX network

Delete a custom performance class from an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param customPerformanceClassId Custom performance class ID
 @return CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest
*/
func (a *CustomPerformanceClassesApiService) DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx context.Context, networkId string, customPerformanceClassId string) CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		customPerformanceClassId: customPerformanceClassId,
	}
}

// Execute executes the request
func (a *CustomPerformanceClassesApiService) DeleteNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r CustomPerformanceClassesApiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPerformanceClassesApiService.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customPerformanceClassId"+"}", url.PathEscape(parameterToString(r.customPerformanceClassId, "")), -1)

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

type CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	ctx context.Context
	ApiService *CustomPerformanceClassesApiService
	networkId string
	customPerformanceClassId string
}

func (r CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Execute() (*InlineResponse20068, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r)
}

/*
GetNetworkApplianceTrafficShapingCustomPerformanceClass Return a custom performance class for an MX network

Return a custom performance class for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param customPerformanceClassId Custom performance class ID
 @return CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest
*/
func (a *CustomPerformanceClassesApiService) GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx context.Context, networkId string, customPerformanceClassId string) CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		customPerformanceClassId: customPerformanceClassId,
	}
}

// Execute executes the request
//  @return InlineResponse20068
func (a *CustomPerformanceClassesApiService) GetNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest) (*InlineResponse20068, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20068
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPerformanceClassesApiService.GetNetworkApplianceTrafficShapingCustomPerformanceClass")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customPerformanceClassId"+"}", url.PathEscape(parameterToString(r.customPerformanceClassId, "")), -1)

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

type CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct {
	ctx context.Context
	ApiService *CustomPerformanceClassesApiService
	networkId string
}

func (r CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest) Execute() ([]InlineResponse20068, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceTrafficShapingCustomPerformanceClassesExecute(r)
}

/*
GetNetworkApplianceTrafficShapingCustomPerformanceClasses List all custom performance classes for an MX network

List all custom performance classes for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest
*/
func (a *CustomPerformanceClassesApiService) GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx context.Context, networkId string) CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest {
	return CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20068
func (a *CustomPerformanceClassesApiService) GetNetworkApplianceTrafficShapingCustomPerformanceClassesExecute(r CustomPerformanceClassesApiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest) ([]InlineResponse20068, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20068
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPerformanceClassesApiService.GetNetworkApplianceTrafficShapingCustomPerformanceClasses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
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

type CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	ctx context.Context
	ApiService *CustomPerformanceClassesApiService
	networkId string
	customPerformanceClassId string
	updateNetworkApplianceTrafficShapingCustomPerformanceClass *InlineObject69
}

func (r CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass InlineObject69) CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	r.updateNetworkApplianceTrafficShapingCustomPerformanceClass = &updateNetworkApplianceTrafficShapingCustomPerformanceClass
	return r
}

func (r CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Execute() (*InlineResponse20068, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r)
}

/*
UpdateNetworkApplianceTrafficShapingCustomPerformanceClass Update a custom performance class for an MX network

Update a custom performance class for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param customPerformanceClassId Custom performance class ID
 @return CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
*/
func (a *CustomPerformanceClassesApiService) UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx context.Context, networkId string, customPerformanceClassId string) CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		customPerformanceClassId: customPerformanceClassId,
	}
}

// Execute executes the request
//  @return InlineResponse20068
func (a *CustomPerformanceClassesApiService) UpdateNetworkApplianceTrafficShapingCustomPerformanceClassExecute(r CustomPerformanceClassesApiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) (*InlineResponse20068, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20068
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPerformanceClassesApiService.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customPerformanceClassId"+"}", url.PathEscape(parameterToString(r.customPerformanceClassId, "")), -1)

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
	localVarPostBody = r.updateNetworkApplianceTrafficShapingCustomPerformanceClass
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
