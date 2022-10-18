/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OptInsApiService OptInsApi service
type OptInsApiService service

type OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx _context.Context
	ApiService *OptInsApiService
	organizationId string
	createOrganizationEarlyAccessFeaturesOptIn *InlineObject194
}

func (r OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest) CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn InlineObject194) OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest {
	r.createOrganizationEarlyAccessFeaturesOptIn = &createOrganizationEarlyAccessFeaturesOptIn
	return r
}

func (r OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
CreateOrganizationEarlyAccessFeaturesOptIn Create a new early access feature opt-in for an organization

Create a new early access feature opt-in for an organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) CreateOrganizationEarlyAccessFeaturesOptIn(ctx _context.Context, organizationId string) OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *OptInsApiService) CreateOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiApiCreateOrganizationEarlyAccessFeaturesOptInRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.CreateOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx _context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
}


func (r OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
DeleteOrganizationEarlyAccessFeaturesOptIn Delete an early access feature opt-in

Delete an early access feature opt-in

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param optInId
 @return OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) DeleteOrganizationEarlyAccessFeaturesOptIn(ctx _context.Context, organizationId string, optInId string) OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
func (a *OptInsApiService) DeleteOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiApiDeleteOrganizationEarlyAccessFeaturesOptInRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.DeleteOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", _neturl.PathEscape(parameterToString(r.optInId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx _context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
}


func (r OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
GetOrganizationEarlyAccessFeaturesOptIn Show an early access feature opt-in for an organization

Show an early access feature opt-in for an organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param optInId
 @return OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptIn(ctx _context.Context, organizationId string, optInId string) OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.GetOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", _neturl.PathEscape(parameterToString(r.optInId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest struct {
	ctx _context.Context
	ApiService *OptInsApiService
	organizationId string
}


func (r OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationEarlyAccessFeaturesOptInsExecute(r)
}

/*
GetOrganizationEarlyAccessFeaturesOptIns List the early access feature opt-ins for an organization

List the early access feature opt-ins for an organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest
*/
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptIns(ctx _context.Context, organizationId string) OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest {
	return OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *OptInsApiService) GetOrganizationEarlyAccessFeaturesOptInsExecute(r OptInsApiApiGetOrganizationEarlyAccessFeaturesOptInsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.GetOrganizationEarlyAccessFeaturesOptIns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest struct {
	ctx _context.Context
	ApiService *OptInsApiService
	organizationId string
	optInId string
	updateOrganizationEarlyAccessFeaturesOptIn *InlineObject195
}

func (r OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn InlineObject195) OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest {
	r.updateOrganizationEarlyAccessFeaturesOptIn = &updateOrganizationEarlyAccessFeaturesOptIn
	return r
}

func (r OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateOrganizationEarlyAccessFeaturesOptInExecute(r)
}

/*
UpdateOrganizationEarlyAccessFeaturesOptIn Update an early access feature opt-in for an organization

Update an early access feature opt-in for an organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param optInId
 @return OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest
*/
func (a *OptInsApiService) UpdateOrganizationEarlyAccessFeaturesOptIn(ctx _context.Context, organizationId string, optInId string) OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest {
	return OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		optInId: optInId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *OptInsApiService) UpdateOrganizationEarlyAccessFeaturesOptInExecute(r OptInsApiApiUpdateOrganizationEarlyAccessFeaturesOptInRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptInsApiService.UpdateOrganizationEarlyAccessFeaturesOptIn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optInId"+"}", _neturl.PathEscape(parameterToString(r.optInId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
