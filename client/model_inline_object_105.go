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

// InlineObject105 struct for InlineObject105
type InlineObject105 struct {
	DefaultSettings *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides *[]NetworksNetworkIdSwitchRoutingMulticastOverrides `json:"overrides,omitempty"`
}

// NewInlineObject105 instantiates a new InlineObject105 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject105() *InlineObject105 {
	this := InlineObject105{}
	return &this
}

// NewInlineObject105WithDefaults instantiates a new InlineObject105 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject105WithDefaults() *InlineObject105 {
	this := InlineObject105{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineObject105) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	if o == nil || o.DefaultSettings == nil {
		var ret NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject105) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool) {
	if o == nil || o.DefaultSettings == nil {
		return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineObject105) HasDefaultSettings() bool {
	if o != nil && o.DefaultSettings != nil {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given NetworksNetworkIdSwitchRoutingMulticastDefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineObject105) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject105) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides {
	if o == nil || o.Overrides == nil {
		var ret []NetworksNetworkIdSwitchRoutingMulticastOverrides
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject105) GetOverridesOk() (*[]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject105) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchRoutingMulticastOverrides and assigns it to the Overrides field.
func (o *InlineObject105) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	o.Overrides = &v
}

func (o InlineObject105) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultSettings != nil {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject105 struct {
	value *InlineObject105
	isSet bool
}

func (v NullableInlineObject105) Get() *InlineObject105 {
	return v.value
}

func (v *NullableInlineObject105) Set(val *InlineObject105) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject105) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject105) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject105(val *InlineObject105) *NullableInlineObject105 {
	return &NullableInlineObject105{value: val, isSet: true}
}

func (v NullableInlineObject105) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject105) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

