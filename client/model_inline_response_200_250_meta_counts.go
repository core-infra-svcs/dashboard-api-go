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

// InlineResponse200250MetaCounts Counts relating to the paginated dataset
type InlineResponse200250MetaCounts struct {
	Items *InlineResponse200250MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200250MetaCounts instantiates a new InlineResponse200250MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200250MetaCounts() *InlineResponse200250MetaCounts {
	this := InlineResponse200250MetaCounts{}
	return &this
}

// NewInlineResponse200250MetaCountsWithDefaults instantiates a new InlineResponse200250MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200250MetaCountsWithDefaults() *InlineResponse200250MetaCounts {
	this := InlineResponse200250MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200250MetaCounts) GetItems() InlineResponse200250MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200250MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200250MetaCounts) GetItemsOk() (*InlineResponse200250MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200250MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200250MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200250MetaCounts) SetItems(v InlineResponse200250MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200250MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200250MetaCounts struct {
	value *InlineResponse200250MetaCounts
	isSet bool
}

func (v NullableInlineResponse200250MetaCounts) Get() *InlineResponse200250MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200250MetaCounts) Set(val *InlineResponse200250MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200250MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200250MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200250MetaCounts(val *InlineResponse200250MetaCounts) *NullableInlineResponse200250MetaCounts {
	return &NullableInlineResponse200250MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200250MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200250MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


