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

// PowerModulesApiService PowerModulesApi service
type PowerModulesApiService service

type PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest struct {
	ctx _context.Context
	ApiService *PowerModulesApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	serials *[]string
	tags *[]string
	tagsFilterType *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) PerPage(perPage int32) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) StartingAfter(startingAfter string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) EndingBefore(endingBefore string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}
// Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) NetworkIds(networkIds []string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}
// Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) ProductTypes(productTypes []string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.productTypes = &productTypes
	return r
}
// Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Serials(serials []string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.serials = &serials
	return r
}
// An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Tags(tags []string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.tags = &tags
	return r
}
// An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected.
func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) TagsFilterType(tagsFilterType string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

func (r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Execute() ([]InlineResponse20070, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationDevicesPowerModulesStatusesByDeviceExecute(r)
}

/*
GetOrganizationDevicesPowerModulesStatusesByDevice List the power status information for devices in an organization

List the power status information for devices in an organization. The data returned by this endpoint is updated every 5 minutes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest
*/
func (a *PowerModulesApiService) GetOrganizationDevicesPowerModulesStatusesByDevice(ctx _context.Context, organizationId string) PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	return PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20070
func (a *PowerModulesApiService) GetOrganizationDevicesPowerModulesStatusesByDeviceExecute(r PowerModulesApiApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) ([]InlineResponse20070, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20070
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerModulesApiService.GetOrganizationDevicesPowerModulesStatusesByDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/powerModules/statuses/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	}
	if r.tagsFilterType != nil {
		localVarQueryParams.Add("tagsFilterType", parameterToString(*r.tagsFilterType, ""))
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
