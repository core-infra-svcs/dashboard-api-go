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

// OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType struct for OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType
type OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType struct {
	// Product type
	ProductType *string `json:"productType,omitempty"`
	Counts *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductTypeWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductTypeWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType{}
	return &this
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) SetProductType(v string) {
	o.ProductType = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) GetCounts() OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) GetCountsOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) SetCounts(v OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


