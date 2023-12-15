/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
	Callback *DevicesSerialLiveToolsPingCallback `json:"callback,omitempty"`
}

// NewInlineObject14 instantiates a new InlineObject14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// NewInlineObject14WithDefaults instantiates a new InlineObject14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14WithDefaults() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject14) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject14) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject14) SetCount(v int32) {
	o.Count = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject14) GetCallback() DevicesSerialLiveToolsPingCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsPingCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetCallbackOk() (*DevicesSerialLiveToolsPingCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject14) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsPingCallback and assigns it to the Callback field.
func (o *InlineObject14) SetCallback(v DevicesSerialLiveToolsPingCallback) {
	o.Callback = &v
}

func (o InlineObject14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject14 struct {
	value *InlineObject14
	isSet bool
}

func (v NullableInlineObject14) Get() *InlineObject14 {
	return v.value
}

func (v *NullableInlineObject14) Set(val *InlineObject14) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14(val *InlineObject14) *NullableInlineObject14 {
	return &NullableInlineObject14{value: val, isSet: true}
}

func (v NullableInlineObject14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


