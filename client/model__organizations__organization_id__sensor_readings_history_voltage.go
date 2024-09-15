/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryVoltage Reading for the 'voltage' metric. This will only be present if the 'metric' property equals 'voltage'.
type OrganizationsOrganizationIdSensorReadingsHistoryVoltage struct {
	// Voltage reading in volts.
	Level *float32 `json:"level,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryVoltage instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryVoltage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryVoltage() *OrganizationsOrganizationIdSensorReadingsHistoryVoltage {
	this := OrganizationsOrganizationIdSensorReadingsHistoryVoltage{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryVoltageWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryVoltage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryVoltageWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryVoltage {
	this := OrganizationsOrganizationIdSensorReadingsHistoryVoltage{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) GetLevel() float32 {
	if o == nil || isNil(o.Level) {
		var ret float32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) GetLevelOk() (*float32, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float32 and assigns it to the Level field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) SetLevel(v float32) {
	o.Level = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryVoltage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryVoltage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) Get() *OrganizationsOrganizationIdSensorReadingsHistoryVoltage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage(val *OrganizationsOrganizationIdSensorReadingsHistoryVoltage) *NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryVoltage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


