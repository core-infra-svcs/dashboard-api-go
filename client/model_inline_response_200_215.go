/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200215 struct for InlineResponse200215
type InlineResponse200215 struct {
	// VPN exclusion rules by network
	Items []InlineResponse20067 `json:"items,omitempty"`
}

// NewInlineResponse200215 instantiates a new InlineResponse200215 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200215() *InlineResponse200215 {
	this := InlineResponse200215{}
	return &this
}

// NewInlineResponse200215WithDefaults instantiates a new InlineResponse200215 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200215WithDefaults() *InlineResponse200215 {
	this := InlineResponse200215{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200215) GetItems() []InlineResponse20067 {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse20067
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200215) GetItemsOk() ([]InlineResponse20067, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200215) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse20067 and assigns it to the Items field.
func (o *InlineResponse200215) SetItems(v []InlineResponse20067) {
	o.Items = v
}

func (o InlineResponse200215) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200215 struct {
	value *InlineResponse200215
	isSet bool
}

func (v NullableInlineResponse200215) Get() *InlineResponse200215 {
	return v.value
}

func (v *NullableInlineResponse200215) Set(val *InlineResponse200215) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200215) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200215) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200215(val *InlineResponse200215) *NullableInlineResponse200215 {
	return &NullableInlineResponse200215{value: val, isSet: true}
}

func (v NullableInlineResponse200215) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200215) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


