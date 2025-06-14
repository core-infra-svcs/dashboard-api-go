/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
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
	"time"
)


// BySwitchApiService BySwitchApi service
type BySwitchApiService service

type BySwitchApiGetOrganizationSwitchPortsBySwitchRequest struct {
	ctx context.Context
	ApiService *BySwitchApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	configurationUpdatedAfter *time.Time
	mac *string
	macs *[]string
	name *string
	networkIds *[]string
	portProfileIds *[]string
	serial *string
	serials *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) PerPage(perPage int32) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) StartingAfter(startingAfter string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) EndingBefore(endingBefore string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) ConfigurationUpdatedAfter(configurationUpdatedAfter time.Time) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.configurationUpdatedAfter = &configurationUpdatedAfter
	return r
}

// Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Mac(mac string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.mac = &mac
	return r
}

// Optional parameter to filter items to switches that have one of the provided MAC addresses.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Macs(macs []string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter items to switches with names that contain the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Name(name string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.name = &name
	return r
}

// Optional parameter to filter items to switches in one of the provided networks.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) NetworkIds(networkIds []string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) PortProfileIds(portProfileIds []string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.portProfileIds = &portProfileIds
	return r
}

// Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Serial(serial string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.serial = &serial
	return r
}

// Optional parameter to filter items to switches that have one of the provided serials.
func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Serials(serials []string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	r.serials = &serials
	return r
}

func (r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) Execute() (*InlineResponse200335, *http.Response, error) {
	return r.ApiService.GetOrganizationSwitchPortsBySwitchExecute(r)
}

/*
GetOrganizationSwitchPortsBySwitch List the switchports in an organization by switch

List the switchports in an organization by switch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BySwitchApiGetOrganizationSwitchPortsBySwitchRequest
*/
func (a *BySwitchApiService) GetOrganizationSwitchPortsBySwitch(ctx context.Context, organizationId string) BySwitchApiGetOrganizationSwitchPortsBySwitchRequest {
	return BySwitchApiGetOrganizationSwitchPortsBySwitchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200335
func (a *BySwitchApiService) GetOrganizationSwitchPortsBySwitchExecute(r BySwitchApiGetOrganizationSwitchPortsBySwitchRequest) (*InlineResponse200335, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200335
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BySwitchApiService.GetOrganizationSwitchPortsBySwitch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/switch/ports/bySwitch"
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
	if r.configurationUpdatedAfter != nil {
		localVarQueryParams.Add("configurationUpdatedAfter", parameterToString(*r.configurationUpdatedAfter, ""))
	}
	if r.mac != nil {
		localVarQueryParams.Add("mac", parameterToString(*r.mac, ""))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.portProfileIds != nil {
		localVarQueryParams.Add("portProfileIds", parameterToString(*r.portProfileIds, "csv"))
	}
	if r.serial != nil {
		localVarQueryParams.Add("serial", parameterToString(*r.serial, ""))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
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

type BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest struct {
	ctx context.Context
	ApiService *BySwitchApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	configurationUpdatedAfter *time.Time
	mac *string
	macs *[]string
	name *string
	networkIds *[]string
	portProfileIds *[]string
	serial *string
	serials *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) PerPage(perPage int32) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) StartingAfter(startingAfter string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) EndingBefore(endingBefore string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) ConfigurationUpdatedAfter(configurationUpdatedAfter time.Time) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.configurationUpdatedAfter = &configurationUpdatedAfter
	return r
}

// Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Mac(mac string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.mac = &mac
	return r
}

// Optional parameter to filter items to switches that have one of the provided MAC addresses.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Macs(macs []string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter items to switches with names that contain the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Name(name string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.name = &name
	return r
}

// Optional parameter to filter items to switches in one of the provided networks.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) NetworkIds(networkIds []string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) PortProfileIds(portProfileIds []string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.portProfileIds = &portProfileIds
	return r
}

// Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Serial(serial string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.serial = &serial
	return r
}

// Optional parameter to filter items to switches that have one of the provided serials.
func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Serials(serials []string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	r.serials = &serials
	return r
}

func (r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) Execute() (*InlineResponse200338, *http.Response, error) {
	return r.ApiService.GetOrganizationSwitchPortsStatusesBySwitchExecute(r)
}

/*
GetOrganizationSwitchPortsStatusesBySwitch List the switchports in an organization

List the switchports in an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest
*/
func (a *BySwitchApiService) GetOrganizationSwitchPortsStatusesBySwitch(ctx context.Context, organizationId string) BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest {
	return BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200338
func (a *BySwitchApiService) GetOrganizationSwitchPortsStatusesBySwitchExecute(r BySwitchApiGetOrganizationSwitchPortsStatusesBySwitchRequest) (*InlineResponse200338, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200338
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BySwitchApiService.GetOrganizationSwitchPortsStatusesBySwitch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/switch/ports/statuses/bySwitch"
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
	if r.configurationUpdatedAfter != nil {
		localVarQueryParams.Add("configurationUpdatedAfter", parameterToString(*r.configurationUpdatedAfter, ""))
	}
	if r.mac != nil {
		localVarQueryParams.Add("mac", parameterToString(*r.mac, ""))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.portProfileIds != nil {
		localVarQueryParams.Add("portProfileIds", parameterToString(*r.portProfileIds, "csv"))
	}
	if r.serial != nil {
		localVarQueryParams.Add("serial", parameterToString(*r.serial, ""))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
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
