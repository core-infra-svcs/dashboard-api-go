/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
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


// SentryApiService SentryApi service
type SentryApiService service

type SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest struct {
	ctx context.Context
	ApiService *SentryApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) PerPage(perPage int32) SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) StartingAfter(startingAfter string) SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) EndingBefore(endingBefore string) SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter Sentry Policies by Network Id
func (r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) NetworkIds(networkIds []string) SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

func (r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) Execute() ([]InlineResponse200289, *http.Response, error) {
	return r.ApiService.GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r)
}

/*
GetOrganizationSmSentryPoliciesAssignmentsByNetwork List the Sentry Policies for an organization ordered in ascending order of priority

List the Sentry Policies for an organization ordered in ascending order of priority

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest
*/
func (a *SentryApiService) GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx context.Context, organizationId string) SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	return SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200289
func (a *SentryApiService) GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r SentryApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) ([]InlineResponse200289, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200289
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SentryApiService.GetOrganizationSmSentryPoliciesAssignmentsByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
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

type SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest struct {
	ctx context.Context
	ApiService *SentryApiService
	organizationId string
	updateOrganizationSmSentryPoliciesAssignments *InlineObject270
}

func (r SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments InlineObject270) SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	r.updateOrganizationSmSentryPoliciesAssignments = &updateOrganizationSmSentryPoliciesAssignments
	return r
}

func (r SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) Execute() (*InlineResponse200288, *http.Response, error) {
	return r.ApiService.UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r)
}

/*
UpdateOrganizationSmSentryPoliciesAssignments Update an Organizations Sentry Policies using the provided list

Update an Organizations Sentry Policies using the provided list. Sentry Policies are ordered in descending order of priority (i.e. highest priority at the bottom, this is opposite the Dashboard UI). Policies not present in the request will be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest
*/
func (a *SentryApiService) UpdateOrganizationSmSentryPoliciesAssignments(ctx context.Context, organizationId string) SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	return SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200288
func (a *SentryApiService) UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r SentryApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) (*InlineResponse200288, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200288
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SentryApiService.UpdateOrganizationSmSentryPoliciesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrganizationSmSentryPoliciesAssignments == nil {
		return localVarReturnValue, nil, reportError("updateOrganizationSmSentryPoliciesAssignments is required and must be specified")
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
	localVarPostBody = r.updateOrganizationSmSentryPoliciesAssignments
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
