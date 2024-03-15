/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	// Name of the custom performance class
	Name *string `json:"name,omitempty"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject63) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject63) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject63) SetName(v string) {
	o.Name = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *InlineObject63) GetMaxLatency() int32 {
	if o == nil || isNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLatency) {
    return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *InlineObject63) HasMaxLatency() bool {
	if o != nil && !isNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *InlineObject63) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *InlineObject63) GetMaxJitter() int32 {
	if o == nil || isNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMaxJitterOk() (*int32, bool) {
	if o == nil || isNil(o.MaxJitter) {
    return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *InlineObject63) HasMaxJitter() bool {
	if o != nil && !isNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *InlineObject63) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *InlineObject63) GetMaxLossPercentage() int32 {
	if o == nil || isNil(o.MaxLossPercentage) {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLossPercentage) {
    return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *InlineObject63) HasMaxLossPercentage() bool {
	if o != nil && !isNil(o.MaxLossPercentage) {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *InlineObject63) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.MaxLatency) {
		toSerialize["maxLatency"] = o.MaxLatency
	}
	if !isNil(o.MaxJitter) {
		toSerialize["maxJitter"] = o.MaxJitter
	}
	if !isNil(o.MaxLossPercentage) {
		toSerialize["maxLossPercentage"] = o.MaxLossPercentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


