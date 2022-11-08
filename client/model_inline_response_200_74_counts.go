/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20074Counts Client count information
type InlineResponse20074Counts struct {
	// Total number of clients with data usage in organization
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse20074Counts instantiates a new InlineResponse20074Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074Counts() *InlineResponse20074Counts {
	this := InlineResponse20074Counts{}
	return &this
}

// NewInlineResponse20074CountsWithDefaults instantiates a new InlineResponse20074Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074CountsWithDefaults() *InlineResponse20074Counts {
	this := InlineResponse20074Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20074Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20074Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse20074Counts) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse20074Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074Counts struct {
	value *InlineResponse20074Counts
	isSet bool
}

func (v NullableInlineResponse20074Counts) Get() *InlineResponse20074Counts {
	return v.value
}

func (v *NullableInlineResponse20074Counts) Set(val *InlineResponse20074Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074Counts(val *InlineResponse20074Counts) *NullableInlineResponse20074Counts {
	return &NullableInlineResponse20074Counts{value: val, isSet: true}
}

func (v NullableInlineResponse20074Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


