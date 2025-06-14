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

// NetworksNetworkIdApplianceFirewallL7FirewallRulesRules struct for NetworksNetworkIdApplianceFirewallL7FirewallRulesRules
type NetworksNetworkIdApplianceFirewallL7FirewallRulesRules struct {
	// 'Deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Type *string `json:"type,omitempty"`
	// The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.
	Value *string `json:"value,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRules instantiates a new NetworksNetworkIdApplianceFirewallL7FirewallRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRules() *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	this := NetworksNetworkIdApplianceFirewallL7FirewallRulesRules{}
	return &this
}

// NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallL7FirewallRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	this := NetworksNetworkIdApplianceFirewallL7FirewallRulesRules{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetValue(v string) {
	o.Value = &v
}

func (o NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules struct {
	value *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) Get() *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) Set(val *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules(val *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) *NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	return &NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallL7FirewallRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


