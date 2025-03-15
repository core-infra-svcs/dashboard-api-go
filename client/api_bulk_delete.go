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


// BulkDeleteApiService BulkDeleteApi service
type BulkDeleteApiService service

type BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest struct {
	ctx context.Context
	ApiService *BulkDeleteApiService
	organizationId string
	createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete *InlineObject220
}

func (r BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest) CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete InlineObject220) BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest {
	r.createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete = &createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete
	return r
}

func (r BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest) Execute() (*InlineResponse200221, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteExecute(r)
}

/*
CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete Unassign the local DNS profile to networks in the organization

Unassign the local DNS profile to networks in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest
*/
func (a *BulkDeleteApiService) CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(ctx context.Context, organizationId string) BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest {
	return BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200221
func (a *BulkDeleteApiService) CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteExecute(r BulkDeleteApiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest) (*InlineResponse200221, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200221
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkDeleteApiService.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkDelete"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete == nil {
		return localVarReturnValue, nil, reportError("createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete is required and must be specified")
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
	localVarPostBody = r.createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete
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

type BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest struct {
	ctx context.Context
	ApiService *BulkDeleteApiService
	organizationId string
	createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete *InlineObject226
}

func (r BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete InlineObject226) BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest {
	r.createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete = &createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete
	return r
}

func (r BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) Execute() (*InlineResponse200225, *http.Response, error) {
	return r.ApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteExecute(r)
}

/*
CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete Unassign the split DNS profile to networks in the organization

Unassign the split DNS profile to networks in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest
*/
func (a *BulkDeleteApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(ctx context.Context, organizationId string) BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest {
	return BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200225
func (a *BulkDeleteApiService) CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteExecute(r BulkDeleteApiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest) (*InlineResponse200225, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200225
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkDeleteApiService.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete")
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
