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

// NetworksNetworkIdWirelessEthernetPortsProfilesPorts struct for NetworksNetworkIdWirelessEthernetPortsProfilesPorts
type NetworksNetworkIdWirelessEthernetPortsProfilesPorts struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Number
	Number *int32 `json:"number,omitempty"`
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Ssid number
	Ssid *int32 `json:"ssid,omitempty"`
	// PSK Group number
	PskGroupId *string `json:"pskGroupId,omitempty"`
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessEthernetPortsProfilesPorts() *NetworksNetworkIdWirelessEthernetPortsProfilesPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesPorts{}
	return &this
}

// NewNetworksNetworkIdWirelessEthernetPortsProfilesPortsWithDefaults instantiates a new NetworksNetworkIdWirelessEthernetPortsProfilesPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessEthernetPortsProfilesPortsWithDefaults() *NetworksNetworkIdWirelessEthernetPortsProfilesPorts {
	this := NetworksNetworkIdWirelessEthernetPortsProfilesPorts{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) SetNumber(v int32) {
	o.Number = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetSsid() int32 {
	if o == nil || isNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetSsidOk() (*int32, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) SetSsid(v int32) {
	o.Ssid = &v
}

// GetPskGroupId returns the PskGroupId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetPskGroupId() string {
	if o == nil || isNil(o.PskGroupId) {
		var ret string
		return ret
	}
	return *o.PskGroupId
}

// GetPskGroupIdOk returns a tuple with the PskGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) GetPskGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.PskGroupId) {
    return nil, false
	}
	return o.PskGroupId, true
}

// HasPskGroupId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) HasPskGroupId() bool {
	if o != nil && !isNil(o.PskGroupId) {
		return true
	}

	return false
}

// SetPskGroupId gets a reference to the given string and assigns it to the PskGroupId field.
func (o *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) SetPskGroupId(v string) {
	o.PskGroupId = &v
}

func (o NetworksNetworkIdWirelessEthernetPortsProfilesPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
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

type NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts struct {
	value *NetworksNetworkIdWirelessEthernetPortsProfilesPorts
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) Get() *NetworksNetworkIdWirelessEthernetPortsProfilesPorts {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) Set(val *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts(val *NetworksNetworkIdWirelessEthernetPortsProfilesPorts) *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts {
	return &NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessEthernetPortsProfilesPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


