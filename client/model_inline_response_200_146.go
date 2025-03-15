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

// InlineResponse200146 struct for InlineResponse200146
type InlineResponse200146 struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// SNMP community string if access is 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// SNMP settings if access is 'users'.
	Users []InlineResponse200146Users `json:"users,omitempty"`
}

// NewInlineResponse200146 instantiates a new InlineResponse200146 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200146() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// NewInlineResponse200146WithDefaults instantiates a new InlineResponse200146 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200146WithDefaults() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200146) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200146) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200146) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *InlineResponse200146) GetCommunityString() string {
	if o == nil || isNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetCommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.CommunityString) {
    return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *InlineResponse200146) HasCommunityString() bool {
	if o != nil && !isNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *InlineResponse200146) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineResponse200146) GetUsers() []InlineResponse200146Users {
	if o == nil || isNil(o.Users) {
		var ret []InlineResponse200146Users
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetUsersOk() ([]InlineResponse200146Users, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineResponse200146) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []InlineResponse200146Users and assigns it to the Users field.
func (o *InlineResponse200146) SetUsers(v []InlineResponse200146Users) {
	o.Users = v
}

func (o InlineResponse200146) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200146 struct {
	value *InlineResponse200146
	isSet bool
}

func (v NullableInlineResponse200146) Get() *InlineResponse200146 {
	return v.value
}

func (v *NullableInlineResponse200146) Set(val *InlineResponse200146) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200146) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200146) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200146(val *InlineResponse200146) *NullableInlineResponse200146 {
	return &NullableInlineResponse200146{value: val, isSet: true}
}

func (v NullableInlineResponse200146) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200146) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


