/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject141 struct for InlineObject141
type InlineObject141 struct {
	// Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
}

// NewInlineObject141 instantiates a new InlineObject141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject141() *InlineObject141 {
	this := InlineObject141{}
	return &this
}

// NewInlineObject141WithDefaults instantiates a new InlineObject141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject141WithDefaults() *InlineObject141 {
	this := InlineObject141{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *InlineObject141) GetBroadcastThreshold() int32 {
	if o == nil || isNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.BroadcastThreshold) {
    return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *InlineObject141) HasBroadcastThreshold() bool {
	if o != nil && !isNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *InlineObject141) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *InlineObject141) GetMulticastThreshold() int32 {
	if o == nil || isNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.MulticastThreshold) {
    return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *InlineObject141) HasMulticastThreshold() bool {
	if o != nil && !isNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *InlineObject141) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *InlineObject141) GetUnknownUnicastThreshold() int32 {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
    return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *InlineObject141) HasUnknownUnicastThreshold() bool {
	if o != nil && !isNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *InlineObject141) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

func (o InlineObject141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BroadcastThreshold) {
		toSerialize["broadcastThreshold"] = o.BroadcastThreshold
	}
	if !isNil(o.MulticastThreshold) {
		toSerialize["multicastThreshold"] = o.MulticastThreshold
	}
	if !isNil(o.UnknownUnicastThreshold) {
		toSerialize["unknownUnicastThreshold"] = o.UnknownUnicastThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject141 struct {
	value *InlineObject141
	isSet bool
}

func (v NullableInlineObject141) Get() *InlineObject141 {
	return v.value
}

func (v *NullableInlineObject141) Set(val *InlineObject141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject141(val *InlineObject141) *NullableInlineObject141 {
	return &NullableInlineObject141{value: val, isSet: true}
}

func (v NullableInlineObject141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


