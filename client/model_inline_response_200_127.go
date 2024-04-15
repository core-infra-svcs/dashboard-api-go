/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200127 struct for InlineResponse200127
type InlineResponse200127 struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse200127 instantiates a new InlineResponse200127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200127() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// NewInlineResponse200127WithDefaults instantiates a new InlineResponse200127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200127WithDefaults() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse200127) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse200127) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse200127) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse200127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200127 struct {
	value *InlineResponse200127
	isSet bool
}

func (v NullableInlineResponse200127) Get() *InlineResponse200127 {
	return v.value
}

func (v *NullableInlineResponse200127) Set(val *InlineResponse200127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200127(val *InlineResponse200127) *NullableInlineResponse200127 {
	return &NullableInlineResponse200127{value: val, isSet: true}
}

func (v NullableInlineResponse200127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


