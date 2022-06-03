/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject180 struct for InlineObject180
type InlineObject180 struct {
	// The ID of the SM license to assign seats from
	LicenseId string `json:"licenseId"`
	// The ID of the SM network to assign the seats to
	NetworkId string `json:"networkId"`
	// The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewInlineObject180 instantiates a new InlineObject180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject180(licenseId string, networkId string, seatCount int32) *InlineObject180 {
	this := InlineObject180{}
	this.LicenseId = licenseId
	this.NetworkId = networkId
	this.SeatCount = seatCount
	return &this
}

// NewInlineObject180WithDefaults instantiates a new InlineObject180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject180WithDefaults() *InlineObject180 {
	this := InlineObject180{}
	return &this
}

// GetLicenseId returns the LicenseId field value
func (o *InlineObject180) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetLicenseIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *InlineObject180) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetNetworkId returns the NetworkId field value
func (o *InlineObject180) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetNetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *InlineObject180) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetSeatCount returns the SeatCount field value
func (o *InlineObject180) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetSeatCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *InlineObject180) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o InlineObject180) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["licenseId"] = o.LicenseId
	}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if true {
		toSerialize["seatCount"] = o.SeatCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject180 struct {
	value *InlineObject180
	isSet bool
}

func (v NullableInlineObject180) Get() *InlineObject180 {
	return v.value
}

func (v *NullableInlineObject180) Set(val *InlineObject180) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject180) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject180(val *InlineObject180) *NullableInlineObject180 {
	return &NullableInlineObject180{value: val, isSet: true}
}

func (v NullableInlineObject180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


