/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200179 struct for InlineResponse200179
type InlineResponse200179 struct {
	// Sentry Group Policies for the Organization keyed by Network Id
	Items []InlineResponse200179Items `json:"items,omitempty"`
}

// NewInlineResponse200179 instantiates a new InlineResponse200179 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200179() *InlineResponse200179 {
	this := InlineResponse200179{}
	return &this
}

// NewInlineResponse200179WithDefaults instantiates a new InlineResponse200179 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200179WithDefaults() *InlineResponse200179 {
	this := InlineResponse200179{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200179) GetItems() []InlineResponse200179Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200179Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetItemsOk() ([]InlineResponse200179Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200179) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200179Items and assigns it to the Items field.
func (o *InlineResponse200179) SetItems(v []InlineResponse200179Items) {
	o.Items = v
}

func (o InlineResponse200179) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200179 struct {
	value *InlineResponse200179
	isSet bool
}

func (v NullableInlineResponse200179) Get() *InlineResponse200179 {
	return v.value
}

func (v *NullableInlineResponse200179) Set(val *InlineResponse200179) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200179) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200179) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200179(val *InlineResponse200179) *NullableInlineResponse200179 {
	return &NullableInlineResponse200179{value: val, isSet: true}
}

func (v NullableInlineResponse200179) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200179) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


