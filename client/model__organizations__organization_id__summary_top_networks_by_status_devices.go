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

// OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices Network device information
type OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices struct {
	// URLs by product type
	ByProductType []OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType `json:"byProductType,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices() *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices{}
	return &this
}

// GetByProductType returns the ByProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) GetByProductType() []OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType {
	if o == nil || isNil(o.ByProductType) {
		var ret []OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType
		return ret
	}
	return o.ByProductType
}

// GetByProductTypeOk returns a tuple with the ByProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) GetByProductTypeOk() ([]OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType, bool) {
	if o == nil || isNil(o.ByProductType) {
    return nil, false
	}
	return o.ByProductType, true
}

// HasByProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) HasByProductType() bool {
	if o != nil && !isNil(o.ByProductType) {
		return true
	}

	return false
}

// SetByProductType gets a reference to the given []OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType and assigns it to the ByProductType field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) SetByProductType(v []OrganizationsOrganizationIdSummaryTopNetworksByStatusDevicesByProductType) {
	o.ByProductType = v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByProductType) {
		toSerialize["byProductType"] = o.ByProductType
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

