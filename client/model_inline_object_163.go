/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject163 struct for InlineObject163
type InlineObject163 struct {
	// AP profile ID
	ProfileId string `json:"profileId"`
}

// NewInlineObject163 instantiates a new InlineObject163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject163(profileId string) *InlineObject163 {
	this := InlineObject163{}
	this.ProfileId = profileId
	return &this
}

// NewInlineObject163WithDefaults instantiates a new InlineObject163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject163WithDefaults() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *InlineObject163) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetProfileIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *InlineObject163) SetProfileId(v string) {
	o.ProfileId = v
}

func (o InlineObject163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject163 struct {
	value *InlineObject163
	isSet bool
}

func (v NullableInlineObject163) Get() *InlineObject163 {
	return v.value
}

func (v *NullableInlineObject163) Set(val *InlineObject163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject163(val *InlineObject163) *NullableInlineObject163 {
	return &NullableInlineObject163{value: val, isSet: true}
}

func (v NullableInlineObject163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


