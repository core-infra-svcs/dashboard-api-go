/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20059Ipv6 IPv6 configuration on the single LAN
type InlineResponse20059Ipv6 struct {
	// Enable IPv6 on single LAN
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the single LAN
	PrefixAssignments []InlineResponse20059Ipv6PrefixAssignments `json:"prefixAssignments,omitempty"`
}

// NewInlineResponse20059Ipv6 instantiates a new InlineResponse20059Ipv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059Ipv6() *InlineResponse20059Ipv6 {
	this := InlineResponse20059Ipv6{}
	return &this
}

// NewInlineResponse20059Ipv6WithDefaults instantiates a new InlineResponse20059Ipv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059Ipv6WithDefaults() *InlineResponse20059Ipv6 {
	this := InlineResponse20059Ipv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20059Ipv6) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059Ipv6) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20059Ipv6) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20059Ipv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *InlineResponse20059Ipv6) GetPrefixAssignments() []InlineResponse20059Ipv6PrefixAssignments {
	if o == nil || isNil(o.PrefixAssignments) {
		var ret []InlineResponse20059Ipv6PrefixAssignments
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059Ipv6) GetPrefixAssignmentsOk() ([]InlineResponse20059Ipv6PrefixAssignments, bool) {
	if o == nil || isNil(o.PrefixAssignments) {
    return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *InlineResponse20059Ipv6) HasPrefixAssignments() bool {
	if o != nil && !isNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []InlineResponse20059Ipv6PrefixAssignments and assigns it to the PrefixAssignments field.
func (o *InlineResponse20059Ipv6) SetPrefixAssignments(v []InlineResponse20059Ipv6PrefixAssignments) {
	o.PrefixAssignments = v
}

func (o InlineResponse20059Ipv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059Ipv6 struct {
	value *InlineResponse20059Ipv6
	isSet bool
}

func (v NullableInlineResponse20059Ipv6) Get() *InlineResponse20059Ipv6 {
	return v.value
}

func (v *NullableInlineResponse20059Ipv6) Set(val *InlineResponse20059Ipv6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059Ipv6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059Ipv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059Ipv6(val *InlineResponse20059Ipv6) *NullableInlineResponse20059Ipv6 {
	return &NullableInlineResponse20059Ipv6{value: val, isSet: true}
}

func (v NullableInlineResponse20059Ipv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059Ipv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


