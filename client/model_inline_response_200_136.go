/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200136 struct for InlineResponse200136
type InlineResponse200136 struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// SNMP community string if access is 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// SNMP settings if access is 'users'.
	Users []InlineResponse200136Users `json:"users,omitempty"`
}

// NewInlineResponse200136 instantiates a new InlineResponse200136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200136() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200136WithDefaults() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200136) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200136) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200136) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *InlineResponse200136) GetCommunityString() string {
	if o == nil || isNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetCommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.CommunityString) {
    return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *InlineResponse200136) HasCommunityString() bool {
	if o != nil && !isNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *InlineResponse200136) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineResponse200136) GetUsers() []InlineResponse200136Users {
	if o == nil || isNil(o.Users) {
		var ret []InlineResponse200136Users
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetUsersOk() ([]InlineResponse200136Users, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineResponse200136) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []InlineResponse200136Users and assigns it to the Users field.
func (o *InlineResponse200136) SetUsers(v []InlineResponse200136Users) {
	o.Users = v
}

func (o InlineResponse200136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.CommunityString) {
		toSerialize["communityString"] = o.CommunityString
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200136 struct {
	value *InlineResponse200136
	isSet bool
}

func (v NullableInlineResponse200136) Get() *InlineResponse200136 {
	return v.value
}

func (v *NullableInlineResponse200136) Set(val *InlineResponse200136) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200136) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200136(val *InlineResponse200136) *NullableInlineResponse200136 {
	return &NullableInlineResponse200136{value: val, isSet: true}
}

func (v NullableInlineResponse200136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


