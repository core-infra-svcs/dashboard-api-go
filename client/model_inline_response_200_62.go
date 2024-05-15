/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20062 struct for InlineResponse20062
type InlineResponse20062 struct {
	BandwidthLimits *InlineResponse20062BandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewInlineResponse20062 instantiates a new InlineResponse20062 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20062() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// NewInlineResponse20062WithDefaults instantiates a new InlineResponse20062 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20062WithDefaults() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineResponse20062) GetBandwidthLimits() InlineResponse20062BandwidthLimits {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret InlineResponse20062BandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetBandwidthLimitsOk() (*InlineResponse20062BandwidthLimits, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineResponse20062) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given InlineResponse20062BandwidthLimits and assigns it to the BandwidthLimits field.
func (o *InlineResponse20062) SetBandwidthLimits(v InlineResponse20062BandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o InlineResponse20062) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20062 struct {
	value *InlineResponse20062
	isSet bool
}

func (v NullableInlineResponse20062) Get() *InlineResponse20062 {
	return v.value
}

func (v *NullableInlineResponse20062) Set(val *InlineResponse20062) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20062) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20062) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20062(val *InlineResponse20062) *NullableInlineResponse20062 {
	return &NullableInlineResponse20062{value: val, isSet: true}
}

func (v NullableInlineResponse20062) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20062) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


