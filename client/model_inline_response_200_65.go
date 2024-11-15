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

// InlineResponse20065 struct for InlineResponse20065
type InlineResponse20065 struct {
	// Whether active-active AutoVPN is enabled
	ActiveActiveAutoVpnEnabled *bool `json:"activeActiveAutoVpnEnabled,omitempty"`
	// The default uplink. Must be one of: 'wan1' or 'wan2'
	DefaultUplink *string `json:"defaultUplink,omitempty"`
	// Whether load balancing is enabled
	LoadBalancingEnabled *bool `json:"loadBalancingEnabled,omitempty"`
	FailoverAndFailback *InlineResponse20065FailoverAndFailback `json:"failoverAndFailback,omitempty"`
	// Uplink preference rules for WAN traffic
	WanTrafficUplinkPreferences []InlineResponse20065WanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
	// Uplink preference rules for VPN traffic
	VpnTrafficUplinkPreferences []InlineResponse20065VpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"`
}

// NewInlineResponse20065 instantiates a new InlineResponse20065 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065WithDefaults() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *InlineResponse20065) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ActiveActiveAutoVpnEnabled) {
    return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *InlineResponse20065) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && !isNil(o.ActiveActiveAutoVpnEnabled) {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *InlineResponse20065) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *InlineResponse20065) GetDefaultUplink() string {
	if o == nil || isNil(o.DefaultUplink) {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || isNil(o.DefaultUplink) {
    return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *InlineResponse20065) HasDefaultUplink() bool {
	if o != nil && !isNil(o.DefaultUplink) {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *InlineResponse20065) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *InlineResponse20065) GetLoadBalancingEnabled() bool {
	if o == nil || isNil(o.LoadBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LoadBalancingEnabled) {
    return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *InlineResponse20065) HasLoadBalancingEnabled() bool {
	if o != nil && !isNil(o.LoadBalancingEnabled) {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *InlineResponse20065) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetFailoverAndFailback returns the FailoverAndFailback field value if set, zero value otherwise.
func (o *InlineResponse20065) GetFailoverAndFailback() InlineResponse20065FailoverAndFailback {
	if o == nil || isNil(o.FailoverAndFailback) {
		var ret InlineResponse20065FailoverAndFailback
		return ret
	}
	return *o.FailoverAndFailback
}

// GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetFailoverAndFailbackOk() (*InlineResponse20065FailoverAndFailback, bool) {
	if o == nil || isNil(o.FailoverAndFailback) {
    return nil, false
	}
	return o.FailoverAndFailback, true
}

// HasFailoverAndFailback returns a boolean if a field has been set.
func (o *InlineResponse20065) HasFailoverAndFailback() bool {
	if o != nil && !isNil(o.FailoverAndFailback) {
		return true
	}

	return false
}

// SetFailoverAndFailback gets a reference to the given InlineResponse20065FailoverAndFailback and assigns it to the FailoverAndFailback field.
func (o *InlineResponse20065) SetFailoverAndFailback(v InlineResponse20065FailoverAndFailback) {
	o.FailoverAndFailback = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineResponse20065) GetWanTrafficUplinkPreferences() []InlineResponse20065WanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []InlineResponse20065WanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetWanTrafficUplinkPreferencesOk() ([]InlineResponse20065WanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineResponse20065) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []InlineResponse20065WanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineResponse20065) SetWanTrafficUplinkPreferences(v []InlineResponse20065WanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineResponse20065) GetVpnTrafficUplinkPreferences() []InlineResponse20065VpnTrafficUplinkPreferences {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
		var ret []InlineResponse20065VpnTrafficUplinkPreferences
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetVpnTrafficUplinkPreferencesOk() ([]InlineResponse20065VpnTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.VpnTrafficUplinkPreferences) {
    return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineResponse20065) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.VpnTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []InlineResponse20065VpnTrafficUplinkPreferences and assigns it to the VpnTrafficUplinkPreferences field.
func (o *InlineResponse20065) SetVpnTrafficUplinkPreferences(v []InlineResponse20065VpnTrafficUplinkPreferences) {
	o.VpnTrafficUplinkPreferences = v
}

func (o InlineResponse20065) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20065 struct {
	value *InlineResponse20065
	isSet bool
}

func (v NullableInlineResponse20065) Get() *InlineResponse20065 {
	return v.value
}

func (v *NullableInlineResponse20065) Set(val *InlineResponse20065) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065(val *InlineResponse20065) *NullableInlineResponse20065 {
	return &NullableInlineResponse20065{value: val, isSet: true}
}

func (v NullableInlineResponse20065) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


