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

// InlineResponse200339MetaCounts Counts
type InlineResponse200339MetaCounts struct {
	Items *InlineResponse200339MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200339MetaCounts instantiates a new InlineResponse200339MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339MetaCounts() *InlineResponse200339MetaCounts {
	this := InlineResponse200339MetaCounts{}
	return &this
}

// NewInlineResponse200339MetaCountsWithDefaults instantiates a new InlineResponse200339MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339MetaCountsWithDefaults() *InlineResponse200339MetaCounts {
	this := InlineResponse200339MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200339MetaCounts) GetItems() InlineResponse200339MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200339MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339MetaCounts) GetItemsOk() (*InlineResponse200339MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200339MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200339MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200339MetaCounts) SetItems(v InlineResponse200339MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200339MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339MetaCounts struct {
	value *InlineResponse200339MetaCounts
	isSet bool
}

func (v NullableInlineResponse200339MetaCounts) Get() *InlineResponse200339MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200339MetaCounts) Set(val *InlineResponse200339MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339MetaCounts(val *InlineResponse200339MetaCounts) *NullableInlineResponse200339MetaCounts {
	return &NullableInlineResponse200339MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200339MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


