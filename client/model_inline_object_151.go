/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject151 struct for InlineObject151
type InlineObject151 struct {
	DefaultSettings *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []NetworksNetworkIdSwitchRoutingMulticastOverrides `json:"overrides,omitempty"`
}

// NewInlineObject151 instantiates a new InlineObject151 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject151() *InlineObject151 {
	this := InlineObject151{}
	return &this
}

// NewInlineObject151WithDefaults instantiates a new InlineObject151 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject151WithDefaults() *InlineObject151 {
	this := InlineObject151{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineObject151) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject151) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineObject151) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given NetworksNetworkIdSwitchRoutingMulticastDefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineObject151) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject151) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides {
	if o == nil || isNil(o.Overrides) {
		var ret []NetworksNetworkIdSwitchRoutingMulticastOverrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject151) GetOverridesOk() ([]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject151) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchRoutingMulticastOverrides and assigns it to the Overrides field.
func (o *InlineObject151) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	o.Overrides = v
}

func (o InlineObject151) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject151 struct {
	value *InlineObject151
	isSet bool
}

func (v NullableInlineObject151) Get() *InlineObject151 {
	return v.value
}

func (v *NullableInlineObject151) Set(val *InlineObject151) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject151) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject151) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject151(val *InlineObject151) *NullableInlineObject151 {
	return &NullableInlineObject151{value: val, isSet: true}
}

func (v NullableInlineObject151) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject151) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


