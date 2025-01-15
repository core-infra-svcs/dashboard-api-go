/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers The DNS servers settings for this address.
type DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers struct {
	// Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
	Addresses []string `json:"addresses,omitempty"`
}

// NewDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers instantiates a new DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers() *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers {
	this := DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers{}
	return &this
}

// NewDevicesSerialWirelessAlternateManagementInterfaceIpv6NameserversWithDefaults instantiates a new DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialWirelessAlternateManagementInterfaceIpv6NameserversWithDefaults() *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers {
	this := DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) GetAddresses() []string {
	if o == nil || isNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) GetAddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) SetAddresses(v []string) {
	o.Addresses = v
}

func (o DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers struct {
	value *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers
	isSet bool
}

func (v NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) Get() *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers {
	return v.value
}

func (v *NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) Set(val *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers(val *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) *NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers {
	return &NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers{value: val, isSet: true}
}

func (v NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


