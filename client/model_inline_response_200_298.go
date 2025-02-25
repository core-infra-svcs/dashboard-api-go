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

// InlineResponse200298 struct for InlineResponse200298
type InlineResponse200298 struct {
	// Array of Limited Access Roles
	Items []InlineResponse200298Items `json:"items,omitempty"`
	Meta *InlineResponse200219Meta `json:"meta,omitempty"`
}

// NewInlineResponse200298 instantiates a new InlineResponse200298 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200298() *InlineResponse200298 {
	this := InlineResponse200298{}
	return &this
}

// NewInlineResponse200298WithDefaults instantiates a new InlineResponse200298 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200298WithDefaults() *InlineResponse200298 {
	this := InlineResponse200298{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200298) GetItems() []InlineResponse200298Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200298Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetItemsOk() ([]InlineResponse200298Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200298) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200298Items and assigns it to the Items field.
func (o *InlineResponse200298) SetItems(v []InlineResponse200298Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200298) GetMeta() InlineResponse200219Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200219Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298) GetMetaOk() (*InlineResponse200219Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200298) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200219Meta and assigns it to the Meta field.
func (o *InlineResponse200298) SetMeta(v InlineResponse200219Meta) {
	o.Meta = &v
}

func (o InlineResponse200298) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200298 struct {
	value *InlineResponse200298
	isSet bool
}

func (v NullableInlineResponse200298) Get() *InlineResponse200298 {
	return v.value
}

func (v *NullableInlineResponse200298) Set(val *InlineResponse200298) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200298) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200298) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200298(val *InlineResponse200298) *NullableInlineResponse200298 {
	return &NullableInlineResponse200298{value: val, isSet: true}
}

func (v NullableInlineResponse200298) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200298) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


