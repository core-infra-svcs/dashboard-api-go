/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20118 struct for InlineResponse20118
type InlineResponse20118 struct {
	// List of packet capture files
	Items []InlineResponse200273Items `json:"items,omitempty"`
}

// NewInlineResponse20118 instantiates a new InlineResponse20118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20118() *InlineResponse20118 {
	this := InlineResponse20118{}
	return &this
}

// NewInlineResponse20118WithDefaults instantiates a new InlineResponse20118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20118WithDefaults() *InlineResponse20118 {
	this := InlineResponse20118{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse20118) GetItems() []InlineResponse200273Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200273Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20118) GetItemsOk() ([]InlineResponse200273Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse20118) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200273Items and assigns it to the Items field.
func (o *InlineResponse20118) SetItems(v []InlineResponse200273Items) {
	o.Items = v
}

func (o InlineResponse20118) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20118 struct {
	value *InlineResponse20118
	isSet bool
}

func (v NullableInlineResponse20118) Get() *InlineResponse20118 {
	return v.value
}

func (v *NullableInlineResponse20118) Set(val *InlineResponse20118) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20118) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20118) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20118(val *InlineResponse20118) *NullableInlineResponse20118 {
	return &NullableInlineResponse20118{value: val, isSet: true}
}

func (v NullableInlineResponse20118) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20118) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


