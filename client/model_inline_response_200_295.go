/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200295 struct for InlineResponse200295
type InlineResponse200295 struct {
	// List of rules
	Items []InlineResponse200295Items `json:"items,omitempty"`
	Meta *InlineResponse200295Meta `json:"meta,omitempty"`
}

// NewInlineResponse200295 instantiates a new InlineResponse200295 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200295() *InlineResponse200295 {
	this := InlineResponse200295{}
	return &this
}

// NewInlineResponse200295WithDefaults instantiates a new InlineResponse200295 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200295WithDefaults() *InlineResponse200295 {
	this := InlineResponse200295{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200295) GetItems() []InlineResponse200295Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200295Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetItemsOk() ([]InlineResponse200295Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200295) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200295Items and assigns it to the Items field.
func (o *InlineResponse200295) SetItems(v []InlineResponse200295Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200295) GetMeta() InlineResponse200295Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200295Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetMetaOk() (*InlineResponse200295Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200295) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200295Meta and assigns it to the Meta field.
func (o *InlineResponse200295) SetMeta(v InlineResponse200295Meta) {
	o.Meta = &v
}

func (o InlineResponse200295) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200295 struct {
	value *InlineResponse200295
	isSet bool
}

func (v NullableInlineResponse200295) Get() *InlineResponse200295 {
	return v.value
}

func (v *NullableInlineResponse200295) Set(val *InlineResponse200295) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200295) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200295) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200295(val *InlineResponse200295) *NullableInlineResponse200295 {
	return &NullableInlineResponse200295{value: val, isSet: true}
}

func (v NullableInlineResponse200295) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200295) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


