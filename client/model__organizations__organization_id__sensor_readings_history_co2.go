/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryCo2 Reading for the 'co2' metric. This will only be present if the 'metric' property equals 'co2'.
type OrganizationsOrganizationIdSensorReadingsHistoryCo2 struct {
	// CO2 reading in parts per million.
	Concentration *int32 `json:"concentration,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryCo2 instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryCo2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryCo2() *OrganizationsOrganizationIdSensorReadingsHistoryCo2 {
	this := OrganizationsOrganizationIdSensorReadingsHistoryCo2{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryCo2WithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryCo2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryCo2WithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryCo2 {
	this := OrganizationsOrganizationIdSensorReadingsHistoryCo2{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCo2) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCo2) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCo2) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCo2) SetConcentration(v int32) {
	o.Concentration = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryCo2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2 struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryCo2
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) Get() *OrganizationsOrganizationIdSensorReadingsHistoryCo2 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryCo2) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryCo2(val *OrganizationsOrganizationIdSensorReadingsHistoryCo2) *NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2 {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCo2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


