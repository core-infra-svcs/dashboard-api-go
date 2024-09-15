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

// OrganizationsOrganizationIdSensorReadingsHistoryPm25 Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.
type OrganizationsOrganizationIdSensorReadingsHistoryPm25 struct {
	// PM2.5 reading in micrograms per cubic meter.
	Concentration *int32 `json:"concentration,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryPm25 instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryPm25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryPm25() *OrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	this := OrganizationsOrganizationIdSensorReadingsHistoryPm25{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryPm25WithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryPm25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryPm25WithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	this := OrganizationsOrganizationIdSensorReadingsHistoryPm25{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPm25) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPm25) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPm25) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPm25) SetConcentration(v int32) {
	o.Concentration = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryPm25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25 struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryPm25
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) Get() *OrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryPm25) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryPm25(val *OrganizationsOrganizationIdSensorReadingsHistoryPm25) *NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25 {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPm25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


