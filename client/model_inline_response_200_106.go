/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200106 struct for InlineResponse200106
type InlineResponse200106 struct {
	// List of metrics that are supported for alerts, based on available sensor devices in the network
	SupportedMetrics []string `json:"supportedMetrics,omitempty"`
	Counts *InlineResponse200106Counts `json:"counts,omitempty"`
}

// NewInlineResponse200106 instantiates a new InlineResponse200106 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200106() *InlineResponse200106 {
	this := InlineResponse200106{}
	return &this
}

// NewInlineResponse200106WithDefaults instantiates a new InlineResponse200106 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200106WithDefaults() *InlineResponse200106 {
	this := InlineResponse200106{}
	return &this
}

// GetSupportedMetrics returns the SupportedMetrics field value if set, zero value otherwise.
func (o *InlineResponse200106) GetSupportedMetrics() []string {
	if o == nil || isNil(o.SupportedMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedMetrics
}

// GetSupportedMetricsOk returns a tuple with the SupportedMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106) GetSupportedMetricsOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedMetrics) {
    return nil, false
	}
	return o.SupportedMetrics, true
}

// HasSupportedMetrics returns a boolean if a field has been set.
func (o *InlineResponse200106) HasSupportedMetrics() bool {
	if o != nil && !isNil(o.SupportedMetrics) {
		return true
	}

	return false
}

// SetSupportedMetrics gets a reference to the given []string and assigns it to the SupportedMetrics field.
func (o *InlineResponse200106) SetSupportedMetrics(v []string) {
	o.SupportedMetrics = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200106) GetCounts() InlineResponse200106Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200106Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106) GetCountsOk() (*InlineResponse200106Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200106) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200106Counts and assigns it to the Counts field.
func (o *InlineResponse200106) SetCounts(v InlineResponse200106Counts) {
	o.Counts = &v
}

func (o InlineResponse200106) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupportedMetrics) {
		toSerialize["supportedMetrics"] = o.SupportedMetrics
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200106 struct {
	value *InlineResponse200106
	isSet bool
}

func (v NullableInlineResponse200106) Get() *InlineResponse200106 {
	return v.value
}

func (v *NullableInlineResponse200106) Set(val *InlineResponse200106) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200106) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200106) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200106(val *InlineResponse200106) *NullableInlineResponse200106 {
	return &NullableInlineResponse200106{value: val, isSet: true}
}

func (v NullableInlineResponse200106) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200106) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


