/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200312 struct for InlineResponse200312
type InlineResponse200312 struct {
	// Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to
	Items []InlineResponse200311Items `json:"items,omitempty"`
	Meta *InlineResponse200222Meta `json:"meta,omitempty"`
}

// NewInlineResponse200312 instantiates a new InlineResponse200312 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200312() *InlineResponse200312 {
	this := InlineResponse200312{}
	return &this
}

// NewInlineResponse200312WithDefaults instantiates a new InlineResponse200312 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200312WithDefaults() *InlineResponse200312 {
	this := InlineResponse200312{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200312) GetItems() []InlineResponse200311Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200311Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200312) GetItemsOk() ([]InlineResponse200311Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200312) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200311Items and assigns it to the Items field.
func (o *InlineResponse200312) SetItems(v []InlineResponse200311Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200312) GetMeta() InlineResponse200222Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200222Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200312) GetMetaOk() (*InlineResponse200222Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200312) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200222Meta and assigns it to the Meta field.
func (o *InlineResponse200312) SetMeta(v InlineResponse200222Meta) {
	o.Meta = &v
}

func (o InlineResponse200312) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200312 struct {
	value *InlineResponse200312
	isSet bool
}

func (v NullableInlineResponse200312) Get() *InlineResponse200312 {
	return v.value
}

func (v *NullableInlineResponse200312) Set(val *InlineResponse200312) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200312) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200312) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200312(val *InlineResponse200312) *NullableInlineResponse200312 {
	return &NullableInlineResponse200312{value: val, isSet: true}
}

func (v NullableInlineResponse200312) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200312) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


