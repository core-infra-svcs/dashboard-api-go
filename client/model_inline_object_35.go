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

// InlineObject35 struct for InlineObject35
type InlineObject35 struct {
	// configured alternate management interface addresses
	Addresses []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses `json:"addresses,omitempty"`
}

// NewInlineObject35 instantiates a new InlineObject35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject35() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// NewInlineObject35WithDefaults instantiates a new InlineObject35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject35WithDefaults() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineObject35) GetAddresses() []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses {
	if o == nil || isNil(o.Addresses) {
		var ret []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetAddressesOk() ([]DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineObject35) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses and assigns it to the Addresses field.
func (o *InlineObject35) SetAddresses(v []DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) {
	o.Addresses = v
}

func (o InlineObject35) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject35 struct {
	value *InlineObject35
	isSet bool
}

func (v NullableInlineObject35) Get() *InlineObject35 {
	return v.value
}

func (v *NullableInlineObject35) Set(val *InlineObject35) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject35) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject35(val *InlineObject35) *NullableInlineObject35 {
	return &NullableInlineObject35{value: val, isSet: true}
}

func (v NullableInlineObject35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


