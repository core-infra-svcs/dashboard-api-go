/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules struct for NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules
type NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules struct {
	// The IP address that will be used to access the internal resource from the WAN
	PublicIp string `json:"publicIp"`
	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	Uplink string `json:"uplink"`
	// An array of associated forwarding rules
	PortRules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules `json:"portRules"`
}

// NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules(publicIp string, uplink string, portRules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	this := NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules{}
	this.PublicIp = publicIp
	this.Uplink = uplink
	this.PortRules = portRules
	return &this
}

// NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	this := NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules{}
	return &this
}

// GetPublicIp returns the PublicIp field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPublicIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPublicIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicIp, true
}

// SetPublicIp sets field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetPublicIp(v string) {
	o.PublicIp = v
}

// GetUplink returns the Uplink field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetUplinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uplink, true
}

// SetUplink sets field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetUplink(v string) {
	o.Uplink = v
}

// GetPortRules returns the PortRules field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPortRules() []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules
		return ret
	}

	return o.PortRules
}

// GetPortRulesOk returns a tuple with the PortRules field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPortRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PortRules, true
}

// SetPortRules sets field value
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetPortRules(v []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) {
	o.PortRules = v
}

func (o NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publicIp"] = o.PublicIp
	}
	if true {
		toSerialize["uplink"] = o.Uplink
	}
	if true {
		toSerialize["portRules"] = o.PortRules
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules struct {
	value *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) Get() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) Set(val *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules(val *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	return &NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


