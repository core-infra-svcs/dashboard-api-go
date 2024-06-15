/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSingleLanMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type NetworksNetworkIdApplianceSingleLanMandatoryDhcp struct {
	// Enable Mandatory DHCP on LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdApplianceSingleLanMandatoryDhcp instantiates a new NetworksNetworkIdApplianceSingleLanMandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSingleLanMandatoryDhcp() *NetworksNetworkIdApplianceSingleLanMandatoryDhcp {
	this := NetworksNetworkIdApplianceSingleLanMandatoryDhcp{}
	return &this
}

// NewNetworksNetworkIdApplianceSingleLanMandatoryDhcpWithDefaults instantiates a new NetworksNetworkIdApplianceSingleLanMandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSingleLanMandatoryDhcpWithDefaults() *NetworksNetworkIdApplianceSingleLanMandatoryDhcp {
	this := NetworksNetworkIdApplianceSingleLanMandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdApplianceSingleLanMandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp struct {
	value *NetworksNetworkIdApplianceSingleLanMandatoryDhcp
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) Get() *NetworksNetworkIdApplianceSingleLanMandatoryDhcp {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) Set(val *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp(val *NetworksNetworkIdApplianceSingleLanMandatoryDhcp) *NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp {
	return &NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSingleLanMandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


