/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200331MetaCounts Counts
type InlineResponse200331MetaCounts struct {
	Items *InlineResponse200331MetaCountsItems `json:"items,omitempty"`
}

// NewInlineResponse200331MetaCounts instantiates a new InlineResponse200331MetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200331MetaCounts() *InlineResponse200331MetaCounts {
	this := InlineResponse200331MetaCounts{}
	return &this
}

// NewInlineResponse200331MetaCountsWithDefaults instantiates a new InlineResponse200331MetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200331MetaCountsWithDefaults() *InlineResponse200331MetaCounts {
	this := InlineResponse200331MetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200331MetaCounts) GetItems() InlineResponse200331MetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret InlineResponse200331MetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200331MetaCounts) GetItemsOk() (*InlineResponse200331MetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200331MetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given InlineResponse200331MetaCountsItems and assigns it to the Items field.
func (o *InlineResponse200331MetaCounts) SetItems(v InlineResponse200331MetaCountsItems) {
	o.Items = &v
}

func (o InlineResponse200331MetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200331MetaCounts struct {
	value *InlineResponse200331MetaCounts
	isSet bool
}

func (v NullableInlineResponse200331MetaCounts) Get() *InlineResponse200331MetaCounts {
	return v.value
}

func (v *NullableInlineResponse200331MetaCounts) Set(val *InlineResponse200331MetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200331MetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200331MetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200331MetaCounts(val *InlineResponse200331MetaCounts) *NullableInlineResponse200331MetaCounts {
	return &NullableInlineResponse200331MetaCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200331MetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200331MetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


