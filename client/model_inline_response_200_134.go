/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200134 struct for InlineResponse200134
type InlineResponse200134 struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse200134 instantiates a new InlineResponse200134 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200134() *InlineResponse200134 {
	this := InlineResponse200134{}
	return &this
}

// NewInlineResponse200134WithDefaults instantiates a new InlineResponse200134 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200134WithDefaults() *InlineResponse200134 {
	this := InlineResponse200134{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse200134) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse200134) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse200134) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse200134) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200134 struct {
	value *InlineResponse200134
	isSet bool
}

func (v NullableInlineResponse200134) Get() *InlineResponse200134 {
	return v.value
}

func (v *NullableInlineResponse200134) Set(val *InlineResponse200134) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200134) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200134) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200134(val *InlineResponse200134) *NullableInlineResponse200134 {
	return &NullableInlineResponse200134{value: val, isSet: true}
}

func (v NullableInlineResponse200134) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200134) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


