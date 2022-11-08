/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject83 struct for InlineObject83
type InlineObject83 struct {
	// Array of Staged Upgrade Groups
	Json []NetworksNetworkIdFirmwareUpgradesStagedStagesJson `json:"_json,omitempty"`
}

// NewInlineObject83 instantiates a new InlineObject83 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject83() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// NewInlineObject83WithDefaults instantiates a new InlineObject83 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject83WithDefaults() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *InlineObject83) GetJson() []NetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	if o == nil || isNil(o.Json) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedStagesJson
		return ret
	}
	return o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetJsonOk() ([]NetworksNetworkIdFirmwareUpgradesStagedStagesJson, bool) {
	if o == nil || isNil(o.Json) {
    return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *InlineObject83) HasJson() bool {
	if o != nil && !isNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedStagesJson and assigns it to the Json field.
func (o *InlineObject83) SetJson(v []NetworksNetworkIdFirmwareUpgradesStagedStagesJson) {
	o.Json = v
}

func (o InlineObject83) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Json) {
		toSerialize["_json"] = o.Json
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject83 struct {
	value *InlineObject83
	isSet bool
}

func (v NullableInlineObject83) Get() *InlineObject83 {
	return v.value
}

func (v *NullableInlineObject83) Set(val *InlineObject83) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject83) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject83) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject83(val *InlineObject83) *NullableInlineObject83 {
	return &NullableInlineObject83{value: val, isSet: true}
}

func (v NullableInlineObject83) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject83) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


