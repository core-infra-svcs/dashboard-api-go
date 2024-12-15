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

// InlineObject251 struct for InlineObject251
type InlineObject251 struct {
	// The ID of the organization to move the SM seats to
	DestOrganizationId string `json:"destOrganizationId"`
	// The ID of the SM license to move the seats from
	LicenseId string `json:"licenseId"`
	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewInlineObject251 instantiates a new InlineObject251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject251(destOrganizationId string, licenseId string, seatCount int32) *InlineObject251 {
	this := InlineObject251{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseId = licenseId
	this.SeatCount = seatCount
	return &this
}

// NewInlineObject251WithDefaults instantiates a new InlineObject251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject251WithDefaults() *InlineObject251 {
	this := InlineObject251{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *InlineObject251) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *InlineObject251) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseId returns the LicenseId field value
func (o *InlineObject251) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetLicenseIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *InlineObject251) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetSeatCount returns the SeatCount field value
func (o *InlineObject251) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetSeatCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *InlineObject251) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o InlineObject251) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject251 struct {
	value *InlineObject251
	isSet bool
}

func (v NullableInlineObject251) Get() *InlineObject251 {
	return v.value
}

func (v *NullableInlineObject251) Set(val *InlineObject251) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject251) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject251(val *InlineObject251) *NullableInlineObject251 {
	return &NullableInlineObject251{value: val, isSet: true}
}

func (v NullableInlineObject251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


