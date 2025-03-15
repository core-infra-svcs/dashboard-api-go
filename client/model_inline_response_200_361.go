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

// InlineResponse200361 struct for InlineResponse200361
type InlineResponse200361 struct {
	// Wireless LAN controller overview
	Items []InlineResponse200361Items `json:"items,omitempty"`
	Meta *InlineResponse200220Meta `json:"meta,omitempty"`
}

// NewInlineResponse200361 instantiates a new InlineResponse200361 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200361() *InlineResponse200361 {
	this := InlineResponse200361{}
	return &this
}

// NewInlineResponse200361WithDefaults instantiates a new InlineResponse200361 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200361WithDefaults() *InlineResponse200361 {
	this := InlineResponse200361{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200361) GetItems() []InlineResponse200361Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200361Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200361) GetItemsOk() ([]InlineResponse200361Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200361) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200361Items and assigns it to the Items field.
func (o *InlineResponse200361) SetItems(v []InlineResponse200361Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200361) GetMeta() InlineResponse200220Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200220Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200361) GetMetaOk() (*InlineResponse200220Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200361) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200220Meta and assigns it to the Meta field.
func (o *InlineResponse200361) SetMeta(v InlineResponse200220Meta) {
	o.Meta = &v
}

func (o InlineResponse200361) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200361 struct {
	value *InlineResponse200361
	isSet bool
}

func (v NullableInlineResponse200361) Get() *InlineResponse200361 {
	return v.value
}

func (v *NullableInlineResponse200361) Set(val *InlineResponse200361) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200361) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200361) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200361(val *InlineResponse200361) *NullableInlineResponse200361 {
	return &NullableInlineResponse200361{value: val, isSet: true}
}

func (v NullableInlineResponse200361) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200361) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


