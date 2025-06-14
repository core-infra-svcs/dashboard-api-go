/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules struct for NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules
type NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules struct {
	// The policy applied to matching traffic. Must be 'deny'.
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
	Type *string `json:"type,omitempty"`
	// The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.
	Value *string `json:"value,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetValue(v string) {
	o.Value = &v
}

func (o NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules struct {
	value *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) Get() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) Set(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
	return &NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


