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

// InlineResponse200230 struct for InlineResponse200230
type InlineResponse200230 struct {
	// List of Cellular Service Provider Communication Plans
	Items []InlineResponse200230Items `json:"items,omitempty"`
	Meta *InlineResponse200230Meta `json:"meta,omitempty"`
}

// NewInlineResponse200230 instantiates a new InlineResponse200230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200230() *InlineResponse200230 {
	this := InlineResponse200230{}
	return &this
}

// NewInlineResponse200230WithDefaults instantiates a new InlineResponse200230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200230WithDefaults() *InlineResponse200230 {
	this := InlineResponse200230{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200230) GetItems() []InlineResponse200230Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200230Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230) GetItemsOk() ([]InlineResponse200230Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200230) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200230Items and assigns it to the Items field.
func (o *InlineResponse200230) SetItems(v []InlineResponse200230Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200230) GetMeta() InlineResponse200230Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200230Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230) GetMetaOk() (*InlineResponse200230Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200230) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200230Meta and assigns it to the Meta field.
func (o *InlineResponse200230) SetMeta(v InlineResponse200230Meta) {
	o.Meta = &v
}

func (o InlineResponse200230) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200230 struct {
	value *InlineResponse200230
	isSet bool
}

func (v NullableInlineResponse200230) Get() *InlineResponse200230 {
	return v.value
}

func (v *NullableInlineResponse200230) Set(val *InlineResponse200230) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200230) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200230(val *InlineResponse200230) *NullableInlineResponse200230 {
	return &NullableInlineResponse200230{value: val, isSet: true}
}

func (v NullableInlineResponse200230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


