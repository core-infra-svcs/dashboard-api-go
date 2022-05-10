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

// CustomAnalyticsApiService CustomAnalyticsApi service
type CustomAnalyticsApiService service

type CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	organizationId string
	createOrganizationCameraCustomAnalyticsArtifact *InlineObject167
}

func (r CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact InlineObject167) CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest {
	r.createOrganizationCameraCustomAnalyticsArtifact = &createOrganizationCameraCustomAnalyticsArtifact
	return r
}

func (r CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
CreateOrganizationCameraCustomAnalyticsArtifact Create custom analytics artifact

Create custom analytics artifact. Returns an artifact upload URL with expiry time. Upload the artifact file with a put request to the returned upload URL before its expiry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *CustomAnalyticsApiService) CreateOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string) CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest {
	return CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CustomAnalyticsApiService) CreateOrganizationCameraCustomAnalyticsArtifactExecute(r CustomAnalyticsApiApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.CreateOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.createOrganizationCameraCustomAnalyticsArtifact
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

type CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	organizationId string
	artifactId string
}


func (r CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
DeleteOrganizationCameraCustomAnalyticsArtifact Delete Custom Analytics Artifact

Delete Custom Analytics Artifact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param artifactId
 @return CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *CustomAnalyticsApiService) DeleteOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string, artifactId string) CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest {
	return CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		artifactId: artifactId,
	}
}

// Execute executes the request
func (a *CustomAnalyticsApiService) DeleteOrganizationCameraCustomAnalyticsArtifactExecute(r CustomAnalyticsApiApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.DeleteOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

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

type CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	serial string
}


func (r CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraCustomAnalyticsExecute(r)
}

/*
GetDeviceCameraCustomAnalytics Return custom analytics settings for a camera

Return custom analytics settings for a camera

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest
*/
func (a *CustomAnalyticsApiService) GetDeviceCameraCustomAnalytics(ctx context.Context, serial string) CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest {
	return CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CustomAnalyticsApiService) GetDeviceCameraCustomAnalyticsExecute(r CustomAnalyticsApiApiGetDeviceCameraCustomAnalyticsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.GetDeviceCameraCustomAnalytics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/customAnalytics"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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

type CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	organizationId string
	artifactId string
}


func (r CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
GetOrganizationCameraCustomAnalyticsArtifact Get Custom Analytics Artifact

Get Custom Analytics Artifact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param artifactId
 @return CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *CustomAnalyticsApiService) GetOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string, artifactId string) CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest {
	return CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CustomAnalyticsApiService) GetOrganizationCameraCustomAnalyticsArtifactExecute(r CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.GetOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

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

type CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	organizationId string
}


func (r CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraCustomAnalyticsArtifactsExecute(r)
}

/*
GetOrganizationCameraCustomAnalyticsArtifacts List Custom Analytics Artifacts

List Custom Analytics Artifacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest
*/
func (a *CustomAnalyticsApiService) GetOrganizationCameraCustomAnalyticsArtifacts(ctx context.Context, organizationId string) CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest {
	return CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *CustomAnalyticsApiService) GetOrganizationCameraCustomAnalyticsArtifactsExecute(r CustomAnalyticsApiApiGetOrganizationCameraCustomAnalyticsArtifactsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.GetOrganizationCameraCustomAnalyticsArtifacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest struct {
	ctx context.Context
	ApiService *CustomAnalyticsApiService
	serial string
	updateDeviceCameraCustomAnalytics *InlineObject2
}

func (r CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest) UpdateDeviceCameraCustomAnalytics(updateDeviceCameraCustomAnalytics InlineObject2) CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest {
	r.updateDeviceCameraCustomAnalytics = &updateDeviceCameraCustomAnalytics
	return r
}

func (r CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateDeviceCameraCustomAnalyticsExecute(r)
}

/*
UpdateDeviceCameraCustomAnalytics Update custom analytics settings for a camera

Update custom analytics settings for a camera

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest
*/
func (a *CustomAnalyticsApiService) UpdateDeviceCameraCustomAnalytics(ctx context.Context, serial string) CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest {
	return CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CustomAnalyticsApiService) UpdateDeviceCameraCustomAnalyticsExecute(r CustomAnalyticsApiApiUpdateDeviceCameraCustomAnalyticsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomAnalyticsApiService.UpdateDeviceCameraCustomAnalytics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/customAnalytics"
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
	localVarPostBody = r.updateDeviceCameraCustomAnalytics
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
