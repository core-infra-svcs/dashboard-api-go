/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200137 struct for InlineResponse200137
type InlineResponse200137 struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// SNMP community string if access is 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// SNMP settings if access is 'users'.
	Users []InlineResponse200137Users `json:"users,omitempty"`
}

// NewInlineResponse200137 instantiates a new InlineResponse200137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200137() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// NewInlineResponse200137WithDefaults instantiates a new InlineResponse200137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200137WithDefaults() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200137) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200137) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200137) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *InlineResponse200137) GetCommunityString() string {
	if o == nil || isNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetCommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.CommunityString) {
    return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *InlineResponse200137) HasCommunityString() bool {
	if o != nil && !isNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *InlineResponse200137) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineResponse200137) GetUsers() []InlineResponse200137Users {
	if o == nil || isNil(o.Users) {
		var ret []InlineResponse200137Users
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetUsersOk() ([]InlineResponse200137Users, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineResponse200137) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []InlineResponse200137Users and assigns it to the Users field.
func (o *InlineResponse200137) SetUsers(v []InlineResponse200137Users) {
	o.Users = v
}

func (o InlineResponse200137) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200137 struct {
	value *InlineResponse200137
	isSet bool
}

func (v NullableInlineResponse200137) Get() *InlineResponse200137 {
	return v.value
}

func (v *NullableInlineResponse200137) Set(val *InlineResponse200137) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200137) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200137(val *InlineResponse200137) *NullableInlineResponse200137 {
	return &NullableInlineResponse200137{value: val, isSet: true}
}

func (v NullableInlineResponse200137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


