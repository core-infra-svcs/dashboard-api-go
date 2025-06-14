/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage Clients usage
type OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage struct {
	// Total data usage by client
	Total *float32 `json:"total,omitempty"`
	// Upstream data usage by client
	Upstream *float32 `json:"upstream,omitempty"`
	// Downstream data usage by client
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


