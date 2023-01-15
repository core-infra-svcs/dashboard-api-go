/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryWater Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
type OrganizationsOrganizationIdSensorReadingsHistoryWater struct {
	// True if water is detected.
	Present *bool `json:"present,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryWater instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryWater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryWater() *OrganizationsOrganizationIdSensorReadingsHistoryWater {
	this := OrganizationsOrganizationIdSensorReadingsHistoryWater{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryWaterWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryWater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryWaterWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryWater {
	this := OrganizationsOrganizationIdSensorReadingsHistoryWater{}
	return &this
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryWater) GetPresent() bool {
	if o == nil || isNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryWater) GetPresentOk() (*bool, bool) {
	if o == nil || isNil(o.Present) {
    return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryWater) HasPresent() bool {
	if o != nil && !isNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryWater) SetPresent(v bool) {
	o.Present = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryWater) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Present) {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryWater struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryWater
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) Get() *OrganizationsOrganizationIdSensorReadingsHistoryWater {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryWater) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryWater(val *OrganizationsOrganizationIdSensorReadingsHistoryWater) *NullableOrganizationsOrganizationIdSensorReadingsHistoryWater {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryWater{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryWater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


