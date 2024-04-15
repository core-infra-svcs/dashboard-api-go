/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20044Muting Mute alerts under certain conditions
type InlineResponse20044Muting struct {
	ByPortSchedules *InlineResponse20044MutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// NewInlineResponse20044Muting instantiates a new InlineResponse20044Muting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044Muting() *InlineResponse20044Muting {
	this := InlineResponse20044Muting{}
	return &this
}

// NewInlineResponse20044MutingWithDefaults instantiates a new InlineResponse20044Muting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044MutingWithDefaults() *InlineResponse20044Muting {
	this := InlineResponse20044Muting{}
	return &this
}

// GetByPortSchedules returns the ByPortSchedules field value if set, zero value otherwise.
func (o *InlineResponse20044Muting) GetByPortSchedules() InlineResponse20044MutingByPortSchedules {
	if o == nil || isNil(o.ByPortSchedules) {
		var ret InlineResponse20044MutingByPortSchedules
		return ret
	}
	return *o.ByPortSchedules
}

// GetByPortSchedulesOk returns a tuple with the ByPortSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Muting) GetByPortSchedulesOk() (*InlineResponse20044MutingByPortSchedules, bool) {
	if o == nil || isNil(o.ByPortSchedules) {
    return nil, false
	}
	return o.ByPortSchedules, true
}

// HasByPortSchedules returns a boolean if a field has been set.
func (o *InlineResponse20044Muting) HasByPortSchedules() bool {
	if o != nil && !isNil(o.ByPortSchedules) {
		return true
	}

	return false
}

// SetByPortSchedules gets a reference to the given InlineResponse20044MutingByPortSchedules and assigns it to the ByPortSchedules field.
func (o *InlineResponse20044Muting) SetByPortSchedules(v InlineResponse20044MutingByPortSchedules) {
	o.ByPortSchedules = &v
}

func (o InlineResponse20044Muting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByPortSchedules) {
		toSerialize["byPortSchedules"] = o.ByPortSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044Muting struct {
	value *InlineResponse20044Muting
	isSet bool
}

func (v NullableInlineResponse20044Muting) Get() *InlineResponse20044Muting {
	return v.value
}

func (v *NullableInlineResponse20044Muting) Set(val *InlineResponse20044Muting) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044Muting) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044Muting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044Muting(val *InlineResponse20044Muting) *NullableInlineResponse20044Muting {
	return &NullableInlineResponse20044Muting{value: val, isSet: true}
}

func (v NullableInlineResponse20044Muting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044Muting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


