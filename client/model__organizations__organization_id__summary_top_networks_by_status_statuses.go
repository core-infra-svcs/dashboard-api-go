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

// OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses Network device statuses
type OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses struct {
	// Overall status of network
	Overall *string `json:"overall,omitempty"`
	// List of status counts by product type
	ByProductType []OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType `json:"byProductType,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses{}
	return &this
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) GetOverall() string {
	if o == nil || isNil(o.Overall) {
		var ret string
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) GetOverallOk() (*string, bool) {
	if o == nil || isNil(o.Overall) {
    return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) HasOverall() bool {
	if o != nil && !isNil(o.Overall) {
		return true
	}

	return false
}

// SetOverall gets a reference to the given string and assigns it to the Overall field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) SetOverall(v string) {
	o.Overall = &v
}

// GetByProductType returns the ByProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) GetByProductType() []OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType {
	if o == nil || isNil(o.ByProductType) {
		var ret []OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType
		return ret
	}
	return o.ByProductType
}

// GetByProductTypeOk returns a tuple with the ByProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) GetByProductTypeOk() ([]OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType, bool) {
	if o == nil || isNil(o.ByProductType) {
    return nil, false
	}
	return o.ByProductType, true
}

// HasByProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) HasByProductType() bool {
	if o != nil && !isNil(o.ByProductType) {
		return true
	}

	return false
}

// SetByProductType gets a reference to the given []OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType and assigns it to the ByProductType field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) SetByProductType(v []OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesByProductType) {
	o.ByProductType = v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Overall) {
		toSerialize["overall"] = o.Overall
	}
	if !isNil(o.ByProductType) {
		toSerialize["byProductType"] = o.ByProductType
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


