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

// NetworksNetworkIdApplianceSsidsNumberDot11w The current setting for Protected Management Frames (802.11w).
type NetworksNetworkIdApplianceSsidsNumberDot11w struct {
	// Whether 802.11w is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11w is required or not.
	Required *bool `json:"required,omitempty"`
}

// NewNetworksNetworkIdApplianceSsidsNumberDot11w instantiates a new NetworksNetworkIdApplianceSsidsNumberDot11w object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSsidsNumberDot11w() *NetworksNetworkIdApplianceSsidsNumberDot11w {
	this := NetworksNetworkIdApplianceSsidsNumberDot11w{}
	return &this
}

// NewNetworksNetworkIdApplianceSsidsNumberDot11wWithDefaults instantiates a new NetworksNetworkIdApplianceSsidsNumberDot11w object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSsidsNumberDot11wWithDefaults() *NetworksNetworkIdApplianceSsidsNumberDot11w {
	this := NetworksNetworkIdApplianceSsidsNumberDot11w{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
    return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *NetworksNetworkIdApplianceSsidsNumberDot11w) SetRequired(v bool) {
	o.Required = &v
}

func (o NetworksNetworkIdApplianceSsidsNumberDot11w) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSsidsNumberDot11w struct {
	value *NetworksNetworkIdApplianceSsidsNumberDot11w
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDot11w) Get() *NetworksNetworkIdApplianceSsidsNumberDot11w {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDot11w) Set(val *NetworksNetworkIdApplianceSsidsNumberDot11w) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDot11w) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDot11w) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSsidsNumberDot11w(val *NetworksNetworkIdApplianceSsidsNumberDot11w) *NullableNetworksNetworkIdApplianceSsidsNumberDot11w {
	return &NullableNetworksNetworkIdApplianceSsidsNumberDot11w{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberDot11w) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberDot11w) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


