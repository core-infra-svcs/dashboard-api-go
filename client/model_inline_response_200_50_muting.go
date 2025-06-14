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

// InlineResponse20050Muting Mute alerts under certain conditions
type InlineResponse20050Muting struct {
	ByPortSchedules *InlineResponse20050MutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// NewInlineResponse20050Muting instantiates a new InlineResponse20050Muting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20050Muting() *InlineResponse20050Muting {
	this := InlineResponse20050Muting{}
	return &this
}

// NewInlineResponse20050MutingWithDefaults instantiates a new InlineResponse20050Muting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20050MutingWithDefaults() *InlineResponse20050Muting {
	this := InlineResponse20050Muting{}
	return &this
}

// GetByPortSchedules returns the ByPortSchedules field value if set, zero value otherwise.
func (o *InlineResponse20050Muting) GetByPortSchedules() InlineResponse20050MutingByPortSchedules {
	if o == nil || isNil(o.ByPortSchedules) {
		var ret InlineResponse20050MutingByPortSchedules
		return ret
	}
	return *o.ByPortSchedules
}

// GetByPortSchedulesOk returns a tuple with the ByPortSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20050Muting) GetByPortSchedulesOk() (*InlineResponse20050MutingByPortSchedules, bool) {
	if o == nil || isNil(o.ByPortSchedules) {
    return nil, false
	}
	return o.ByPortSchedules, true
}

// HasByPortSchedules returns a boolean if a field has been set.
func (o *InlineResponse20050Muting) HasByPortSchedules() bool {
	if o != nil && !isNil(o.ByPortSchedules) {
		return true
	}

	return false
}

// SetByPortSchedules gets a reference to the given InlineResponse20050MutingByPortSchedules and assigns it to the ByPortSchedules field.
func (o *InlineResponse20050Muting) SetByPortSchedules(v InlineResponse20050MutingByPortSchedules) {
	o.ByPortSchedules = &v
}

func (o InlineResponse20050Muting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByPortSchedules) {
		toSerialize["byPortSchedules"] = o.ByPortSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20050Muting struct {
	value *InlineResponse20050Muting
	isSet bool
}

func (v NullableInlineResponse20050Muting) Get() *InlineResponse20050Muting {
	return v.value
}

func (v *NullableInlineResponse20050Muting) Set(val *InlineResponse20050Muting) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20050Muting) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20050Muting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20050Muting(val *InlineResponse20050Muting) *NullableInlineResponse20050Muting {
	return &NullableInlineResponse20050Muting{value: val, isSet: true}
}

func (v NullableInlineResponse20050Muting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20050Muting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


