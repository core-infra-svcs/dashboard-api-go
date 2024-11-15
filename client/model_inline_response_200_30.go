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

// InlineResponse20030 struct for InlineResponse20030
type InlineResponse20030 struct {
	Livestream *InlineResponse20030Livestream `json:"livestream,omitempty"`
}

// NewInlineResponse20030 instantiates a new InlineResponse20030 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030() *InlineResponse20030 {
	this := InlineResponse20030{}
	return &this
}

// NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030WithDefaults() *InlineResponse20030 {
	this := InlineResponse20030{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *InlineResponse20030) GetLivestream() InlineResponse20030Livestream {
	if o == nil || isNil(o.Livestream) {
		var ret InlineResponse20030Livestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetLivestreamOk() (*InlineResponse20030Livestream, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *InlineResponse20030) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given InlineResponse20030Livestream and assigns it to the Livestream field.
func (o *InlineResponse20030) SetLivestream(v InlineResponse20030Livestream) {
	o.Livestream = &v
}

func (o InlineResponse20030) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030 struct {
	value *InlineResponse20030
	isSet bool
}

func (v NullableInlineResponse20030) Get() *InlineResponse20030 {
	return v.value
}

func (v *NullableInlineResponse20030) Set(val *InlineResponse20030) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030(val *InlineResponse20030) *NullableInlineResponse20030 {
	return &NullableInlineResponse20030{value: val, isSet: true}
}

func (v NullableInlineResponse20030) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


