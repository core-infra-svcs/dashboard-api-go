/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping     The firewall and traffic shaping rules and settings for your policy. 
type NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping struct {
	// How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	TrafficShapingRules []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules `json:"trafficShapingRules,omitempty"`
	// An ordered array of the L3 firewall rules
	L3FirewallRules []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules `json:"l3FirewallRules,omitempty"`
	// An ordered array of L7 firewall rules
	L7FirewallRules []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules `json:"l7FirewallRules,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetSettings(v string) {
	o.Settings = &v
}

// GetTrafficShapingRules returns the TrafficShapingRules field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetTrafficShapingRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
	if o == nil || isNil(o.TrafficShapingRules) {
		var ret []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules
		return ret
	}
	return o.TrafficShapingRules
}

// GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetTrafficShapingRulesOk() ([]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules, bool) {
	if o == nil || isNil(o.TrafficShapingRules) {
    return nil, false
	}
	return o.TrafficShapingRules, true
}

// HasTrafficShapingRules returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasTrafficShapingRules() bool {
	if o != nil && !isNil(o.TrafficShapingRules) {
		return true
	}

	return false
}

// SetTrafficShapingRules gets a reference to the given []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules and assigns it to the TrafficShapingRules field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetTrafficShapingRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) {
	o.TrafficShapingRules = v
}

// GetL3FirewallRules returns the L3FirewallRules field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL3FirewallRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
	if o == nil || isNil(o.L3FirewallRules) {
		var ret []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules
		return ret
	}
	return o.L3FirewallRules
}

// GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL3FirewallRulesOk() ([]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules, bool) {
	if o == nil || isNil(o.L3FirewallRules) {
    return nil, false
	}
	return o.L3FirewallRules, true
}

// HasL3FirewallRules returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasL3FirewallRules() bool {
	if o != nil && !isNil(o.L3FirewallRules) {
		return true
	}

	return false
}

// SetL3FirewallRules gets a reference to the given []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules and assigns it to the L3FirewallRules field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetL3FirewallRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) {
	o.L3FirewallRules = v
}

// GetL7FirewallRules returns the L7FirewallRules field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL7FirewallRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
	if o == nil || isNil(o.L7FirewallRules) {
		var ret []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules
		return ret
	}
	return o.L7FirewallRules
}

// GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL7FirewallRulesOk() ([]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules, bool) {
	if o == nil || isNil(o.L7FirewallRules) {
    return nil, false
	}
	return o.L7FirewallRules, true
}

// HasL7FirewallRules returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasL7FirewallRules() bool {
	if o != nil && !isNil(o.L7FirewallRules) {
		return true
	}

	return false
}

// SetL7FirewallRules gets a reference to the given []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules and assigns it to the L7FirewallRules field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetL7FirewallRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) {
	o.L7FirewallRules = v
}

func (o NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.TrafficShapingRules) {
		toSerialize["trafficShapingRules"] = o.TrafficShapingRules
	}
	if !isNil(o.L3FirewallRules) {
		toSerialize["l3FirewallRules"] = o.L3FirewallRules
	}
	if !isNil(o.L7FirewallRules) {
		toSerialize["l7FirewallRules"] = o.L7FirewallRules
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping struct {
	value *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) Get() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) Set(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping {
	return &NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


