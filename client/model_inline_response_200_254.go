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

// InlineResponse200254 struct for InlineResponse200254
type InlineResponse200254 struct {
	// Resulting licenses from the move
	ResultingLicenses []InlineResponse200253 `json:"resultingLicenses,omitempty"`
}

// NewInlineResponse200254 instantiates a new InlineResponse200254 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200254() *InlineResponse200254 {
	this := InlineResponse200254{}
	return &this
}

// NewInlineResponse200254WithDefaults instantiates a new InlineResponse200254 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200254WithDefaults() *InlineResponse200254 {
	this := InlineResponse200254{}
	return &this
}

// GetResultingLicenses returns the ResultingLicenses field value if set, zero value otherwise.
func (o *InlineResponse200254) GetResultingLicenses() []InlineResponse200253 {
	if o == nil || isNil(o.ResultingLicenses) {
		var ret []InlineResponse200253
		return ret
	}
	return o.ResultingLicenses
}

// GetResultingLicensesOk returns a tuple with the ResultingLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200254) GetResultingLicensesOk() ([]InlineResponse200253, bool) {
	if o == nil || isNil(o.ResultingLicenses) {
    return nil, false
	}
	return o.ResultingLicenses, true
}

// HasResultingLicenses returns a boolean if a field has been set.
func (o *InlineResponse200254) HasResultingLicenses() bool {
	if o != nil && !isNil(o.ResultingLicenses) {
		return true
	}

	return false
}

// SetResultingLicenses gets a reference to the given []InlineResponse200253 and assigns it to the ResultingLicenses field.
func (o *InlineResponse200254) SetResultingLicenses(v []InlineResponse200253) {
	o.ResultingLicenses = v
}

func (o InlineResponse200254) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingLicenses) {
		toSerialize["resultingLicenses"] = o.ResultingLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200254 struct {
	value *InlineResponse200254
	isSet bool
}

func (v NullableInlineResponse200254) Get() *InlineResponse200254 {
	return v.value
}

func (v *NullableInlineResponse200254) Set(val *InlineResponse200254) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200254) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200254) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200254(val *InlineResponse200254) *NullableInlineResponse200254 {
	return &NullableInlineResponse200254{value: val, isSet: true}
}

func (v NullableInlineResponse200254) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200254) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


