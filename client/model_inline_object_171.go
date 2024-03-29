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

// InlineObject171 struct for InlineObject171
type InlineObject171 struct {
	// List of AP serials
	Serials []string `json:"serials"`
	// AP profile ID
	ProfileId string `json:"profileId"`
}

// NewInlineObject171 instantiates a new InlineObject171 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject171(serials []string, profileId string) *InlineObject171 {
	this := InlineObject171{}
	this.Serials = serials
	this.ProfileId = profileId
	return &this
}

// NewInlineObject171WithDefaults instantiates a new InlineObject171 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject171WithDefaults() *InlineObject171 {
	this := InlineObject171{}
	return &this
}

// GetSerials returns the Serials field value
func (o *InlineObject171) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject171) SetSerials(v []string) {
	o.Serials = v
}

// GetProfileId returns the ProfileId field value
func (o *InlineObject171) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetProfileIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *InlineObject171) SetProfileId(v string) {
	o.ProfileId = v
}

func (o InlineObject171) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serials"] = o.Serials
	}
	if true {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject171 struct {
	value *InlineObject171
	isSet bool
}

func (v NullableInlineObject171) Get() *InlineObject171 {
	return v.value
}

func (v *NullableInlineObject171) Set(val *InlineObject171) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject171) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject171) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject171(val *InlineObject171) *NullableInlineObject171 {
	return &NullableInlineObject171{value: val, isSet: true}
}

func (v NullableInlineObject171) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject171) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


