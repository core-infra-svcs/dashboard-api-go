/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules struct for NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules
type NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules struct {
	// A descriptive name for the rule
	Name *string `json:"name,omitempty"`
	// The IP address that will be used to access the internal resource from the WAN
	PublicIp *string `json:"publicIp,omitempty"`
	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LanIp string `json:"lanIp"`
	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	Uplink *string `json:"uplink,omitempty"`
	// The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource
	AllowedInbound []NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound `json:"allowedInbound,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules instantiates a new NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules(lanIp string) *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	this := NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules{}
	this.LanIp = lanIp
	return &this
}

// NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	this := NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) SetName(v string) {
	o.Name = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetLanIp returns the LanIp field value
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetLanIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetLanIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LanIp, true
}

// SetLanIp sets field value
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) SetLanIp(v string) {
	o.LanIp = v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetUplink() string {
	if o == nil || isNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetUplinkOk() (*string, bool) {
	if o == nil || isNil(o.Uplink) {
    return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) HasUplink() bool {
	if o != nil && !isNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) SetUplink(v string) {
	o.Uplink = &v
}

// GetAllowedInbound returns the AllowedInbound field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetAllowedInbound() []NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound {
	if o == nil || isNil(o.AllowedInbound) {
		var ret []NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound
		return ret
	}
	return o.AllowedInbound
}

// GetAllowedInboundOk returns a tuple with the AllowedInbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) GetAllowedInboundOk() ([]NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound, bool) {
	if o == nil || isNil(o.AllowedInbound) {
    return nil, false
	}
	return o.AllowedInbound, true
}

// HasAllowedInbound returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) HasAllowedInbound() bool {
	if o != nil && !isNil(o.AllowedInbound) {
		return true
	}

	return false
}

// SetAllowedInbound gets a reference to the given []NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound and assigns it to the AllowedInbound field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) SetAllowedInbound(v []NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) {
	o.AllowedInbound = v
}

func (o NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if true {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.Uplink) {
		toSerialize["uplink"] = o.Uplink
	}
	if !isNil(o.AllowedInbound) {
		toSerialize["allowedInbound"] = o.AllowedInbound
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules struct {
	value *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) Get() *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) Set(val *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules(val *NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	return &NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


