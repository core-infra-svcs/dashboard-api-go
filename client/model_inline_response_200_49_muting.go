/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20049Muting Mute alerts under certain conditions
type InlineResponse20049Muting struct {
	ByPortSchedules *InlineResponse20049MutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// NewInlineResponse20049Muting instantiates a new InlineResponse20049Muting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049Muting() *InlineResponse20049Muting {
	this := InlineResponse20049Muting{}
	return &this
}

// NewInlineResponse20049MutingWithDefaults instantiates a new InlineResponse20049Muting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049MutingWithDefaults() *InlineResponse20049Muting {
	this := InlineResponse20049Muting{}
	return &this
}

// GetByPortSchedules returns the ByPortSchedules field value if set, zero value otherwise.
func (o *InlineResponse20049Muting) GetByPortSchedules() InlineResponse20049MutingByPortSchedules {
	if o == nil || isNil(o.ByPortSchedules) {
		var ret InlineResponse20049MutingByPortSchedules
		return ret
	}
	return *o.ByPortSchedules
}

// GetByPortSchedulesOk returns a tuple with the ByPortSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049Muting) GetByPortSchedulesOk() (*InlineResponse20049MutingByPortSchedules, bool) {
	if o == nil || isNil(o.ByPortSchedules) {
    return nil, false
	}
	return o.ByPortSchedules, true
}

// HasByPortSchedules returns a boolean if a field has been set.
func (o *InlineResponse20049Muting) HasByPortSchedules() bool {
	if o != nil && !isNil(o.ByPortSchedules) {
		return true
	}

	return false
}

// SetByPortSchedules gets a reference to the given InlineResponse20049MutingByPortSchedules and assigns it to the ByPortSchedules field.
func (o *InlineResponse20049Muting) SetByPortSchedules(v InlineResponse20049MutingByPortSchedules) {
	o.ByPortSchedules = &v
}

func (o InlineResponse20049Muting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByPortSchedules) {
		toSerialize["byPortSchedules"] = o.ByPortSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049Muting struct {
	value *InlineResponse20049Muting
	isSet bool
}

func (v NullableInlineResponse20049Muting) Get() *InlineResponse20049Muting {
	return v.value
}

func (v *NullableInlineResponse20049Muting) Set(val *InlineResponse20049Muting) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049Muting) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049Muting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049Muting(val *InlineResponse20049Muting) *NullableInlineResponse20049Muting {
	return &NullableInlineResponse20049Muting{value: val, isSet: true}
}

func (v NullableInlineResponse20049Muting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049Muting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


