/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject275 struct for InlineObject275
type InlineObject275 struct {
	// The ID of the organization to move the SM seats to
	DestOrganizationId string `json:"destOrganizationId"`
	// The ID of the SM license to move the seats from
	LicenseId string `json:"licenseId"`
	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewInlineObject275 instantiates a new InlineObject275 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject275(destOrganizationId string, licenseId string, seatCount int32) *InlineObject275 {
	this := InlineObject275{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseId = licenseId
	this.SeatCount = seatCount
	return &this
}

// NewInlineObject275WithDefaults instantiates a new InlineObject275 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject275WithDefaults() *InlineObject275 {
	this := InlineObject275{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *InlineObject275) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *InlineObject275) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *InlineObject275) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseId returns the LicenseId field value
func (o *InlineObject275) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *InlineObject275) GetLicenseIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *InlineObject275) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetSeatCount returns the SeatCount field value
func (o *InlineObject275) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *InlineObject275) GetSeatCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *InlineObject275) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o InlineObject275) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseId"] = o.LicenseId
	}
	if true {
		toSerialize["seatCount"] = o.SeatCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject275 struct {
	value *InlineObject275
	isSet bool
}

func (v NullableInlineObject275) Get() *InlineObject275 {
	return v.value
}

func (v *NullableInlineObject275) Set(val *InlineObject275) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject275) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject275) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject275(val *InlineObject275) *NullableInlineObject275 {
	return &NullableInlineObject275{value: val, isSet: true}
}

func (v NullableInlineObject275) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject275) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


