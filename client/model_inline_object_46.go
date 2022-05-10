/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject46 struct for InlineObject46
type InlineObject46 struct {
	// Toggle for enabling or disabling active-active AutoVPN
	ActiveActiveAutoVpnEnabled *bool `json:"activeActiveAutoVpnEnabled,omitempty"`
	// The default uplink. Must be one of: 'wan1' or 'wan2'
	DefaultUplink *string `json:"defaultUplink,omitempty"`
	// Toggle for enabling or disabling load balancing
	LoadBalancingEnabled *bool `json:"loadBalancingEnabled,omitempty"`
	// Array of uplink preference rules for WAN traffic
	WanTrafficUplinkPreferences []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
	// Array of uplink preference rules for VPN traffic
	VpnTrafficUplinkPreferences []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"`
}

// NewInlineObject46 instantiates a new InlineObject46 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject46() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// NewInlineObject46WithDefaults instantiates a new InlineObject46 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject46WithDefaults() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *InlineObject46) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || o.ActiveActiveAutoVpnEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || o.ActiveActiveAutoVpnEnabled == nil {
		return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *InlineObject46) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && o.ActiveActiveAutoVpnEnabled != nil {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *InlineObject46) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *InlineObject46) GetDefaultUplink() string {
	if o == nil || o.DefaultUplink == nil {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || o.DefaultUplink == nil {
		return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *InlineObject46) HasDefaultUplink() bool {
	if o != nil && o.DefaultUplink != nil {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *InlineObject46) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject46) GetLoadBalancingEnabled() bool {
	if o == nil || o.LoadBalancingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || o.LoadBalancingEnabled == nil {
		return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject46) HasLoadBalancingEnabled() bool {
	if o != nil && o.LoadBalancingEnabled != nil {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *InlineObject46) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject46) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	if o == nil || o.WanTrafficUplinkPreferences == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, bool) {
	if o == nil || o.WanTrafficUplinkPreferences == nil {
		return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject46) HasWanTrafficUplinkPreferences() bool {
	if o != nil && o.WanTrafficUplinkPreferences != nil {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineObject46) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject46) GetVpnTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	if o == nil || o.VpnTrafficUplinkPreferences == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetVpnTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, bool) {
	if o == nil || o.VpnTrafficUplinkPreferences == nil {
		return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject46) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && o.VpnTrafficUplinkPreferences != nil {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences and assigns it to the VpnTrafficUplinkPreferences field.
func (o *InlineObject46) SetVpnTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) {
	o.VpnTrafficUplinkPreferences = v
}

func (o InlineObject46) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveActiveAutoVpnEnabled != nil {
		toSerialize["activeActiveAutoVpnEnabled"] = o.ActiveActiveAutoVpnEnabled
	}
	if o.DefaultUplink != nil {
		toSerialize["defaultUplink"] = o.DefaultUplink
	}
	if o.LoadBalancingEnabled != nil {
		toSerialize["loadBalancingEnabled"] = o.LoadBalancingEnabled
	}
	if o.WanTrafficUplinkPreferences != nil {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	if o.VpnTrafficUplinkPreferences != nil {
		toSerialize["vpnTrafficUplinkPreferences"] = o.VpnTrafficUplinkPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject46 struct {
	value *InlineObject46
	isSet bool
}

func (v NullableInlineObject46) Get() *InlineObject46 {
	return v.value
}

func (v *NullableInlineObject46) Set(val *InlineObject46) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject46) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject46) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject46(val *InlineObject46) *NullableInlineObject46 {
	return &NullableInlineObject46{value: val, isSet: true}
}

func (v NullableInlineObject46) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject46) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


