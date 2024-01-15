/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200157 struct for InlineResponse200157
type InlineResponse200157 struct {
	// Resulting licenses from the move
	ResultingLicenses []InlineResponse200156 `json:"resultingLicenses,omitempty"`
}

// NewInlineResponse200157 instantiates a new InlineResponse200157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200157() *InlineResponse200157 {
	this := InlineResponse200157{}
	return &this
}

// NewInlineResponse200157WithDefaults instantiates a new InlineResponse200157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200157WithDefaults() *InlineResponse200157 {
	this := InlineResponse200157{}
	return &this
}

// GetResultingLicenses returns the ResultingLicenses field value if set, zero value otherwise.
func (o *InlineResponse200157) GetResultingLicenses() []InlineResponse200156 {
	if o == nil || isNil(o.ResultingLicenses) {
		var ret []InlineResponse200156
		return ret
	}
	return o.ResultingLicenses
}

// GetResultingLicensesOk returns a tuple with the ResultingLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157) GetResultingLicensesOk() ([]InlineResponse200156, bool) {
	if o == nil || isNil(o.ResultingLicenses) {
    return nil, false
	}
	return o.ResultingLicenses, true
}

// HasResultingLicenses returns a boolean if a field has been set.
func (o *InlineResponse200157) HasResultingLicenses() bool {
	if o != nil && !isNil(o.ResultingLicenses) {
		return true
	}

	return false
}

// SetResultingLicenses gets a reference to the given []InlineResponse200156 and assigns it to the ResultingLicenses field.
func (o *InlineResponse200157) SetResultingLicenses(v []InlineResponse200156) {
	o.ResultingLicenses = v
}

func (o InlineResponse200157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingLicenses) {
		toSerialize["resultingLicenses"] = o.ResultingLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200157 struct {
	value *InlineResponse200157
	isSet bool
}

func (v NullableInlineResponse200157) Get() *InlineResponse200157 {
	return v.value
}

func (v *NullableInlineResponse200157) Set(val *InlineResponse200157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200157(val *InlineResponse200157) *NullableInlineResponse200157 {
	return &NullableInlineResponse200157{value: val, isSet: true}
}

func (v NullableInlineResponse200157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


