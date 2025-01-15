/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryCurrent Reading for the 'current' metric. This will only be present if the 'metric' property equals 'current'.
type OrganizationsOrganizationIdSensorReadingsHistoryCurrent struct {
	// Electrical current reading in amperes.
	Draw *float32 `json:"draw,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryCurrent instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryCurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryCurrent() *OrganizationsOrganizationIdSensorReadingsHistoryCurrent {
	this := OrganizationsOrganizationIdSensorReadingsHistoryCurrent{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryCurrentWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryCurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryCurrentWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryCurrent {
	this := OrganizationsOrganizationIdSensorReadingsHistoryCurrent{}
	return &this
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) GetDraw() float32 {
	if o == nil || isNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) GetDrawOk() (*float32, bool) {
	if o == nil || isNil(o.Draw) {
    return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) HasDraw() bool {
	if o != nil && !isNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) SetDraw(v float32) {
	o.Draw = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryCurrent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryCurrent
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) Get() *OrganizationsOrganizationIdSensorReadingsHistoryCurrent {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent(val *OrganizationsOrganizationIdSensorReadingsHistoryCurrent) *NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryCurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


