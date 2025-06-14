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

// OrganizationsOrganizationIdSensorReadingsHistoryApparentPower Reading for the 'apparentPower' metric. This will only be present if the 'metric' property equals 'apparentPower'.
type OrganizationsOrganizationIdSensorReadingsHistoryApparentPower struct {
	// Apparent power reading in volt-amperes.
	Draw *float32 `json:"draw,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryApparentPower instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryApparentPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryApparentPower() *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower {
	this := OrganizationsOrganizationIdSensorReadingsHistoryApparentPower{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryApparentPowerWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryApparentPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryApparentPowerWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower {
	this := OrganizationsOrganizationIdSensorReadingsHistoryApparentPower{}
	return &this
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) GetDraw() float32 {
	if o == nil || isNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) GetDrawOk() (*float32, bool) {
	if o == nil || isNil(o.Draw) {
    return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) HasDraw() bool {
	if o != nil && !isNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) SetDraw(v float32) {
	o.Draw = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) Get() *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower(val *OrganizationsOrganizationIdSensorReadingsHistoryApparentPower) *NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryApparentPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


