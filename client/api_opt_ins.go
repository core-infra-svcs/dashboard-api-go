/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
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


// OptInsApiService OptInsApi service
type OptInsApiService service

type OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx context.Context
	ApiService *OptInsApiService
	organizationId string
	createOrganizationEarlyAccessFeaturesOptIn *InlineObject227
}

func (r OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest) CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn InlineObject227) OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest {
	r.createOrganizationEarlyAccessFeaturesOptIn = &createOrganizationEarlyAccessFeaturesOptIn
	return r
}

func (r OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest) Execute() (*InlineResponse200243, *http.Response, error) {
	return r.ApiService.CreateOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
CreateOrganizationEarlyAccessFeaturesOptIn Create a new early access feature opt-in for an organization

Create a new early access feature opt-in for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) CreateOrganizationEarlyAccessFeaturesOptIn(ctx context.Context, organizationId string) OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200243
func (a *OptInsApiService) CreateOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiCreateOrganizationEarlyAccessFeaturesOptInRequest) (*InlineResponse200243, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200243
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.CreateOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationEarlyAccessFeaturesOptIn == nil {
		return localVarReturnValue, nil, reportError("createOrganizationEarlyAccessFeaturesOptIn is required and must be specified")
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
	localVarPostBody = r.createOrganizationEarlyAccessFeaturesOptIn
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

type OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
}

func (r OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
DeleteOrganizationEarlyAccessFeaturesOptIn Delete an early access feature opt-in

Delete an early access feature opt-in

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param optInId Opt in ID
 @return OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) DeleteOrganizationEarlyAccessFeaturesOptIn(ctx context.Context, organizationId string, optInId string) OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
func (a *OptInsApiService) DeleteOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiDeleteOrganizationEarlyAccessFeaturesOptInRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.DeleteOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", url.PathEscape(parameterToString(r.optInId, "")), -1)

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

type OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
}

func (r OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest) Execute() (*InlineResponse200243, *http.Response, error) {
	return r.ApiService.GetOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
GetOrganizationEarlyAccessFeaturesOptIn Show an early access feature opt-in for an organization

Show an early access feature opt-in for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param optInId Opt in ID
 @return OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptIn(ctx context.Context, organizationId string, optInId string) OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
//  @return InlineResponse200243
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiGetOrganizationEarlyAccessFeaturesOptInRequest) (*InlineResponse200243, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200243
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.GetOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", url.PathEscape(parameterToString(r.optInId, "")), -1)

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

type OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest struct {
	ctx context.Context
	ApiService *OptInsApiService
	organizationId string
}

func (r OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest) Execute() (*InlineResponse200243, *http.Response, error) {
	return r.ApiService.GetOrganizationEarlyAccessFeaturesOptInsExecute(r)
}

/*
GetOrganizationEarlyAccessFeaturesOptIns List the early access feature opt-ins for an organization

List the early access feature opt-ins for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest
*/
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptIns(ctx context.Context, organizationId string) OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest {
	return OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200243
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptInsExecute(r OptInsApiGetOrganizationEarlyAccessFeaturesOptInsRequest) (*InlineResponse200243, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200243
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.GetOrganizationEarlyAccessFeaturesOptIns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns"
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

type OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
	updateOrganizationEarlyAccessFeaturesOptIn *InlineObject228
}

func (r OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn InlineObject228) OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest {
	r.updateOrganizationEarlyAccessFeaturesOptIn = &updateOrganizationEarlyAccessFeaturesOptIn
	return r
}

func (r OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) Execute() (*InlineResponse200243, *http.Response, error) {
	return r.ApiService.UpdateOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
UpdateOrganizationEarlyAccessFeaturesOptIn Update an early access feature opt-in for an organization

Update an early access feature opt-in for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param optInId Opt in ID
 @return OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) UpdateOrganizationEarlyAccessFeaturesOptIn(ctx context.Context, organizationId string, optInId string) OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
//  @return InlineResponse200243
func (a *OptInsApiService) UpdateOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) (*InlineResponse200243, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200243
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.UpdateOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", url.PathEscape(parameterToString(r.optInId, "")), -1)

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
	localVarPostBody = r.updateOrganizationEarlyAccessFeaturesOptIn
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
