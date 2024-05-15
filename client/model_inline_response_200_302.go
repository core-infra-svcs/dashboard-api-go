/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200302 struct for InlineResponse200302
type InlineResponse200302 struct {
	// The top-level propery containing all status data.
	Items []InlineResponse200302Items `json:"items,omitempty"`
	Meta *InlineResponse200302Meta `json:"meta,omitempty"`
}

// NewInlineResponse200302 instantiates a new InlineResponse200302 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200302() *InlineResponse200302 {
	this := InlineResponse200302{}
	return &this
}

// NewInlineResponse200302WithDefaults instantiates a new InlineResponse200302 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200302WithDefaults() *InlineResponse200302 {
	this := InlineResponse200302{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200302) GetItems() []InlineResponse200302Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200302Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200302) GetItemsOk() ([]InlineResponse200302Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200302) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200302Items and assigns it to the Items field.
func (o *InlineResponse200302) SetItems(v []InlineResponse200302Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200302) GetMeta() InlineResponse200302Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200302Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200302) GetMetaOk() (*InlineResponse200302Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200302) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200302Meta and assigns it to the Meta field.
func (o *InlineResponse200302) SetMeta(v InlineResponse200302Meta) {
	o.Meta = &v
}

func (o InlineResponse200302) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200302 struct {
	value *InlineResponse200302
	isSet bool
}

func (v NullableInlineResponse200302) Get() *InlineResponse200302 {
	return v.value
}

func (v *NullableInlineResponse200302) Set(val *InlineResponse200302) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200302) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200302) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200302(val *InlineResponse200302) *NullableInlineResponse200302 {
	return &NullableInlineResponse200302{value: val, isSet: true}
}

func (v NullableInlineResponse200302) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200302) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


