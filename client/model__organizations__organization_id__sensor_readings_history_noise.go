/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryNoise Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
type OrganizationsOrganizationIdSensorReadingsHistoryNoise struct {
	Ambient *OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient `json:"ambient,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNoise instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryNoise() *OrganizationsOrganizationIdSensorReadingsHistoryNoise {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNoise{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryNoiseWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryNoise {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNoise{}
	return &this
}

// GetAmbient returns the Ambient field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoise) GetAmbient() OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient {
	if o == nil || isNil(o.Ambient) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient
		return ret
	}
	return *o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoise) GetAmbientOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient, bool) {
	if o == nil || isNil(o.Ambient) {
    return nil, false
	}
	return o.Ambient, true
}

// HasAmbient returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoise) HasAmbient() bool {
	if o != nil && !isNil(o.Ambient) {
		return true
	}

	return false
}

// SetAmbient gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient and assigns it to the Ambient field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNoise) SetAmbient(v OrganizationsOrganizationIdSensorReadingsHistoryNoiseAmbient) {
	o.Ambient = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryNoise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ambient) {
		toSerialize["ambient"] = o.Ambient
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryNoise
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) Get() *OrganizationsOrganizationIdSensorReadingsHistoryNoise {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryNoise(val *OrganizationsOrganizationIdSensorReadingsHistoryNoise) *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


