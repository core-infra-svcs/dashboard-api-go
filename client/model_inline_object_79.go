/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject79 struct for InlineObject79
type InlineObject79 struct {
	// All firmware upgrade stages in the network with their start time.
	Stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages `json:"stages"`
}

// NewInlineObject79 instantiates a new InlineObject79 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject79(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *InlineObject79 {
	this := InlineObject79{}
	this.Stages = stages
	return &this
}

// NewInlineObject79WithDefaults instantiates a new InlineObject79 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject79WithDefaults() *InlineObject79 {
	this := InlineObject79{}
	return &this
}

// GetStages returns the Stages field value
func (o *InlineObject79) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedEventsStages
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetStagesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InlineObject79) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	o.Stages = v
}

func (o InlineObject79) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject79 struct {
	value *InlineObject79
	isSet bool
}

func (v NullableInlineObject79) Get() *InlineObject79 {
	return v.value
}

func (v *NullableInlineObject79) Set(val *InlineObject79) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject79) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject79) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject79(val *InlineObject79) *NullableInlineObject79 {
	return &NullableInlineObject79{value: val, isSet: true}
}

func (v NullableInlineObject79) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject79) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


