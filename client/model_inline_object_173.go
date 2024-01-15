/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject173 struct for InlineObject173
type InlineObject173 struct {
	// If true, the SSID device type group policies are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of device type policies.
	DeviceTypePolicies []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"`
}

// NewInlineObject173 instantiates a new InlineObject173 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject173() *InlineObject173 {
	this := InlineObject173{}
	return &this
}

// NewInlineObject173WithDefaults instantiates a new InlineObject173 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject173WithDefaults() *InlineObject173 {
	this := InlineObject173{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject173) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject173) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject173) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeviceTypePolicies returns the DeviceTypePolicies field value if set, zero value otherwise.
func (o *InlineObject173) GetDeviceTypePolicies() []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	if o == nil || isNil(o.DeviceTypePolicies) {
		var ret []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
		return ret
	}
	return o.DeviceTypePolicies
}

// GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetDeviceTypePoliciesOk() ([]NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies, bool) {
	if o == nil || isNil(o.DeviceTypePolicies) {
    return nil, false
	}
	return o.DeviceTypePolicies, true
}

// HasDeviceTypePolicies returns a boolean if a field has been set.
func (o *InlineObject173) HasDeviceTypePolicies() bool {
	if o != nil && !isNil(o.DeviceTypePolicies) {
		return true
	}

	return false
}

// SetDeviceTypePolicies gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies and assigns it to the DeviceTypePolicies field.
func (o *InlineObject173) SetDeviceTypePolicies(v []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) {
	o.DeviceTypePolicies = v
}

func (o InlineObject173) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DeviceTypePolicies) {
		toSerialize["deviceTypePolicies"] = o.DeviceTypePolicies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject173 struct {
	value *InlineObject173
	isSet bool
}

func (v NullableInlineObject173) Get() *InlineObject173 {
	return v.value
}

func (v *NullableInlineObject173) Set(val *InlineObject173) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject173) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject173) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject173(val *InlineObject173) *NullableInlineObject173 {
	return &NullableInlineObject173{value: val, isSet: true}
}

func (v NullableInlineObject173) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject173) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


