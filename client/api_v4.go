/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
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


// V4ApiService V4Api service
type V4ApiService service

type V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest struct {
	ctx context.Context
	ApiService *V4ApiService
	networkId string
	t0 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) T0(t0 string) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) Timespan(timespan float32) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) PerPage(perPage int32) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) StartingAfter(startingAfter string) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) EndingBefore(endingBefore string) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) Execute() ([]InlineResponse200148, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpV4ServersSeenExecute(r)
}

/*
GetNetworkSwitchDhcpV4ServersSeen Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)

Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest
*/
func (a *V4ApiService) GetNetworkSwitchDhcpV4ServersSeen(ctx context.Context, networkId string) V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	return V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200148
func (a *V4ApiService) GetNetworkSwitchDhcpV4ServersSeenExecute(r V4ApiGetNetworkSwitchDhcpV4ServersSeenRequest) ([]InlineResponse200148, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200148
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V4ApiService.GetNetworkSwitchDhcpV4ServersSeen")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcp/v4/servers/seen"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
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
