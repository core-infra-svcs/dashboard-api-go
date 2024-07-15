/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryTvoc Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
type OrganizationsOrganizationIdSensorReadingsHistoryTvoc struct {
	// TVOC reading in micrograms per cubic meter.
	Concentration *int32 `json:"concentration,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryTvoc instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryTvoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryTvoc() *OrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	this := OrganizationsOrganizationIdSensorReadingsHistoryTvoc{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryTvocWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryTvoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryTvocWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	this := OrganizationsOrganizationIdSensorReadingsHistoryTvoc{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) SetConcentration(v int32) {
	o.Concentration = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryTvoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryTvoc
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) Get() *OrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc(val *OrganizationsOrganizationIdSensorReadingsHistoryTvoc) *NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryTvoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


