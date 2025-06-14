/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069BandwidthLimitsWan2 uplink wan2 configured limits [optional]
type InlineResponse20069BandwidthLimitsWan2 struct {
	// configured UP limit for the uplink (in Kbps).  Null indicated unlimited
	LimitUp *int32 `json:"limitUp,omitempty"`
	// configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewInlineResponse20069BandwidthLimitsWan2 instantiates a new InlineResponse20069BandwidthLimitsWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069BandwidthLimitsWan2() *InlineResponse20069BandwidthLimitsWan2 {
	this := InlineResponse20069BandwidthLimitsWan2{}
	return &this
}

// NewInlineResponse20069BandwidthLimitsWan2WithDefaults instantiates a new InlineResponse20069BandwidthLimitsWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069BandwidthLimitsWan2WithDefaults() *InlineResponse20069BandwidthLimitsWan2 {
	this := InlineResponse20069BandwidthLimitsWan2{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *InlineResponse20069BandwidthLimitsWan2) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069BandwidthLimitsWan2) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *InlineResponse20069BandwidthLimitsWan2) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *InlineResponse20069BandwidthLimitsWan2) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *InlineResponse20069BandwidthLimitsWan2) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069BandwidthLimitsWan2) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *InlineResponse20069BandwidthLimitsWan2) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *InlineResponse20069BandwidthLimitsWan2) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o InlineResponse20069BandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069BandwidthLimitsWan2 struct {
	value *InlineResponse20069BandwidthLimitsWan2
	isSet bool
}

func (v NullableInlineResponse20069BandwidthLimitsWan2) Get() *InlineResponse20069BandwidthLimitsWan2 {
	return v.value
}

func (v *NullableInlineResponse20069BandwidthLimitsWan2) Set(val *InlineResponse20069BandwidthLimitsWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069BandwidthLimitsWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069BandwidthLimitsWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069BandwidthLimitsWan2(val *InlineResponse20069BandwidthLimitsWan2) *NullableInlineResponse20069BandwidthLimitsWan2 {
	return &NullableInlineResponse20069BandwidthLimitsWan2{value: val, isSet: true}
}

func (v NullableInlineResponse20069BandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069BandwidthLimitsWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


