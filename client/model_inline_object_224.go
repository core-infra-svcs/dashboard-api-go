/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject224 struct for InlineObject224
type InlineObject224 struct {
	// Status the eSIM will be updated to
	Status *string `json:"status,omitempty"`
}

// NewInlineObject224 instantiates a new InlineObject224 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject224() *InlineObject224 {
	this := InlineObject224{}
	return &this
}

// NewInlineObject224WithDefaults instantiates a new InlineObject224 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject224WithDefaults() *InlineObject224 {
	this := InlineObject224{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineObject224) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineObject224) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineObject224) SetStatus(v string) {
	o.Status = &v
}

func (o InlineObject224) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject224 struct {
	value *InlineObject224
	isSet bool
}

func (v NullableInlineObject224) Get() *InlineObject224 {
	return v.value
}

func (v *NullableInlineObject224) Set(val *InlineObject224) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject224) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject224) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject224(val *InlineObject224) *NullableInlineObject224 {
	return &NullableInlineObject224{value: val, isSet: true}
}

func (v NullableInlineObject224) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject224) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


