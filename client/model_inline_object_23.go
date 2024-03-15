/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	Livestream *DevicesSerialSensorRelationshipsLivestream `json:"livestream,omitempty"`
}

// NewInlineObject23 instantiates a new InlineObject23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject23() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// NewInlineObject23WithDefaults instantiates a new InlineObject23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject23WithDefaults() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *InlineObject23) GetLivestream() DevicesSerialSensorRelationshipsLivestream {
	if o == nil || isNil(o.Livestream) {
		var ret DevicesSerialSensorRelationshipsLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetLivestreamOk() (*DevicesSerialSensorRelationshipsLivestream, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *InlineObject23) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given DevicesSerialSensorRelationshipsLivestream and assigns it to the Livestream field.
func (o *InlineObject23) SetLivestream(v DevicesSerialSensorRelationshipsLivestream) {
	o.Livestream = &v
}

func (o InlineObject23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject23 struct {
	value *InlineObject23
	isSet bool
}

func (v NullableInlineObject23) Get() *InlineObject23 {
	return v.value
}

func (v *NullableInlineObject23) Set(val *InlineObject23) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject23) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject23(val *InlineObject23) *NullableInlineObject23 {
	return &NullableInlineObject23{value: val, isSet: true}
}

func (v NullableInlineObject23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


