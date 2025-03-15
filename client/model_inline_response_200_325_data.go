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

// InlineResponse200325Data A breakdown of how many kilobytes have passed through this port during the interval timespan.
type InlineResponse200325Data struct {
	Usage *InlineResponse200325DataUsage `json:"usage,omitempty"`
}

// NewInlineResponse200325Data instantiates a new InlineResponse200325Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200325Data() *InlineResponse200325Data {
	this := InlineResponse200325Data{}
	return &this
}

// NewInlineResponse200325DataWithDefaults instantiates a new InlineResponse200325Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200325DataWithDefaults() *InlineResponse200325Data {
	this := InlineResponse200325Data{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200325Data) GetUsage() InlineResponse200325DataUsage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200325DataUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Data) GetUsageOk() (*InlineResponse200325DataUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200325Data) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200325DataUsage and assigns it to the Usage field.
func (o *InlineResponse200325Data) SetUsage(v InlineResponse200325DataUsage) {
	o.Usage = &v
}

func (o InlineResponse200325Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200325Data struct {
	value *InlineResponse200325Data
	isSet bool
}

func (v NullableInlineResponse200325Data) Get() *InlineResponse200325Data {
	return v.value
}

func (v *NullableInlineResponse200325Data) Set(val *InlineResponse200325Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200325Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200325Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200325Data(val *InlineResponse200325Data) *NullableInlineResponse200325Data {
	return &NullableInlineResponse200325Data{value: val, isSet: true}
}

func (v NullableInlineResponse200325Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200325Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


