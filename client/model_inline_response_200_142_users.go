/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200142Users struct for InlineResponse200142Users
type InlineResponse200142Users struct {
	// The username for the SNMP user.
	Username *string `json:"username,omitempty"`
	// The passphrase for the SNMP user.
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewInlineResponse200142Users instantiates a new InlineResponse200142Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200142Users() *InlineResponse200142Users {
	this := InlineResponse200142Users{}
	return &this
}

// NewInlineResponse200142UsersWithDefaults instantiates a new InlineResponse200142Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200142UsersWithDefaults() *InlineResponse200142Users {
	this := InlineResponse200142Users{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse200142Users) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142Users) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse200142Users) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse200142Users) SetUsername(v string) {
	o.Username = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineResponse200142Users) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142Users) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineResponse200142Users) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineResponse200142Users) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o InlineResponse200142Users) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200142Users struct {
	value *InlineResponse200142Users
	isSet bool
}

func (v NullableInlineResponse200142Users) Get() *InlineResponse200142Users {
	return v.value
}

func (v *NullableInlineResponse200142Users) Set(val *InlineResponse200142Users) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200142Users) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200142Users) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200142Users(val *InlineResponse200142Users) *NullableInlineResponse200142Users {
	return &NullableInlineResponse200142Users{value: val, isSet: true}
}

func (v NullableInlineResponse200142Users) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200142Users) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


