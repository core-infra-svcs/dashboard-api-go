/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject25 struct for InlineObject25
type InlineObject25 struct {
	// configured alternate management interface addresses
	Addresses []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses `json:"addresses,omitempty"`
}

// NewInlineObject25 instantiates a new InlineObject25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject25() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// NewInlineObject25WithDefaults instantiates a new InlineObject25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject25WithDefaults() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineObject25) GetAddresses() []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses {
	if o == nil || isNil(o.Addresses) {
		var ret []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject25) GetAddressesOk() ([]DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineObject25) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses and assigns it to the Addresses field.
func (o *InlineObject25) SetAddresses(v []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) {
	o.Addresses = v
}

func (o InlineObject25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject25 struct {
	value *InlineObject25
	isSet bool
}

func (v NullableInlineObject25) Get() *InlineObject25 {
	return v.value
}

func (v *NullableInlineObject25) Set(val *InlineObject25) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject25) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject25(val *InlineObject25) *NullableInlineObject25 {
	return &NullableInlineObject25{value: val, isSet: true}
}

func (v NullableInlineObject25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


