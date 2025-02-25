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

// InlineResponse200328MetaCounts Counts
type InlineResponse200328MetaCounts struct {
	Items *InlineResponse200328MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200328MetaCounts instantiates a new InlineResponse200328MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200328MetaCounts() *InlineResponse200328MetaCounts {
	this := InlineResponse200328MetaCounts{}
	return &this
}

// NewInlineResponse200328MetaCountsWithDefaults instantiates a new InlineResponse200328MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200328MetaCountsWithDefaults() *InlineResponse200328MetaCounts {
	this := InlineResponse200328MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200328MetaCounts) GetItems() InlineResponse200328MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200328MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200328MetaCounts) GetItemsOk() (*InlineResponse200328MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200328MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200328MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200328MetaCounts) SetItems(v InlineResponse200328MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200328MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200328MetaCounts struct {
	value *InlineResponse200328MetaCounts
	isSet bool
}

func (v NullableInlineResponse200328MetaCounts) Get() *InlineResponse200328MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200328MetaCounts) Set(val *InlineResponse200328MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200328MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200328MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200328MetaCounts(val *InlineResponse200328MetaCounts) *NullableInlineResponse200328MetaCounts {
	return &NullableInlineResponse200328MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200328MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200328MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


