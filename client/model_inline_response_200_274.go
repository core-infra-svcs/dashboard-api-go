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

// InlineResponse200274 struct for InlineResponse200274
type InlineResponse200274 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200273 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200273 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200274 instantiates a new InlineResponse200274 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// NewInlineResponse200274WithDefaults instantiates a new InlineResponse200274 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274WithDefaults() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200274) GetRemainderLicenses() []InlineResponse200273 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200273
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetRemainderLicensesOk() ([]InlineResponse200273, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200274) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200273 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200274) SetRemainderLicenses(v []InlineResponse200273) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200274) GetMovedLicenses() []InlineResponse200273 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200273
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetMovedLicensesOk() ([]InlineResponse200273, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200274) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200273 and assigns it to the MovedLicenses field.
func (o *InlineResponse200274) SetMovedLicenses(v []InlineResponse200273) {
	o.MovedLicenses = v
}

func (o InlineResponse200274) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274 struct {
	value *InlineResponse200274
	isSet bool
}

func (v NullableInlineResponse200274) Get() *InlineResponse200274 {
	return v.value
}

func (v *NullableInlineResponse200274) Set(val *InlineResponse200274) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274(val *InlineResponse200274) *NullableInlineResponse200274 {
	return &NullableInlineResponse200274{value: val, isSet: true}
}

func (v NullableInlineResponse200274) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


