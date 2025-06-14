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

// InlineResponse20068 struct for InlineResponse20068
type InlineResponse20068 struct {
	// Name of the custom performance class
	Name *string `json:"name,omitempty"`
	// ID of the custom performance class
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewInlineResponse20068 instantiates a new InlineResponse20068 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// NewInlineResponse20068WithDefaults instantiates a new InlineResponse20068 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068WithDefaults() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20068) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20068) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20068) SetName(v string) {
	o.Name = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *InlineResponse20068) GetCustomPerformanceClassId() string {
	if o == nil || isNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomPerformanceClassId) {
    return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *InlineResponse20068) HasCustomPerformanceClassId() bool {
	if o != nil && !isNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *InlineResponse20068) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *InlineResponse20068) GetMaxLatency() int32 {
	if o == nil || isNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLatency) {
    return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *InlineResponse20068) HasMaxLatency() bool {
	if o != nil && !isNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *InlineResponse20068) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *InlineResponse20068) GetMaxJitter() int32 {
	if o == nil || isNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetMaxJitterOk() (*int32, bool) {
	if o == nil || isNil(o.MaxJitter) {
    return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *InlineResponse20068) HasMaxJitter() bool {
	if o != nil && !isNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *InlineResponse20068) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *InlineResponse20068) GetMaxLossPercentage() int32 {
	if o == nil || isNil(o.MaxLossPercentage) {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLossPercentage) {
    return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *InlineResponse20068) HasMaxLossPercentage() bool {
	if o != nil && !isNil(o.MaxLossPercentage) {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *InlineResponse20068) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o InlineResponse20068) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
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

type NullableInlineResponse20068 struct {
	value *InlineResponse20068
	isSet bool
}

func (v NullableInlineResponse20068) Get() *InlineResponse20068 {
	return v.value
}

func (v *NullableInlineResponse20068) Set(val *InlineResponse20068) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068(val *InlineResponse20068) *NullableInlineResponse20068 {
	return &NullableInlineResponse20068{value: val, isSet: true}
}

func (v NullableInlineResponse20068) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


