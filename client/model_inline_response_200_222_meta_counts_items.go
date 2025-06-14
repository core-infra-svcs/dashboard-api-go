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

// InlineResponse200222MetaCountsItems Counts relating to the paginated items
type InlineResponse200222MetaCountsItems struct {
	// The total number of items in the dataset
	Total *int32 `json:"total,omitempty"`
	// The number of items in the dataset that are available on subsequent pages
	Remaining *int32 `json:"remaining,omitempty"`
}

// NewInlineResponse200222MetaCountsItems instantiates a new InlineResponse200222MetaCountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200222MetaCountsItems() *InlineResponse200222MetaCountsItems {
	this := InlineResponse200222MetaCountsItems{}
	return &this
}

// NewInlineResponse200222MetaCountsItemsWithDefaults instantiates a new InlineResponse200222MetaCountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200222MetaCountsItemsWithDefaults() *InlineResponse200222MetaCountsItems {
	this := InlineResponse200222MetaCountsItems{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200222MetaCountsItems) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200222MetaCountsItems) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200222MetaCountsItems) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200222MetaCountsItems) SetTotal(v int32) {
	o.Total = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *InlineResponse200222MetaCountsItems) GetRemaining() int32 {
	if o == nil || isNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200222MetaCountsItems) GetRemainingOk() (*int32, bool) {
	if o == nil || isNil(o.Remaining) {
    return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *InlineResponse200222MetaCountsItems) HasRemaining() bool {
	if o != nil && !isNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *InlineResponse200222MetaCountsItems) SetRemaining(v int32) {
	o.Remaining = &v
}

func (o InlineResponse200222MetaCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200222MetaCountsItems struct {
	value *InlineResponse200222MetaCountsItems
	isSet bool
}

func (v NullableInlineResponse200222MetaCountsItems) Get() *InlineResponse200222MetaCountsItems {
	return v.value
}

func (v *NullableInlineResponse200222MetaCountsItems) Set(val *InlineResponse200222MetaCountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200222MetaCountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200222MetaCountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200222MetaCountsItems(val *InlineResponse200222MetaCountsItems) *NullableInlineResponse200222MetaCountsItems {
	return &NullableInlineResponse200222MetaCountsItems{value: val, isSet: true}
}

func (v NullableInlineResponse200222MetaCountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200222MetaCountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


