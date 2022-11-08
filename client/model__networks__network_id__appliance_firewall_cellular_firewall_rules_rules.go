/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules struct for NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
type NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol string `json:"protocol"`
	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort *string `json:"srcPort,omitempty"`
	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcCidr string `json:"srcCidr"`
	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestCidr string `json:"destCidr"`
	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled *bool `json:"syslogEnabled,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules instantiates a new NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules(policy string, protocol string, srcCidr string, destCidr string) *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	this := NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules{}
	this.Policy = policy
	this.Protocol = protocol
	this.SrcCidr = srcCidr
	this.DestCidr = destCidr
	return &this
}

// NewNetworksNetworkIdApplianceFirewallCellularFirewallRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallCellularFirewallRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	this := NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetProtocol(v string) {
	o.Protocol = v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetSrcCidr returns the SrcCidr field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSrcCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SrcCidr
}

// GetSrcCidrOk returns a tuple with the SrcCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSrcCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SrcCidr, true
}

// SetSrcCidr sets field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetSrcCidr(v string) {
	o.SrcCidr = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetDestCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetDestCidr(v string) {
	o.DestCidr = v
}

// GetSyslogEnabled returns the SyslogEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSyslogEnabled() bool {
	if o == nil || isNil(o.SyslogEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogEnabled
}

// GetSyslogEnabledOk returns a tuple with the SyslogEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) GetSyslogEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogEnabled) {
    return nil, false
	}
	return o.SyslogEnabled, true
}

// HasSyslogEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) HasSyslogEnabled() bool {
	if o != nil && !isNil(o.SyslogEnabled) {
		return true
	}

	return false
}

// SetSyslogEnabled gets a reference to the given bool and assigns it to the SyslogEnabled field.
func (o *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) SetSyslogEnabled(v bool) {
	o.SyslogEnabled = &v
}

func (o NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if true {
		toSerialize["srcCidr"] = o.SrcCidr
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if true {
		toSerialize["destCidr"] = o.DestCidr
	}
	if !isNil(o.SyslogEnabled) {
		toSerialize["syslogEnabled"] = o.SyslogEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules struct {
	value *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) Get() *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) Set(val *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules(val *NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) *NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	return &NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


