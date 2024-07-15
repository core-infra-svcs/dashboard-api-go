/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2007InterfacesWan1VlanTagging VLAN tagging settings.
type InlineResponse2007InterfacesWan1VlanTagging struct {
	// Whether VLAN tagging is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the VLAN to use for VLAN tagging.
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewInlineResponse2007InterfacesWan1VlanTagging instantiates a new InlineResponse2007InterfacesWan1VlanTagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007InterfacesWan1VlanTagging() *InlineResponse2007InterfacesWan1VlanTagging {
	this := InlineResponse2007InterfacesWan1VlanTagging{}
	return &this
}

// NewInlineResponse2007InterfacesWan1VlanTaggingWithDefaults instantiates a new InlineResponse2007InterfacesWan1VlanTagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007InterfacesWan1VlanTaggingWithDefaults() *InlineResponse2007InterfacesWan1VlanTagging {
	this := InlineResponse2007InterfacesWan1VlanTagging{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1VlanTagging) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1VlanTagging) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1VlanTagging) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2007InterfacesWan1VlanTagging) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1VlanTagging) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1VlanTagging) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1VlanTagging) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineResponse2007InterfacesWan1VlanTagging) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o InlineResponse2007InterfacesWan1VlanTagging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007InterfacesWan1VlanTagging struct {
	value *InlineResponse2007InterfacesWan1VlanTagging
	isSet bool
}

func (v NullableInlineResponse2007InterfacesWan1VlanTagging) Get() *InlineResponse2007InterfacesWan1VlanTagging {
	return v.value
}

func (v *NullableInlineResponse2007InterfacesWan1VlanTagging) Set(val *InlineResponse2007InterfacesWan1VlanTagging) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007InterfacesWan1VlanTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007InterfacesWan1VlanTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007InterfacesWan1VlanTagging(val *InlineResponse2007InterfacesWan1VlanTagging) *NullableInlineResponse2007InterfacesWan1VlanTagging {
	return &NullableInlineResponse2007InterfacesWan1VlanTagging{value: val, isSet: true}
}

func (v NullableInlineResponse2007InterfacesWan1VlanTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007InterfacesWan1VlanTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


