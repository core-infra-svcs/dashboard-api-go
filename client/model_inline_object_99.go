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

// InlineObject99 struct for InlineObject99
type InlineObject99 struct {
	// Array of Staged Upgrade Groups
	Json []NetworksNetworkIdFirmwareUpgradesStagedStagesJson `json:"_json,omitempty"`
}

// NewInlineObject99 instantiates a new InlineObject99 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject99() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// NewInlineObject99WithDefaults instantiates a new InlineObject99 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject99WithDefaults() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *InlineObject99) GetJson() []NetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	if o == nil || isNil(o.Json) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedStagesJson
		return ret
	}
	return o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetJsonOk() ([]NetworksNetworkIdFirmwareUpgradesStagedStagesJson, bool) {
	if o == nil || isNil(o.Json) {
    return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *InlineObject99) HasJson() bool {
	if o != nil && !isNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedStagesJson and assigns it to the Json field.
func (o *InlineObject99) SetJson(v []NetworksNetworkIdFirmwareUpgradesStagedStagesJson) {
	o.Json = v
}

func (o InlineObject99) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Json) {
		toSerialize["_json"] = o.Json
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject99 struct {
	value *InlineObject99
	isSet bool
}

func (v NullableInlineObject99) Get() *InlineObject99 {
	return v.value
}

func (v *NullableInlineObject99) Set(val *InlineObject99) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject99) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject99) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject99(val *InlineObject99) *NullableInlineObject99 {
	return &NullableInlineObject99{value: val, isSet: true}
}

func (v NullableInlineObject99) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject99) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


