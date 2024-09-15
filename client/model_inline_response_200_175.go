/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200175 struct for InlineResponse200175
type InlineResponse200175 struct {
	// The latency history bucket start time in seconds
	T0 *int32 `json:"t0,omitempty"`
	// The latency history bucket end time in seconds
	T1 *int32 `json:"t1,omitempty"`
	LatencyBinsByCategory *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory `json:"latencyBinsByCategory,omitempty"`
}

// NewInlineResponse200175 instantiates a new InlineResponse200175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200175() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// NewInlineResponse200175WithDefaults instantiates a new InlineResponse200175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200175WithDefaults() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// GetT0 returns the T0 field value if set, zero value otherwise.
func (o *InlineResponse200175) GetT0() int32 {
	if o == nil || isNil(o.T0) {
		var ret int32
		return ret
	}
	return *o.T0
}

// GetT0Ok returns a tuple with the T0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetT0Ok() (*int32, bool) {
	if o == nil || isNil(o.T0) {
    return nil, false
	}
	return o.T0, true
}

// HasT0 returns a boolean if a field has been set.
func (o *InlineResponse200175) HasT0() bool {
	if o != nil && !isNil(o.T0) {
		return true
	}

	return false
}

// SetT0 gets a reference to the given int32 and assigns it to the T0 field.
func (o *InlineResponse200175) SetT0(v int32) {
	o.T0 = &v
}

// GetT1 returns the T1 field value if set, zero value otherwise.
func (o *InlineResponse200175) GetT1() int32 {
	if o == nil || isNil(o.T1) {
		var ret int32
		return ret
	}
	return *o.T1
}

// GetT1Ok returns a tuple with the T1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetT1Ok() (*int32, bool) {
	if o == nil || isNil(o.T1) {
    return nil, false
	}
	return o.T1, true
}

// HasT1 returns a boolean if a field has been set.
func (o *InlineResponse200175) HasT1() bool {
	if o != nil && !isNil(o.T1) {
		return true
	}

	return false
}

// SetT1 gets a reference to the given int32 and assigns it to the T1 field.
func (o *InlineResponse200175) SetT1(v int32) {
	o.T1 = &v
}

// GetLatencyBinsByCategory returns the LatencyBinsByCategory field value if set, zero value otherwise.
func (o *InlineResponse200175) GetLatencyBinsByCategory() NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory {
	if o == nil || isNil(o.LatencyBinsByCategory) {
		var ret NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory
		return ret
	}
	return *o.LatencyBinsByCategory
}

// GetLatencyBinsByCategoryOk returns a tuple with the LatencyBinsByCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetLatencyBinsByCategoryOk() (*NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory, bool) {
	if o == nil || isNil(o.LatencyBinsByCategory) {
    return nil, false
	}
	return o.LatencyBinsByCategory, true
}

// HasLatencyBinsByCategory returns a boolean if a field has been set.
func (o *InlineResponse200175) HasLatencyBinsByCategory() bool {
	if o != nil && !isNil(o.LatencyBinsByCategory) {
		return true
	}

	return false
}

// SetLatencyBinsByCategory gets a reference to the given NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory and assigns it to the LatencyBinsByCategory field.
func (o *InlineResponse200175) SetLatencyBinsByCategory(v NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) {
	o.LatencyBinsByCategory = &v
}

func (o InlineResponse200175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.T0) {
		toSerialize["t0"] = o.T0
	}
	if !isNil(o.T1) {
		toSerialize["t1"] = o.T1
	}
	if !isNil(o.LatencyBinsByCategory) {
		toSerialize["latencyBinsByCategory"] = o.LatencyBinsByCategory
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200175 struct {
	value *InlineResponse200175
	isSet bool
}

func (v NullableInlineResponse200175) Get() *InlineResponse200175 {
	return v.value
}

func (v *NullableInlineResponse200175) Set(val *InlineResponse200175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200175(val *InlineResponse200175) *NullableInlineResponse200175 {
	return &NullableInlineResponse200175{value: val, isSet: true}
}

func (v NullableInlineResponse200175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


