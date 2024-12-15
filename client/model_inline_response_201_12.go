/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20112 struct for InlineResponse20112
type InlineResponse20112 struct {
	// List of updated AP serials
	Serials []string `json:"serials,omitempty"`
	// AP profile ID
	ProfileId *string `json:"profileId,omitempty"`
}

// NewInlineResponse20112 instantiates a new InlineResponse20112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20112() *InlineResponse20112 {
	this := InlineResponse20112{}
	return &this
}

// NewInlineResponse20112WithDefaults instantiates a new InlineResponse20112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20112WithDefaults() *InlineResponse20112 {
	this := InlineResponse20112{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse20112) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20112) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse20112) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse20112) SetSerials(v []string) {
	o.Serials = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *InlineResponse20112) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20112) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *InlineResponse20112) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *InlineResponse20112) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o InlineResponse20112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20112 struct {
	value *InlineResponse20112
	isSet bool
}

func (v NullableInlineResponse20112) Get() *InlineResponse20112 {
	return v.value
}

func (v *NullableInlineResponse20112) Set(val *InlineResponse20112) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20112) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20112(val *InlineResponse20112) *NullableInlineResponse20112 {
	return &NullableInlineResponse20112{value: val, isSet: true}
}

func (v NullableInlineResponse20112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


