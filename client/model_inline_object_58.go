/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject58 struct for InlineObject58
type InlineObject58 struct {
	// Toggle for enabling or disabling active-active AutoVPN
	ActiveActiveAutoVpnEnabled *bool `json:"activeActiveAutoVpnEnabled,omitempty"`
	// The default uplink. Must be one of: 'wan1' or 'wan2'
	DefaultUplink *string `json:"defaultUplink,omitempty"`
	// Toggle for enabling or disabling load balancing
	LoadBalancingEnabled *bool `json:"loadBalancingEnabled,omitempty"`
	FailoverAndFailback *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback `json:"failoverAndFailback,omitempty"`
	// Array of uplink preference rules for WAN traffic
	WanTrafficUplinkPreferences []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
	// Array of uplink preference rules for VPN traffic
	VpnTrafficUplinkPreferences []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"`
}

// NewInlineObject58 instantiates a new InlineObject58 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject58() *InlineObject58 {
	this := InlineObject58{}
	return &this
}

// NewInlineObject58WithDefaults instantiates a new InlineObject58 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject58WithDefaults() *InlineObject58 {
	this := InlineObject58{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *InlineObject58) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
    return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *InlineObject58) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && !isNil(o.ActiveActiveAutoVpnEnabled) {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *InlineObject58) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *InlineObject58) GetDefaultUplink() string {
	if o == nil || isNil(o.DefaultUplink) {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || isNil(o.DefaultUplink) {
    return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *InlineObject58) HasDefaultUplink() bool {
	if o != nil && !isNil(o.DefaultUplink) {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *InlineObject58) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject58) GetLoadBalancingEnabled() bool {
	if o == nil || isNil(o.LoadBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LoadBalancingEnabled) {
    return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject58) HasLoadBalancingEnabled() bool {
	if o != nil && !isNil(o.LoadBalancingEnabled) {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *InlineObject58) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetFailoverAndFailback returns the FailoverAndFailback field value if set, zero value otherwise.
func (o *InlineObject58) GetFailoverAndFailback() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	if o == nil || isNil(o.FailoverAndFailback) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback
		return ret
	}
	return *o.FailoverAndFailback
}

// GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetFailoverAndFailbackOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback, bool) {
	if o == nil || isNil(o.FailoverAndFailback) {
    return nil, false
	}
	return o.FailoverAndFailback, true
}

// HasFailoverAndFailback returns a boolean if a field has been set.
func (o *InlineObject58) HasFailoverAndFailback() bool {
	if o != nil && !isNil(o.FailoverAndFailback) {
		return true
	}

	return false
}

// SetFailoverAndFailback gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback and assigns it to the FailoverAndFailback field.
func (o *InlineObject58) SetFailoverAndFailback(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) {
	o.FailoverAndFailback = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject58) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject58) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineObject58) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject58) GetVpnTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetVpnTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
    return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject58) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.VpnTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences and assigns it to the VpnTrafficUplinkPreferences field.
func (o *InlineObject58) SetVpnTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) {
	o.VpnTrafficUplinkPreferences = v
}

func (o InlineObject58) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActiveActiveAutoVpnEnabled) {
		toSerialize["activeActiveAutoVpnEnabled"] = o.ActiveActiveAutoVpnEnabled
	}
	if !isNil(o.DefaultUplink) {
		toSerialize["defaultUplink"] = o.DefaultUplink
	}
	if !isNil(o.LoadBalancingEnabled) {
		toSerialize["loadBalancingEnabled"] = o.LoadBalancingEnabled
	}
	if !isNil(o.FailoverAndFailback) {
		toSerialize["failoverAndFailback"] = o.FailoverAndFailback
	}
	if !isNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	if !isNil(o.VpnTrafficUplinkPreferences) {
		toSerialize["vpnTrafficUplinkPreferences"] = o.VpnTrafficUplinkPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject58 struct {
	value *InlineObject58
	isSet bool
}

func (v NullableInlineObject58) Get() *InlineObject58 {
	return v.value
}

func (v *NullableInlineObject58) Set(val *InlineObject58) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject58) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject58) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject58(val *InlineObject58) *NullableInlineObject58 {
	return &NullableInlineObject58{value: val, isSet: true}
}

func (v NullableInlineObject58) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject58) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


