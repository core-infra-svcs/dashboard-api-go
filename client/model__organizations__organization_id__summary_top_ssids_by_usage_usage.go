/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage Date usage of the SSID, in megabytes
type OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage struct {
	// Total usage of the SSID
	Total *float32 `json:"total,omitempty"`
	// Downstream usage of the SSID
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream usage of the SSID
	Upstream *float32 `json:"upstream,omitempty"`
	// Percentage usage of the SSID
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage() *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


