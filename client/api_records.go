/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
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


// RecordsApiService RecordsApi service
type RecordsApiService service

type RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest struct {
	ctx context.Context
	ApiService *RecordsApiService
	organizationId string
	createOrganizationApplianceDnsLocalRecord *InlineObject222
}

func (r RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest) CreateOrganizationApplianceDnsLocalRecord(createOrganizationApplianceDnsLocalRecord InlineObject222) RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest {
	r.createOrganizationApplianceDnsLocalRecord = &createOrganizationApplianceDnsLocalRecord
	return r
}

func (r RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest) Execute() ([]InlineResponse200222, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsLocalRecordExecute(r)
}

/*
CreateOrganizationApplianceDnsLocalRecord Create a new local DNS record

Create a new local DNS record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest
*/
func (a *RecordsApiService) CreateOrganizationApplianceDnsLocalRecord(ctx context.Context, organizationId string) RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest {
	return RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200222
func (a *RecordsApiService) CreateOrganizationApplianceDnsLocalRecordExecute(r RecordsApiCreateOrganizationApplianceDnsLocalRecordRequest) ([]InlineResponse200222, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200222
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.CreateOrganizationApplianceDnsLocalRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/local/records"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationApplianceDnsLocalRecord == nil {
		return localVarReturnValue, nil, reportError("createOrganizationApplianceDnsLocalRecord is required and must be specified")
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
	localVarPostBody = r.createOrganizationApplianceDnsLocalRecord
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

type RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest struct {
	ctx context.Context
	ApiService *RecordsApiService
	organizationId string
	recordId string
}

func (r RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationApplianceDnsLocalRecordExecute(r)
}

/*
DeleteOrganizationApplianceDnsLocalRecord Deletes a local DNS record

Deletes a local DNS record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param recordId Record ID
 @return RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest
*/
func (a *RecordsApiService) DeleteOrganizationApplianceDnsLocalRecord(ctx context.Context, organizationId string, recordId string) RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest {
	return RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		recordId: recordId,
	}
}

// Execute executes the request
func (a *RecordsApiService) DeleteOrganizationApplianceDnsLocalRecordExecute(r RecordsApiDeleteOrganizationApplianceDnsLocalRecordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.DeleteOrganizationApplianceDnsLocalRecord")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/local/records/{recordId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterToString(r.recordId, "")), -1)

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

type RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest struct {
	ctx context.Context
	ApiService *RecordsApiService
	organizationId string
	profileIds *[]string
}

// Optional parameter to filter the results by profile IDs
func (r RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest) ProfileIds(profileIds []string) RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest {
	r.profileIds = &profileIds
	return r
}

func (r RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest) Execute() ([]InlineResponse200222, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceDnsLocalRecordsExecute(r)
}

/*
GetOrganizationApplianceDnsLocalRecords Fetch the DNS records used in local DNS profiles

Fetch the DNS records used in local DNS profiles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest
*/
func (a *RecordsApiService) GetOrganizationApplianceDnsLocalRecords(ctx context.Context, organizationId string) RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest {
	return RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200222
func (a *RecordsApiService) GetOrganizationApplianceDnsLocalRecordsExecute(r RecordsApiGetOrganizationApplianceDnsLocalRecordsRequest) ([]InlineResponse200222, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200222
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.GetOrganizationApplianceDnsLocalRecords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/local/records"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.profileIds != nil {
		localVarQueryParams.Add("profileIds", parameterToString(*r.profileIds, "csv"))
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

type RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest struct {
	ctx context.Context
	ApiService *RecordsApiService
	organizationId string
	recordId string
	updateOrganizationApplianceDnsLocalRecord *InlineObject223
}

func (r RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest) UpdateOrganizationApplianceDnsLocalRecord(updateOrganizationApplianceDnsLocalRecord InlineObject223) RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest {
	r.updateOrganizationApplianceDnsLocalRecord = &updateOrganizationApplianceDnsLocalRecord
	return r
}

func (r RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest) Execute() (*InlineResponse200222, *http.Response, error) {
	return r.ApiService.UpdateOrganizationApplianceDnsLocalRecordExecute(r)
}

/*
UpdateOrganizationApplianceDnsLocalRecord Updates a local DNS record

Updates a local DNS record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param recordId Record ID
 @return RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest
*/
func (a *RecordsApiService) UpdateOrganizationApplianceDnsLocalRecord(ctx context.Context, organizationId string, recordId string) RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest {
	return RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		recordId: recordId,
	}
}

// Execute executes the request
//  @return InlineResponse200222
func (a *RecordsApiService) UpdateOrganizationApplianceDnsLocalRecordExecute(r RecordsApiUpdateOrganizationApplianceDnsLocalRecordRequest) (*InlineResponse200222, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200222
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.UpdateOrganizationApplianceDnsLocalRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/local/records/{recordId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterToString(r.recordId, "")), -1)

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
	localVarPostBody = r.updateOrganizationApplianceDnsLocalRecord
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
