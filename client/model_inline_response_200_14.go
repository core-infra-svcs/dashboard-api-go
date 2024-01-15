/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20014 struct for InlineResponse20014
type InlineResponse20014 struct {
	// configured alternate management interface addresses
	Addresses []InlineResponse20014Addresses `json:"addresses,omitempty"`
}

// NewInlineResponse20014 instantiates a new InlineResponse20014 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014WithDefaults() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineResponse20014) GetAddresses() []InlineResponse20014Addresses {
	if o == nil || isNil(o.Addresses) {
		var ret []InlineResponse20014Addresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetAddressesOk() ([]InlineResponse20014Addresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineResponse20014) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []InlineResponse20014Addresses and assigns it to the Addresses field.
func (o *InlineResponse20014) SetAddresses(v []InlineResponse20014Addresses) {
	o.Addresses = v
}

func (o InlineResponse20014) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014 struct {
	value *InlineResponse20014
	isSet bool
}

func (v NullableInlineResponse20014) Get() *InlineResponse20014 {
	return v.value
}

func (v *NullableInlineResponse20014) Set(val *InlineResponse20014) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014(val *InlineResponse20014) *NullableInlineResponse20014 {
	return &NullableInlineResponse20014{value: val, isSet: true}
}

func (v NullableInlineResponse20014) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


