/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200280 struct for InlineResponse200280
type InlineResponse200280 struct {
	// Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to
	Items []InlineResponse200279Items `json:"items,omitempty"`
	Meta *InlineResponse200277Meta `json:"meta,omitempty"`
}

// NewInlineResponse200280 instantiates a new InlineResponse200280 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200280() *InlineResponse200280 {
	this := InlineResponse200280{}
	return &this
}

// NewInlineResponse200280WithDefaults instantiates a new InlineResponse200280 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200280WithDefaults() *InlineResponse200280 {
	this := InlineResponse200280{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200280) GetItems() []InlineResponse200279Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200279Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280) GetItemsOk() ([]InlineResponse200279Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200280) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200279Items and assigns it to the Items field.
func (o *InlineResponse200280) SetItems(v []InlineResponse200279Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200280) GetMeta() InlineResponse200277Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200277Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280) GetMetaOk() (*InlineResponse200277Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200280) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200277Meta and assigns it to the Meta field.
func (o *InlineResponse200280) SetMeta(v InlineResponse200277Meta) {
	o.Meta = &v
}

func (o InlineResponse200280) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200280 struct {
	value *InlineResponse200280
	isSet bool
}

func (v NullableInlineResponse200280) Get() *InlineResponse200280 {
	return v.value
}

func (v *NullableInlineResponse200280) Set(val *InlineResponse200280) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200280) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200280) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200280(val *InlineResponse200280) *NullableInlineResponse200280 {
	return &NullableInlineResponse200280{value: val, isSet: true}
}

func (v NullableInlineResponse200280) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200280) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


