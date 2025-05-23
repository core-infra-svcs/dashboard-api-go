/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject164 struct for InlineObject164
type InlineObject164 struct {
	// Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
	// Grouped traffic types
	TreatTheseTrafficTypesAsOneThreshold []string `json:"treatTheseTrafficTypesAsOneThreshold,omitempty"`
}

// NewInlineObject164 instantiates a new InlineObject164 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject164() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// NewInlineObject164WithDefaults instantiates a new InlineObject164 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject164WithDefaults() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *InlineObject164) GetBroadcastThreshold() int32 {
	if o == nil || isNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.BroadcastThreshold) {
    return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *InlineObject164) HasBroadcastThreshold() bool {
	if o != nil && !isNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *InlineObject164) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *InlineObject164) GetMulticastThreshold() int32 {
	if o == nil || isNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.MulticastThreshold) {
    return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *InlineObject164) HasMulticastThreshold() bool {
	if o != nil && !isNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *InlineObject164) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *InlineObject164) GetUnknownUnicastThreshold() int32 {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
    return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *InlineObject164) HasUnknownUnicastThreshold() bool {
	if o != nil && !isNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *InlineObject164) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

// GetTreatTheseTrafficTypesAsOneThreshold returns the TreatTheseTrafficTypesAsOneThreshold field value if set, zero value otherwise.
func (o *InlineObject164) GetTreatTheseTrafficTypesAsOneThreshold() []string {
	if o == nil || isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
		var ret []string
		return ret
	}
	return o.TreatTheseTrafficTypesAsOneThreshold
}

// GetTreatTheseTrafficTypesAsOneThresholdOk returns a tuple with the TreatTheseTrafficTypesAsOneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetTreatTheseTrafficTypesAsOneThresholdOk() ([]string, bool) {
	if o == nil || isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
    return nil, false
	}
	return o.TreatTheseTrafficTypesAsOneThreshold, true
}

// HasTreatTheseTrafficTypesAsOneThreshold returns a boolean if a field has been set.
func (o *InlineObject164) HasTreatTheseTrafficTypesAsOneThreshold() bool {
	if o != nil && !isNil(o.TreatTheseTrafficTypesAsOneThreshold) {
		return true
	}

	return false
}

// SetTreatTheseTrafficTypesAsOneThreshold gets a reference to the given []string and assigns it to the TreatTheseTrafficTypesAsOneThreshold field.
func (o *InlineObject164) SetTreatTheseTrafficTypesAsOneThreshold(v []string) {
	o.TreatTheseTrafficTypesAsOneThreshold = v
}

func (o InlineObject164) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject164 struct {
	value *InlineObject164
	isSet bool
}

func (v NullableInlineObject164) Get() *InlineObject164 {
	return v.value
}

func (v *NullableInlineObject164) Set(val *InlineObject164) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject164) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject164) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject164(val *InlineObject164) *NullableInlineObject164 {
	return &NullableInlineObject164{value: val, isSet: true}
}

func (v NullableInlineObject164) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject164) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


