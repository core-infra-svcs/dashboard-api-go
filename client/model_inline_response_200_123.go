/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200123 struct for InlineResponse200123
type InlineResponse200123 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200122 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200122 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200123 instantiates a new InlineResponse200123 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200123() *InlineResponse200123 {
	this := InlineResponse200123{}
	return &this
}

// NewInlineResponse200123WithDefaults instantiates a new InlineResponse200123 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200123WithDefaults() *InlineResponse200123 {
	this := InlineResponse200123{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200123) GetRemainderLicenses() []InlineResponse200122 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200122
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200123) GetRemainderLicensesOk() ([]InlineResponse200122, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200123) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200122 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200123) SetRemainderLicenses(v []InlineResponse200122) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200123) GetMovedLicenses() []InlineResponse200122 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200122
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200123) GetMovedLicensesOk() ([]InlineResponse200122, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200123) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200122 and assigns it to the MovedLicenses field.
func (o *InlineResponse200123) SetMovedLicenses(v []InlineResponse200122) {
	o.MovedLicenses = v
}

func (o InlineResponse200123) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200123 struct {
	value *InlineResponse200123
	isSet bool
}

func (v NullableInlineResponse200123) Get() *InlineResponse200123 {
	return v.value
}

func (v *NullableInlineResponse200123) Set(val *InlineResponse200123) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200123) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200123) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200123(val *InlineResponse200123) *NullableInlineResponse200123 {
	return &NullableInlineResponse200123{value: val, isSet: true}
}

func (v NullableInlineResponse200123) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200123) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


