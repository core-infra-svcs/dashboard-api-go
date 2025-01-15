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

// OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts Network client counts
type OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts struct {
	// Total count of clients in network
	Total *int32 `json:"total,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCountsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCountsWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) SetTotal(v int32) {
	o.Total = &v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusClientsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


