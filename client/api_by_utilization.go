/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
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


// ByUtilizationApiService ByUtilizationApi service
type ByUtilizationApiService service

type ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest struct {
	ctx context.Context
	ApiService *ByUtilizationApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) T0(t0 string) ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) T1(t1 string) ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 31 days. The default is 1 day.
func (r ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) Timespan(timespan float32) ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.timespan = &timespan
	return r
}

func (r ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) Execute() ([]InlineResponse200169, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopAppliancesByUtilizationExecute(r)
}

/*
GetOrganizationSummaryTopAppliancesByUtilization Return the top 10 appliances sorted by utilization over given time range.

Return the top 10 appliances sorted by utilization over given time range.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest
*/
func (a *ByUtilizationApiService) GetOrganizationSummaryTopAppliancesByUtilization(ctx context.Context, organizationId string) ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	return ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200169
func (a *ByUtilizationApiService) GetOrganizationSummaryTopAppliancesByUtilizationExecute(r ByUtilizationApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) ([]InlineResponse200169, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200169
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByUtilizationApiService.GetOrganizationSummaryTopAppliancesByUtilization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/appliances/byUtilization"
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
