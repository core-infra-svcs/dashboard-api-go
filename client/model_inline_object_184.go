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

// InlineObject184 struct for InlineObject184
type InlineObject184 struct {
	// AP profile ID
	ProfileId string `json:"profileId"`
}

// NewInlineObject184 instantiates a new InlineObject184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject184(profileId string) *InlineObject184 {
	this := InlineObject184{}
	this.ProfileId = profileId
	return &this
}

// NewInlineObject184WithDefaults instantiates a new InlineObject184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject184WithDefaults() *InlineObject184 {
	this := InlineObject184{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *InlineObject184) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetProfileIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *InlineObject184) SetProfileId(v string) {
	o.ProfileId = v
}

func (o InlineObject184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject184 struct {
	value *InlineObject184
	isSet bool
}

func (v NullableInlineObject184) Get() *InlineObject184 {
	return v.value
}

func (v *NullableInlineObject184) Set(val *InlineObject184) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject184) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject184(val *InlineObject184) *NullableInlineObject184 {
	return &NullableInlineObject184{value: val, isSet: true}
}

func (v NullableInlineObject184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


