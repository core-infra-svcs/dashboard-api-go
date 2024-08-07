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

// InlineResponse200258SystemsManager Aggregated data for Systems Manager licenses (Per-device licensing only)
type InlineResponse200258SystemsManager struct {
	Counts *InlineResponse200258SystemsManagerCounts `json:"counts,omitempty"`
}

// NewInlineResponse200258SystemsManager instantiates a new InlineResponse200258SystemsManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200258SystemsManager() *InlineResponse200258SystemsManager {
	this := InlineResponse200258SystemsManager{}
	return &this
}

// NewInlineResponse200258SystemsManagerWithDefaults instantiates a new InlineResponse200258SystemsManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200258SystemsManagerWithDefaults() *InlineResponse200258SystemsManager {
	this := InlineResponse200258SystemsManager{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200258SystemsManager) GetCounts() InlineResponse200258SystemsManagerCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200258SystemsManagerCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258SystemsManager) GetCountsOk() (*InlineResponse200258SystemsManagerCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200258SystemsManager) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200258SystemsManagerCounts and assigns it to the Counts field.
func (o *InlineResponse200258SystemsManager) SetCounts(v InlineResponse200258SystemsManagerCounts) {
	o.Counts = &v
}

func (o InlineResponse200258SystemsManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200258SystemsManager struct {
	value *InlineResponse200258SystemsManager
	isSet bool
}

func (v NullableInlineResponse200258SystemsManager) Get() *InlineResponse200258SystemsManager {
	return v.value
}

func (v *NullableInlineResponse200258SystemsManager) Set(val *InlineResponse200258SystemsManager) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200258SystemsManager) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200258SystemsManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200258SystemsManager(val *InlineResponse200258SystemsManager) *NullableInlineResponse200258SystemsManager {
	return &NullableInlineResponse200258SystemsManager{value: val, isSet: true}
}

func (v NullableInlineResponse200258SystemsManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200258SystemsManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


