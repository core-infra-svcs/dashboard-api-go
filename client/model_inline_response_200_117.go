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

// InlineResponse200117 struct for InlineResponse200117
type InlineResponse200117 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
}

// NewInlineResponse200117 instantiates a new InlineResponse200117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200117() *InlineResponse200117 {
	this := InlineResponse200117{}
	return &this
}

// NewInlineResponse200117WithDefaults instantiates a new InlineResponse200117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200117WithDefaults() *InlineResponse200117 {
	this := InlineResponse200117{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse200117) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200117) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse200117) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse200117) SetIds(v []string) {
	o.Ids = v
}

func (o InlineResponse200117) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200117 struct {
	value *InlineResponse200117
	isSet bool
}

func (v NullableInlineResponse200117) Get() *InlineResponse200117 {
	return v.value
}

func (v *NullableInlineResponse200117) Set(val *InlineResponse200117) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200117) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200117(val *InlineResponse200117) *NullableInlineResponse200117 {
	return &NullableInlineResponse200117{value: val, isSet: true}
}

func (v NullableInlineResponse200117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


