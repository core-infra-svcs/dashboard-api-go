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

// InlineResponse200330 struct for InlineResponse200330
type InlineResponse200330 struct {
	// Switches
	Items []InlineResponse200330Items `json:"items,omitempty"`
	Meta *InlineResponse200222Meta `json:"meta,omitempty"`
}

// NewInlineResponse200330 instantiates a new InlineResponse200330 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200330() *InlineResponse200330 {
	this := InlineResponse200330{}
	return &this
}

// NewInlineResponse200330WithDefaults instantiates a new InlineResponse200330 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200330WithDefaults() *InlineResponse200330 {
	this := InlineResponse200330{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200330) GetItems() []InlineResponse200330Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200330Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330) GetItemsOk() ([]InlineResponse200330Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200330) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200330Items and assigns it to the Items field.
func (o *InlineResponse200330) SetItems(v []InlineResponse200330Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200330) GetMeta() InlineResponse200222Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200222Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330) GetMetaOk() (*InlineResponse200222Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200330) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200222Meta and assigns it to the Meta field.
func (o *InlineResponse200330) SetMeta(v InlineResponse200222Meta) {
	o.Meta = &v
}

func (o InlineResponse200330) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200330 struct {
	value *InlineResponse200330
	isSet bool
}

func (v NullableInlineResponse200330) Get() *InlineResponse200330 {
	return v.value
}

func (v *NullableInlineResponse200330) Set(val *InlineResponse200330) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200330) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200330) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200330(val *InlineResponse200330) *NullableInlineResponse200330 {
	return &NullableInlineResponse200330{value: val, isSet: true}
}

func (v NullableInlineResponse200330) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200330) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


