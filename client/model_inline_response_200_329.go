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

// InlineResponse200329 struct for InlineResponse200329
type InlineResponse200329 struct {
	// List of settings
	Items []InlineResponse200177 `json:"items,omitempty"`
	Meta *InlineResponse200329Meta `json:"meta,omitempty"`
}

// NewInlineResponse200329 instantiates a new InlineResponse200329 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200329() *InlineResponse200329 {
	this := InlineResponse200329{}
	return &this
}

// NewInlineResponse200329WithDefaults instantiates a new InlineResponse200329 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200329WithDefaults() *InlineResponse200329 {
	this := InlineResponse200329{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200329) GetItems() []InlineResponse200177 {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200177
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329) GetItemsOk() ([]InlineResponse200177, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200329) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200177 and assigns it to the Items field.
func (o *InlineResponse200329) SetItems(v []InlineResponse200177) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200329) GetMeta() InlineResponse200329Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200329Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329) GetMetaOk() (*InlineResponse200329Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200329) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200329Meta and assigns it to the Meta field.
func (o *InlineResponse200329) SetMeta(v InlineResponse200329Meta) {
	o.Meta = &v
}

func (o InlineResponse200329) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200329 struct {
	value *InlineResponse200329
	isSet bool
}

func (v NullableInlineResponse200329) Get() *InlineResponse200329 {
	return v.value
}

func (v *NullableInlineResponse200329) Set(val *InlineResponse200329) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200329) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200329) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200329(val *InlineResponse200329) *NullableInlineResponse200329 {
	return &NullableInlineResponse200329{value: val, isSet: true}
}

func (v NullableInlineResponse200329) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200329) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


