/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject141 struct for InlineObject141
type InlineObject141 struct {
	DefaultSettings *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []NetworksNetworkIdSwitchRoutingMulticastOverrides `json:"overrides,omitempty"`
}

// NewInlineObject141 instantiates a new InlineObject141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject141() *InlineObject141 {
	this := InlineObject141{}
	return &this
}

// NewInlineObject141WithDefaults instantiates a new InlineObject141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject141WithDefaults() *InlineObject141 {
	this := InlineObject141{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineObject141) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineObject141) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given NetworksNetworkIdSwitchRoutingMulticastDefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineObject141) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject141) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides {
	if o == nil || isNil(o.Overrides) {
		var ret []NetworksNetworkIdSwitchRoutingMulticastOverrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetOverridesOk() ([]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject141) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchRoutingMulticastOverrides and assigns it to the Overrides field.
func (o *InlineObject141) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	o.Overrides = v
}

func (o InlineObject141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject141 struct {
	value *InlineObject141
	isSet bool
}

func (v NullableInlineObject141) Get() *InlineObject141 {
	return v.value
}

func (v *NullableInlineObject141) Set(val *InlineObject141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject141(val *InlineObject141) *NullableInlineObject141 {
	return &NullableInlineObject141{value: val, isSet: true}
}

func (v NullableInlineObject141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


