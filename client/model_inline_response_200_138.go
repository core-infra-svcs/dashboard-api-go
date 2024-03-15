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

// InlineResponse200138 struct for InlineResponse200138
type InlineResponse200138 struct {
	ResponseCodeCounts *InlineResponse200138ResponseCodeCounts `json:"responseCodeCounts,omitempty"`
}

// NewInlineResponse200138 instantiates a new InlineResponse200138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200138() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// NewInlineResponse200138WithDefaults instantiates a new InlineResponse200138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200138WithDefaults() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// GetResponseCodeCounts returns the ResponseCodeCounts field value if set, zero value otherwise.
func (o *InlineResponse200138) GetResponseCodeCounts() InlineResponse200138ResponseCodeCounts {
	if o == nil || isNil(o.ResponseCodeCounts) {
		var ret InlineResponse200138ResponseCodeCounts
		return ret
	}
	return *o.ResponseCodeCounts
}

// GetResponseCodeCountsOk returns a tuple with the ResponseCodeCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200138) GetResponseCodeCountsOk() (*InlineResponse200138ResponseCodeCounts, bool) {
	if o == nil || isNil(o.ResponseCodeCounts) {
    return nil, false
	}
	return o.ResponseCodeCounts, true
}

// HasResponseCodeCounts returns a boolean if a field has been set.
func (o *InlineResponse200138) HasResponseCodeCounts() bool {
	if o != nil && !isNil(o.ResponseCodeCounts) {
		return true
	}

	return false
}

// SetResponseCodeCounts gets a reference to the given InlineResponse200138ResponseCodeCounts and assigns it to the ResponseCodeCounts field.
func (o *InlineResponse200138) SetResponseCodeCounts(v InlineResponse200138ResponseCodeCounts) {
	o.ResponseCodeCounts = &v
}

func (o InlineResponse200138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResponseCodeCounts) {
		toSerialize["responseCodeCounts"] = o.ResponseCodeCounts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200138 struct {
	value *InlineResponse200138
	isSet bool
}

func (v NullableInlineResponse200138) Get() *InlineResponse200138 {
	return v.value
}

func (v *NullableInlineResponse200138) Set(val *InlineResponse200138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200138(val *InlineResponse200138) *NullableInlineResponse200138 {
	return &NullableInlineResponse200138{value: val, isSet: true}
}

func (v NullableInlineResponse200138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


