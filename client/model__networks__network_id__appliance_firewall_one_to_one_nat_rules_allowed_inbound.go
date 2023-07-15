/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound struct for NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound
type NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound struct {
	// Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	// An array of ports or port ranges that will be forwarded to the host on the LAN
	DestinationPorts []string `json:"destinationPorts,omitempty"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or 'any'
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound instantiates a new NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound() *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound {
	this := NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound{}
	return &this
}

// NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInboundWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInboundWithDefaults() *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound {
	this := NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDestinationPorts returns the DestinationPorts field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetDestinationPorts() []string {
	if o == nil || isNil(o.DestinationPorts) {
		var ret []string
		return ret
	}
	return o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetDestinationPortsOk() ([]string, bool) {
	if o == nil || isNil(o.DestinationPorts) {
    return nil, false
	}
	return o.DestinationPorts, true
}

// HasDestinationPorts returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) HasDestinationPorts() bool {
	if o != nil && !isNil(o.DestinationPorts) {
		return true
	}

	return false
}

// SetDestinationPorts gets a reference to the given []string and assigns it to the DestinationPorts field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) SetDestinationPorts(v []string) {
	o.DestinationPorts = v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.DestinationPorts) {
		toSerialize["destinationPorts"] = o.DestinationPorts
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound struct {
	value *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) Get() *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) Set(val *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound(val *NetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound {
	return &NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToOneNatRulesAllowedInbound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


