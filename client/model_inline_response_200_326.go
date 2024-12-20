/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200326 struct for InlineResponse200326
type InlineResponse200326 struct {
	// Estimated time of completion.
	EstimatedCompletedAt *string `json:"estimatedCompletedAt,omitempty"`
}

// NewInlineResponse200326 instantiates a new InlineResponse200326 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200326() *InlineResponse200326 {
	this := InlineResponse200326{}
	return &this
}

// NewInlineResponse200326WithDefaults instantiates a new InlineResponse200326 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200326WithDefaults() *InlineResponse200326 {
	this := InlineResponse200326{}
	return &this
}

// GetEstimatedCompletedAt returns the EstimatedCompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200326) GetEstimatedCompletedAt() string {
	if o == nil || isNil(o.EstimatedCompletedAt) {
		var ret string
		return ret
	}
	return *o.EstimatedCompletedAt
}

// GetEstimatedCompletedAtOk returns a tuple with the EstimatedCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200326) GetEstimatedCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.EstimatedCompletedAt) {
    return nil, false
	}
	return o.EstimatedCompletedAt, true
}

// HasEstimatedCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200326) HasEstimatedCompletedAt() bool {
	if o != nil && !isNil(o.EstimatedCompletedAt) {
		return true
	}

	return false
}

// SetEstimatedCompletedAt gets a reference to the given string and assigns it to the EstimatedCompletedAt field.
func (o *InlineResponse200326) SetEstimatedCompletedAt(v string) {
	o.EstimatedCompletedAt = &v
}

func (o InlineResponse200326) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EstimatedCompletedAt) {
		toSerialize["estimatedCompletedAt"] = o.EstimatedCompletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200326 struct {
	value *InlineResponse200326
	isSet bool
}

func (v NullableInlineResponse200326) Get() *InlineResponse200326 {
	return v.value
}

func (v *NullableInlineResponse200326) Set(val *InlineResponse200326) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200326) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200326) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200326(val *InlineResponse200326) *NullableInlineResponse200326 {
	return &NullableInlineResponse200326{value: val, isSet: true}
}

func (v NullableInlineResponse200326) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200326) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


