/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200174BandwidthLimits The uplink bandwidth settings for the pricing plan.
type InlineResponse200174BandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps).
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps).
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewInlineResponse200174BandwidthLimits instantiates a new InlineResponse200174BandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200174BandwidthLimits() *InlineResponse200174BandwidthLimits {
	this := InlineResponse200174BandwidthLimits{}
	return &this
}

// NewInlineResponse200174BandwidthLimitsWithDefaults instantiates a new InlineResponse200174BandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200174BandwidthLimitsWithDefaults() *InlineResponse200174BandwidthLimits {
	this := InlineResponse200174BandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *InlineResponse200174BandwidthLimits) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200174BandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *InlineResponse200174BandwidthLimits) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *InlineResponse200174BandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *InlineResponse200174BandwidthLimits) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200174BandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *InlineResponse200174BandwidthLimits) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *InlineResponse200174BandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o InlineResponse200174BandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200174BandwidthLimits struct {
	value *InlineResponse200174BandwidthLimits
	isSet bool
}

func (v NullableInlineResponse200174BandwidthLimits) Get() *InlineResponse200174BandwidthLimits {
	return v.value
}

func (v *NullableInlineResponse200174BandwidthLimits) Set(val *InlineResponse200174BandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200174BandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200174BandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200174BandwidthLimits(val *InlineResponse200174BandwidthLimits) *NullableInlineResponse200174BandwidthLimits {
	return &NullableInlineResponse200174BandwidthLimits{value: val, isSet: true}
}

func (v NullableInlineResponse200174BandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200174BandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

