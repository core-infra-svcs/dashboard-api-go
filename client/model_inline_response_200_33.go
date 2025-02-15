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

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	Livestream *InlineResponse20033Livestream `json:"livestream,omitempty"`
}

// NewInlineResponse20033 instantiates a new InlineResponse20033 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20033() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20033WithDefaults() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *InlineResponse20033) GetLivestream() InlineResponse20033Livestream {
	if o == nil || isNil(o.Livestream) {
		var ret InlineResponse20033Livestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetLivestreamOk() (*InlineResponse20033Livestream, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *InlineResponse20033) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given InlineResponse20033Livestream and assigns it to the Livestream field.
func (o *InlineResponse20033) SetLivestream(v InlineResponse20033Livestream) {
	o.Livestream = &v
}

func (o InlineResponse20033) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20033 struct {
	value *InlineResponse20033
	isSet bool
}

func (v NullableInlineResponse20033) Get() *InlineResponse20033 {
	return v.value
}

func (v *NullableInlineResponse20033) Set(val *InlineResponse20033) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20033) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20033) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20033(val *InlineResponse20033) *NullableInlineResponse20033 {
	return &NullableInlineResponse20033{value: val, isSet: true}
}

func (v NullableInlineResponse20033) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20033) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


