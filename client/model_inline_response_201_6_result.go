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

// InlineResponse2016Result Result of the throughput test request
type InlineResponse2016Result struct {
	Speeds *InlineResponse2016ResultSpeeds `json:"speeds,omitempty"`
}

// NewInlineResponse2016Result instantiates a new InlineResponse2016Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2016Result() *InlineResponse2016Result {
	this := InlineResponse2016Result{}
	return &this
}

// NewInlineResponse2016ResultWithDefaults instantiates a new InlineResponse2016Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2016ResultWithDefaults() *InlineResponse2016Result {
	this := InlineResponse2016Result{}
	return &this
}

// GetSpeeds returns the Speeds field value if set, zero value otherwise.
func (o *InlineResponse2016Result) GetSpeeds() InlineResponse2016ResultSpeeds {
	if o == nil || isNil(o.Speeds) {
		var ret InlineResponse2016ResultSpeeds
		return ret
	}
	return *o.Speeds
}

// GetSpeedsOk returns a tuple with the Speeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016Result) GetSpeedsOk() (*InlineResponse2016ResultSpeeds, bool) {
	if o == nil || isNil(o.Speeds) {
    return nil, false
	}
	return o.Speeds, true
}

// HasSpeeds returns a boolean if a field has been set.
func (o *InlineResponse2016Result) HasSpeeds() bool {
	if o != nil && !isNil(o.Speeds) {
		return true
	}

	return false
}

// SetSpeeds gets a reference to the given InlineResponse2016ResultSpeeds and assigns it to the Speeds field.
func (o *InlineResponse2016Result) SetSpeeds(v InlineResponse2016ResultSpeeds) {
	o.Speeds = &v
}

func (o InlineResponse2016Result) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Speeds) {
		toSerialize["speeds"] = o.Speeds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2016Result struct {
	value *InlineResponse2016Result
	isSet bool
}

func (v NullableInlineResponse2016Result) Get() *InlineResponse2016Result {
	return v.value
}

func (v *NullableInlineResponse2016Result) Set(val *InlineResponse2016Result) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2016Result) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2016Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2016Result(val *InlineResponse2016Result) *NullableInlineResponse2016Result {
	return &NullableInlineResponse2016Result{value: val, isSet: true}
}

func (v NullableInlineResponse2016Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2016Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


