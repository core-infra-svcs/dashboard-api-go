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

// InlineObject69 struct for InlineObject69
type InlineObject69 struct {
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

// NewInlineObject69 instantiates a new InlineObject69 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject69() *InlineObject69 {
	this := InlineObject69{}
	return &this
}

// NewInlineObject69WithDefaults instantiates a new InlineObject69 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject69WithDefaults() *InlineObject69 {
	this := InlineObject69{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *InlineObject69) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
    return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *InlineObject69) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && !isNil(o.ActiveActiveAutoVpnEnabled) {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *InlineObject69) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *InlineObject69) GetDefaultUplink() string {
	if o == nil || isNil(o.DefaultUplink) {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || isNil(o.DefaultUplink) {
    return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *InlineObject69) HasDefaultUplink() bool {
	if o != nil && !isNil(o.DefaultUplink) {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *InlineObject69) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject69) GetLoadBalancingEnabled() bool {
	if o == nil || isNil(o.LoadBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LoadBalancingEnabled) {
    return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject69) HasLoadBalancingEnabled() bool {
	if o != nil && !isNil(o.LoadBalancingEnabled) {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *InlineObject69) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetFailoverAndFailback returns the FailoverAndFailback field value if set, zero value otherwise.
func (o *InlineObject69) GetFailoverAndFailback() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	if o == nil || isNil(o.FailoverAndFailback) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback
		return ret
	}
	return *o.FailoverAndFailback
}

// GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetFailoverAndFailbackOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback, bool) {
	if o == nil || isNil(o.FailoverAndFailback) {
    return nil, false
	}
	return o.FailoverAndFailback, true
}

// HasFailoverAndFailback returns a boolean if a field has been set.
func (o *InlineObject69) HasFailoverAndFailback() bool {
	if o != nil && !isNil(o.FailoverAndFailback) {
		return true
	}

	return false
}

// SetFailoverAndFailback gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback and assigns it to the FailoverAndFailback field.
func (o *InlineObject69) SetFailoverAndFailback(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) {
	o.FailoverAndFailback = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject69) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject69) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineObject69) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject69) GetVpnTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetVpnTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
    return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject69) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.VpnTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences and assigns it to the VpnTrafficUplinkPreferences field.
func (o *InlineObject69) SetVpnTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) {
	o.VpnTrafficUplinkPreferences = v
}

func (o InlineObject69) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject69 struct {
	value *InlineObject69
	isSet bool
}

func (v NullableInlineObject69) Get() *InlineObject69 {
	return v.value
}

func (v *NullableInlineObject69) Set(val *InlineObject69) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject69) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject69) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject69(val *InlineObject69) *NullableInlineObject69 {
	return &NullableInlineObject69{value: val, isSet: true}
}

func (v NullableInlineObject69) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject69) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


