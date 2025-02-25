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

// InlineResponse200301 struct for InlineResponse200301
type InlineResponse200301 struct {
	// Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to
	Items []InlineResponse200300Items `json:"items,omitempty"`
	Meta *InlineResponse200219Meta `json:"meta,omitempty"`
}

// NewInlineResponse200301 instantiates a new InlineResponse200301 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200301() *InlineResponse200301 {
	this := InlineResponse200301{}
	return &this
}

// NewInlineResponse200301WithDefaults instantiates a new InlineResponse200301 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200301WithDefaults() *InlineResponse200301 {
	this := InlineResponse200301{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200301) GetItems() []InlineResponse200300Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200300Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200301) GetItemsOk() ([]InlineResponse200300Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200301) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200300Items and assigns it to the Items field.
func (o *InlineResponse200301) SetItems(v []InlineResponse200300Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200301) GetMeta() InlineResponse200219Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200219Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200301) GetMetaOk() (*InlineResponse200219Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200301) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200219Meta and assigns it to the Meta field.
func (o *InlineResponse200301) SetMeta(v InlineResponse200219Meta) {
	o.Meta = &v
}

func (o InlineResponse200301) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200301 struct {
	value *InlineResponse200301
	isSet bool
}

func (v NullableInlineResponse200301) Get() *InlineResponse200301 {
	return v.value
}

func (v *NullableInlineResponse200301) Set(val *InlineResponse200301) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200301) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200301) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200301(val *InlineResponse200301) *NullableInlineResponse200301 {
	return &NullableInlineResponse200301{value: val, isSet: true}
}

func (v NullableInlineResponse200301) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200301) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


