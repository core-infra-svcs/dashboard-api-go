/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200239 struct for InlineResponse200239
type InlineResponse200239 struct {
	// Organization Alert counts by type
	Items []InlineResponse200239Items `json:"items"`
	Meta InlineResponse200238Meta `json:"meta"`
}

// NewInlineResponse200239 instantiates a new InlineResponse200239 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200239(items []InlineResponse200239Items, meta InlineResponse200238Meta) *InlineResponse200239 {
	this := InlineResponse200239{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewInlineResponse200239WithDefaults instantiates a new InlineResponse200239 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200239WithDefaults() *InlineResponse200239 {
	this := InlineResponse200239{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineResponse200239) GetItems() []InlineResponse200239Items {
	if o == nil {
		var ret []InlineResponse200239Items
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200239) GetItemsOk() ([]InlineResponse200239Items, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InlineResponse200239) SetItems(v []InlineResponse200239Items) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *InlineResponse200239) GetMeta() InlineResponse200238Meta {
	if o == nil {
		var ret InlineResponse200238Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200239) GetMetaOk() (*InlineResponse200238Meta, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineResponse200239) SetMeta(v InlineResponse200238Meta) {
	o.Meta = v
}

func (o InlineResponse200239) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200239 struct {
	value *InlineResponse200239
	isSet bool
}

func (v NullableInlineResponse200239) Get() *InlineResponse200239 {
	return v.value
}

func (v *NullableInlineResponse200239) Set(val *InlineResponse200239) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200239) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200239) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200239(val *InlineResponse200239) *NullableInlineResponse200239 {
	return &NullableInlineResponse200239{value: val, isSet: true}
}

func (v NullableInlineResponse200239) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200239) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


