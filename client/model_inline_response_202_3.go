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

// InlineResponse2023 struct for InlineResponse2023
type InlineResponse2023 struct {
	// Shows the success of the reboot
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse2023 instantiates a new InlineResponse2023 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2023() *InlineResponse2023 {
	this := InlineResponse2023{}
	return &this
}

// NewInlineResponse2023WithDefaults instantiates a new InlineResponse2023 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2023WithDefaults() *InlineResponse2023 {
	this := InlineResponse2023{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse2023) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2023) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse2023) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse2023) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse2023) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2023 struct {
	value *InlineResponse2023
	isSet bool
}

func (v NullableInlineResponse2023) Get() *InlineResponse2023 {
	return v.value
}

func (v *NullableInlineResponse2023) Set(val *InlineResponse2023) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2023) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2023) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2023(val *InlineResponse2023) *NullableInlineResponse2023 {
	return &NullableInlineResponse2023{value: val, isSet: true}
}

func (v NullableInlineResponse2023) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2023) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


