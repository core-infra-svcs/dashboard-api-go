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

// NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules struct for NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules
type NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol string `json:"protocol"`
	// Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.
	DestCidr string `json:"destCidr"`
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules(policy string, protocol string, destCidr string) *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules{}
	this.Policy = policy
	this.Protocol = protocol
	this.DestCidr = destCidr
	return &this
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetDestCidr(v string) {
	o.DestCidr = v
}

func (o NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if true {
		toSerialize["destCidr"] = o.DestCidr
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules struct {
	value *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) Get() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) Set(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
	return &NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


