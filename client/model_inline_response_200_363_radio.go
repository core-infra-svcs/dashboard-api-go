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

// InlineResponse200363Radio Add scanning API Radio
type InlineResponse200363Radio struct {
	// Radio Type. Valid types are: Wi-Fi, Bluetooth
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse200363Radio instantiates a new InlineResponse200363Radio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200363Radio() *InlineResponse200363Radio {
	this := InlineResponse200363Radio{}
	return &this
}

// NewInlineResponse200363RadioWithDefaults instantiates a new InlineResponse200363Radio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200363RadioWithDefaults() *InlineResponse200363Radio {
	this := InlineResponse200363Radio{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200363Radio) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200363Radio) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200363Radio) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200363Radio) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse200363Radio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200363Radio struct {
	value *InlineResponse200363Radio
	isSet bool
}

func (v NullableInlineResponse200363Radio) Get() *InlineResponse200363Radio {
	return v.value
}

func (v *NullableInlineResponse200363Radio) Set(val *InlineResponse200363Radio) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200363Radio) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200363Radio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200363Radio(val *InlineResponse200363Radio) *NullableInlineResponse200363Radio {
	return &NullableInlineResponse200363Radio{value: val, isSet: true}
}

func (v NullableInlineResponse200363Radio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200363Radio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


