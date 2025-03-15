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

// InlineResponse200344 struct for InlineResponse200344
type InlineResponse200344 struct {
	// Estimated time of completion.
	EstimatedCompletedAt *string `json:"estimatedCompletedAt,omitempty"`
}

// NewInlineResponse200344 instantiates a new InlineResponse200344 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200344() *InlineResponse200344 {
	this := InlineResponse200344{}
	return &this
}

// NewInlineResponse200344WithDefaults instantiates a new InlineResponse200344 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200344WithDefaults() *InlineResponse200344 {
	this := InlineResponse200344{}
	return &this
}

// GetEstimatedCompletedAt returns the EstimatedCompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200344) GetEstimatedCompletedAt() string {
	if o == nil || isNil(o.EstimatedCompletedAt) {
		var ret string
		return ret
	}
	return *o.EstimatedCompletedAt
}

// GetEstimatedCompletedAtOk returns a tuple with the EstimatedCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200344) GetEstimatedCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.EstimatedCompletedAt) {
    return nil, false
	}
	return o.EstimatedCompletedAt, true
}

// HasEstimatedCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200344) HasEstimatedCompletedAt() bool {
	if o != nil && !isNil(o.EstimatedCompletedAt) {
		return true
	}

	return false
}

// SetEstimatedCompletedAt gets a reference to the given string and assigns it to the EstimatedCompletedAt field.
func (o *InlineResponse200344) SetEstimatedCompletedAt(v string) {
	o.EstimatedCompletedAt = &v
}

func (o InlineResponse200344) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EstimatedCompletedAt) {
		toSerialize["estimatedCompletedAt"] = o.EstimatedCompletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200344 struct {
	value *InlineResponse200344
	isSet bool
}

func (v NullableInlineResponse200344) Get() *InlineResponse200344 {
	return v.value
}

func (v *NullableInlineResponse200344) Set(val *InlineResponse200344) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200344) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200344) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200344(val *InlineResponse200344) *NullableInlineResponse200344 {
	return &NullableInlineResponse200344{value: val, isSet: true}
}

func (v NullableInlineResponse200344) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200344) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


