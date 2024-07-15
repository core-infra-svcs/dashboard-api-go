/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200307MetaCounts Count metadata related to this result set.
type InlineResponse200307MetaCounts struct {
	Items *InlineResponse200307MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200307MetaCounts instantiates a new InlineResponse200307MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200307MetaCounts() *InlineResponse200307MetaCounts {
	this := InlineResponse200307MetaCounts{}
	return &this
}

// NewInlineResponse200307MetaCountsWithDefaults instantiates a new InlineResponse200307MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200307MetaCountsWithDefaults() *InlineResponse200307MetaCounts {
	this := InlineResponse200307MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200307MetaCounts) GetItems() InlineResponse200307MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200307MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307MetaCounts) GetItemsOk() (*InlineResponse200307MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200307MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200307MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200307MetaCounts) SetItems(v InlineResponse200307MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200307MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200307MetaCounts struct {
	value *InlineResponse200307MetaCounts
	isSet bool
}

func (v NullableInlineResponse200307MetaCounts) Get() *InlineResponse200307MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200307MetaCounts) Set(val *InlineResponse200307MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200307MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200307MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200307MetaCounts(val *InlineResponse200307MetaCounts) *NullableInlineResponse200307MetaCounts {
	return &NullableInlineResponse200307MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200307MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200307MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


