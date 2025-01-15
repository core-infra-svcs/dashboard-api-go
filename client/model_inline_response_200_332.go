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

// InlineResponse200332 struct for InlineResponse200332
type InlineResponse200332 struct {
	// Overview history of wireless LAN controllers
	Items []InlineResponse200332Items `json:"items,omitempty"`
	Meta *InlineResponse200250Meta `json:"meta,omitempty"`
}

// NewInlineResponse200332 instantiates a new InlineResponse200332 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200332() *InlineResponse200332 {
	this := InlineResponse200332{}
	return &this
}

// NewInlineResponse200332WithDefaults instantiates a new InlineResponse200332 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200332WithDefaults() *InlineResponse200332 {
	this := InlineResponse200332{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200332) GetItems() []InlineResponse200332Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200332Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200332) GetItemsOk() ([]InlineResponse200332Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200332) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200332Items and assigns it to the Items field.
func (o *InlineResponse200332) SetItems(v []InlineResponse200332Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200332) GetMeta() InlineResponse200250Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200250Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200332) GetMetaOk() (*InlineResponse200250Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200332) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200250Meta and assigns it to the Meta field.
func (o *InlineResponse200332) SetMeta(v InlineResponse200250Meta) {
	o.Meta = &v
}

func (o InlineResponse200332) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200332 struct {
	value *InlineResponse200332
	isSet bool
}

func (v NullableInlineResponse200332) Get() *InlineResponse200332 {
	return v.value
}

func (v *NullableInlineResponse200332) Set(val *InlineResponse200332) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200332) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200332) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200332(val *InlineResponse200332) *NullableInlineResponse200332 {
	return &NullableInlineResponse200332{value: val, isSet: true}
}

func (v NullableInlineResponse200332) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200332) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


