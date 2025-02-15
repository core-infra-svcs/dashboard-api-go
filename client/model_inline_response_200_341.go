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

// InlineResponse200341 struct for InlineResponse200341
type InlineResponse200341 struct {
	// List of Catalyst access points information
	Items []InlineResponse200341Items `json:"items,omitempty"`
	Meta *InlineResponse200219Meta `json:"meta,omitempty"`
}

// NewInlineResponse200341 instantiates a new InlineResponse200341 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200341() *InlineResponse200341 {
	this := InlineResponse200341{}
	return &this
}

// NewInlineResponse200341WithDefaults instantiates a new InlineResponse200341 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200341WithDefaults() *InlineResponse200341 {
	this := InlineResponse200341{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200341) GetItems() []InlineResponse200341Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200341Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341) GetItemsOk() ([]InlineResponse200341Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200341) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200341Items and assigns it to the Items field.
func (o *InlineResponse200341) SetItems(v []InlineResponse200341Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200341) GetMeta() InlineResponse200219Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200219Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341) GetMetaOk() (*InlineResponse200219Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200341) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200219Meta and assigns it to the Meta field.
func (o *InlineResponse200341) SetMeta(v InlineResponse200219Meta) {
	o.Meta = &v
}

func (o InlineResponse200341) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200341 struct {
	value *InlineResponse200341
	isSet bool
}

func (v NullableInlineResponse200341) Get() *InlineResponse200341 {
	return v.value
}

func (v *NullableInlineResponse200341) Set(val *InlineResponse200341) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200341) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200341) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200341(val *InlineResponse200341) *NullableInlineResponse200341 {
	return &NullableInlineResponse200341{value: val, isSet: true}
}

func (v NullableInlineResponse200341) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200341) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


