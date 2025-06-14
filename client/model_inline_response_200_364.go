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

// InlineResponse200364 struct for InlineResponse200364
type InlineResponse200364 struct {
	// Estimated time of completion.
	EstimatedCompletedAt *string `json:"estimatedCompletedAt,omitempty"`
}

// NewInlineResponse200364 instantiates a new InlineResponse200364 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200364() *InlineResponse200364 {
	this := InlineResponse200364{}
	return &this
}

// NewInlineResponse200364WithDefaults instantiates a new InlineResponse200364 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200364WithDefaults() *InlineResponse200364 {
	this := InlineResponse200364{}
	return &this
}

// GetEstimatedCompletedAt returns the EstimatedCompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200364) GetEstimatedCompletedAt() string {
	if o == nil || isNil(o.EstimatedCompletedAt) {
		var ret string
		return ret
	}
	return *o.EstimatedCompletedAt
}

// GetEstimatedCompletedAtOk returns a tuple with the EstimatedCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200364) GetEstimatedCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.EstimatedCompletedAt) {
    return nil, false
	}
	return o.EstimatedCompletedAt, true
}

// HasEstimatedCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200364) HasEstimatedCompletedAt() bool {
	if o != nil && !isNil(o.EstimatedCompletedAt) {
		return true
	}

	return false
}

// SetEstimatedCompletedAt gets a reference to the given string and assigns it to the EstimatedCompletedAt field.
func (o *InlineResponse200364) SetEstimatedCompletedAt(v string) {
	o.EstimatedCompletedAt = &v
}

func (o InlineResponse200364) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EstimatedCompletedAt) {
		toSerialize["estimatedCompletedAt"] = o.EstimatedCompletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200364 struct {
	value *InlineResponse200364
	isSet bool
}

func (v NullableInlineResponse200364) Get() *InlineResponse200364 {
	return v.value
}

func (v *NullableInlineResponse200364) Set(val *InlineResponse200364) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200364) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200364) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200364(val *InlineResponse200364) *NullableInlineResponse200364 {
	return &NullableInlineResponse200364{value: val, isSet: true}
}

func (v NullableInlineResponse200364) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200364) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


