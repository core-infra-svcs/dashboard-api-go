/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200260 struct for InlineResponse200260
type InlineResponse200260 struct {
	// Serials of the devices that were released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse200260 instantiates a new InlineResponse200260 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200260() *InlineResponse200260 {
	this := InlineResponse200260{}
	return &this
}

// NewInlineResponse200260WithDefaults instantiates a new InlineResponse200260 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200260WithDefaults() *InlineResponse200260 {
	this := InlineResponse200260{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200260) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200260) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200260) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200260) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200260 struct {
	value *InlineResponse200260
	isSet bool
}

func (v NullableInlineResponse200260) Get() *InlineResponse200260 {
	return v.value
}

func (v *NullableInlineResponse200260) Set(val *InlineResponse200260) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200260) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200260) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200260(val *InlineResponse200260) *NullableInlineResponse200260 {
	return &NullableInlineResponse200260{value: val, isSet: true}
}

func (v NullableInlineResponse200260) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200260) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


