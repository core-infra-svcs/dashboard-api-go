/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200122 struct for InlineResponse200122
type InlineResponse200122 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
}

// NewInlineResponse200122 instantiates a new InlineResponse200122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200122() *InlineResponse200122 {
	this := InlineResponse200122{}
	return &this
}

// NewInlineResponse200122WithDefaults instantiates a new InlineResponse200122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200122WithDefaults() *InlineResponse200122 {
	this := InlineResponse200122{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse200122) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse200122) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse200122) SetIds(v []string) {
	o.Ids = v
}

func (o InlineResponse200122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200122 struct {
	value *InlineResponse200122
	isSet bool
}

func (v NullableInlineResponse200122) Get() *InlineResponse200122 {
	return v.value
}

func (v *NullableInlineResponse200122) Set(val *InlineResponse200122) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200122) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200122(val *InlineResponse200122) *NullableInlineResponse200122 {
	return &NullableInlineResponse200122{value: val, isSet: true}
}

func (v NullableInlineResponse200122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


