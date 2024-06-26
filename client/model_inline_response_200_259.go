/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200259 struct for InlineResponse200259
type InlineResponse200259 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200258 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200258 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200259 instantiates a new InlineResponse200259 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200259() *InlineResponse200259 {
	this := InlineResponse200259{}
	return &this
}

// NewInlineResponse200259WithDefaults instantiates a new InlineResponse200259 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200259WithDefaults() *InlineResponse200259 {
	this := InlineResponse200259{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200259) GetRemainderLicenses() []InlineResponse200258 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200258
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetRemainderLicensesOk() ([]InlineResponse200258, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200259) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200258 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200259) SetRemainderLicenses(v []InlineResponse200258) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200259) GetMovedLicenses() []InlineResponse200258 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200258
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetMovedLicensesOk() ([]InlineResponse200258, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200259) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200258 and assigns it to the MovedLicenses field.
func (o *InlineResponse200259) SetMovedLicenses(v []InlineResponse200258) {
	o.MovedLicenses = v
}

func (o InlineResponse200259) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200259 struct {
	value *InlineResponse200259
	isSet bool
}

func (v NullableInlineResponse200259) Get() *InlineResponse200259 {
	return v.value
}

func (v *NullableInlineResponse200259) Set(val *InlineResponse200259) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200259) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200259) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200259(val *InlineResponse200259) *NullableInlineResponse200259 {
	return &NullableInlineResponse200259{value: val, isSet: true}
}

func (v NullableInlineResponse200259) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200259) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


