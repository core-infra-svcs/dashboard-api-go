/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryFrequency Reading for the 'frequency' metric. This will only be present if the 'metric' property equals 'frequency'.
type OrganizationsOrganizationIdSensorReadingsHistoryFrequency struct {
	// Electrical current frequency reading in hertz.
	Level *float32 `json:"level,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryFrequency instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryFrequency() *OrganizationsOrganizationIdSensorReadingsHistoryFrequency {
	this := OrganizationsOrganizationIdSensorReadingsHistoryFrequency{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryFrequencyWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryFrequencyWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryFrequency {
	this := OrganizationsOrganizationIdSensorReadingsHistoryFrequency{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) GetLevel() float32 {
	if o == nil || isNil(o.Level) {
		var ret float32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) GetLevelOk() (*float32, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float32 and assigns it to the Level field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) SetLevel(v float32) {
	o.Level = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryFrequency
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) Get() *OrganizationsOrganizationIdSensorReadingsHistoryFrequency {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency(val *OrganizationsOrganizationIdSensorReadingsHistoryFrequency) *NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


