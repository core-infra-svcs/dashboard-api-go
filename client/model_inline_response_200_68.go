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

// InlineResponse20068 struct for InlineResponse20068
type InlineResponse20068 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
}

// NewInlineResponse20068 instantiates a new InlineResponse20068 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// NewInlineResponse20068WithDefaults instantiates a new InlineResponse20068 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068WithDefaults() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse20068) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse20068) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse20068) SetIds(v []string) {
	o.Ids = v
}

func (o InlineResponse20068) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20068 struct {
	value *InlineResponse20068
	isSet bool
}

func (v NullableInlineResponse20068) Get() *InlineResponse20068 {
	return v.value
}

func (v *NullableInlineResponse20068) Set(val *InlineResponse20068) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068(val *InlineResponse20068) *NullableInlineResponse20068 {
	return &NullableInlineResponse20068{value: val, isSet: true}
}

func (v NullableInlineResponse20068) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


