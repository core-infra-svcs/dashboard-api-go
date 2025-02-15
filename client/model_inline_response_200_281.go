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

// InlineResponse200281 struct for InlineResponse200281
type InlineResponse200281 struct {
	// Serials of the devices that were released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse200281 instantiates a new InlineResponse200281 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200281() *InlineResponse200281 {
	this := InlineResponse200281{}
	return &this
}

// NewInlineResponse200281WithDefaults instantiates a new InlineResponse200281 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200281WithDefaults() *InlineResponse200281 {
	this := InlineResponse200281{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200281) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200281) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200281) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200281) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200281) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200281 struct {
	value *InlineResponse200281
	isSet bool
}

func (v NullableInlineResponse200281) Get() *InlineResponse200281 {
	return v.value
}

func (v *NullableInlineResponse200281) Set(val *InlineResponse200281) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200281) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200281) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200281(val *InlineResponse200281) *NullableInlineResponse200281 {
	return &NullableInlineResponse200281{value: val, isSet: true}
}

func (v NullableInlineResponse200281) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200281) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


