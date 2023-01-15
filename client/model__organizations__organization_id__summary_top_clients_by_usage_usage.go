/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopClientsByUsageUsage Data usage information
type OrganizationsOrganizationIdSummaryTopClientsByUsageUsage struct {
	// Total data usage by client
	Total *float32 `json:"total,omitempty"`
	// Upstream data usage by client
	Upstream *float32 `json:"upstream,omitempty"`
	// Downstream data usage by client
	Downstream *float32 `json:"downstream,omitempty"`
	// Percentage of total data usage by client
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopClientsByUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopClientsByUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopClientsByUsageUsage() *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopClientsByUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopClientsByUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopClientsByUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopClientsByUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopClientsByUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage(val *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


