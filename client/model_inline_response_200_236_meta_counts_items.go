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

// InlineResponse200236MetaCountsItems Count of Communication Plans available
type InlineResponse200236MetaCountsItems struct {
	// Total number of Communication Plans
	Total *int32 `json:"total,omitempty"`
	// Remaining number of Communication Plans
	Remaining *int32 `json:"remaining,omitempty"`
}

// NewInlineResponse200236MetaCountsItems instantiates a new InlineResponse200236MetaCountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200236MetaCountsItems() *InlineResponse200236MetaCountsItems {
	this := InlineResponse200236MetaCountsItems{}
	return &this
}

// NewInlineResponse200236MetaCountsItemsWithDefaults instantiates a new InlineResponse200236MetaCountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200236MetaCountsItemsWithDefaults() *InlineResponse200236MetaCountsItems {
	this := InlineResponse200236MetaCountsItems{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200236MetaCountsItems) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200236MetaCountsItems) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200236MetaCountsItems) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200236MetaCountsItems) SetTotal(v int32) {
	o.Total = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *InlineResponse200236MetaCountsItems) GetRemaining() int32 {
	if o == nil || isNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200236MetaCountsItems) GetRemainingOk() (*int32, bool) {
	if o == nil || isNil(o.Remaining) {
    return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *InlineResponse200236MetaCountsItems) HasRemaining() bool {
	if o != nil && !isNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *InlineResponse200236MetaCountsItems) SetRemaining(v int32) {
	o.Remaining = &v
}

func (o InlineResponse200236MetaCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200236MetaCountsItems struct {
	value *InlineResponse200236MetaCountsItems
	isSet bool
}

func (v NullableInlineResponse200236MetaCountsItems) Get() *InlineResponse200236MetaCountsItems {
	return v.value
}

func (v *NullableInlineResponse200236MetaCountsItems) Set(val *InlineResponse200236MetaCountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200236MetaCountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200236MetaCountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200236MetaCountsItems(val *InlineResponse200236MetaCountsItems) *NullableInlineResponse200236MetaCountsItems {
	return &NullableInlineResponse200236MetaCountsItems{value: val, isSet: true}
}

func (v NullableInlineResponse200236MetaCountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200236MetaCountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


