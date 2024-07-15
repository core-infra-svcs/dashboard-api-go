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

// InlineResponse20063 struct for InlineResponse20063
type InlineResponse20063 struct {
	BandwidthLimits *InlineResponse20063BandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewInlineResponse20063 instantiates a new InlineResponse20063 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063WithDefaults() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineResponse20063) GetBandwidthLimits() InlineResponse20063BandwidthLimits {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret InlineResponse20063BandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetBandwidthLimitsOk() (*InlineResponse20063BandwidthLimits, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineResponse20063) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given InlineResponse20063BandwidthLimits and assigns it to the BandwidthLimits field.
func (o *InlineResponse20063) SetBandwidthLimits(v InlineResponse20063BandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o InlineResponse20063) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20063 struct {
	value *InlineResponse20063
	isSet bool
}

func (v NullableInlineResponse20063) Get() *InlineResponse20063 {
	return v.value
}

func (v *NullableInlineResponse20063) Set(val *InlineResponse20063) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063(val *InlineResponse20063) *NullableInlineResponse20063 {
	return &NullableInlineResponse20063{value: val, isSet: true}
}

func (v NullableInlineResponse20063) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


