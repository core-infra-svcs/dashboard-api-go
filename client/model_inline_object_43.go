/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject43 struct for InlineObject43
type InlineObject43 struct {
	// Name of the custom performance class
	Name *string `json:"name,omitempty"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewInlineObject43 instantiates a new InlineObject43 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject43() *InlineObject43 {
	this := InlineObject43{}
	return &this
}

// NewInlineObject43WithDefaults instantiates a new InlineObject43 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject43WithDefaults() *InlineObject43 {
	this := InlineObject43{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject43) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject43) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject43) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject43) SetName(v string) {
	o.Name = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *InlineObject43) GetMaxLatency() int32 {
	if o == nil || o.MaxLatency == nil {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject43) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || o.MaxLatency == nil {
		return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *InlineObject43) HasMaxLatency() bool {
	if o != nil && o.MaxLatency != nil {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *InlineObject43) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *InlineObject43) GetMaxJitter() int32 {
	if o == nil || o.MaxJitter == nil {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject43) GetMaxJitterOk() (*int32, bool) {
	if o == nil || o.MaxJitter == nil {
		return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *InlineObject43) HasMaxJitter() bool {
	if o != nil && o.MaxJitter != nil {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *InlineObject43) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *InlineObject43) GetMaxLossPercentage() int32 {
	if o == nil || o.MaxLossPercentage == nil {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject43) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || o.MaxLossPercentage == nil {
		return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *InlineObject43) HasMaxLossPercentage() bool {
	if o != nil && o.MaxLossPercentage != nil {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *InlineObject43) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o InlineObject43) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MaxLatency != nil {
		toSerialize["maxLatency"] = o.MaxLatency
	}
	if o.MaxJitter != nil {
		toSerialize["maxJitter"] = o.MaxJitter
	}
	if o.MaxLossPercentage != nil {
		toSerialize["maxLossPercentage"] = o.MaxLossPercentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject43 struct {
	value *InlineObject43
	isSet bool
}

func (v NullableInlineObject43) Get() *InlineObject43 {
	return v.value
}

func (v *NullableInlineObject43) Set(val *InlineObject43) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject43) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject43) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject43(val *InlineObject43) *NullableInlineObject43 {
	return &NullableInlineObject43{value: val, isSet: true}
}

func (v NullableInlineObject43) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject43) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

