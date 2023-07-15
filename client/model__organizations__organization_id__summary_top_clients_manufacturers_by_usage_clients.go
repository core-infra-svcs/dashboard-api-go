/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients Clients info
type OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients struct {
	Counts *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsWithDefaults() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) GetCounts() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) GetCountsOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) SetCounts(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients struct {
	value *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) Get() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) Set(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	return &NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


