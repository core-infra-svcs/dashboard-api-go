/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200220Profile The profile the network is attached to
type InlineResponse200220Profile struct {
	// ID of the profile
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200220Profile instantiates a new InlineResponse200220Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200220Profile() *InlineResponse200220Profile {
	this := InlineResponse200220Profile{}
	return &this
}

// NewInlineResponse200220ProfileWithDefaults instantiates a new InlineResponse200220Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200220ProfileWithDefaults() *InlineResponse200220Profile {
	this := InlineResponse200220Profile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200220Profile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200220Profile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200220Profile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200220Profile) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200220Profile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200220Profile struct {
	value *InlineResponse200220Profile
	isSet bool
}

func (v NullableInlineResponse200220Profile) Get() *InlineResponse200220Profile {
	return v.value
}

func (v *NullableInlineResponse200220Profile) Set(val *InlineResponse200220Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200220Profile) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200220Profile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200220Profile(val *InlineResponse200220Profile) *NullableInlineResponse200220Profile {
	return &NullableInlineResponse200220Profile{value: val, isSet: true}
}

func (v NullableInlineResponse200220Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200220Profile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


