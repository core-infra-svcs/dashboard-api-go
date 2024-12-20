/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject252 struct for InlineObject252
type InlineObject252 struct {
	// The ID of the SM license to renew. This license must already be assigned to an SM network
	LicenseIdToRenew string `json:"licenseIdToRenew"`
	// The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'
	UnusedLicenseId string `json:"unusedLicenseId"`
}

// NewInlineObject252 instantiates a new InlineObject252 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject252(licenseIdToRenew string, unusedLicenseId string) *InlineObject252 {
	this := InlineObject252{}
	this.LicenseIdToRenew = licenseIdToRenew
	this.UnusedLicenseId = unusedLicenseId
	return &this
}

// NewInlineObject252WithDefaults instantiates a new InlineObject252 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject252WithDefaults() *InlineObject252 {
	this := InlineObject252{}
	return &this
}

// GetLicenseIdToRenew returns the LicenseIdToRenew field value
func (o *InlineObject252) GetLicenseIdToRenew() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseIdToRenew
}

// GetLicenseIdToRenewOk returns a tuple with the LicenseIdToRenew field value
// and a boolean to check if the value has been set.
func (o *InlineObject252) GetLicenseIdToRenewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LicenseIdToRenew, true
}

// SetLicenseIdToRenew sets field value
func (o *InlineObject252) SetLicenseIdToRenew(v string) {
	o.LicenseIdToRenew = v
}

// GetUnusedLicenseId returns the UnusedLicenseId field value
func (o *InlineObject252) GetUnusedLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnusedLicenseId
}

// GetUnusedLicenseIdOk returns a tuple with the UnusedLicenseId field value
// and a boolean to check if the value has been set.
func (o *InlineObject252) GetUnusedLicenseIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnusedLicenseId, true
}

// SetUnusedLicenseId sets field value
func (o *InlineObject252) SetUnusedLicenseId(v string) {
	o.UnusedLicenseId = v
}

func (o InlineObject252) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["licenseIdToRenew"] = o.LicenseIdToRenew
	}
	if true {
		toSerialize["unusedLicenseId"] = o.UnusedLicenseId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject252 struct {
	value *InlineObject252
	isSet bool
}

func (v NullableInlineObject252) Get() *InlineObject252 {
	return v.value
}

func (v *NullableInlineObject252) Set(val *InlineObject252) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject252) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject252) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject252(val *InlineObject252) *NullableInlineObject252 {
	return &NullableInlineObject252{value: val, isSet: true}
}

func (v NullableInlineObject252) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject252) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


