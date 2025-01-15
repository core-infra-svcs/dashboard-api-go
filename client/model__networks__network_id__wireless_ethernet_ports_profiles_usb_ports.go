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

// NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts struct for NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts
type NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Ssid number
	Ssid *int32 `json:"ssid,omitempty"`
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts() *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts{}
	return &this
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesUsbPortsWithDefaults instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessEthernetPortsProfilesUsbPortsWithDefaults() *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetSsid() int32 {
	if o == nil || isNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) GetSsidOk() (*int32, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) SetSsid(v int32) {
	o.Ssid = &v
}

func (o NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts struct {
	value *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) Get() *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) Set(val *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts(val *NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) *NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts {
	return &NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


