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

// InlineResponse200297MetaCountsItems Items
type InlineResponse200297MetaCountsItems struct {
	// Total number of items
	Total *int32 `json:"total,omitempty"`
	// Remaining number of items
	Remaining *int32 `json:"remaining,omitempty"`
}

// NewInlineResponse200297MetaCountsItems instantiates a new InlineResponse200297MetaCountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200297MetaCountsItems() *InlineResponse200297MetaCountsItems {
	this := InlineResponse200297MetaCountsItems{}
	return &this
}

// NewInlineResponse200297MetaCountsItemsWithDefaults instantiates a new InlineResponse200297MetaCountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200297MetaCountsItemsWithDefaults() *InlineResponse200297MetaCountsItems {
	this := InlineResponse200297MetaCountsItems{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200297MetaCountsItems) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200297MetaCountsItems) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200297MetaCountsItems) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200297MetaCountsItems) SetTotal(v int32) {
	o.Total = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *InlineResponse200297MetaCountsItems) GetRemaining() int32 {
	if o == nil || isNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200297MetaCountsItems) GetRemainingOk() (*int32, bool) {
	if o == nil || isNil(o.Remaining) {
    return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *InlineResponse200297MetaCountsItems) HasRemaining() bool {
	if o != nil && !isNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *InlineResponse200297MetaCountsItems) SetRemaining(v int32) {
	o.Remaining = &v
}

func (o InlineResponse200297MetaCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200297MetaCountsItems struct {
	value *InlineResponse200297MetaCountsItems
	isSet bool
}

func (v NullableInlineResponse200297MetaCountsItems) Get() *InlineResponse200297MetaCountsItems {
	return v.value
}

func (v *NullableInlineResponse200297MetaCountsItems) Set(val *InlineResponse200297MetaCountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200297MetaCountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200297MetaCountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200297MetaCountsItems(val *InlineResponse200297MetaCountsItems) *NullableInlineResponse200297MetaCountsItems {
	return &NullableInlineResponse200297MetaCountsItems{value: val, isSet: true}
}

func (v NullableInlineResponse200297MetaCountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200297MetaCountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

