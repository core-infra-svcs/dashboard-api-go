/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200311Data A breakdown of how many kilobytes have passed through this port during the interval timespan.
type InlineResponse200311Data struct {
	Usage *InlineResponse200311DataUsage `json:"usage,omitempty"`
}

// NewInlineResponse200311Data instantiates a new InlineResponse200311Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200311Data() *InlineResponse200311Data {
	this := InlineResponse200311Data{}
	return &this
}

// NewInlineResponse200311DataWithDefaults instantiates a new InlineResponse200311Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200311DataWithDefaults() *InlineResponse200311Data {
	this := InlineResponse200311Data{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200311Data) GetUsage() InlineResponse200311DataUsage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200311DataUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200311Data) GetUsageOk() (*InlineResponse200311DataUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200311Data) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200311DataUsage and assigns it to the Usage field.
func (o *InlineResponse200311Data) SetUsage(v InlineResponse200311DataUsage) {
	o.Usage = &v
}

func (o InlineResponse200311Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200311Data struct {
	value *InlineResponse200311Data
	isSet bool
}

func (v NullableInlineResponse200311Data) Get() *InlineResponse200311Data {
	return v.value
}

func (v *NullableInlineResponse200311Data) Set(val *InlineResponse200311Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200311Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200311Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200311Data(val *InlineResponse200311Data) *NullableInlineResponse200311Data {
	return &NullableInlineResponse200311Data{value: val, isSet: true}
}

func (v NullableInlineResponse200311Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200311Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

