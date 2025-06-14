/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200257 struct for InlineResponse200257
type InlineResponse200257 struct {
	// List of Cellular Service Provider Rate Plans
	Items []InlineResponse200257Items `json:"items,omitempty"`
	Meta *InlineResponse200257Meta `json:"meta,omitempty"`
}

// NewInlineResponse200257 instantiates a new InlineResponse200257 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200257() *InlineResponse200257 {
	this := InlineResponse200257{}
	return &this
}

// NewInlineResponse200257WithDefaults instantiates a new InlineResponse200257 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200257WithDefaults() *InlineResponse200257 {
	this := InlineResponse200257{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200257) GetItems() []InlineResponse200257Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200257Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200257) GetItemsOk() ([]InlineResponse200257Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200257) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200257Items and assigns it to the Items field.
func (o *InlineResponse200257) SetItems(v []InlineResponse200257Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200257) GetMeta() InlineResponse200257Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200257Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200257) GetMetaOk() (*InlineResponse200257Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200257) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200257Meta and assigns it to the Meta field.
func (o *InlineResponse200257) SetMeta(v InlineResponse200257Meta) {
	o.Meta = &v
}

func (o InlineResponse200257) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200257 struct {
	value *InlineResponse200257
	isSet bool
}

func (v NullableInlineResponse200257) Get() *InlineResponse200257 {
	return v.value
}

func (v *NullableInlineResponse200257) Set(val *InlineResponse200257) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200257) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200257) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200257(val *InlineResponse200257) *NullableInlineResponse200257 {
	return &NullableInlineResponse200257{value: val, isSet: true}
}

func (v NullableInlineResponse200257) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200257) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


