/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200296 struct for InlineResponse200296
type InlineResponse200296 struct {
	// List of rules
	Items []InlineResponse200296Items `json:"items,omitempty"`
	Meta *InlineResponse200296Meta `json:"meta,omitempty"`
}

// NewInlineResponse200296 instantiates a new InlineResponse200296 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200296() *InlineResponse200296 {
	this := InlineResponse200296{}
	return &this
}

// NewInlineResponse200296WithDefaults instantiates a new InlineResponse200296 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200296WithDefaults() *InlineResponse200296 {
	this := InlineResponse200296{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200296) GetItems() []InlineResponse200296Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200296Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetItemsOk() ([]InlineResponse200296Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200296) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200296Items and assigns it to the Items field.
func (o *InlineResponse200296) SetItems(v []InlineResponse200296Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200296) GetMeta() InlineResponse200296Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200296Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetMetaOk() (*InlineResponse200296Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200296) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200296Meta and assigns it to the Meta field.
func (o *InlineResponse200296) SetMeta(v InlineResponse200296Meta) {
	o.Meta = &v
}

func (o InlineResponse200296) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200296 struct {
	value *InlineResponse200296
	isSet bool
}

func (v NullableInlineResponse200296) Get() *InlineResponse200296 {
	return v.value
}

func (v *NullableInlineResponse200296) Set(val *InlineResponse200296) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200296) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200296) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200296(val *InlineResponse200296) *NullableInlineResponse200296 {
	return &NullableInlineResponse200296{value: val, isSet: true}
}

func (v NullableInlineResponse200296) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200296) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


