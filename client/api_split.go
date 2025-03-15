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


// SplitApiService SplitApi service
type SplitApiService service

type SplitApiCreateOrganizationApplianceDnsSplitProfileRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	createOrganizationApplianceDnsSplitProfile *InlineObject224
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfileRequest) CreateOrganizationApplianceDnsSplitProfile(createOrganizationApplianceDnsSplitProfile InlineObject224) SplitApiCreateOrganizationApplianceDnsSplitProfileRequest {
	r.createOrganizationApplianceDnsSplitProfile = &createOrganizationApplianceDnsSplitProfile
	return r
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfileRequest) Execute() (*InlineResponse200223, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsSplitProfileExecute(r)
}

/*
CreateOrganizationApplianceDnsSplitProfile Create a new split DNS profile

Create a new split DNS profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SplitApiCreateOrganizationApplianceDnsSplitProfileRequest
*/
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfile(ctx context.Context, organizationId string) SplitApiCreateOrganizationApplianceDnsSplitProfileRequest {
	return SplitApiCreateOrganizationApplianceDnsSplitProfileRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200223
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfileExecute(r SplitApiCreateOrganizationApplianceDnsSplitProfileRequest) (*InlineResponse200223, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200223
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.CreateOrganizationApplianceDnsSplitProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationApplianceDnsSplitProfile == nil {
		return localVarReturnValue, nil, reportError("createOrganizationApplianceDnsSplitProfile is required and must be specified")
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
	localVarPostBody = r.createOrganizationApplianceDnsSplitProfile
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

type SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate *InlineObject225
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate InlineObject225) SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest {
	r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate = &createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate
	return r
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest) Execute() (*InlineResponse200225, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateExecute(r)
}

/*
CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate Assign the split DNS profile to networks in the organization

Assign the split DNS profile to networks in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest
*/
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(ctx context.Context, organizationId string) SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest {
	return SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200225
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateExecute(r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest) (*InlineResponse200225, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200225
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkCreate"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate == nil {
		return localVarReturnValue, nil, reportError("createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate is required and must be specified")
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
	localVarPostBody = r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate
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

type SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete *InlineObject226
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete InlineObject226) SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest {
	r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete = &createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete
	return r
}

func (r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) Execute() (*InlineResponse200225, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteExecute(r)
}

/*
CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete Unassign the split DNS profile to networks in the organization

Unassign the split DNS profile to networks in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest
*/
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(ctx context.Context, organizationId string) SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest {
	return SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200225
func (a *SplitApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteExecute(r SplitApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) (*InlineResponse200225, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200225
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkDelete"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete == nil {
		return localVarReturnValue, nil, reportError("createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete is required and must be specified")
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
	localVarPostBody = r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete
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

type SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	profileId string
}

func (r SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationApplianceDnsSplitProfileExecute(r)
}

/*
DeleteOrganizationApplianceDnsSplitProfile Deletes a split DNS profile

Deletes a split DNS profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param profileId Profile ID
 @return SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest
*/
func (a *SplitApiService) DeleteOrganizationApplianceDnsSplitProfile(ctx context.Context, organizationId string, profileId string) SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest {
	return SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		profileId: profileId,
	}
}

// Execute executes the request
func (a *SplitApiService) DeleteOrganizationApplianceDnsSplitProfileExecute(r SplitApiDeleteOrganizationApplianceDnsSplitProfileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.DeleteOrganizationApplianceDnsSplitProfile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles/{profileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterToString(r.profileId, "")), -1)

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

type SplitApiGetOrganizationApplianceDnsSplitProfilesRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	profileIds *[]string
}

// Optional parameter to filter the results by profile IDs
func (r SplitApiGetOrganizationApplianceDnsSplitProfilesRequest) ProfileIds(profileIds []string) SplitApiGetOrganizationApplianceDnsSplitProfilesRequest {
	r.profileIds = &profileIds
	return r
}

func (r SplitApiGetOrganizationApplianceDnsSplitProfilesRequest) Execute() ([]InlineResponse200223, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceDnsSplitProfilesExecute(r)
}

/*
GetOrganizationApplianceDnsSplitProfiles Fetch the split DNS profiles used in the organization

Fetch the split DNS profiles used in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SplitApiGetOrganizationApplianceDnsSplitProfilesRequest
*/
func (a *SplitApiService) GetOrganizationApplianceDnsSplitProfiles(ctx context.Context, organizationId string) SplitApiGetOrganizationApplianceDnsSplitProfilesRequest {
	return SplitApiGetOrganizationApplianceDnsSplitProfilesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200223
func (a *SplitApiService) GetOrganizationApplianceDnsSplitProfilesExecute(r SplitApiGetOrganizationApplianceDnsSplitProfilesRequest) ([]InlineResponse200223, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200223
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.GetOrganizationApplianceDnsSplitProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles"
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

type SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	profileIds *[]string
	networkIds *[]string
}

// Optional parameter to filter the results by profile IDs
func (r SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest) ProfileIds(profileIds []string) SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest {
	r.profileIds = &profileIds
	return r
}

// Optional parameter to filter the results by network IDs
func (r SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest) NetworkIds(networkIds []string) SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest {
	r.networkIds = &networkIds
	return r
}

func (r SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest) Execute() (*InlineResponse200224, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceDnsSplitProfilesAssignmentsExecute(r)
}

/*
GetOrganizationApplianceDnsSplitProfilesAssignments Fetch the split DNS profile assignments in the organization

Fetch the split DNS profile assignments in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest
*/
func (a *SplitApiService) GetOrganizationApplianceDnsSplitProfilesAssignments(ctx context.Context, organizationId string) SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest {
	return SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200224
func (a *SplitApiService) GetOrganizationApplianceDnsSplitProfilesAssignmentsExecute(r SplitApiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest) (*InlineResponse200224, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200224
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.GetOrganizationApplianceDnsSplitProfilesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.profileIds != nil {
		localVarQueryParams.Add("profileIds", parameterToString(*r.profileIds, "csv"))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
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

type SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest struct {
	ctx context.Context
	ApiService *SplitApiService
	organizationId string
	profileId string
	updateOrganizationApplianceDnsSplitProfile *InlineObject227
}

func (r SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest) UpdateOrganizationApplianceDnsSplitProfile(updateOrganizationApplianceDnsSplitProfile InlineObject227) SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest {
	r.updateOrganizationApplianceDnsSplitProfile = &updateOrganizationApplianceDnsSplitProfile
	return r
}

func (r SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest) Execute() (*InlineResponse200223, *http.Response, error) {
	return r.ApiService.UpdateOrganizationApplianceDnsSplitProfileExecute(r)
}

/*
UpdateOrganizationApplianceDnsSplitProfile Update a split DNS profile

Update a split DNS profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param profileId Profile ID
 @return SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest
*/
func (a *SplitApiService) UpdateOrganizationApplianceDnsSplitProfile(ctx context.Context, organizationId string, profileId string) SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest {
	return SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		profileId: profileId,
	}
}

// Execute executes the request
//  @return InlineResponse200223
func (a *SplitApiService) UpdateOrganizationApplianceDnsSplitProfileExecute(r SplitApiUpdateOrganizationApplianceDnsSplitProfileRequest) (*InlineResponse200223, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200223
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SplitApiService.UpdateOrganizationApplianceDnsSplitProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/split/profiles/{profileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterToString(r.profileId, "")), -1)

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
	localVarPostBody = r.updateOrganizationApplianceDnsSplitProfile
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
