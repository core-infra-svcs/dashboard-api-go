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

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	Group *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup `json:"group,omitempty"`
}

// NewInlineResponse20035 instantiates a new InlineResponse20035 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035WithDefaults() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20035) GetGroup() NetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	if o == nil || isNil(o.Group) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedStagesGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetGroupOk() (*NetworksNetworkIdFirmwareUpgradesStagedStagesGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20035) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedStagesGroup and assigns it to the Group field.
func (o *InlineResponse20035) SetGroup(v NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) {
	o.Group = &v
}

func (o InlineResponse20035) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035 struct {
	value *InlineResponse20035
	isSet bool
}

func (v NullableInlineResponse20035) Get() *InlineResponse20035 {
	return v.value
}

func (v *NullableInlineResponse20035) Set(val *InlineResponse20035) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035(val *InlineResponse20035) *NullableInlineResponse20035 {
	return &NullableInlineResponse20035{value: val, isSet: true}
}

func (v NullableInlineResponse20035) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


