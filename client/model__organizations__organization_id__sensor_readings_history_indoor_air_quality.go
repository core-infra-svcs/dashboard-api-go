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

// OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
type OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality struct {
	// Indoor air quality score between 0 and 100.
	Score *int32 `json:"score,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality() *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	this := OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQualityWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQualityWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	this := OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) GetScore() int32 {
	if o == nil || isNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) GetScoreOk() (*int32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) SetScore(v int32) {
	o.Score = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) Get() *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality(val *OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) *NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


