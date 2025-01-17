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

// InlineResponse20016Apns struct for InlineResponse20016Apns
type InlineResponse20016Apns struct {
	// APN name.
	Name string `json:"name"`
	// IP versions to support (permitted values include 'ipv4', 'ipv6').
	AllowedIpTypes []string `json:"allowedIpTypes"`
	Authentication *InlineResponse20016Authentication `json:"authentication,omitempty"`
}

// NewInlineResponse20016Apns instantiates a new InlineResponse20016Apns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016Apns(name string, allowedIpTypes []string) *InlineResponse20016Apns {
	this := InlineResponse20016Apns{}
	this.Name = name
	this.AllowedIpTypes = allowedIpTypes
	return &this
}

// NewInlineResponse20016ApnsWithDefaults instantiates a new InlineResponse20016Apns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016ApnsWithDefaults() *InlineResponse20016Apns {
	this := InlineResponse20016Apns{}
	return &this
}

// GetName returns the Name field value
func (o *InlineResponse20016Apns) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Apns) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20016Apns) SetName(v string) {
	o.Name = v
}

// GetAllowedIpTypes returns the AllowedIpTypes field value
func (o *InlineResponse20016Apns) GetAllowedIpTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIpTypes
}

// GetAllowedIpTypesOk returns a tuple with the AllowedIpTypes field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Apns) GetAllowedIpTypesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedIpTypes, true
}

// SetAllowedIpTypes sets field value
func (o *InlineResponse20016Apns) SetAllowedIpTypes(v []string) {
	o.AllowedIpTypes = v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse20016Apns) GetAuthentication() InlineResponse20016Authentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse20016Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Apns) GetAuthenticationOk() (*InlineResponse20016Authentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse20016Apns) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse20016Authentication and assigns it to the Authentication field.
func (o *InlineResponse20016Apns) SetAuthentication(v InlineResponse20016Authentication) {
	o.Authentication = &v
}

func (o InlineResponse20016Apns) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20016Apns struct {
	value *InlineResponse20016Apns
	isSet bool
}

func (v NullableInlineResponse20016Apns) Get() *InlineResponse20016Apns {
	return v.value
}

func (v *NullableInlineResponse20016Apns) Set(val *InlineResponse20016Apns) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016Apns) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016Apns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016Apns(val *InlineResponse20016Apns) *NullableInlineResponse20016Apns {
	return &NullableInlineResponse20016Apns{value: val, isSet: true}
}

func (v NullableInlineResponse20016Apns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016Apns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


