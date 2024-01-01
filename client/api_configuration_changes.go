/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
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


// ConfigurationChangesApiService ConfigurationChangesApi service
type ConfigurationChangesApiService service

type ConfigurationChangesApiGetOrganizationConfigurationChangesRequest struct {
	ctx context.Context
	ApiService *ConfigurationChangesApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkId *string
	adminId *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) T0(t0 string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) T1(t1 string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) Timespan(timespan float32) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 5000. Default is 5000.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) PerPage(perPage int32) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) StartingAfter(startingAfter string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) EndingBefore(endingBefore string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Filters on the given network
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) NetworkId(networkId string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.networkId = &networkId
	return r
}

// Filters on the given Admin
func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) AdminId(adminId string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	r.adminId = &adminId
	return r
}

func (r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) Execute() ([]InlineResponse200139, *http.Response, error) {
	return r.ApiService.GetOrganizationConfigurationChangesExecute(r)
}

/*
GetOrganizationConfigurationChanges View the Change Log for your organization

View the Change Log for your organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ConfigurationChangesApiGetOrganizationConfigurationChangesRequest
*/
func (a *ConfigurationChangesApiService) GetOrganizationConfigurationChanges(ctx context.Context, organizationId string) ConfigurationChangesApiGetOrganizationConfigurationChangesRequest {
	return ConfigurationChangesApiGetOrganizationConfigurationChangesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200139
func (a *ConfigurationChangesApiService) GetOrganizationConfigurationChangesExecute(r ConfigurationChangesApiGetOrganizationConfigurationChangesRequest) ([]InlineResponse200139, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200139
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationChangesApiService.GetOrganizationConfigurationChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/configurationChanges"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.networkId != nil {
		localVarQueryParams.Add("networkId", parameterToString(*r.networkId, ""))
	}
	if r.adminId != nil {
		localVarQueryParams.Add("adminId", parameterToString(*r.adminId, ""))
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
