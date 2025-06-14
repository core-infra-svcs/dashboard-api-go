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

// InlineResponse200252MetaCounts Count metadata related to this result set.
type InlineResponse200252MetaCounts struct {
	Items *InlineResponse200252MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200252MetaCounts instantiates a new InlineResponse200252MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200252MetaCounts() *InlineResponse200252MetaCounts {
	this := InlineResponse200252MetaCounts{}
	return &this
}

// NewInlineResponse200252MetaCountsWithDefaults instantiates a new InlineResponse200252MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200252MetaCountsWithDefaults() *InlineResponse200252MetaCounts {
	this := InlineResponse200252MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200252MetaCounts) GetItems() InlineResponse200252MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200252MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200252MetaCounts) GetItemsOk() (*InlineResponse200252MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200252MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200252MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200252MetaCounts) SetItems(v InlineResponse200252MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200252MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200252MetaCounts struct {
	value *InlineResponse200252MetaCounts
	isSet bool
}

func (v NullableInlineResponse200252MetaCounts) Get() *InlineResponse200252MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200252MetaCounts) Set(val *InlineResponse200252MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200252MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200252MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200252MetaCounts(val *InlineResponse200252MetaCounts) *NullableInlineResponse200252MetaCounts {
	return &NullableInlineResponse200252MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200252MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200252MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


