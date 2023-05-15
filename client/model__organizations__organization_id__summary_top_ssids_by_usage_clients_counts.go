/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts Counts of the clients
type OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts struct {
	// Total counts of the clients
	Total *int32 `json:"total,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCountsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCountsWithDefaults() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) SetTotal(v int32) {
	o.Total = &v
}

func (o OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts struct {
	value *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) Get() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) Set(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts {
	return &NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


