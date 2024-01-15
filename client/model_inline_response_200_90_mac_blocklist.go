/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20090MacBlocklist MAC blocklist
type InlineResponse20090MacBlocklist struct {
	// Enable MAC blocklist for switches in the network
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20090MacBlocklist instantiates a new InlineResponse20090MacBlocklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090MacBlocklist() *InlineResponse20090MacBlocklist {
	this := InlineResponse20090MacBlocklist{}
	return &this
}

// NewInlineResponse20090MacBlocklistWithDefaults instantiates a new InlineResponse20090MacBlocklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090MacBlocklistWithDefaults() *InlineResponse20090MacBlocklist {
	this := InlineResponse20090MacBlocklist{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20090MacBlocklist) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090MacBlocklist) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20090MacBlocklist) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20090MacBlocklist) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20090MacBlocklist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090MacBlocklist struct {
	value *InlineResponse20090MacBlocklist
	isSet bool
}

func (v NullableInlineResponse20090MacBlocklist) Get() *InlineResponse20090MacBlocklist {
	return v.value
}

func (v *NullableInlineResponse20090MacBlocklist) Set(val *InlineResponse20090MacBlocklist) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090MacBlocklist) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090MacBlocklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090MacBlocklist(val *InlineResponse20090MacBlocklist) *NullableInlineResponse20090MacBlocklist {
	return &NullableInlineResponse20090MacBlocklist{value: val, isSet: true}
}

func (v NullableInlineResponse20090MacBlocklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090MacBlocklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

