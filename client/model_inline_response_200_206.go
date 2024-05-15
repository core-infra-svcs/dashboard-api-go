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

// InlineResponse200206 struct for InlineResponse200206
type InlineResponse200206 struct {
	// VPN exclusion rules by network
	Items []InlineResponse20064 `json:"items,omitempty"`
}

// NewInlineResponse200206 instantiates a new InlineResponse200206 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200206() *InlineResponse200206 {
	this := InlineResponse200206{}
	return &this
}

// NewInlineResponse200206WithDefaults instantiates a new InlineResponse200206 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200206WithDefaults() *InlineResponse200206 {
	this := InlineResponse200206{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200206) GetItems() []InlineResponse20064 {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse20064
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200206) GetItemsOk() ([]InlineResponse20064, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200206) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse20064 and assigns it to the Items field.
func (o *InlineResponse200206) SetItems(v []InlineResponse20064) {
	o.Items = v
}

func (o InlineResponse200206) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200206 struct {
	value *InlineResponse200206
	isSet bool
}

func (v NullableInlineResponse200206) Get() *InlineResponse200206 {
	return v.value
}

func (v *NullableInlineResponse200206) Set(val *InlineResponse200206) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200206) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200206) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200206(val *InlineResponse200206) *NullableInlineResponse200206 {
	return &NullableInlineResponse200206{value: val, isSet: true}
}

func (v NullableInlineResponse200206) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200206) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


