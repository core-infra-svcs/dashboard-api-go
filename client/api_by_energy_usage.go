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

// ByEnergyUsageApiService ByEnergyUsageApi service
type ByEnergyUsageApiService service

type ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest struct {
	ctx _context.Context
	ApiService *ByEnergyUsageApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T0(t0 string) ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T1(t1 string) ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Timespan(timespan float32) ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Execute() ([]InlineResponse20097, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r)
}

/*
GetOrganizationSummaryTopSwitchesByEnergyUsage Return metrics for organization's top 10 switches by energy usage over given time range

Return metrics for organization's top 10 switches by energy usage over given time range. Default unit is joules.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest
*/
func (a *ByEnergyUsageApiService) GetOrganizationSummaryTopSwitchesByEnergyUsage(ctx _context.Context, organizationId string) ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	return ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20097
func (a *ByEnergyUsageApiService) GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r ByEnergyUsageApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) ([]InlineResponse20097, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20097
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByEnergyUsageApiService.GetOrganizationSummaryTopSwitchesByEnergyUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/switches/byEnergyUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
