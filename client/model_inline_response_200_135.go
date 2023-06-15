/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200135 struct for InlineResponse200135
type InlineResponse200135 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200134 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200134 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200135 instantiates a new InlineResponse200135 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200135() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// NewInlineResponse200135WithDefaults instantiates a new InlineResponse200135 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200135WithDefaults() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200135) GetRemainderLicenses() []InlineResponse200134 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200134
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetRemainderLicensesOk() ([]InlineResponse200134, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200135) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200134 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200135) SetRemainderLicenses(v []InlineResponse200134) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200135) GetMovedLicenses() []InlineResponse200134 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200134
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetMovedLicensesOk() ([]InlineResponse200134, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200135) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200134 and assigns it to the MovedLicenses field.
func (o *InlineResponse200135) SetMovedLicenses(v []InlineResponse200134) {
	o.MovedLicenses = v
}

func (o InlineResponse200135) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200135 struct {
	value *InlineResponse200135
	isSet bool
}

func (v NullableInlineResponse200135) Get() *InlineResponse200135 {
	return v.value
}

func (v *NullableInlineResponse200135) Set(val *InlineResponse200135) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200135) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200135) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200135(val *InlineResponse200135) *NullableInlineResponse200135 {
	return &NullableInlineResponse200135{value: val, isSet: true}
}

func (v NullableInlineResponse200135) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200135) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


