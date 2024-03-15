/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessControlListsRules struct for NetworksNetworkIdSwitchAccessControlListsRules
type NetworksNetworkIdSwitchAccessControlListsRules struct {
	// Description of the rule (optional).
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule.
	Policy string `json:"policy"`
	// IP address version (must be 'any', 'ipv4' or 'ipv6'). Applicable only if network supports IPv6. Default value is 'ipv4'.
	IpVersion *string `json:"ipVersion,omitempty"`
	// The type of protocol (must be 'tcp', 'udp', or 'any').
	Protocol string `json:"protocol"`
	// Source IP address (in IP or CIDR notation) or 'any'.
	SrcCidr string `json:"srcCidr"`
	// Source port. Must be in the range  of 1-65535 or 'any'. Default is 'any'.
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination IP address (in IP or CIDR notation) or 'any'.
	DstCidr string `json:"dstCidr"`
	// Destination port. Must be in the range of 1-65535 or 'any'. Default is 'any'.
	DstPort *string `json:"dstPort,omitempty"`
	// Incoming traffic VLAN. Must be in the range of 1-4095 or 'any'. Default is 'any'.
	Vlan *string `json:"vlan,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessControlListsRules instantiates a new NetworksNetworkIdSwitchAccessControlListsRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessControlListsRules(policy string, protocol string, srcCidr string, dstCidr string) *NetworksNetworkIdSwitchAccessControlListsRules {
	this := NetworksNetworkIdSwitchAccessControlListsRules{}
	this.Policy = policy
	this.Protocol = protocol
	this.SrcCidr = srcCidr
	this.DstCidr = dstCidr
	return &this
}

// NewNetworksNetworkIdSwitchAccessControlListsRulesWithDefaults instantiates a new NetworksNetworkIdSwitchAccessControlListsRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessControlListsRulesWithDefaults() *NetworksNetworkIdSwitchAccessControlListsRules {
	this := NetworksNetworkIdSwitchAccessControlListsRules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetPolicy(v string) {
	o.Policy = v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetIpVersion() string {
	if o == nil || isNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetIpVersionOk() (*string, bool) {
	if o == nil || isNil(o.IpVersion) {
    return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) HasIpVersion() bool {
	if o != nil && !isNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetProtocol returns the Protocol field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetProtocol(v string) {
	o.Protocol = v
}

// GetSrcCidr returns the SrcCidr field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetSrcCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SrcCidr
}

// GetSrcCidrOk returns a tuple with the SrcCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetSrcCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SrcCidr, true
}

// SetSrcCidr sets field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetSrcCidr(v string) {
	o.SrcCidr = v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstCidr returns the DstCidr field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetDstCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DstCidr
}

// GetDstCidrOk returns a tuple with the DstCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetDstCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DstCidr, true
}

// SetDstCidr sets field value
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetDstCidr(v string) {
	o.DstCidr = v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetDstPort() string {
	if o == nil || isNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetDstPortOk() (*string, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetDstPort(v string) {
	o.DstPort = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetVlan() string {
	if o == nil || isNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) GetVlanOk() (*string, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *NetworksNetworkIdSwitchAccessControlListsRules) SetVlan(v string) {
	o.Vlan = &v
}

func (o NetworksNetworkIdSwitchAccessControlListsRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["srcCidr"] = o.SrcCidr
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if true {
		toSerialize["dstCidr"] = o.DstCidr
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessControlListsRules struct {
	value *NetworksNetworkIdSwitchAccessControlListsRules
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessControlListsRules) Get() *NetworksNetworkIdSwitchAccessControlListsRules {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessControlListsRules) Set(val *NetworksNetworkIdSwitchAccessControlListsRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessControlListsRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessControlListsRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessControlListsRules(val *NetworksNetworkIdSwitchAccessControlListsRules) *NullableNetworksNetworkIdSwitchAccessControlListsRules {
	return &NullableNetworksNetworkIdSwitchAccessControlListsRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessControlListsRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessControlListsRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


