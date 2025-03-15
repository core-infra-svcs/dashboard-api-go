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

// InlineResponse200139 struct for InlineResponse200139
type InlineResponse200139 struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse200139 instantiates a new InlineResponse200139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200139() *InlineResponse200139 {
	this := InlineResponse200139{}
	return &this
}

// NewInlineResponse200139WithDefaults instantiates a new InlineResponse200139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200139WithDefaults() *InlineResponse200139 {
	this := InlineResponse200139{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse200139) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200139) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse200139) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse200139) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse200139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200139 struct {
	value *InlineResponse200139
	isSet bool
}

func (v NullableInlineResponse200139) Get() *InlineResponse200139 {
	return v.value
}

func (v *NullableInlineResponse200139) Set(val *InlineResponse200139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200139(val *InlineResponse200139) *NullableInlineResponse200139 {
	return &NullableInlineResponse200139{value: val, isSet: true}
}

func (v NullableInlineResponse200139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


