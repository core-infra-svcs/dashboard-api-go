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

// InlineResponse200374 struct for InlineResponse200374
type InlineResponse200374 struct {
	// Wireless LAN controller L3 interfaces
	Items []InlineResponse200374Items `json:"items,omitempty"`
	Meta *InlineResponse200222Meta `json:"meta,omitempty"`
}

// NewInlineResponse200374 instantiates a new InlineResponse200374 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200374() *InlineResponse200374 {
	this := InlineResponse200374{}
	return &this
}

// NewInlineResponse200374WithDefaults instantiates a new InlineResponse200374 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200374WithDefaults() *InlineResponse200374 {
	this := InlineResponse200374{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200374) GetItems() []InlineResponse200374Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200374Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200374) GetItemsOk() ([]InlineResponse200374Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200374) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200374Items and assigns it to the Items field.
func (o *InlineResponse200374) SetItems(v []InlineResponse200374Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200374) GetMeta() InlineResponse200222Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200222Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200374) GetMetaOk() (*InlineResponse200222Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200374) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200222Meta and assigns it to the Meta field.
func (o *InlineResponse200374) SetMeta(v InlineResponse200222Meta) {
	o.Meta = &v
}

func (o InlineResponse200374) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200374 struct {
	value *InlineResponse200374
	isSet bool
}

func (v NullableInlineResponse200374) Get() *InlineResponse200374 {
	return v.value
}

func (v *NullableInlineResponse200374) Set(val *InlineResponse200374) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200374) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200374) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200374(val *InlineResponse200374) *NullableInlineResponse200374 {
	return &NullableInlineResponse200374{value: val, isSet: true}
}

func (v NullableInlineResponse200374) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200374) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


