/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20055 struct for InlineResponse20055
type InlineResponse20055 struct {
	Group *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup `json:"group,omitempty"`
}

// NewInlineResponse20055 instantiates a new InlineResponse20055 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20055() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// NewInlineResponse20055WithDefaults instantiates a new InlineResponse20055 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20055WithDefaults() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20055) GetGroup() NetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	if o == nil || isNil(o.Group) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedStagesGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055) GetGroupOk() (*NetworksNetworkIdFirmwareUpgradesStagedStagesGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20055) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedStagesGroup and assigns it to the Group field.
func (o *InlineResponse20055) SetGroup(v NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) {
	o.Group = &v
}

func (o InlineResponse20055) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20055 struct {
	value *InlineResponse20055
	isSet bool
}

func (v NullableInlineResponse20055) Get() *InlineResponse20055 {
	return v.value
}

func (v *NullableInlineResponse20055) Set(val *InlineResponse20055) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20055) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20055) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20055(val *InlineResponse20055) *NullableInlineResponse20055 {
	return &NullableInlineResponse20055{value: val, isSet: true}
}

func (v NullableInlineResponse20055) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20055) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


