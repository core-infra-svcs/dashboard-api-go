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

// InlineResponse200237MetaCounts Counts of involved entities
type InlineResponse200237MetaCounts struct {
	Items *InlineResponse200237MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200237MetaCounts instantiates a new InlineResponse200237MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200237MetaCounts() *InlineResponse200237MetaCounts {
	this := InlineResponse200237MetaCounts{}
	return &this
}

// NewInlineResponse200237MetaCountsWithDefaults instantiates a new InlineResponse200237MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200237MetaCountsWithDefaults() *InlineResponse200237MetaCounts {
	this := InlineResponse200237MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200237MetaCounts) GetItems() InlineResponse200237MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200237MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237MetaCounts) GetItemsOk() (*InlineResponse200237MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200237MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200237MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200237MetaCounts) SetItems(v InlineResponse200237MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200237MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200237MetaCounts struct {
	value *InlineResponse200237MetaCounts
	isSet bool
}

func (v NullableInlineResponse200237MetaCounts) Get() *InlineResponse200237MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200237MetaCounts) Set(val *InlineResponse200237MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200237MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200237MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200237MetaCounts(val *InlineResponse200237MetaCounts) *NullableInlineResponse200237MetaCounts {
	return &NullableInlineResponse200237MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200237MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200237MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


