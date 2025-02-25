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

// InlineResponse20085Counts The number of clients on a network over a given time range
type InlineResponse20085Counts struct {
	// The total number of clients on a network
	Total *int32 `json:"total,omitempty"`
	// The total number of clients with heavy usage on a network
	WithHeavyUsage *int32 `json:"withHeavyUsage,omitempty"`
}

// NewInlineResponse20085Counts instantiates a new InlineResponse20085Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20085Counts() *InlineResponse20085Counts {
	this := InlineResponse20085Counts{}
	return &this
}

// NewInlineResponse20085CountsWithDefaults instantiates a new InlineResponse20085Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20085CountsWithDefaults() *InlineResponse20085Counts {
	this := InlineResponse20085Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20085Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20085Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse20085Counts) SetTotal(v int32) {
	o.Total = &v
}

// GetWithHeavyUsage returns the WithHeavyUsage field value if set, zero value otherwise.
func (o *InlineResponse20085Counts) GetWithHeavyUsage() int32 {
	if o == nil || isNil(o.WithHeavyUsage) {
		var ret int32
		return ret
	}
	return *o.WithHeavyUsage
}

// GetWithHeavyUsageOk returns a tuple with the WithHeavyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Counts) GetWithHeavyUsageOk() (*int32, bool) {
	if o == nil || isNil(o.WithHeavyUsage) {
    return nil, false
	}
	return o.WithHeavyUsage, true
}

// HasWithHeavyUsage returns a boolean if a field has been set.
func (o *InlineResponse20085Counts) HasWithHeavyUsage() bool {
	if o != nil && !isNil(o.WithHeavyUsage) {
		return true
	}

	return false
}

// SetWithHeavyUsage gets a reference to the given int32 and assigns it to the WithHeavyUsage field.
func (o *InlineResponse20085Counts) SetWithHeavyUsage(v int32) {
	o.WithHeavyUsage = &v
}

func (o InlineResponse20085Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.WithHeavyUsage) {
		toSerialize["withHeavyUsage"] = o.WithHeavyUsage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20085Counts struct {
	value *InlineResponse20085Counts
	isSet bool
}

func (v NullableInlineResponse20085Counts) Get() *InlineResponse20085Counts {
	return v.value
}

func (v *NullableInlineResponse20085Counts) Set(val *InlineResponse20085Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20085Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20085Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20085Counts(val *InlineResponse20085Counts) *NullableInlineResponse20085Counts {
	return &NullableInlineResponse20085Counts{value: val, isSet: true}
}

func (v NullableInlineResponse20085Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20085Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


