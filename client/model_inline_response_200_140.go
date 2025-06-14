/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200140 struct for InlineResponse200140
type InlineResponse200140 struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse200140 instantiates a new InlineResponse200140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200140() *InlineResponse200140 {
	this := InlineResponse200140{}
	return &this
}

// NewInlineResponse200140WithDefaults instantiates a new InlineResponse200140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200140WithDefaults() *InlineResponse200140 {
	this := InlineResponse200140{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse200140) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200140) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse200140) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse200140) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse200140) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200140 struct {
	value *InlineResponse200140
	isSet bool
}

func (v NullableInlineResponse200140) Get() *InlineResponse200140 {
	return v.value
}

func (v *NullableInlineResponse200140) Set(val *InlineResponse200140) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200140) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200140(val *InlineResponse200140) *NullableInlineResponse200140 {
	return &NullableInlineResponse200140{value: val, isSet: true}
}

func (v NullableInlineResponse200140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


