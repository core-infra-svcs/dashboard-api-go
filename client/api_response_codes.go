/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
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


// ResponseCodesApiService ResponseCodesApi service
type ResponseCodesApiService service

type ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct {
	ctx context.Context
	ApiService *ResponseCodesApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
	version *int32
	operationIds *[]string
	sourceIps *[]string
	adminIds *[]string
	userAgent *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T0(t0 string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T1(t1 string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Timespan(timespan float32) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Interval(interval int32) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.interval = &interval
	return r
}

// Filter by API version of the endpoint. Allowable values are: [0, 1]
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Version(version int32) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.version = &version
	return r
}

// Filter by operation ID of the endpoint
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) OperationIds(operationIds []string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.operationIds = &operationIds
	return r
}

// Filter by source IP that made the API request
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) SourceIps(sourceIps []string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.sourceIps = &sourceIps
	return r
}

// Filter by admin ID of user that made the API request
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) AdminIds(adminIds []string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.adminIds = &adminIds
	return r
}

// Filter by user agent string for API request. This will filter by a complete or partial match.
func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) UserAgent(userAgent string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.userAgent = &userAgent
	return r
}

func (r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Execute() ([]InlineResponse200120, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r)
}

/*
GetOrganizationApiRequestsOverviewResponseCodesByInterval Tracks organizations' API requests by response code across a given time period

Tracks organizations' API requests by response code across a given time period

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest
*/
func (a *ResponseCodesApiService) GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx context.Context, organizationId string) ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	return ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200120
func (a *ResponseCodesApiService) GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r ResponseCodesApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) ([]InlineResponse200120, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200120
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResponseCodesApiService.GetOrganizationApiRequestsOverviewResponseCodesByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval"
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
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.operationIds != nil {
		localVarQueryParams.Add("operationIds", parameterToString(*r.operationIds, "csv"))
	}
	if r.sourceIps != nil {
		localVarQueryParams.Add("sourceIps", parameterToString(*r.sourceIps, "csv"))
	}
	if r.adminIds != nil {
		localVarQueryParams.Add("adminIds", parameterToString(*r.adminIds, "csv"))
	}
	if r.userAgent != nil {
		localVarQueryParams.Add("userAgent", parameterToString(*r.userAgent, ""))
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
