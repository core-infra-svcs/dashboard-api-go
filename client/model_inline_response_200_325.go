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

// InlineResponse200325 struct for InlineResponse200325
type InlineResponse200325 struct {
	// Switches
	Items []InlineResponse200325Items `json:"items,omitempty"`
	Meta *InlineResponse200220Meta `json:"meta,omitempty"`
}

// NewInlineResponse200325 instantiates a new InlineResponse200325 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200325() *InlineResponse200325 {
	this := InlineResponse200325{}
	return &this
}

// NewInlineResponse200325WithDefaults instantiates a new InlineResponse200325 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200325WithDefaults() *InlineResponse200325 {
	this := InlineResponse200325{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200325) GetItems() []InlineResponse200325Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200325Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325) GetItemsOk() ([]InlineResponse200325Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200325) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200325Items and assigns it to the Items field.
func (o *InlineResponse200325) SetItems(v []InlineResponse200325Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200325) GetMeta() InlineResponse200220Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200220Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325) GetMetaOk() (*InlineResponse200220Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200325) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200220Meta and assigns it to the Meta field.
func (o *InlineResponse200325) SetMeta(v InlineResponse200220Meta) {
	o.Meta = &v
}

func (o InlineResponse200325) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200325 struct {
	value *InlineResponse200325
	isSet bool
}

func (v NullableInlineResponse200325) Get() *InlineResponse200325 {
	return v.value
}

func (v *NullableInlineResponse200325) Set(val *InlineResponse200325) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200325) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200325) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200325(val *InlineResponse200325) *NullableInlineResponse200325 {
	return &NullableInlineResponse200325{value: val, isSet: true}
}

func (v NullableInlineResponse200325) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200325) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


