/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20052 struct for InlineResponse20052
type InlineResponse20052 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
}

// NewInlineResponse20052 instantiates a new InlineResponse20052 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20052() *InlineResponse20052 {
	this := InlineResponse20052{}
	return &this
}

// NewInlineResponse20052WithDefaults instantiates a new InlineResponse20052 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20052WithDefaults() *InlineResponse20052 {
	this := InlineResponse20052{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse20052) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse20052) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse20052) SetIds(v []string) {
	o.Ids = v
}

func (o InlineResponse20052) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20052 struct {
	value *InlineResponse20052
	isSet bool
}

func (v NullableInlineResponse20052) Get() *InlineResponse20052 {
	return v.value
}

func (v *NullableInlineResponse20052) Set(val *InlineResponse20052) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20052) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20052) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20052(val *InlineResponse20052) *NullableInlineResponse20052 {
	return &NullableInlineResponse20052{value: val, isSet: true}
}

func (v NullableInlineResponse20052) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20052) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


