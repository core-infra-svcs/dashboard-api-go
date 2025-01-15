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
	"time"
)


// DiscoveryApiService DiscoveryApi service
type DiscoveryApiService service

type DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest struct {
	ctx context.Context
	ApiService *DiscoveryApiService
	organizationId string
	t0 *string
	timespan *float32
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

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) T0(t0 string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Timespan(timespan float32) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) PerPage(perPage int32) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) StartingAfter(startingAfter string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) EndingBefore(endingBefore string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) ConfigurationUpdatedAfter(configurationUpdatedAfter time.Time) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.configurationUpdatedAfter = &configurationUpdatedAfter
	return r
}

// Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Mac(mac string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.mac = &mac
	return r
}

// Optional parameter to filter items to switches that have one of the provided MAC addresses.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Macs(macs []string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter items to switches with names that contain the search term or are an exact match.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Name(name string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.name = &name
	return r
}

// Optional parameter to filter items to switches in one of the provided networks.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) NetworkIds(networkIds []string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) PortProfileIds(portProfileIds []string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.portProfileIds = &portProfileIds
	return r
}

// Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Serial(serial string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.serial = &serial
	return r
}

// Optional parameter to filter items to switches that have one of the provided serials.
func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Serials(serials []string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	r.serials = &serials
	return r
}

func (r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) Execute() (*InlineResponse200310, *http.Response, error) {
	return r.ApiService.GetOrganizationSwitchPortsTopologyDiscoveryByDeviceExecute(r)
}

/*
GetOrganizationSwitchPortsTopologyDiscoveryByDevice List most recently seen LLDP/CDP discovery and topology information per switch port in an organization.

List most recently seen LLDP/CDP discovery and topology information per switch port in an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest
*/
func (a *DiscoveryApiService) GetOrganizationSwitchPortsTopologyDiscoveryByDevice(ctx context.Context, organizationId string) DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest {
	return DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200310
func (a *DiscoveryApiService) GetOrganizationSwitchPortsTopologyDiscoveryByDeviceExecute(r DiscoveryApiGetOrganizationSwitchPortsTopologyDiscoveryByDeviceRequest) (*InlineResponse200310, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200310
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryApiService.GetOrganizationSwitchPortsTopologyDiscoveryByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/switch/ports/topology/discovery/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
