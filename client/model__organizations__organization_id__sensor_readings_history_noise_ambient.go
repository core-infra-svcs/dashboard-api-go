/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient Ambient noise reading.
type OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient struct {
	// Ambient noise reading in adjusted decibels.
	Level *int32 `json:"level,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient() *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbientWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbientWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) GetLevel() int32 {
	if o == nil || isNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) GetLevelOk() (*int32, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) SetLevel(v int32) {
	o.Level = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) Get() *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient(val *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


