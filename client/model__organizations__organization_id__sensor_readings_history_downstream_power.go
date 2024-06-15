/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower Reading for the 'downstreamPower' metric. This will only be present if the 'metric' property equals 'downstreamPower'.
type OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower struct {
	// True if power is turned on to the device that is connected downstream of the MT40 power monitor.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower() *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower {
	this := OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPowerWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPowerWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower {
	this := OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) Get() *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower(val *OrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) *NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDownstreamPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


