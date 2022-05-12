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

// InlineObject135 struct for InlineObject135
type InlineObject135 struct {
	// If true, the SSID device type group policies are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of device type policies.
	DeviceTypePolicies *[]NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"`
}

// NewInlineObject135 instantiates a new InlineObject135 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject135() *InlineObject135 {
	this := InlineObject135{}
	return &this
}

// NewInlineObject135WithDefaults instantiates a new InlineObject135 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject135WithDefaults() *InlineObject135 {
	this := InlineObject135{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject135) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject135) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject135) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeviceTypePolicies returns the DeviceTypePolicies field value if set, zero value otherwise.
func (o *InlineObject135) GetDeviceTypePolicies() []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	if o == nil || o.DeviceTypePolicies == nil {
		var ret []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
		return ret
	}
	return *o.DeviceTypePolicies
}

// GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDeviceTypePoliciesOk() (*[]NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies, bool) {
	if o == nil || o.DeviceTypePolicies == nil {
		return nil, false
	}
	return o.DeviceTypePolicies, true
}

// HasDeviceTypePolicies returns a boolean if a field has been set.
func (o *InlineObject135) HasDeviceTypePolicies() bool {
	if o != nil && o.DeviceTypePolicies != nil {
		return true
	}

	return false
}

// SetDeviceTypePolicies gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies and assigns it to the DeviceTypePolicies field.
func (o *InlineObject135) SetDeviceTypePolicies(v []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) {
	o.DeviceTypePolicies = &v
}

func (o InlineObject135) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.DeviceTypePolicies != nil {
		toSerialize["deviceTypePolicies"] = o.DeviceTypePolicies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject135 struct {
	value *InlineObject135
	isSet bool
}

func (v NullableInlineObject135) Get() *InlineObject135 {
	return v.value
}

func (v *NullableInlineObject135) Set(val *InlineObject135) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject135) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject135) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject135(val *InlineObject135) *NullableInlineObject135 {
	return &NullableInlineObject135{value: val, isSet: true}
}

func (v NullableInlineObject135) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject135) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

