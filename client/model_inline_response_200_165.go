/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200165 struct for InlineResponse200165
type InlineResponse200165 struct {
	VlanProfile *InlineResponse200165VlanProfile `json:"vlanProfile,omitempty"`
	// Array of Device Serials
	Serials []string `json:"serials,omitempty"`
	// Array of Switch Stack IDs
	StackIds []string `json:"stackIds,omitempty"`
}

// NewInlineResponse200165 instantiates a new InlineResponse200165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200165() *InlineResponse200165 {
	this := InlineResponse200165{}
	return &this
}

// NewInlineResponse200165WithDefaults instantiates a new InlineResponse200165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200165WithDefaults() *InlineResponse200165 {
	this := InlineResponse200165{}
	return &this
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *InlineResponse200165) GetVlanProfile() InlineResponse200165VlanProfile {
	if o == nil || isNil(o.VlanProfile) {
		var ret InlineResponse200165VlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200165) GetVlanProfileOk() (*InlineResponse200165VlanProfile, bool) {
	if o == nil || isNil(o.VlanProfile) {
    return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *InlineResponse200165) HasVlanProfile() bool {
	if o != nil && !isNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given InlineResponse200165VlanProfile and assigns it to the VlanProfile field.
func (o *InlineResponse200165) SetVlanProfile(v InlineResponse200165VlanProfile) {
	o.VlanProfile = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200165) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200165) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200165) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200165) SetSerials(v []string) {
	o.Serials = v
}

// GetStackIds returns the StackIds field value if set, zero value otherwise.
func (o *InlineResponse200165) GetStackIds() []string {
	if o == nil || isNil(o.StackIds) {
		var ret []string
		return ret
	}
	return o.StackIds
}

// GetStackIdsOk returns a tuple with the StackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200165) GetStackIdsOk() ([]string, bool) {
	if o == nil || isNil(o.StackIds) {
    return nil, false
	}
	return o.StackIds, true
}

// HasStackIds returns a boolean if a field has been set.
func (o *InlineResponse200165) HasStackIds() bool {
	if o != nil && !isNil(o.StackIds) {
		return true
	}

	return false
}

// SetStackIds gets a reference to the given []string and assigns it to the StackIds field.
func (o *InlineResponse200165) SetStackIds(v []string) {
	o.StackIds = v
}

func (o InlineResponse200165) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200165 struct {
	value *InlineResponse200165
	isSet bool
}

func (v NullableInlineResponse200165) Get() *InlineResponse200165 {
	return v.value
}

func (v *NullableInlineResponse200165) Set(val *InlineResponse200165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200165(val *InlineResponse200165) *NullableInlineResponse200165 {
	return &NullableInlineResponse200165{value: val, isSet: true}
}

func (v NullableInlineResponse200165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


