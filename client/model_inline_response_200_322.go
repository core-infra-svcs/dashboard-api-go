/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200322 struct for InlineResponse200322
type InlineResponse200322 struct {
	// Estimated time of completion.
	EstimatedCompletedAt *string `json:"estimatedCompletedAt,omitempty"`
}

// NewInlineResponse200322 instantiates a new InlineResponse200322 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200322() *InlineResponse200322 {
	this := InlineResponse200322{}
	return &this
}

// NewInlineResponse200322WithDefaults instantiates a new InlineResponse200322 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200322WithDefaults() *InlineResponse200322 {
	this := InlineResponse200322{}
	return &this
}

// GetEstimatedCompletedAt returns the EstimatedCompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200322) GetEstimatedCompletedAt() string {
	if o == nil || isNil(o.EstimatedCompletedAt) {
		var ret string
		return ret
	}
	return *o.EstimatedCompletedAt
}

// GetEstimatedCompletedAtOk returns a tuple with the EstimatedCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322) GetEstimatedCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.EstimatedCompletedAt) {
    return nil, false
	}
	return o.EstimatedCompletedAt, true
}

// HasEstimatedCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200322) HasEstimatedCompletedAt() bool {
	if o != nil && !isNil(o.EstimatedCompletedAt) {
		return true
	}

	return false
}

// SetEstimatedCompletedAt gets a reference to the given string and assigns it to the EstimatedCompletedAt field.
func (o *InlineResponse200322) SetEstimatedCompletedAt(v string) {
	o.EstimatedCompletedAt = &v
}

func (o InlineResponse200322) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EstimatedCompletedAt) {
		toSerialize["estimatedCompletedAt"] = o.EstimatedCompletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200322 struct {
	value *InlineResponse200322
	isSet bool
}

func (v NullableInlineResponse200322) Get() *InlineResponse200322 {
	return v.value
}

func (v *NullableInlineResponse200322) Set(val *InlineResponse200322) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200322) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200322) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200322(val *InlineResponse200322) *NullableInlineResponse200322 {
	return &NullableInlineResponse200322{value: val, isSet: true}
}

func (v NullableInlineResponse200322) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200322) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


