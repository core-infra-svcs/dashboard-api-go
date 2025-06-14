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

// InlineResponse200292MetaCounts Counts relating to the paginated dataset
type InlineResponse200292MetaCounts struct {
	Items *InlineResponse200292MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200292MetaCounts instantiates a new InlineResponse200292MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200292MetaCounts() *InlineResponse200292MetaCounts {
	this := InlineResponse200292MetaCounts{}
	return &this
}

// NewInlineResponse200292MetaCountsWithDefaults instantiates a new InlineResponse200292MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200292MetaCountsWithDefaults() *InlineResponse200292MetaCounts {
	this := InlineResponse200292MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200292MetaCounts) GetItems() InlineResponse200292MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200292MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292MetaCounts) GetItemsOk() (*InlineResponse200292MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200292MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200292MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200292MetaCounts) SetItems(v InlineResponse200292MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200292MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200292MetaCounts struct {
	value *InlineResponse200292MetaCounts
	isSet bool
}

func (v NullableInlineResponse200292MetaCounts) Get() *InlineResponse200292MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200292MetaCounts) Set(val *InlineResponse200292MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200292MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200292MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200292MetaCounts(val *InlineResponse200292MetaCounts) *NullableInlineResponse200292MetaCounts {
	return &NullableInlineResponse200292MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200292MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200292MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


