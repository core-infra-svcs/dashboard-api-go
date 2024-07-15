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

// NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 struct for NetworksNetworkIdWirelessEthernetPortsProfilesPorts1
type NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 struct {
	// AP port name
	Name string `json:"name"`
	// AP port enabled
	Enabled *bool `json:"enabled,omitempty"`
	// AP port ssid number
	Ssid *int32 `json:"ssid,omitempty"`
	// AP port PSK Group ID
	PskGroupId *string `json:"pskGroupId,omitempty"`
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts1 instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts1(name string) *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesPorts1{}
	this.Name = name
	return &this
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts1WithDefaults instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts1WithDefaults() *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesPorts1{}
	return &this
}

// GetName returns the Name field value
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetSsid() int32 {
	if o == nil || isNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetSsidOk() (*int32, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) SetSsid(v int32) {
	o.Ssid = &v
}

// GetPskGroupId returns the PskGroupId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetPskGroupId() string {
	if o == nil || isNil(o.PskGroupId) {
		var ret string
		return ret
	}
	return *o.PskGroupId
}

// GetPskGroupIdOk returns a tuple with the PskGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) GetPskGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.PskGroupId) {
    return nil, false
	}
	return o.PskGroupId, true
}

// HasPskGroupId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) HasPskGroupId() bool {
	if o != nil && !isNil(o.PskGroupId) {
		return true
	}

	return false
}

// SetPskGroupId gets a reference to the given string and assigns it to the PskGroupId field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) SetPskGroupId(v string) {
	o.PskGroupId = &v
}

func (o NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1 struct {
	value *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) Get() *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) Set(val *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1(val *NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1 {
	return &NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


