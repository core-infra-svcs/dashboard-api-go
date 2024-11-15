/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules struct for NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules
type NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules struct {
	// Protocol for this split tunnel rule.
	Protocol *string `json:"protocol,omitempty"`
	// Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or 'any'.
	DestCidr string `json:"destCidr"`
	// Destination port for this split tunnel rule, (integer in the range 1-65535), or 'any'.
	DestPort *string `json:"destPort,omitempty"`
	// Traffic policy specified for this split tunnel rule, 'allow' or 'deny'.
	Policy string `json:"policy"`
	// Description for this split tunnel rule (optional).
	Comment *string `json:"comment,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules(destCidr string, policy string) *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules {
	this := NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules{}
	this.DestCidr = destCidr
	this.Policy = policy
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRulesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRulesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules {
	this := NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDestCidr returns the DestCidr field value
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetDestCidr(v string) {
	o.DestCidr = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetDestPort(v string) {
	o.DestPort = &v
}

// GetPolicy returns the Policy field value
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetPolicy(v string) {
	o.Policy = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetComment(v string) {
	o.Comment = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["destCidr"] = o.DestCidr
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules struct {
	value *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) Get() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) Set(val *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules(val *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules {
	return &NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


