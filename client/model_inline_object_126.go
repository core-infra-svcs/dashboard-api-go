/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject126 struct for InlineObject126
type InlineObject126 struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// The SNMP community string. Only relevant if 'access' is set to 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// The list of SNMP users. Only relevant if 'access' is set to 'users'.
	Users []NetworksNetworkIdSnmpUsers `json:"users,omitempty"`
}

// NewInlineObject126 instantiates a new InlineObject126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject126() *InlineObject126 {
	this := InlineObject126{}
	return &this
}

// NewInlineObject126WithDefaults instantiates a new InlineObject126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject126WithDefaults() *InlineObject126 {
	this := InlineObject126{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineObject126) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineObject126) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineObject126) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *InlineObject126) GetCommunityString() string {
	if o == nil || isNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetCommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.CommunityString) {
    return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *InlineObject126) HasCommunityString() bool {
	if o != nil && !isNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *InlineObject126) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineObject126) GetUsers() []NetworksNetworkIdSnmpUsers {
	if o == nil || isNil(o.Users) {
		var ret []NetworksNetworkIdSnmpUsers
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetUsersOk() ([]NetworksNetworkIdSnmpUsers, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineObject126) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []NetworksNetworkIdSnmpUsers and assigns it to the Users field.
func (o *InlineObject126) SetUsers(v []NetworksNetworkIdSnmpUsers) {
	o.Users = v
}

func (o InlineObject126) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject126 struct {
	value *InlineObject126
	isSet bool
}

func (v NullableInlineObject126) Get() *InlineObject126 {
	return v.value
}

func (v *NullableInlineObject126) Set(val *InlineObject126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject126(val *InlineObject126) *NullableInlineObject126 {
	return &NullableInlineObject126{value: val, isSet: true}
}

func (v NullableInlineObject126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


