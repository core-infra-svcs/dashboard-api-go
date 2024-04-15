/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts struct for NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts
type NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts struct {
	// AP port name
	Name string `json:"name"`
	// AP port enabled
	Enabled *bool `json:"enabled,omitempty"`
	// AP port ssid number
	Ssid *int32 `json:"ssid,omitempty"`
	// AP port PSK Group number
	PskGroupId *string `json:"pskGroupId,omitempty"`
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts(name string) *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts{}
	this.Name = name
	return &this
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPortsWithDefaults instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPortsWithDefaults() *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts{}
	return &this
}

// GetName returns the Name field value
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetSsid() int32 {
	if o == nil || isNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetSsidOk() (*int32, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) SetSsid(v int32) {
	o.Ssid = &v
}

// GetPskGroupId returns the PskGroupId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetPskGroupId() string {
	if o == nil || isNil(o.PskGroupId) {
		var ret string
		return ret
	}
	return *o.PskGroupId
}

// GetPskGroupIdOk returns a tuple with the PskGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) GetPskGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.PskGroupId) {
    return nil, false
	}
	return o.PskGroupId, true
}

// HasPskGroupId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) HasPskGroupId() bool {
	if o != nil && !isNil(o.PskGroupId) {
		return true
	}

	return false
}

// SetPskGroupId gets a reference to the given string and assigns it to the PskGroupId field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) SetPskGroupId(v string) {
	o.PskGroupId = &v
}

func (o NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.PskGroupId) {
		toSerialize["pskGroupId"] = o.PskGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts struct {
	value *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) Get() *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) Set(val *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts(val *NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) *NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts {
	return &NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


