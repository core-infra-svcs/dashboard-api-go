/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject143 struct for InlineObject143
type InlineObject143 struct {
	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize *int32 `json:"defaultMtuSize,omitempty"`
	// Override MTU size for individual switches or switch templates. An empty array will clear overrides.
	Overrides []InlineResponse200157Overrides `json:"overrides,omitempty"`
}

// NewInlineObject143 instantiates a new InlineObject143 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject143() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// NewInlineObject143WithDefaults instantiates a new InlineObject143 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject143WithDefaults() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// GetDefaultMtuSize returns the DefaultMtuSize field value if set, zero value otherwise.
func (o *InlineObject143) GetDefaultMtuSize() int32 {
	if o == nil || isNil(o.DefaultMtuSize) {
		var ret int32
		return ret
	}
	return *o.DefaultMtuSize
}

// GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetDefaultMtuSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DefaultMtuSize) {
    return nil, false
	}
	return o.DefaultMtuSize, true
}

// HasDefaultMtuSize returns a boolean if a field has been set.
func (o *InlineObject143) HasDefaultMtuSize() bool {
	if o != nil && !isNil(o.DefaultMtuSize) {
		return true
	}

	return false
}

// SetDefaultMtuSize gets a reference to the given int32 and assigns it to the DefaultMtuSize field.
func (o *InlineObject143) SetDefaultMtuSize(v int32) {
	o.DefaultMtuSize = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineObject143) GetOverrides() []InlineResponse200157Overrides {
	if o == nil || isNil(o.Overrides) {
		var ret []InlineResponse200157Overrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetOverridesOk() ([]InlineResponse200157Overrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineObject143) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []InlineResponse200157Overrides and assigns it to the Overrides field.
func (o *InlineObject143) SetOverrides(v []InlineResponse200157Overrides) {
	o.Overrides = v
}

func (o InlineObject143) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultMtuSize) {
		toSerialize["defaultMtuSize"] = o.DefaultMtuSize
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject143 struct {
	value *InlineObject143
	isSet bool
}

func (v NullableInlineObject143) Get() *InlineObject143 {
	return v.value
}

func (v *NullableInlineObject143) Set(val *InlineObject143) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject143) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject143) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject143(val *InlineObject143) *NullableInlineObject143 {
	return &NullableInlineObject143{value: val, isSet: true}
}

func (v NullableInlineObject143) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject143) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


