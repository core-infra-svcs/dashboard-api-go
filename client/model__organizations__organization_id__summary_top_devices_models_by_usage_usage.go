/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage Usage info in megabytes
type OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage struct {
	// Total usage in megabytes
	Total *float32 `json:"total,omitempty"`
	// Average usage in megabytes
	Average *float32 `json:"average,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage() *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) GetAverage() float32 {
	if o == nil || isNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) GetAverageOk() (*float32, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) SetAverage(v float32) {
	o.Average = &v
}

func (o OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage(val *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


