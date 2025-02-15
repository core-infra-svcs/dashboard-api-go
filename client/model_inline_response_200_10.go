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

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	Interfaces *InlineResponse20010Interfaces `json:"interfaces,omitempty"`
}

// NewInlineResponse20010 instantiates a new InlineResponse20010 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20010() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20010WithDefaults() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *InlineResponse20010) GetInterfaces() InlineResponse20010Interfaces {
	if o == nil || isNil(o.Interfaces) {
		var ret InlineResponse20010Interfaces
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetInterfacesOk() (*InlineResponse20010Interfaces, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *InlineResponse20010) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given InlineResponse20010Interfaces and assigns it to the Interfaces field.
func (o *InlineResponse20010) SetInterfaces(v InlineResponse20010Interfaces) {
	o.Interfaces = &v
}

func (o InlineResponse20010) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20010 struct {
	value *InlineResponse20010
	isSet bool
}

func (v NullableInlineResponse20010) Get() *InlineResponse20010 {
	return v.value
}

func (v *NullableInlineResponse20010) Set(val *InlineResponse20010) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20010) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20010) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20010(val *InlineResponse20010) *NullableInlineResponse20010 {
	return &NullableInlineResponse20010{value: val, isSet: true}
}

func (v NullableInlineResponse20010) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20010) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


