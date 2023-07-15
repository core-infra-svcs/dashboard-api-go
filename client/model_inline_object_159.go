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

// InlineObject159 struct for InlineObject159
type InlineObject159 struct {
	// If true, the SSID device type group policies are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of device type policies.
	DeviceTypePolicies []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"`
}

// NewInlineObject159 instantiates a new InlineObject159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject159() *InlineObject159 {
	this := InlineObject159{}
	return &this
}

// NewInlineObject159WithDefaults instantiates a new InlineObject159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject159WithDefaults() *InlineObject159 {
	this := InlineObject159{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject159) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject159) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject159) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeviceTypePolicies returns the DeviceTypePolicies field value if set, zero value otherwise.
func (o *InlineObject159) GetDeviceTypePolicies() []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	if o == nil || isNil(o.DeviceTypePolicies) {
		var ret []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
		return ret
	}
	return o.DeviceTypePolicies
}

// GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetDeviceTypePoliciesOk() ([]NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies, bool) {
	if o == nil || isNil(o.DeviceTypePolicies) {
    return nil, false
	}
	return o.DeviceTypePolicies, true
}

// HasDeviceTypePolicies returns a boolean if a field has been set.
func (o *InlineObject159) HasDeviceTypePolicies() bool {
	if o != nil && !isNil(o.DeviceTypePolicies) {
		return true
	}

	return false
}

// SetDeviceTypePolicies gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies and assigns it to the DeviceTypePolicies field.
func (o *InlineObject159) SetDeviceTypePolicies(v []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) {
	o.DeviceTypePolicies = v
}

func (o InlineObject159) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DeviceTypePolicies) {
		toSerialize["deviceTypePolicies"] = o.DeviceTypePolicies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject159 struct {
	value *InlineObject159
	isSet bool
}

func (v NullableInlineObject159) Get() *InlineObject159 {
	return v.value
}

func (v *NullableInlineObject159) Set(val *InlineObject159) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject159) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject159(val *InlineObject159) *NullableInlineObject159 {
	return &NullableInlineObject159{value: val, isSet: true}
}

func (v NullableInlineObject159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


