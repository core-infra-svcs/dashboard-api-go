/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200271 struct for InlineResponse200271
type InlineResponse200271 struct {
	// The ID of the organization to move the SM seats to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// The ID of the SM license to move the seats from
	LicenseId *string `json:"licenseId,omitempty"`
	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	SeatCount *int32 `json:"seatCount,omitempty"`
}

// NewInlineResponse200271 instantiates a new InlineResponse200271 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200271() *InlineResponse200271 {
	this := InlineResponse200271{}
	return &this
}

// NewInlineResponse200271WithDefaults instantiates a new InlineResponse200271 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200271WithDefaults() *InlineResponse200271 {
	this := InlineResponse200271{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200271) GetDestOrganizationId() string {
	if o == nil || isNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200271) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.DestOrganizationId) {
    return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200271) HasDestOrganizationId() bool {
	if o != nil && !isNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *InlineResponse200271) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *InlineResponse200271) GetLicenseId() string {
	if o == nil || isNil(o.LicenseId) {
		var ret string
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200271) GetLicenseIdOk() (*string, bool) {
	if o == nil || isNil(o.LicenseId) {
    return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *InlineResponse200271) HasLicenseId() bool {
	if o != nil && !isNil(o.LicenseId) {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given string and assigns it to the LicenseId field.
func (o *InlineResponse200271) SetLicenseId(v string) {
	o.LicenseId = &v
}

// GetSeatCount returns the SeatCount field value if set, zero value otherwise.
func (o *InlineResponse200271) GetSeatCount() int32 {
	if o == nil || isNil(o.SeatCount) {
		var ret int32
		return ret
	}
	return *o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200271) GetSeatCountOk() (*int32, bool) {
	if o == nil || isNil(o.SeatCount) {
    return nil, false
	}
	return o.SeatCount, true
}

// HasSeatCount returns a boolean if a field has been set.
func (o *InlineResponse200271) HasSeatCount() bool {
	if o != nil && !isNil(o.SeatCount) {
		return true
	}

	return false
}

// SetSeatCount gets a reference to the given int32 and assigns it to the SeatCount field.
func (o *InlineResponse200271) SetSeatCount(v int32) {
	o.SeatCount = &v
}

func (o InlineResponse200271) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !isNil(o.LicenseId) {
		toSerialize["licenseId"] = o.LicenseId
	}
	if !isNil(o.SeatCount) {
		toSerialize["seatCount"] = o.SeatCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200271 struct {
	value *InlineResponse200271
	isSet bool
}

func (v NullableInlineResponse200271) Get() *InlineResponse200271 {
	return v.value
}

func (v *NullableInlineResponse200271) Set(val *InlineResponse200271) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200271) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200271) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200271(val *InlineResponse200271) *NullableInlineResponse200271 {
	return &NullableInlineResponse200271{value: val, isSet: true}
}

func (v NullableInlineResponse200271) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200271) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


