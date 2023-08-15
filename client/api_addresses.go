/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
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


// AddressesApiService AddressesApi service
type AddressesApiService service

type AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest struct {
	ctx context.Context
	ApiService *AddressesApiService
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
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) PerPage(perPage int32) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) StartingAfter(startingAfter string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) EndingBefore(endingBefore string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) NetworkIds(networkIds []string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) ProductTypes(productTypes []string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Serials(serials []string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.serials = &serials
	return r
}

// An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Tags(tags []string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.tags = &tags
	return r
}

// An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected.
func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) TagsFilterType(tagsFilterType string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

func (r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Execute() ([]InlineResponse200129, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesUplinksAddressesByDeviceExecute(r)
}

/*
GetOrganizationDevicesUplinksAddressesByDevice List the current uplink addresses for devices in an organization.

List the current uplink addresses for devices in an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest
*/
func (a *AddressesApiService) GetOrganizationDevicesUplinksAddressesByDevice(ctx context.Context, organizationId string) AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	return AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200129
func (a *AddressesApiService) GetOrganizationDevicesUplinksAddressesByDeviceExecute(r AddressesApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) ([]InlineResponse200129, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200129
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesApiService.GetOrganizationDevicesUplinksAddressesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/uplinks/addresses/byDevice"
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
