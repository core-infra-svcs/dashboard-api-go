/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20038 struct for InlineResponse20038
type InlineResponse20038 struct {
	// configured alternate management interface addresses
	Addresses []InlineResponse20038Addresses `json:"addresses,omitempty"`
}

// NewInlineResponse20038 instantiates a new InlineResponse20038 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038WithDefaults() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineResponse20038) GetAddresses() []InlineResponse20038Addresses {
	if o == nil || isNil(o.Addresses) {
		var ret []InlineResponse20038Addresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetAddressesOk() ([]InlineResponse20038Addresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineResponse20038) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []InlineResponse20038Addresses and assigns it to the Addresses field.
func (o *InlineResponse20038) SetAddresses(v []InlineResponse20038Addresses) {
	o.Addresses = v
}

func (o InlineResponse20038) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038 struct {
	value *InlineResponse20038
	isSet bool
}

func (v NullableInlineResponse20038) Get() *InlineResponse20038 {
	return v.value
}

func (v *NullableInlineResponse20038) Set(val *InlineResponse20038) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038(val *InlineResponse20038) *NullableInlineResponse20038 {
	return &NullableInlineResponse20038{value: val, isSet: true}
}

func (v NullableInlineResponse20038) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


