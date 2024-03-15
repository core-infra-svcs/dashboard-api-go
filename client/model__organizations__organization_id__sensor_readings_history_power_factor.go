/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor Reading for the 'powerFactor' metric. This will only be present if the 'metric' property equals 'powerFactor'.
type OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor struct {
	// Power factor reading as a percentage.
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor() *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor {
	this := OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryPowerFactorWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryPowerFactorWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor {
	this := OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) GetPercentage() int32 {
	if o == nil || isNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) GetPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) Get() *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor(val *OrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) *NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryPowerFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


