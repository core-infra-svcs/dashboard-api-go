/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication DHCP Enforced Deauthentication enables the disassociation of wireless clients in addition to Mandatory DHCP. This param is only valid on firmware versions >= MX 17.0 where the associated LAN has Mandatory DHCP Enabled 
type NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication struct {
	// Enable DCHP Enforced Deauthentication on the SSID.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication instantiates a new NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication() *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication {
	this := NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication{}
	return &this
}

// NewNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthenticationWithDefaults instantiates a new NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthenticationWithDefaults() *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication {
	this := NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication struct {
	value *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) Get() *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) Set(val *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication(val *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) *NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication {
	return &NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


