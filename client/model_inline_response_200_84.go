/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20084 struct for InlineResponse20084
type InlineResponse20084 struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewInlineResponse20084 instantiates a new InlineResponse20084 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20084() *InlineResponse20084 {
	this := InlineResponse20084{}
	return &this
}

// NewInlineResponse20084WithDefaults instantiates a new InlineResponse20084 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20084WithDefaults() *InlineResponse20084 {
	this := InlineResponse20084{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse20084) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20084) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse20084) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *InlineResponse20084) SetSuccess(v bool) {
	o.Success = &v
}

func (o InlineResponse20084) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20084 struct {
	value *InlineResponse20084
	isSet bool
}

func (v NullableInlineResponse20084) Get() *InlineResponse20084 {
	return v.value
}

func (v *NullableInlineResponse20084) Set(val *InlineResponse20084) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20084) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20084) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20084(val *InlineResponse20084) *NullableInlineResponse20084 {
	return &NullableInlineResponse20084{value: val, isSet: true}
}

func (v NullableInlineResponse20084) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20084) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


