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

// InlineObject228 struct for InlineObject228
type InlineObject228 struct {
	// The ID of the SM license to assign seats from
	LicenseId string `json:"licenseId"`
	// The ID of the SM network to assign the seats to
	NetworkId string `json:"networkId"`
	// The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewInlineObject228 instantiates a new InlineObject228 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject228(licenseId string, networkId string, seatCount int32) *InlineObject228 {
	this := InlineObject228{}
	this.LicenseId = licenseId
	this.NetworkId = networkId
	this.SeatCount = seatCount
	return &this
}

// NewInlineObject228WithDefaults instantiates a new InlineObject228 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject228WithDefaults() *InlineObject228 {
	this := InlineObject228{}
	return &this
}

// GetLicenseId returns the LicenseId field value
func (o *InlineObject228) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetLicenseIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *InlineObject228) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetNetworkId returns the NetworkId field value
func (o *InlineObject228) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *InlineObject228) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetSeatCount returns the SeatCount field value
func (o *InlineObject228) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetSeatCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *InlineObject228) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o InlineObject228) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject228 struct {
	value *InlineObject228
	isSet bool
}

func (v NullableInlineObject228) Get() *InlineObject228 {
	return v.value
}

func (v *NullableInlineObject228) Set(val *InlineObject228) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject228) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject228) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject228(val *InlineObject228) *NullableInlineObject228 {
	return &NullableInlineObject228{value: val, isSet: true}
}

func (v NullableInlineObject228) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject228) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


