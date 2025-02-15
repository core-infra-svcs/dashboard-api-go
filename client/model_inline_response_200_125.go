/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200125 struct for InlineResponse200125
type InlineResponse200125 struct {
	// The Meraki Ids of the set of endpoints.
	Ids []string `json:"ids,omitempty"`
}

// NewInlineResponse200125 instantiates a new InlineResponse200125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200125() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200125WithDefaults() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse200125) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200125) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse200125) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse200125) SetIds(v []string) {
	o.Ids = v
}

func (o InlineResponse200125) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200125 struct {
	value *InlineResponse200125
	isSet bool
}

func (v NullableInlineResponse200125) Get() *InlineResponse200125 {
	return v.value
}

func (v *NullableInlineResponse200125) Set(val *InlineResponse200125) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200125) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200125(val *InlineResponse200125) *NullableInlineResponse200125 {
	return &NullableInlineResponse200125{value: val, isSet: true}
}

func (v NullableInlineResponse200125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


