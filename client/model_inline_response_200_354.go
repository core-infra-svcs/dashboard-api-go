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

// InlineResponse200354 struct for InlineResponse200354
type InlineResponse200354 struct {
	// Paginated list of scanning api receivers by network ID.
	Items []InlineResponse200354Items `json:"items,omitempty"`
	Meta *InlineResponse200222Meta `json:"meta,omitempty"`
}

// NewInlineResponse200354 instantiates a new InlineResponse200354 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200354() *InlineResponse200354 {
	this := InlineResponse200354{}
	return &this
}

// NewInlineResponse200354WithDefaults instantiates a new InlineResponse200354 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200354WithDefaults() *InlineResponse200354 {
	this := InlineResponse200354{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200354) GetItems() []InlineResponse200354Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200354Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetItemsOk() ([]InlineResponse200354Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200354) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200354Items and assigns it to the Items field.
func (o *InlineResponse200354) SetItems(v []InlineResponse200354Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200354) GetMeta() InlineResponse200222Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200222Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetMetaOk() (*InlineResponse200222Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200354) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200222Meta and assigns it to the Meta field.
func (o *InlineResponse200354) SetMeta(v InlineResponse200222Meta) {
	o.Meta = &v
}

func (o InlineResponse200354) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200354 struct {
	value *InlineResponse200354
	isSet bool
}

func (v NullableInlineResponse200354) Get() *InlineResponse200354 {
	return v.value
}

func (v *NullableInlineResponse200354) Set(val *InlineResponse200354) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200354) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200354) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200354(val *InlineResponse200354) *NullableInlineResponse200354 {
	return &NullableInlineResponse200354{value: val, isSet: true}
}

func (v NullableInlineResponse200354) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200354) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


