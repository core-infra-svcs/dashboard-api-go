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

// OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts Counts of clients
type OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts struct {
	// Total counts of clients
	Total *int32 `json:"total,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCountsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCountsWithDefaults() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) SetTotal(v int32) {
	o.Total = &v
}

func (o OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts struct {
	value *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) Get() *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) Set(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts(val *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts {
	return &NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClientsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


