/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage Data usage of the device
type OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage struct {
	// Total data usage of the device
	Total *float32 `json:"total,omitempty"`
	// Data usage of the device by percentage
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage() *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopDevicesByUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopDevicesByUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) GetPercentage() float32 {
	if o == nil || o.Percentage == nil {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) GetPercentageOk() (*float32, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage(val *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


