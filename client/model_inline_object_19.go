/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject19 struct for InlineObject19
type InlineObject19 struct {
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
	Callback *DevicesSerialLiveToolsArpTableCallback `json:"callback,omitempty"`
}

// NewInlineObject19 instantiates a new InlineObject19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject19() *InlineObject19 {
	this := InlineObject19{}
	return &this
}

// NewInlineObject19WithDefaults instantiates a new InlineObject19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject19WithDefaults() *InlineObject19 {
	this := InlineObject19{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject19) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject19) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject19) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject19) SetCount(v int32) {
	o.Count = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject19) GetCallback() DevicesSerialLiveToolsArpTableCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsArpTableCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject19) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject19) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsArpTableCallback and assigns it to the Callback field.
func (o *InlineObject19) SetCallback(v DevicesSerialLiveToolsArpTableCallback) {
	o.Callback = &v
}

func (o InlineObject19) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject19 struct {
	value *InlineObject19
	isSet bool
}

func (v NullableInlineObject19) Get() *InlineObject19 {
	return v.value
}

func (v *NullableInlineObject19) Set(val *InlineObject19) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject19) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject19) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject19(val *InlineObject19) *NullableInlineObject19 {
	return &NullableInlineObject19{value: val, isSet: true}
}

func (v NullableInlineObject19) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject19) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


