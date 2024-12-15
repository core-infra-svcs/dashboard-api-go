/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject66 struct for InlineObject66
type InlineObject66 struct {
	// Name of the custom performance class
	Name string `json:"name"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewInlineObject66 instantiates a new InlineObject66 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject66(name string) *InlineObject66 {
	this := InlineObject66{}
	this.Name = name
	return &this
}

// NewInlineObject66WithDefaults instantiates a new InlineObject66 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject66WithDefaults() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject66) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject66) SetName(v string) {
	o.Name = v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *InlineObject66) GetMaxLatency() int32 {
	if o == nil || isNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLatency) {
    return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *InlineObject66) HasMaxLatency() bool {
	if o != nil && !isNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *InlineObject66) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *InlineObject66) GetMaxJitter() int32 {
	if o == nil || isNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetMaxJitterOk() (*int32, bool) {
	if o == nil || isNil(o.MaxJitter) {
    return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *InlineObject66) HasMaxJitter() bool {
	if o != nil && !isNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *InlineObject66) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *InlineObject66) GetMaxLossPercentage() int32 {
	if o == nil || isNil(o.MaxLossPercentage) {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLossPercentage) {
    return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *InlineObject66) HasMaxLossPercentage() bool {
	if o != nil && !isNil(o.MaxLossPercentage) {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *InlineObject66) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o InlineObject66) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableInlineObject66 struct {
	value *InlineObject66
	isSet bool
}

func (v NullableInlineObject66) Get() *InlineObject66 {
	return v.value
}

func (v *NullableInlineObject66) Set(val *InlineObject66) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject66) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject66) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject66(val *InlineObject66) *NullableInlineObject66 {
	return &NullableInlineObject66{value: val, isSet: true}
}

func (v NullableInlineObject66) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject66) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


