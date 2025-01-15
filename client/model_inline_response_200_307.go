/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200307 struct for InlineResponse200307
type InlineResponse200307 struct {
	// Switches
	Items []InlineResponse200307Items `json:"items,omitempty"`
	Meta *InlineResponse200250Meta `json:"meta,omitempty"`
}

// NewInlineResponse200307 instantiates a new InlineResponse200307 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200307() *InlineResponse200307 {
	this := InlineResponse200307{}
	return &this
}

// NewInlineResponse200307WithDefaults instantiates a new InlineResponse200307 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200307WithDefaults() *InlineResponse200307 {
	this := InlineResponse200307{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200307) GetItems() []InlineResponse200307Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200307Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307) GetItemsOk() ([]InlineResponse200307Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200307) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200307Items and assigns it to the Items field.
func (o *InlineResponse200307) SetItems(v []InlineResponse200307Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200307) GetMeta() InlineResponse200250Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200250Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307) GetMetaOk() (*InlineResponse200250Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200307) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200250Meta and assigns it to the Meta field.
func (o *InlineResponse200307) SetMeta(v InlineResponse200250Meta) {
	o.Meta = &v
}

func (o InlineResponse200307) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200307 struct {
	value *InlineResponse200307
	isSet bool
}

func (v NullableInlineResponse200307) Get() *InlineResponse200307 {
	return v.value
}

func (v *NullableInlineResponse200307) Set(val *InlineResponse200307) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200307) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200307) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200307(val *InlineResponse200307) *NullableInlineResponse200307 {
	return &NullableInlineResponse200307{value: val, isSet: true}
}

func (v NullableInlineResponse200307) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200307) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


