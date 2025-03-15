/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200346 struct for InlineResponse200346
type InlineResponse200346 struct {
	// The top-level propery containing all status data.
	Items []InlineResponse200346Items `json:"items,omitempty"`
	Meta *InlineResponse200346Meta `json:"meta,omitempty"`
}

// NewInlineResponse200346 instantiates a new InlineResponse200346 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200346() *InlineResponse200346 {
	this := InlineResponse200346{}
	return &this
}

// NewInlineResponse200346WithDefaults instantiates a new InlineResponse200346 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200346WithDefaults() *InlineResponse200346 {
	this := InlineResponse200346{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200346) GetItems() []InlineResponse200346Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200346Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346) GetItemsOk() ([]InlineResponse200346Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200346) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200346Items and assigns it to the Items field.
func (o *InlineResponse200346) SetItems(v []InlineResponse200346Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200346) GetMeta() InlineResponse200346Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200346Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346) GetMetaOk() (*InlineResponse200346Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200346) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200346Meta and assigns it to the Meta field.
func (o *InlineResponse200346) SetMeta(v InlineResponse200346Meta) {
	o.Meta = &v
}

func (o InlineResponse200346) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200346 struct {
	value *InlineResponse200346
	isSet bool
}

func (v NullableInlineResponse200346) Get() *InlineResponse200346 {
	return v.value
}

func (v *NullableInlineResponse200346) Set(val *InlineResponse200346) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200346) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200346) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200346(val *InlineResponse200346) *NullableInlineResponse200346 {
	return &NullableInlineResponse200346{value: val, isSet: true}
}

func (v NullableInlineResponse200346) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200346) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


