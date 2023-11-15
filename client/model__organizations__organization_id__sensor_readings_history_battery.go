/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryBattery Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
type OrganizationsOrganizationIdSensorReadingsHistoryBattery struct {
	// Remaining battery life.
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryBattery instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryBattery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryBattery() *OrganizationsOrganizationIdSensorReadingsHistoryBattery {
	this := OrganizationsOrganizationIdSensorReadingsHistoryBattery{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryBatteryWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryBattery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryBatteryWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryBattery {
	this := OrganizationsOrganizationIdSensorReadingsHistoryBattery{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryBattery) GetPercentage() int32 {
	if o == nil || isNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryBattery) GetPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryBattery) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryBattery) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryBattery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryBattery
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) Get() *OrganizationsOrganizationIdSensorReadingsHistoryBattery {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryBattery) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryBattery(val *OrganizationsOrganizationIdSensorReadingsHistoryBattery) *NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryBattery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


