/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200247SystemsManagerCounts Aggregated license count data for Systems Manager
type InlineResponse200247SystemsManagerCounts struct {
	// The total number of Systems Manager seats
	TotalSeats *int32 `json:"totalSeats,omitempty"`
	// The number of Systems Manager seats in use
	ActiveSeats *int32 `json:"activeSeats,omitempty"`
	// The number of unused Systems Manager seats
	UnassignedSeats *int32 `json:"unassignedSeats,omitempty"`
	// The total number of enrolled Systems Manager devices
	OrgwideEnrolledDevices *int32 `json:"orgwideEnrolledDevices,omitempty"`
}

// NewInlineResponse200247SystemsManagerCounts instantiates a new InlineResponse200247SystemsManagerCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200247SystemsManagerCounts() *InlineResponse200247SystemsManagerCounts {
	this := InlineResponse200247SystemsManagerCounts{}
	return &this
}

// NewInlineResponse200247SystemsManagerCountsWithDefaults instantiates a new InlineResponse200247SystemsManagerCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200247SystemsManagerCountsWithDefaults() *InlineResponse200247SystemsManagerCounts {
	this := InlineResponse200247SystemsManagerCounts{}
	return &this
}

// GetTotalSeats returns the TotalSeats field value if set, zero value otherwise.
func (o *InlineResponse200247SystemsManagerCounts) GetTotalSeats() int32 {
	if o == nil || isNil(o.TotalSeats) {
		var ret int32
		return ret
	}
	return *o.TotalSeats
}

// GetTotalSeatsOk returns a tuple with the TotalSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200247SystemsManagerCounts) GetTotalSeatsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalSeats) {
    return nil, false
	}
	return o.TotalSeats, true
}

// HasTotalSeats returns a boolean if a field has been set.
func (o *InlineResponse200247SystemsManagerCounts) HasTotalSeats() bool {
	if o != nil && !isNil(o.TotalSeats) {
		return true
	}

	return false
}

// SetTotalSeats gets a reference to the given int32 and assigns it to the TotalSeats field.
func (o *InlineResponse200247SystemsManagerCounts) SetTotalSeats(v int32) {
	o.TotalSeats = &v
}

// GetActiveSeats returns the ActiveSeats field value if set, zero value otherwise.
func (o *InlineResponse200247SystemsManagerCounts) GetActiveSeats() int32 {
	if o == nil || isNil(o.ActiveSeats) {
		var ret int32
		return ret
	}
	return *o.ActiveSeats
}

// GetActiveSeatsOk returns a tuple with the ActiveSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200247SystemsManagerCounts) GetActiveSeatsOk() (*int32, bool) {
	if o == nil || isNil(o.ActiveSeats) {
    return nil, false
	}
	return o.ActiveSeats, true
}

// HasActiveSeats returns a boolean if a field has been set.
func (o *InlineResponse200247SystemsManagerCounts) HasActiveSeats() bool {
	if o != nil && !isNil(o.ActiveSeats) {
		return true
	}

	return false
}

// SetActiveSeats gets a reference to the given int32 and assigns it to the ActiveSeats field.
func (o *InlineResponse200247SystemsManagerCounts) SetActiveSeats(v int32) {
	o.ActiveSeats = &v
}

// GetUnassignedSeats returns the UnassignedSeats field value if set, zero value otherwise.
func (o *InlineResponse200247SystemsManagerCounts) GetUnassignedSeats() int32 {
	if o == nil || isNil(o.UnassignedSeats) {
		var ret int32
		return ret
	}
	return *o.UnassignedSeats
}

// GetUnassignedSeatsOk returns a tuple with the UnassignedSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200247SystemsManagerCounts) GetUnassignedSeatsOk() (*int32, bool) {
	if o == nil || isNil(o.UnassignedSeats) {
    return nil, false
	}
	return o.UnassignedSeats, true
}

// HasUnassignedSeats returns a boolean if a field has been set.
func (o *InlineResponse200247SystemsManagerCounts) HasUnassignedSeats() bool {
	if o != nil && !isNil(o.UnassignedSeats) {
		return true
	}

	return false
}

// SetUnassignedSeats gets a reference to the given int32 and assigns it to the UnassignedSeats field.
func (o *InlineResponse200247SystemsManagerCounts) SetUnassignedSeats(v int32) {
	o.UnassignedSeats = &v
}

// GetOrgwideEnrolledDevices returns the OrgwideEnrolledDevices field value if set, zero value otherwise.
func (o *InlineResponse200247SystemsManagerCounts) GetOrgwideEnrolledDevices() int32 {
	if o == nil || isNil(o.OrgwideEnrolledDevices) {
		var ret int32
		return ret
	}
	return *o.OrgwideEnrolledDevices
}

// GetOrgwideEnrolledDevicesOk returns a tuple with the OrgwideEnrolledDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200247SystemsManagerCounts) GetOrgwideEnrolledDevicesOk() (*int32, bool) {
	if o == nil || isNil(o.OrgwideEnrolledDevices) {
    return nil, false
	}
	return o.OrgwideEnrolledDevices, true
}

// HasOrgwideEnrolledDevices returns a boolean if a field has been set.
func (o *InlineResponse200247SystemsManagerCounts) HasOrgwideEnrolledDevices() bool {
	if o != nil && !isNil(o.OrgwideEnrolledDevices) {
		return true
	}

	return false
}

// SetOrgwideEnrolledDevices gets a reference to the given int32 and assigns it to the OrgwideEnrolledDevices field.
func (o *InlineResponse200247SystemsManagerCounts) SetOrgwideEnrolledDevices(v int32) {
	o.OrgwideEnrolledDevices = &v
}

func (o InlineResponse200247SystemsManagerCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TotalSeats) {
		toSerialize["totalSeats"] = o.TotalSeats
	}
	if !isNil(o.ActiveSeats) {
		toSerialize["activeSeats"] = o.ActiveSeats
	}
	if !isNil(o.UnassignedSeats) {
		toSerialize["unassignedSeats"] = o.UnassignedSeats
	}
	if !isNil(o.OrgwideEnrolledDevices) {
		toSerialize["orgwideEnrolledDevices"] = o.OrgwideEnrolledDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200247SystemsManagerCounts struct {
	value *InlineResponse200247SystemsManagerCounts
	isSet bool
}

func (v NullableInlineResponse200247SystemsManagerCounts) Get() *InlineResponse200247SystemsManagerCounts {
	return v.value
}

func (v *NullableInlineResponse200247SystemsManagerCounts) Set(val *InlineResponse200247SystemsManagerCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200247SystemsManagerCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200247SystemsManagerCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200247SystemsManagerCounts(val *InlineResponse200247SystemsManagerCounts) *NullableInlineResponse200247SystemsManagerCounts {
	return &NullableInlineResponse200247SystemsManagerCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200247SystemsManagerCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200247SystemsManagerCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


