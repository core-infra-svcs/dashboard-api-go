/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200158 struct for InlineResponse200158
type InlineResponse200158 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200157 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200157 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200158 instantiates a new InlineResponse200158 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200158() *InlineResponse200158 {
	this := InlineResponse200158{}
	return &this
}

// NewInlineResponse200158WithDefaults instantiates a new InlineResponse200158 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200158WithDefaults() *InlineResponse200158 {
	this := InlineResponse200158{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200158) GetRemainderLicenses() []InlineResponse200157 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200157
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200158) GetRemainderLicensesOk() ([]InlineResponse200157, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200158) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200157 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200158) SetRemainderLicenses(v []InlineResponse200157) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200158) GetMovedLicenses() []InlineResponse200157 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200157
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200158) GetMovedLicensesOk() ([]InlineResponse200157, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200158) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200157 and assigns it to the MovedLicenses field.
func (o *InlineResponse200158) SetMovedLicenses(v []InlineResponse200157) {
	o.MovedLicenses = v
}

func (o InlineResponse200158) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200158 struct {
	value *InlineResponse200158
	isSet bool
}

func (v NullableInlineResponse200158) Get() *InlineResponse200158 {
	return v.value
}

func (v *NullableInlineResponse200158) Set(val *InlineResponse200158) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200158) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200158) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200158(val *InlineResponse200158) *NullableInlineResponse200158 {
	return &NullableInlineResponse200158{value: val, isSet: true}
}

func (v NullableInlineResponse200158) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200158) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


