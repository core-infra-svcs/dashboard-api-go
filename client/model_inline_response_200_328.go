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

// InlineResponse200328 struct for InlineResponse200328
type InlineResponse200328 struct {
	// List of rules
	Items []InlineResponse200328Items `json:"items,omitempty"`
	Meta *InlineResponse200328Meta `json:"meta,omitempty"`
}

// NewInlineResponse200328 instantiates a new InlineResponse200328 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200328() *InlineResponse200328 {
	this := InlineResponse200328{}
	return &this
}

// NewInlineResponse200328WithDefaults instantiates a new InlineResponse200328 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200328WithDefaults() *InlineResponse200328 {
	this := InlineResponse200328{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200328) GetItems() []InlineResponse200328Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200328Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200328) GetItemsOk() ([]InlineResponse200328Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200328) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200328Items and assigns it to the Items field.
func (o *InlineResponse200328) SetItems(v []InlineResponse200328Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200328) GetMeta() InlineResponse200328Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200328Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200328) GetMetaOk() (*InlineResponse200328Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200328) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200328Meta and assigns it to the Meta field.
func (o *InlineResponse200328) SetMeta(v InlineResponse200328Meta) {
	o.Meta = &v
}

func (o InlineResponse200328) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200328 struct {
	value *InlineResponse200328
	isSet bool
}

func (v NullableInlineResponse200328) Get() *InlineResponse200328 {
	return v.value
}

func (v *NullableInlineResponse200328) Set(val *InlineResponse200328) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200328) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200328) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200328(val *InlineResponse200328) *NullableInlineResponse200328 {
	return &NullableInlineResponse200328{value: val, isSet: true}
}

func (v NullableInlineResponse200328) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200328) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


