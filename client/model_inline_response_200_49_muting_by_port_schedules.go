/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20049MutingByPortSchedules Mute wireless unreachable alerts based on switch port schedules
type InlineResponse20049MutingByPortSchedules struct {
	// If true, then wireless unreachable alerts will be muted when caused by a port schedule
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20049MutingByPortSchedules instantiates a new InlineResponse20049MutingByPortSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049MutingByPortSchedules() *InlineResponse20049MutingByPortSchedules {
	this := InlineResponse20049MutingByPortSchedules{}
	return &this
}

// NewInlineResponse20049MutingByPortSchedulesWithDefaults instantiates a new InlineResponse20049MutingByPortSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049MutingByPortSchedulesWithDefaults() *InlineResponse20049MutingByPortSchedules {
	this := InlineResponse20049MutingByPortSchedules{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20049MutingByPortSchedules) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049MutingByPortSchedules) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20049MutingByPortSchedules) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20049MutingByPortSchedules) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20049MutingByPortSchedules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049MutingByPortSchedules struct {
	value *InlineResponse20049MutingByPortSchedules
	isSet bool
}

func (v NullableInlineResponse20049MutingByPortSchedules) Get() *InlineResponse20049MutingByPortSchedules {
	return v.value
}

func (v *NullableInlineResponse20049MutingByPortSchedules) Set(val *InlineResponse20049MutingByPortSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049MutingByPortSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049MutingByPortSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049MutingByPortSchedules(val *InlineResponse20049MutingByPortSchedules) *NullableInlineResponse20049MutingByPortSchedules {
	return &NullableInlineResponse20049MutingByPortSchedules{value: val, isSet: true}
}

func (v NullableInlineResponse20049MutingByPortSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049MutingByPortSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


