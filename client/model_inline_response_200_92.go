/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20092 struct for InlineResponse20092
type InlineResponse20092 struct {
	VlanProfile *InlineResponse20092VlanProfile `json:"vlanProfile,omitempty"`
	// Array of Device Serials
	Serials []string `json:"serials,omitempty"`
	// Array of Switch Stack IDs
	StackIds []string `json:"stackIds,omitempty"`
}

// NewInlineResponse20092 instantiates a new InlineResponse20092 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092WithDefaults() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *InlineResponse20092) GetVlanProfile() InlineResponse20092VlanProfile {
	if o == nil || isNil(o.VlanProfile) {
		var ret InlineResponse20092VlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetVlanProfileOk() (*InlineResponse20092VlanProfile, bool) {
	if o == nil || isNil(o.VlanProfile) {
    return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *InlineResponse20092) HasVlanProfile() bool {
	if o != nil && !isNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given InlineResponse20092VlanProfile and assigns it to the VlanProfile field.
func (o *InlineResponse20092) SetVlanProfile(v InlineResponse20092VlanProfile) {
	o.VlanProfile = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse20092) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse20092) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse20092) SetSerials(v []string) {
	o.Serials = v
}

// GetStackIds returns the StackIds field value if set, zero value otherwise.
func (o *InlineResponse20092) GetStackIds() []string {
	if o == nil || isNil(o.StackIds) {
		var ret []string
		return ret
	}
	return o.StackIds
}

// GetStackIdsOk returns a tuple with the StackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetStackIdsOk() ([]string, bool) {
	if o == nil || isNil(o.StackIds) {
    return nil, false
	}
	return o.StackIds, true
}

// HasStackIds returns a boolean if a field has been set.
func (o *InlineResponse20092) HasStackIds() bool {
	if o != nil && !isNil(o.StackIds) {
		return true
	}

	return false
}

// SetStackIds gets a reference to the given []string and assigns it to the StackIds field.
func (o *InlineResponse20092) SetStackIds(v []string) {
	o.StackIds = v
}

func (o InlineResponse20092) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20092 struct {
	value *InlineResponse20092
	isSet bool
}

func (v NullableInlineResponse20092) Get() *InlineResponse20092 {
	return v.value
}

func (v *NullableInlineResponse20092) Set(val *InlineResponse20092) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092(val *InlineResponse20092) *NullableInlineResponse20092 {
	return &NullableInlineResponse20092{value: val, isSet: true}
}

func (v NullableInlineResponse20092) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


