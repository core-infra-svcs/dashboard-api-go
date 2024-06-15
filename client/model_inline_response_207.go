/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse207 struct for InlineResponse207
type InlineResponse207 struct {
	// The ID of the job that was used to create all of the device swaps.
	JobId *string `json:"jobId,omitempty"`
	// An array of recent swap requests and their statuses.
	Swaps []InlineResponse207Swaps `json:"swaps,omitempty"`
}

// NewInlineResponse207 instantiates a new InlineResponse207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// NewInlineResponse207WithDefaults instantiates a new InlineResponse207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207WithDefaults() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *InlineResponse207) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
    return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *InlineResponse207) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *InlineResponse207) SetJobId(v string) {
	o.JobId = &v
}

// GetSwaps returns the Swaps field value if set, zero value otherwise.
func (o *InlineResponse207) GetSwaps() []InlineResponse207Swaps {
	if o == nil || isNil(o.Swaps) {
		var ret []InlineResponse207Swaps
		return ret
	}
	return o.Swaps
}

// GetSwapsOk returns a tuple with the Swaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetSwapsOk() ([]InlineResponse207Swaps, bool) {
	if o == nil || isNil(o.Swaps) {
    return nil, false
	}
	return o.Swaps, true
}

// HasSwaps returns a boolean if a field has been set.
func (o *InlineResponse207) HasSwaps() bool {
	if o != nil && !isNil(o.Swaps) {
		return true
	}

	return false
}

// SetSwaps gets a reference to the given []InlineResponse207Swaps and assigns it to the Swaps field.
func (o *InlineResponse207) SetSwaps(v []InlineResponse207Swaps) {
	o.Swaps = v
}

func (o InlineResponse207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !isNil(o.Swaps) {
		toSerialize["swaps"] = o.Swaps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207 struct {
	value *InlineResponse207
	isSet bool
}

func (v NullableInlineResponse207) Get() *InlineResponse207 {
	return v.value
}

func (v *NullableInlineResponse207) Set(val *InlineResponse207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207(val *InlineResponse207) *NullableInlineResponse207 {
	return &NullableInlineResponse207{value: val, isSet: true}
}

func (v NullableInlineResponse207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


