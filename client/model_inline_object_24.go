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

// InlineObject24 struct for InlineObject24
type InlineObject24 struct {
	// Operation to run on the sensor. 'enableDownstreamPower', 'disableDownstreamPower', and 'cycleDownstreamPower' turn power on/off to the device that is connected downstream of an MT40 power monitor. 'refreshData' causes an MT15 or MT40 device to upload its latest readings so that they are immediately available in the Dashboard API.
	Operation string `json:"operation"`
}

// NewInlineObject24 instantiates a new InlineObject24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject24(operation string) *InlineObject24 {
	this := InlineObject24{}
	this.Operation = operation
	return &this
}

// NewInlineObject24WithDefaults instantiates a new InlineObject24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject24WithDefaults() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// GetOperation returns the Operation field value
func (o *InlineObject24) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetOperationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *InlineObject24) SetOperation(v string) {
	o.Operation = v
}

func (o InlineObject24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject24 struct {
	value *InlineObject24
	isSet bool
}

func (v NullableInlineObject24) Get() *InlineObject24 {
	return v.value
}

func (v *NullableInlineObject24) Set(val *InlineObject24) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject24) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject24(val *InlineObject24) *NullableInlineObject24 {
	return &NullableInlineObject24{value: val, isSet: true}
}

func (v NullableInlineObject24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


