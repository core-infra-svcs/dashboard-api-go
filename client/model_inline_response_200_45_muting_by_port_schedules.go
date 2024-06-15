/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20045MutingByPortSchedules Mute wireless unreachable alerts based on switch port schedules
type InlineResponse20045MutingByPortSchedules struct {
	// If true, then wireless unreachable alerts will be muted when caused by a port schedule
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20045MutingByPortSchedules instantiates a new InlineResponse20045MutingByPortSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20045MutingByPortSchedules() *InlineResponse20045MutingByPortSchedules {
	this := InlineResponse20045MutingByPortSchedules{}
	return &this
}

// NewInlineResponse20045MutingByPortSchedulesWithDefaults instantiates a new InlineResponse20045MutingByPortSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20045MutingByPortSchedulesWithDefaults() *InlineResponse20045MutingByPortSchedules {
	this := InlineResponse20045MutingByPortSchedules{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20045MutingByPortSchedules) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20045MutingByPortSchedules) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20045MutingByPortSchedules) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20045MutingByPortSchedules) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20045MutingByPortSchedules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20045MutingByPortSchedules struct {
	value *InlineResponse20045MutingByPortSchedules
	isSet bool
}

func (v NullableInlineResponse20045MutingByPortSchedules) Get() *InlineResponse20045MutingByPortSchedules {
	return v.value
}

func (v *NullableInlineResponse20045MutingByPortSchedules) Set(val *InlineResponse20045MutingByPortSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20045MutingByPortSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20045MutingByPortSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20045MutingByPortSchedules(val *InlineResponse20045MutingByPortSchedules) *NullableInlineResponse20045MutingByPortSchedules {
	return &NullableInlineResponse20045MutingByPortSchedules{value: val, isSet: true}
}

func (v NullableInlineResponse20045MutingByPortSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20045MutingByPortSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


