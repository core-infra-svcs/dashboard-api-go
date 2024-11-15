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

// InlineResponse200212 struct for InlineResponse200212
type InlineResponse200212 struct {
	ResponseCodeCounts *InlineResponse200212ResponseCodeCounts `json:"responseCodeCounts,omitempty"`
}

// NewInlineResponse200212 instantiates a new InlineResponse200212 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200212() *InlineResponse200212 {
	this := InlineResponse200212{}
	return &this
}

// NewInlineResponse200212WithDefaults instantiates a new InlineResponse200212 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200212WithDefaults() *InlineResponse200212 {
	this := InlineResponse200212{}
	return &this
}

// GetResponseCodeCounts returns the ResponseCodeCounts field value if set, zero value otherwise.
func (o *InlineResponse200212) GetResponseCodeCounts() InlineResponse200212ResponseCodeCounts {
	if o == nil || isNil(o.ResponseCodeCounts) {
		var ret InlineResponse200212ResponseCodeCounts
		return ret
	}
	return *o.ResponseCodeCounts
}

// GetResponseCodeCountsOk returns a tuple with the ResponseCodeCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200212) GetResponseCodeCountsOk() (*InlineResponse200212ResponseCodeCounts, bool) {
	if o == nil || isNil(o.ResponseCodeCounts) {
    return nil, false
	}
	return o.ResponseCodeCounts, true
}

// HasResponseCodeCounts returns a boolean if a field has been set.
func (o *InlineResponse200212) HasResponseCodeCounts() bool {
	if o != nil && !isNil(o.ResponseCodeCounts) {
		return true
	}

	return false
}

// SetResponseCodeCounts gets a reference to the given InlineResponse200212ResponseCodeCounts and assigns it to the ResponseCodeCounts field.
func (o *InlineResponse200212) SetResponseCodeCounts(v InlineResponse200212ResponseCodeCounts) {
	o.ResponseCodeCounts = &v
}

func (o InlineResponse200212) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResponseCodeCounts) {
		toSerialize["responseCodeCounts"] = o.ResponseCodeCounts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200212 struct {
	value *InlineResponse200212
	isSet bool
}

func (v NullableInlineResponse200212) Get() *InlineResponse200212 {
	return v.value
}

func (v *NullableInlineResponse200212) Set(val *InlineResponse200212) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200212) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200212) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200212(val *InlineResponse200212) *NullableInlineResponse200212 {
	return &NullableInlineResponse200212{value: val, isSet: true}
}

func (v NullableInlineResponse200212) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200212) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


