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

// InlineResponse200359 struct for InlineResponse200359
type InlineResponse200359 struct {
	// Wireless LAN controller overview
	Items []InlineResponse200359Items `json:"items,omitempty"`
	Meta *InlineResponse200219Meta `json:"meta,omitempty"`
}

// NewInlineResponse200359 instantiates a new InlineResponse200359 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200359() *InlineResponse200359 {
	this := InlineResponse200359{}
	return &this
}

// NewInlineResponse200359WithDefaults instantiates a new InlineResponse200359 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200359WithDefaults() *InlineResponse200359 {
	this := InlineResponse200359{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200359) GetItems() []InlineResponse200359Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200359Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359) GetItemsOk() ([]InlineResponse200359Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200359) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200359Items and assigns it to the Items field.
func (o *InlineResponse200359) SetItems(v []InlineResponse200359Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200359) GetMeta() InlineResponse200219Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200219Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359) GetMetaOk() (*InlineResponse200219Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200359) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200219Meta and assigns it to the Meta field.
func (o *InlineResponse200359) SetMeta(v InlineResponse200219Meta) {
	o.Meta = &v
}

func (o InlineResponse200359) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200359 struct {
	value *InlineResponse200359
	isSet bool
}

func (v NullableInlineResponse200359) Get() *InlineResponse200359 {
	return v.value
}

func (v *NullableInlineResponse200359) Set(val *InlineResponse200359) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200359) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200359) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200359(val *InlineResponse200359) *NullableInlineResponse200359 {
	return &NullableInlineResponse200359{value: val, isSet: true}
}

func (v NullableInlineResponse200359) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200359) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


