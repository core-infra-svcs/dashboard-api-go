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

// InlineObject149 struct for InlineObject149
type InlineObject149 struct {
	DefaultSettings *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []NetworksNetworkIdSwitchRoutingMulticastOverrides `json:"overrides,omitempty"`
}

// NewInlineObject149 instantiates a new InlineObject149 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject149() *InlineObject149 {
	this := InlineObject149{}
	return &this
}

// NewInlineObject149WithDefaults instantiates a new InlineObject149 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject149WithDefaults() *InlineObject149 {
	this := InlineObject149{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineObject149) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject149) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineObject149) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given NetworksNetworkIdSwitchRoutingMulticastDefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineObject149) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject149) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides {
	if o == nil || isNil(o.Overrides) {
		var ret []NetworksNetworkIdSwitchRoutingMulticastOverrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject149) GetOverridesOk() ([]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject149) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchRoutingMulticastOverrides and assigns it to the Overrides field.
func (o *InlineObject149) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	o.Overrides = v
}

func (o InlineObject149) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject149 struct {
	value *InlineObject149
	isSet bool
}

func (v NullableInlineObject149) Get() *InlineObject149 {
	return v.value
}

func (v *NullableInlineObject149) Set(val *InlineObject149) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject149) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject149) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject149(val *InlineObject149) *NullableInlineObject149 {
	return &NullableInlineObject149{value: val, isSet: true}
}

func (v NullableInlineObject149) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject149) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


