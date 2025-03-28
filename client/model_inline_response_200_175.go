/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200175 struct for InlineResponse200175
type InlineResponse200175 struct {
	VlanProfile *InlineResponse200175VlanProfile `json:"vlanProfile,omitempty"`
	// Array of Device Serials
	Serials []string `json:"serials,omitempty"`
	// Array of Switch Stack IDs
	StackIds []string `json:"stackIds,omitempty"`
}

// NewInlineResponse200175 instantiates a new InlineResponse200175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200175() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// NewInlineResponse200175WithDefaults instantiates a new InlineResponse200175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200175WithDefaults() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *InlineResponse200175) GetVlanProfile() InlineResponse200175VlanProfile {
	if o == nil || isNil(o.VlanProfile) {
		var ret InlineResponse200175VlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetVlanProfileOk() (*InlineResponse200175VlanProfile, bool) {
	if o == nil || isNil(o.VlanProfile) {
    return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *InlineResponse200175) HasVlanProfile() bool {
	if o != nil && !isNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given InlineResponse200175VlanProfile and assigns it to the VlanProfile field.
func (o *InlineResponse200175) SetVlanProfile(v InlineResponse200175VlanProfile) {
	o.VlanProfile = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200175) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200175) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200175) SetSerials(v []string) {
	o.Serials = v
}

// GetStackIds returns the StackIds field value if set, zero value otherwise.
func (o *InlineResponse200175) GetStackIds() []string {
	if o == nil || isNil(o.StackIds) {
		var ret []string
		return ret
	}
	return o.StackIds
}

// GetStackIdsOk returns a tuple with the StackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetStackIdsOk() ([]string, bool) {
	if o == nil || isNil(o.StackIds) {
    return nil, false
	}
	return o.StackIds, true
}

// HasStackIds returns a boolean if a field has been set.
func (o *InlineResponse200175) HasStackIds() bool {
	if o != nil && !isNil(o.StackIds) {
		return true
	}

	return false
}

// SetStackIds gets a reference to the given []string and assigns it to the StackIds field.
func (o *InlineResponse200175) SetStackIds(v []string) {
	o.StackIds = v
}

func (o InlineResponse200175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VlanProfile) {
		toSerialize["vlanProfile"] = o.VlanProfile
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.StackIds) {
		toSerialize["stackIds"] = o.StackIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200175 struct {
	value *InlineResponse200175
	isSet bool
}

func (v NullableInlineResponse200175) Get() *InlineResponse200175 {
	return v.value
}

func (v *NullableInlineResponse200175) Set(val *InlineResponse200175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200175(val *InlineResponse200175) *NullableInlineResponse200175 {
	return &NullableInlineResponse200175{value: val, isSet: true}
}

func (v NullableInlineResponse200175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


