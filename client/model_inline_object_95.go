/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject95 struct for InlineObject95
type InlineObject95 struct {
	// All firmware upgrade stages in the network with their start time.
	Stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages `json:"stages"`
}

// NewInlineObject95 instantiates a new InlineObject95 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject95(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *InlineObject95 {
	this := InlineObject95{}
	this.Stages = stages
	return &this
}

// NewInlineObject95WithDefaults instantiates a new InlineObject95 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject95WithDefaults() *InlineObject95 {
	this := InlineObject95{}
	return &this
}

// GetStages returns the Stages field value
func (o *InlineObject95) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedEventsStages
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetStagesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InlineObject95) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	o.Stages = v
}

func (o InlineObject95) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject95 struct {
	value *InlineObject95
	isSet bool
}

func (v NullableInlineObject95) Get() *InlineObject95 {
	return v.value
}

func (v *NullableInlineObject95) Set(val *InlineObject95) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject95) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject95) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject95(val *InlineObject95) *NullableInlineObject95 {
	return &NullableInlineObject95{value: val, isSet: true}
}

func (v NullableInlineObject95) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject95) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


