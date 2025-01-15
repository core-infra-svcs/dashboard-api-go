/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200161MacBlocklist MAC blocklist
type InlineResponse200161MacBlocklist struct {
	// Enable MAC blocklist for switches in the network
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200161MacBlocklist instantiates a new InlineResponse200161MacBlocklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200161MacBlocklist() *InlineResponse200161MacBlocklist {
	this := InlineResponse200161MacBlocklist{}
	return &this
}

// NewInlineResponse200161MacBlocklistWithDefaults instantiates a new InlineResponse200161MacBlocklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200161MacBlocklistWithDefaults() *InlineResponse200161MacBlocklist {
	this := InlineResponse200161MacBlocklist{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200161MacBlocklist) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161MacBlocklist) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200161MacBlocklist) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200161MacBlocklist) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200161MacBlocklist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200161MacBlocklist struct {
	value *InlineResponse200161MacBlocklist
	isSet bool
}

func (v NullableInlineResponse200161MacBlocklist) Get() *InlineResponse200161MacBlocklist {
	return v.value
}

func (v *NullableInlineResponse200161MacBlocklist) Set(val *InlineResponse200161MacBlocklist) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200161MacBlocklist) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200161MacBlocklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200161MacBlocklist(val *InlineResponse200161MacBlocklist) *NullableInlineResponse200161MacBlocklist {
	return &NullableInlineResponse200161MacBlocklist{value: val, isSet: true}
}

func (v NullableInlineResponse200161MacBlocklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200161MacBlocklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


