/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules struct for NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules
type NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules struct {
	// A description of the rule
	Name *string `json:"name,omitempty"`
	// 'tcp' or 'udp'
	Protocol *string `json:"protocol,omitempty"`
	// Destination port of the traffic that is arriving on the WAN
	PublicPort *string `json:"publicPort,omitempty"`
	// Local IP address to which traffic will be forwarded
	LocalIp *string `json:"localIp,omitempty"`
	// Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port
	LocalPort *string `json:"localPort,omitempty"`
	// Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or 'any'
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules {
	this := NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules{}
	return &this
}

// NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRulesWithDefaults() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules {
	this := NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPublicPort returns the PublicPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetPublicPort() string {
	if o == nil || isNil(o.PublicPort) {
		var ret string
		return ret
	}
	return *o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetPublicPortOk() (*string, bool) {
	if o == nil || isNil(o.PublicPort) {
    return nil, false
	}
	return o.PublicPort, true
}

// HasPublicPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasPublicPort() bool {
	if o != nil && !isNil(o.PublicPort) {
		return true
	}

	return false
}

// SetPublicPort gets a reference to the given string and assigns it to the PublicPort field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetPublicPort(v string) {
	o.PublicPort = &v
}

// GetLocalIp returns the LocalIp field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalIp() string {
	if o == nil || isNil(o.LocalIp) {
		var ret string
		return ret
	}
	return *o.LocalIp
}

// GetLocalIpOk returns a tuple with the LocalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalIpOk() (*string, bool) {
	if o == nil || isNil(o.LocalIp) {
    return nil, false
	}
	return o.LocalIp, true
}

// HasLocalIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasLocalIp() bool {
	if o != nil && !isNil(o.LocalIp) {
		return true
	}

	return false
}

// SetLocalIp gets a reference to the given string and assigns it to the LocalIp field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetLocalIp(v string) {
	o.LocalIp = &v
}

// GetLocalPort returns the LocalPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalPort() string {
	if o == nil || isNil(o.LocalPort) {
		var ret string
		return ret
	}
	return *o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalPortOk() (*string, bool) {
	if o == nil || isNil(o.LocalPort) {
    return nil, false
	}
	return o.LocalPort, true
}

// HasLocalPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasLocalPort() bool {
	if o != nil && !isNil(o.LocalPort) {
		return true
	}

	return false
}

// SetLocalPort gets a reference to the given string and assigns it to the LocalPort field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetLocalPort(v string) {
	o.LocalPort = &v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.PublicPort) {
		toSerialize["publicPort"] = o.PublicPort
	}
	if !isNil(o.LocalIp) {
		toSerialize["localIp"] = o.LocalIp
	}
	if !isNil(o.LocalPort) {
		toSerialize["localPort"] = o.LocalPort
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules struct {
	value *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) Get() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) Set(val *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules(val *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules {
	return &NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


