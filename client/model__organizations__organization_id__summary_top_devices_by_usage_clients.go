/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopDevicesByUsageClients Clients
type OrganizationsOrganizationIdSummaryTopDevicesByUsageClients struct {
	Counts *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopDevicesByUsageClients instantiates a new OrganizationsOrganizationIdSummaryTopDevicesByUsageClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopDevicesByUsageClients() *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopDevicesByUsageClients{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopDevicesByUsageClientsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopDevicesByUsageClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopDevicesByUsageClientsWithDefaults() *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopDevicesByUsageClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) GetCounts() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) GetCountsOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) SetCounts(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients struct {
	value *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) Get() *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) Set(val *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients(val *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients {
	return &NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


