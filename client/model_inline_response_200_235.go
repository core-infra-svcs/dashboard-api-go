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

// InlineResponse200235 struct for InlineResponse200235
type InlineResponse200235 struct {
	// Organization Alert counts by type
	Items []InlineResponse200235Items `json:"items"`
	Meta InlineResponse200234Meta `json:"meta"`
}

// NewInlineResponse200235 instantiates a new InlineResponse200235 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200235(items []InlineResponse200235Items, meta InlineResponse200234Meta) *InlineResponse200235 {
	this := InlineResponse200235{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewInlineResponse200235WithDefaults instantiates a new InlineResponse200235 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200235WithDefaults() *InlineResponse200235 {
	this := InlineResponse200235{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineResponse200235) GetItems() []InlineResponse200235Items {
	if o == nil {
		var ret []InlineResponse200235Items
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200235) GetItemsOk() ([]InlineResponse200235Items, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InlineResponse200235) SetItems(v []InlineResponse200235Items) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *InlineResponse200235) GetMeta() InlineResponse200234Meta {
	if o == nil {
		var ret InlineResponse200234Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200235) GetMetaOk() (*InlineResponse200234Meta, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineResponse200235) SetMeta(v InlineResponse200234Meta) {
	o.Meta = v
}

func (o InlineResponse200235) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200235 struct {
	value *InlineResponse200235
	isSet bool
}

func (v NullableInlineResponse200235) Get() *InlineResponse200235 {
	return v.value
}

func (v *NullableInlineResponse200235) Set(val *InlineResponse200235) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200235) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200235) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200235(val *InlineResponse200235) *NullableInlineResponse200235 {
	return &NullableInlineResponse200235{value: val, isSet: true}
}

func (v NullableInlineResponse200235) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200235) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


