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

// InlineResponse20079BandwidthLimits The bandwidth settings for the 'cellular' uplink
type InlineResponse20079BandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps). 'null' indicates no limit.
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). 'null' indicates no limit.
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewInlineResponse20079BandwidthLimits instantiates a new InlineResponse20079BandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079BandwidthLimits() *InlineResponse20079BandwidthLimits {
	this := InlineResponse20079BandwidthLimits{}
	return &this
}

// NewInlineResponse20079BandwidthLimitsWithDefaults instantiates a new InlineResponse20079BandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079BandwidthLimitsWithDefaults() *InlineResponse20079BandwidthLimits {
	this := InlineResponse20079BandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *InlineResponse20079BandwidthLimits) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079BandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *InlineResponse20079BandwidthLimits) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *InlineResponse20079BandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *InlineResponse20079BandwidthLimits) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079BandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *InlineResponse20079BandwidthLimits) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *InlineResponse20079BandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o InlineResponse20079BandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079BandwidthLimits struct {
	value *InlineResponse20079BandwidthLimits
	isSet bool
}

func (v NullableInlineResponse20079BandwidthLimits) Get() *InlineResponse20079BandwidthLimits {
	return v.value
}

func (v *NullableInlineResponse20079BandwidthLimits) Set(val *InlineResponse20079BandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079BandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079BandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079BandwidthLimits(val *InlineResponse20079BandwidthLimits) *NullableInlineResponse20079BandwidthLimits {
	return &NullableInlineResponse20079BandwidthLimits{value: val, isSet: true}
}

func (v NullableInlineResponse20079BandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079BandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

