/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCellularSimsAuthentication APN authentication configurations.
type DevicesSerialCellularSimsAuthentication struct {
	// APN auth type.
	Type *string `json:"type,omitempty"`
	// APN username, if type is set.
	Username *string `json:"username,omitempty"`
	// APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Password *string `json:"password,omitempty"`
}

// NewDevicesSerialCellularSimsAuthentication instantiates a new DevicesSerialCellularSimsAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularSimsAuthentication() *DevicesSerialCellularSimsAuthentication {
	this := DevicesSerialCellularSimsAuthentication{}
	var type_ string = "none"
	this.Type = &type_
	return &this
}

// NewDevicesSerialCellularSimsAuthenticationWithDefaults instantiates a new DevicesSerialCellularSimsAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularSimsAuthenticationWithDefaults() *DevicesSerialCellularSimsAuthentication {
	this := DevicesSerialCellularSimsAuthentication{}
	var type_ string = "none"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsAuthentication) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsAuthentication) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsAuthentication) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevicesSerialCellularSimsAuthentication) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsAuthentication) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsAuthentication) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DevicesSerialCellularSimsAuthentication) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsAuthentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsAuthentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DevicesSerialCellularSimsAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o DevicesSerialCellularSimsAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularSimsAuthentication struct {
	value *DevicesSerialCellularSimsAuthentication
	isSet bool
}

func (v NullableDevicesSerialCellularSimsAuthentication) Get() *DevicesSerialCellularSimsAuthentication {
	return v.value
}

func (v *NullableDevicesSerialCellularSimsAuthentication) Set(val *DevicesSerialCellularSimsAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularSimsAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularSimsAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularSimsAuthentication(val *DevicesSerialCellularSimsAuthentication) *NullableDevicesSerialCellularSimsAuthentication {
	return &NullableDevicesSerialCellularSimsAuthentication{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularSimsAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularSimsAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


