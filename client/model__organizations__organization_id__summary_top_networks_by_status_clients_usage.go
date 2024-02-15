/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage Network client usage data
type OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage struct {
	// Total upstream usage in network, in KB
	Upstream *float32 `json:"upstream,omitempty"`
	// Total downstream usage in network, in KB
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage{}
	return &this
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


