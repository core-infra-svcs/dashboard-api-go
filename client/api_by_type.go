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


// ByTypeApiService ByTypeApi service
type ByTypeApiService service

type ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest struct {
	ctx context.Context
	ApiService *ByTypeApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	sortOrder *string
	networkId *string
	severity *string
	types *[]string
	tsStart *time.Time
	tsEnd *time.Time
	category *string
	sortBy *string
	serials *[]string
	deviceTypes *[]string
	deviceTags *[]string
	active *bool
	dismissed *bool
	resolved *bool
	suppressAlertsForOfflineNodes *bool
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) PerPage(perPage int32) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) StartingAfter(startingAfter string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) EndingBefore(endingBefore string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.endingBefore = &endingBefore
	return r
}

// Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) SortOrder(sortOrder string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.sortOrder = &sortOrder
	return r
}

// Optional parameter to filter alerts overview by network ids.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) NetworkId(networkId string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.networkId = &networkId
	return r
}

// Optional parameter to filter alerts overview by severity type.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Severity(severity string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.severity = &severity
	return r
}

// Optional parameter to filter by alert type.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Types(types []string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.types = &types
	return r
}

// Optional parameter to filter by starting timestamp
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) TsStart(tsStart time.Time) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.tsStart = &tsStart
	return r
}

// Optional parameter to filter by end timestamp
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) TsEnd(tsEnd time.Time) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.tsEnd = &tsEnd
	return r
}

// Optional parameter to filter by category.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Category(category string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.category = &category
	return r
}

// Optional parameter to set column to sort by.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) SortBy(sortBy string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.sortBy = &sortBy
	return r
}

// Optional parameter to filter by primary device serial
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Serials(serials []string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter by device types
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) DeviceTypes(deviceTypes []string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.deviceTypes = &deviceTypes
	return r
}

// Optional parameter to filter by device tags
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) DeviceTags(deviceTags []string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.deviceTags = &deviceTags
	return r
}

// Optional parameter to filter by active alerts defaults to true
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Active(active bool) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.active = &active
	return r
}

// Optional parameter to filter by dismissed alerts defaults to false
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Dismissed(dismissed bool) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.dismissed = &dismissed
	return r
}

// Optional parameter to filter by resolved alerts defaults to false
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Resolved(resolved bool) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.resolved = &resolved
	return r
}

// When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false.
func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes bool) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	r.suppressAlertsForOfflineNodes = &suppressAlertsForOfflineNodes
	return r
}

func (r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) Execute() (*InlineResponse200241, *http.Response, error) {
	return r.ApiService.GetOrganizationAssuranceAlertsOverviewByTypeExecute(r)
}

/*
GetOrganizationAssuranceAlertsOverviewByType Return a Summary of Alerts grouped by type and severity

Return a Summary of Alerts grouped by type and severity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest
*/
func (a *ByTypeApiService) GetOrganizationAssuranceAlertsOverviewByType(ctx context.Context, organizationId string) ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest {
	return ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200241
func (a *ByTypeApiService) GetOrganizationAssuranceAlertsOverviewByTypeExecute(r ByTypeApiGetOrganizationAssuranceAlertsOverviewByTypeRequest) (*InlineResponse200241, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200241
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByTypeApiService.GetOrganizationAssuranceAlertsOverviewByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/assurance/alerts/overview/byType"
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
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	if r.networkId != nil {
		localVarQueryParams.Add("networkId", parameterToString(*r.networkId, ""))
	}
	if r.severity != nil {
		localVarQueryParams.Add("severity", parameterToString(*r.severity, ""))
	}
	if r.types != nil {
		localVarQueryParams.Add("types", parameterToString(*r.types, "csv"))
	}
	if r.tsStart != nil {
		localVarQueryParams.Add("tsStart", parameterToString(*r.tsStart, ""))
	}
	if r.tsEnd != nil {
		localVarQueryParams.Add("tsEnd", parameterToString(*r.tsEnd, ""))
	}
	if r.category != nil {
		localVarQueryParams.Add("category", parameterToString(*r.category, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sortBy", parameterToString(*r.sortBy, ""))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.deviceTypes != nil {
		localVarQueryParams.Add("deviceTypes", parameterToString(*r.deviceTypes, "csv"))
	}
	if r.deviceTags != nil {
		localVarQueryParams.Add("deviceTags", parameterToString(*r.deviceTags, "csv"))
	}
	if r.active != nil {
		localVarQueryParams.Add("active", parameterToString(*r.active, ""))
	}
	if r.dismissed != nil {
		localVarQueryParams.Add("dismissed", parameterToString(*r.dismissed, ""))
	}
	if r.resolved != nil {
		localVarQueryParams.Add("resolved", parameterToString(*r.resolved, ""))
	}
	if r.suppressAlertsForOfflineNodes != nil {
		localVarQueryParams.Add("suppressAlertsForOfflineNodes", parameterToString(*r.suppressAlertsForOfflineNodes, ""))
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
