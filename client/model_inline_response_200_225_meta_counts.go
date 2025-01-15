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

// InlineResponse200225MetaCounts Counts
type InlineResponse200225MetaCounts struct {
	// Total Segments
	Items int32 `json:"items"`
}

// NewInlineResponse200225MetaCounts instantiates a new InlineResponse200225MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200225MetaCounts(items int32) *InlineResponse200225MetaCounts {
	this := InlineResponse200225MetaCounts{}
	this.Items = items
	return &this
}

// NewInlineResponse200225MetaCountsWithDefaults instantiates a new InlineResponse200225MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200225MetaCountsWithDefaults() *InlineResponse200225MetaCounts {
	this := InlineResponse200225MetaCounts{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineResponse200225MetaCounts) GetItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200225MetaCounts) GetItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *InlineResponse200225MetaCounts) SetItems(v int32) {
	o.Items = v
}

func (o InlineResponse200225MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200225MetaCounts struct {
	value *InlineResponse200225MetaCounts
	isSet bool
}

func (v NullableInlineResponse200225MetaCounts) Get() *InlineResponse200225MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200225MetaCounts) Set(val *InlineResponse200225MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200225MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200225MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200225MetaCounts(val *InlineResponse200225MetaCounts) *NullableInlineResponse200225MetaCounts {
	return &NullableInlineResponse200225MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200225MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200225MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

