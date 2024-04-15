/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200175 struct for InlineResponse200175
type InlineResponse200175 struct {
	// AP profile ID
	ProfileId *string `json:"profileId,omitempty"`
}

// NewInlineResponse200175 instantiates a new InlineResponse200175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200175() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// NewInlineResponse200175WithDefaults instantiates a new InlineResponse200175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200175WithDefaults() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *InlineResponse200175) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *InlineResponse200175) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *InlineResponse200175) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o InlineResponse200175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200175 struct {
	value *InlineResponse200175
	isSet bool
}

func (v NullableInlineResponse200175) Get() *InlineResponse200175 {
	return v.value
}

func (v *NullableInlineResponse200175) Set(val *InlineResponse200175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200175(val *InlineResponse200175) *NullableInlineResponse200175 {
	return &NullableInlineResponse200175{value: val, isSet: true}
}

func (v NullableInlineResponse200175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


