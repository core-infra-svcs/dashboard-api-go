/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject92 struct for InlineObject92
type InlineObject92 struct {
	// The name for your group policy. Required.
	Name string `json:"name"`
	Scheduling *NetworksNetworkIdGroupPoliciesScheduling `json:"scheduling,omitempty"`
	Bandwidth *NetworksNetworkIdGroupPoliciesBandwidth `json:"bandwidth,omitempty"`
	FirewallAndTrafficShaping *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"`
	ContentFiltering *NetworksNetworkIdGroupPoliciesContentFiltering `json:"contentFiltering,omitempty"`
	// Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
	SplashAuthSettings *string `json:"splashAuthSettings,omitempty"`
	VlanTagging *NetworksNetworkIdGroupPoliciesVlanTagging `json:"vlanTagging,omitempty"`
	BonjourForwarding *NetworksNetworkIdGroupPoliciesBonjourForwarding `json:"bonjourForwarding,omitempty"`
}

// NewInlineObject92 instantiates a new InlineObject92 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject92(name string) *InlineObject92 {
	this := InlineObject92{}
	this.Name = name
	return &this
}

// NewInlineObject92WithDefaults instantiates a new InlineObject92 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject92WithDefaults() *InlineObject92 {
	this := InlineObject92{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject92) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject92) SetName(v string) {
	o.Name = v
}

// GetScheduling returns the Scheduling field value if set, zero value otherwise.
func (o *InlineObject92) GetScheduling() NetworksNetworkIdGroupPoliciesScheduling {
	if o == nil || isNil(o.Scheduling) {
		var ret NetworksNetworkIdGroupPoliciesScheduling
		return ret
	}
	return *o.Scheduling
}

// GetSchedulingOk returns a tuple with the Scheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetSchedulingOk() (*NetworksNetworkIdGroupPoliciesScheduling, bool) {
	if o == nil || isNil(o.Scheduling) {
    return nil, false
	}
	return o.Scheduling, true
}

// HasScheduling returns a boolean if a field has been set.
func (o *InlineObject92) HasScheduling() bool {
	if o != nil && !isNil(o.Scheduling) {
		return true
	}

	return false
}

// SetScheduling gets a reference to the given NetworksNetworkIdGroupPoliciesScheduling and assigns it to the Scheduling field.
func (o *InlineObject92) SetScheduling(v NetworksNetworkIdGroupPoliciesScheduling) {
	o.Scheduling = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *InlineObject92) GetBandwidth() NetworksNetworkIdGroupPoliciesBandwidth {
	if o == nil || isNil(o.Bandwidth) {
		var ret NetworksNetworkIdGroupPoliciesBandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetBandwidthOk() (*NetworksNetworkIdGroupPoliciesBandwidth, bool) {
	if o == nil || isNil(o.Bandwidth) {
    return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *InlineObject92) HasBandwidth() bool {
	if o != nil && !isNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given NetworksNetworkIdGroupPoliciesBandwidth and assigns it to the Bandwidth field.
func (o *InlineObject92) SetBandwidth(v NetworksNetworkIdGroupPoliciesBandwidth) {
	o.Bandwidth = &v
}

// GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field value if set, zero value otherwise.
func (o *InlineObject92) GetFirewallAndTrafficShaping() NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping {
	if o == nil || isNil(o.FirewallAndTrafficShaping) {
		var ret NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping
		return ret
	}
	return *o.FirewallAndTrafficShaping
}

// GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetFirewallAndTrafficShapingOk() (*NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping, bool) {
	if o == nil || isNil(o.FirewallAndTrafficShaping) {
    return nil, false
	}
	return o.FirewallAndTrafficShaping, true
}

// HasFirewallAndTrafficShaping returns a boolean if a field has been set.
func (o *InlineObject92) HasFirewallAndTrafficShaping() bool {
	if o != nil && !isNil(o.FirewallAndTrafficShaping) {
		return true
	}

	return false
}

// SetFirewallAndTrafficShaping gets a reference to the given NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping and assigns it to the FirewallAndTrafficShaping field.
func (o *InlineObject92) SetFirewallAndTrafficShaping(v NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) {
	o.FirewallAndTrafficShaping = &v
}

// GetContentFiltering returns the ContentFiltering field value if set, zero value otherwise.
func (o *InlineObject92) GetContentFiltering() NetworksNetworkIdGroupPoliciesContentFiltering {
	if o == nil || isNil(o.ContentFiltering) {
		var ret NetworksNetworkIdGroupPoliciesContentFiltering
		return ret
	}
	return *o.ContentFiltering
}

// GetContentFilteringOk returns a tuple with the ContentFiltering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetContentFilteringOk() (*NetworksNetworkIdGroupPoliciesContentFiltering, bool) {
	if o == nil || isNil(o.ContentFiltering) {
    return nil, false
	}
	return o.ContentFiltering, true
}

// HasContentFiltering returns a boolean if a field has been set.
func (o *InlineObject92) HasContentFiltering() bool {
	if o != nil && !isNil(o.ContentFiltering) {
		return true
	}

	return false
}

// SetContentFiltering gets a reference to the given NetworksNetworkIdGroupPoliciesContentFiltering and assigns it to the ContentFiltering field.
func (o *InlineObject92) SetContentFiltering(v NetworksNetworkIdGroupPoliciesContentFiltering) {
	o.ContentFiltering = &v
}

// GetSplashAuthSettings returns the SplashAuthSettings field value if set, zero value otherwise.
func (o *InlineObject92) GetSplashAuthSettings() string {
	if o == nil || isNil(o.SplashAuthSettings) {
		var ret string
		return ret
	}
	return *o.SplashAuthSettings
}

// GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetSplashAuthSettingsOk() (*string, bool) {
	if o == nil || isNil(o.SplashAuthSettings) {
    return nil, false
	}
	return o.SplashAuthSettings, true
}

// HasSplashAuthSettings returns a boolean if a field has been set.
func (o *InlineObject92) HasSplashAuthSettings() bool {
	if o != nil && !isNil(o.SplashAuthSettings) {
		return true
	}

	return false
}

// SetSplashAuthSettings gets a reference to the given string and assigns it to the SplashAuthSettings field.
func (o *InlineObject92) SetSplashAuthSettings(v string) {
	o.SplashAuthSettings = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *InlineObject92) GetVlanTagging() NetworksNetworkIdGroupPoliciesVlanTagging {
	if o == nil || isNil(o.VlanTagging) {
		var ret NetworksNetworkIdGroupPoliciesVlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetVlanTaggingOk() (*NetworksNetworkIdGroupPoliciesVlanTagging, bool) {
	if o == nil || isNil(o.VlanTagging) {
    return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *InlineObject92) HasVlanTagging() bool {
	if o != nil && !isNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given NetworksNetworkIdGroupPoliciesVlanTagging and assigns it to the VlanTagging field.
func (o *InlineObject92) SetVlanTagging(v NetworksNetworkIdGroupPoliciesVlanTagging) {
	o.VlanTagging = &v
}

// GetBonjourForwarding returns the BonjourForwarding field value if set, zero value otherwise.
func (o *InlineObject92) GetBonjourForwarding() NetworksNetworkIdGroupPoliciesBonjourForwarding {
	if o == nil || isNil(o.BonjourForwarding) {
		var ret NetworksNetworkIdGroupPoliciesBonjourForwarding
		return ret
	}
	return *o.BonjourForwarding
}

// GetBonjourForwardingOk returns a tuple with the BonjourForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetBonjourForwardingOk() (*NetworksNetworkIdGroupPoliciesBonjourForwarding, bool) {
	if o == nil || isNil(o.BonjourForwarding) {
    return nil, false
	}
	return o.BonjourForwarding, true
}

// HasBonjourForwarding returns a boolean if a field has been set.
func (o *InlineObject92) HasBonjourForwarding() bool {
	if o != nil && !isNil(o.BonjourForwarding) {
		return true
	}

	return false
}

// SetBonjourForwarding gets a reference to the given NetworksNetworkIdGroupPoliciesBonjourForwarding and assigns it to the BonjourForwarding field.
func (o *InlineObject92) SetBonjourForwarding(v NetworksNetworkIdGroupPoliciesBonjourForwarding) {
	o.BonjourForwarding = &v
}

func (o InlineObject92) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Scheduling) {
		toSerialize["scheduling"] = o.Scheduling
	}
	if !isNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !isNil(o.FirewallAndTrafficShaping) {
		toSerialize["firewallAndTrafficShaping"] = o.FirewallAndTrafficShaping
	}
	if !isNil(o.ContentFiltering) {
		toSerialize["contentFiltering"] = o.ContentFiltering
	}
	if !isNil(o.SplashAuthSettings) {
		toSerialize["splashAuthSettings"] = o.SplashAuthSettings
	}
	if !isNil(o.VlanTagging) {
		toSerialize["vlanTagging"] = o.VlanTagging
	}
	if !isNil(o.BonjourForwarding) {
		toSerialize["bonjourForwarding"] = o.BonjourForwarding
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject92 struct {
	value *InlineObject92
	isSet bool
}

func (v NullableInlineObject92) Get() *InlineObject92 {
	return v.value
}

func (v *NullableInlineObject92) Set(val *InlineObject92) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject92) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject92) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject92(val *InlineObject92) *NullableInlineObject92 {
	return &NullableInlineObject92{value: val, isSet: true}
}

func (v NullableInlineObject92) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject92) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


