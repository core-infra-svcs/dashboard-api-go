/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCellularSimsApns struct for DevicesSerialCellularSimsApns
type DevicesSerialCellularSimsApns struct {
	// APN name.
	Name string `json:"name"`
	// IP versions to support (permitted values include 'ipv4', 'ipv6').
	AllowedIpTypes []string `json:"allowedIpTypes"`
	Authentication *DevicesSerialCellularSimsAuthentication `json:"authentication,omitempty"`
}

// NewDevicesSerialCellularSimsApns instantiates a new DevicesSerialCellularSimsApns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularSimsApns(name string, allowedIpTypes []string) *DevicesSerialCellularSimsApns {
	this := DevicesSerialCellularSimsApns{}
	this.Name = name
	this.AllowedIpTypes = allowedIpTypes
	return &this
}

// NewDevicesSerialCellularSimsApnsWithDefaults instantiates a new DevicesSerialCellularSimsApns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularSimsApnsWithDefaults() *DevicesSerialCellularSimsApns {
	this := DevicesSerialCellularSimsApns{}
	return &this
}

// GetName returns the Name field value
func (o *DevicesSerialCellularSimsApns) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsApns) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DevicesSerialCellularSimsApns) SetName(v string) {
	o.Name = v
}

// GetAllowedIpTypes returns the AllowedIpTypes field value
func (o *DevicesSerialCellularSimsApns) GetAllowedIpTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIpTypes
}

// GetAllowedIpTypesOk returns a tuple with the AllowedIpTypes field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsApns) GetAllowedIpTypesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedIpTypes, true
}

// SetAllowedIpTypes sets field value
func (o *DevicesSerialCellularSimsApns) SetAllowedIpTypes(v []string) {
	o.AllowedIpTypes = v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsApns) GetAuthentication() DevicesSerialCellularSimsAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret DevicesSerialCellularSimsAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsApns) GetAuthenticationOk() (*DevicesSerialCellularSimsAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsApns) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given DevicesSerialCellularSimsAuthentication and assigns it to the Authentication field.
func (o *DevicesSerialCellularSimsApns) SetAuthentication(v DevicesSerialCellularSimsAuthentication) {
	o.Authentication = &v
}

func (o DevicesSerialCellularSimsApns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["allowedIpTypes"] = o.AllowedIpTypes
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularSimsApns struct {
	value *DevicesSerialCellularSimsApns
	isSet bool
}

func (v NullableDevicesSerialCellularSimsApns) Get() *DevicesSerialCellularSimsApns {
	return v.value
}

func (v *NullableDevicesSerialCellularSimsApns) Set(val *DevicesSerialCellularSimsApns) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularSimsApns) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularSimsApns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularSimsApns(val *DevicesSerialCellularSimsApns) *NullableDevicesSerialCellularSimsApns {
	return &NullableDevicesSerialCellularSimsApns{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularSimsApns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularSimsApns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


