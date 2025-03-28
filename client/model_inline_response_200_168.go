/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200168 struct for InlineResponse200168
type InlineResponse200168 struct {
	// Broadcast threshold.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Multicast threshold.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Unknown Unicast threshold.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
	// Grouped traffic types
	TreatTheseTrafficTypesAsOneThreshold []string `json:"treatTheseTrafficTypesAsOneThreshold,omitempty"`
}

// NewInlineResponse200168 instantiates a new InlineResponse200168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200168() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// NewInlineResponse200168WithDefaults instantiates a new InlineResponse200168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200168WithDefaults() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *InlineResponse200168) GetBroadcastThreshold() int32 {
	if o == nil || isNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.BroadcastThreshold) {
    return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *InlineResponse200168) HasBroadcastThreshold() bool {
	if o != nil && !isNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *InlineResponse200168) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *InlineResponse200168) GetMulticastThreshold() int32 {
	if o == nil || isNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.MulticastThreshold) {
    return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *InlineResponse200168) HasMulticastThreshold() bool {
	if o != nil && !isNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *InlineResponse200168) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *InlineResponse200168) GetUnknownUnicastThreshold() int32 {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
    return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *InlineResponse200168) HasUnknownUnicastThreshold() bool {
	if o != nil && !isNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *InlineResponse200168) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

// GetTreatTheseTrafficTypesAsOneThreshold returns the TreatTheseTrafficTypesAsOneThreshold field value if set, zero value otherwise.
func (o *InlineResponse200168) GetTreatTheseTrafficTypesAsOneThreshold() []string {
	if o == nil || isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
		var ret []string
		return ret
	}
	return o.TreatTheseTrafficTypesAsOneThreshold
}

// GetTreatTheseTrafficTypesAsOneThresholdOk returns a tuple with the TreatTheseTrafficTypesAsOneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetTreatTheseTrafficTypesAsOneThresholdOk() ([]string, bool) {
	if o == nil || isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
    return nil, false
	}
	return o.TreatTheseTrafficTypesAsOneThreshold, true
}

// HasTreatTheseTrafficTypesAsOneThreshold returns a boolean if a field has been set.
func (o *InlineResponse200168) HasTreatTheseTrafficTypesAsOneThreshold() bool {
	if o != nil && !isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
		return true
	}

	return false
}

// SetTreatTheseTrafficTypesAsOneThreshold gets a reference to the given []string and assigns it to the TreatTheseTrafficTypesAsOneThreshold field.
func (o *InlineResponse200168) SetTreatTheseTrafficTypesAsOneThreshold(v []string) {
	o.TreatTheseTrafficTypesAsOneThreshold = v
}

func (o InlineResponse200168) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
		toSerialize["treatTheseTrafficTypesAsOneThreshold"] = o.TreatTheseTrafficTypesAsOneThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200168 struct {
	value *InlineResponse200168
	isSet bool
}

func (v NullableInlineResponse200168) Get() *InlineResponse200168 {
	return v.value
}

func (v *NullableInlineResponse200168) Set(val *InlineResponse200168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200168(val *InlineResponse200168) *NullableInlineResponse200168 {
	return &NullableInlineResponse200168{value: val, isSet: true}
}

func (v NullableInlineResponse200168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


