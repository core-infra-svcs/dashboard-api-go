/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject117 struct for InlineObject117
type InlineObject117 struct {
	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize *int32 `json:"defaultMtuSize,omitempty"`
	// Override MTU size for individual switches or switch profiles. An empty array will clear overrides.
	Overrides []NetworksNetworkIdSwitchMtuOverrides `json:"overrides,omitempty"`
}

// NewInlineObject117 instantiates a new InlineObject117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject117() *InlineObject117 {
	this := InlineObject117{}
	return &this
}

// NewInlineObject117WithDefaults instantiates a new InlineObject117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject117WithDefaults() *InlineObject117 {
	this := InlineObject117{}
	return &this
}

// GetDefaultMtuSize returns the DefaultMtuSize field value if set, zero value otherwise.
func (o *InlineObject117) GetDefaultMtuSize() int32 {
	if o == nil || isNil(o.DefaultMtuSize) {
		var ret int32
		return ret
	}
	return *o.DefaultMtuSize
}

// GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject117) GetDefaultMtuSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DefaultMtuSize) {
    return nil, false
	}
	return o.DefaultMtuSize, true
}

// HasDefaultMtuSize returns a boolean if a field has been set.
func (o *InlineObject117) HasDefaultMtuSize() bool {
	if o != nil && !isNil(o.DefaultMtuSize) {
		return true
	}

	return false
}

// SetDefaultMtuSize gets a reference to the given int32 and assigns it to the DefaultMtuSize field.
func (o *InlineObject117) SetDefaultMtuSize(v int32) {
	o.DefaultMtuSize = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject117) GetOverrides() []NetworksNetworkIdSwitchMtuOverrides {
	if o == nil || isNil(o.Overrides) {
		var ret []NetworksNetworkIdSwitchMtuOverrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject117) GetOverridesOk() ([]NetworksNetworkIdSwitchMtuOverrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject117) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []NetworksNetworkIdSwitchMtuOverrides and assigns it to the Overrides field.
func (o *InlineObject117) SetOverrides(v []NetworksNetworkIdSwitchMtuOverrides) {
	o.Overrides = v
}

func (o InlineObject117) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultMtuSize) {
		toSerialize["defaultMtuSize"] = o.DefaultMtuSize
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject117 struct {
	value *InlineObject117
	isSet bool
}

func (v NullableInlineObject117) Get() *InlineObject117 {
	return v.value
}

func (v *NullableInlineObject117) Set(val *InlineObject117) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject117) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject117(val *InlineObject117) *NullableInlineObject117 {
	return &NullableInlineObject117{value: val, isSet: true}
}

func (v NullableInlineObject117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


