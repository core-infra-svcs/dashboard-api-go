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

// InlineObject132 struct for InlineObject132
type InlineObject132 struct {
	DefaultSettings *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []NetworksNetworkIdSwitchRoutingMulticastOverrides `json:"overrides,omitempty"`
}

// NewInlineObject132 instantiates a new InlineObject132 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject132() *InlineObject132 {
	this := InlineObject132{}
	return &this
}

// NewInlineObject132WithDefaults instantiates a new InlineObject132 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject132WithDefaults() *InlineObject132 {
	this := InlineObject132{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineObject132) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineObject132) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given NetworksNetworkIdSwitchRoutingMulticastDefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineObject132) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject132) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides {
	if o == nil || isNil(o.Overrides) {
		var ret []NetworksNetworkIdSwitchRoutingMulticastOverrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetOverridesOk() ([]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject132) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchRoutingMulticastOverrides and assigns it to the Overrides field.
func (o *InlineObject132) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	o.Overrides = v
}

func (o InlineObject132) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject132 struct {
	value *InlineObject132
	isSet bool
}

func (v NullableInlineObject132) Get() *InlineObject132 {
	return v.value
}

func (v *NullableInlineObject132) Set(val *InlineObject132) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject132) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject132) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject132(val *InlineObject132) *NullableInlineObject132 {
	return &NullableInlineObject132{value: val, isSet: true}
}

func (v NullableInlineObject132) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject132) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


