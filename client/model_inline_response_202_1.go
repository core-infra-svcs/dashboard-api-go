/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2021 struct for InlineResponse2021
type InlineResponse2021 struct {
	// Shows the success of the reboot
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse2021 instantiates a new InlineResponse2021 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2021() *InlineResponse2021 {
	this := InlineResponse2021{}
	return &this
}

// NewInlineResponse2021WithDefaults instantiates a new InlineResponse2021 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2021WithDefaults() *InlineResponse2021 {
	this := InlineResponse2021{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse2021) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2021) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse2021) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse2021) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse2021) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2021 struct {
	value *InlineResponse2021
	isSet bool
}

func (v NullableInlineResponse2021) Get() *InlineResponse2021 {
	return v.value
}

func (v *NullableInlineResponse2021) Set(val *InlineResponse2021) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2021) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2021) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2021(val *InlineResponse2021) *NullableInlineResponse2021 {
	return &NullableInlineResponse2021{value: val, isSet: true}
}

func (v NullableInlineResponse2021) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2021) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


