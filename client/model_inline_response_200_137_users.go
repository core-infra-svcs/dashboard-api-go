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

// InlineResponse200137Users struct for InlineResponse200137Users
type InlineResponse200137Users struct {
	// The username for the SNMP user.
	Username *string `json:"username,omitempty"`
	// The passphrase for the SNMP user.
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewInlineResponse200137Users instantiates a new InlineResponse200137Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200137Users() *InlineResponse200137Users {
	this := InlineResponse200137Users{}
	return &this
}

// NewInlineResponse200137UsersWithDefaults instantiates a new InlineResponse200137Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200137UsersWithDefaults() *InlineResponse200137Users {
	this := InlineResponse200137Users{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse200137Users) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137Users) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse200137Users) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse200137Users) SetUsername(v string) {
	o.Username = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineResponse200137Users) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137Users) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineResponse200137Users) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineResponse200137Users) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o InlineResponse200137Users) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200137Users struct {
	value *InlineResponse200137Users
	isSet bool
}

func (v NullableInlineResponse200137Users) Get() *InlineResponse200137Users {
	return v.value
}

func (v *NullableInlineResponse200137Users) Set(val *InlineResponse200137Users) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200137Users) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200137Users) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200137Users(val *InlineResponse200137Users) *NullableInlineResponse200137Users {
	return &NullableInlineResponse200137Users{value: val, isSet: true}
}

func (v NullableInlineResponse200137Users) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200137Users) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


