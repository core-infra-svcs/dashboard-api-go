/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20045 struct for InlineResponse20045
type InlineResponse20045 struct {
	// List of metrics that are supported for alerts, based on available sensor devices in the network
	SupportedMetrics []string `json:"supportedMetrics,omitempty"`
	Counts *InlineResponse20045Counts `json:"counts,omitempty"`
}

// NewInlineResponse20045 instantiates a new InlineResponse20045 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20045() *InlineResponse20045 {
	this := InlineResponse20045{}
	return &this
}

// NewInlineResponse20045WithDefaults instantiates a new InlineResponse20045 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20045WithDefaults() *InlineResponse20045 {
	this := InlineResponse20045{}
	return &this
}

// GetSupportedMetrics returns the SupportedMetrics field value if set, zero value otherwise.
func (o *InlineResponse20045) GetSupportedMetrics() []string {
	if o == nil || isNil(o.SupportedMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedMetrics
}

// GetSupportedMetricsOk returns a tuple with the SupportedMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20045) GetSupportedMetricsOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedMetrics) {
    return nil, false
	}
	return o.SupportedMetrics, true
}

// HasSupportedMetrics returns a boolean if a field has been set.
func (o *InlineResponse20045) HasSupportedMetrics() bool {
	if o != nil && !isNil(o.SupportedMetrics) {
		return true
	}

	return false
}

// SetSupportedMetrics gets a reference to the given []string and assigns it to the SupportedMetrics field.
func (o *InlineResponse20045) SetSupportedMetrics(v []string) {
	o.SupportedMetrics = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse20045) GetCounts() InlineResponse20045Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse20045Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20045) GetCountsOk() (*InlineResponse20045Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse20045) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse20045Counts and assigns it to the Counts field.
func (o *InlineResponse20045) SetCounts(v InlineResponse20045Counts) {
	o.Counts = &v
}

func (o InlineResponse20045) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupportedMetrics) {
		toSerialize["supportedMetrics"] = o.SupportedMetrics
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20045 struct {
	value *InlineResponse20045
	isSet bool
}

func (v NullableInlineResponse20045) Get() *InlineResponse20045 {
	return v.value
}

func (v *NullableInlineResponse20045) Set(val *InlineResponse20045) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20045) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20045) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20045(val *InlineResponse20045) *NullableInlineResponse20045 {
	return &NullableInlineResponse20045{value: val, isSet: true}
}

func (v NullableInlineResponse20045) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20045) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


