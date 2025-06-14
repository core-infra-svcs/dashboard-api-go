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

// OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed Details about the failed unit
type OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed struct {
	Chassis *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis `json:"chassis,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedWithDefaults instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedWithDefaults() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed{}
	return &this
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) GetChassis() OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis {
	if o == nil || isNil(o.Chassis) {
		var ret OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) GetChassisOk() (*OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis, bool) {
	if o == nil || isNil(o.Chassis) {
    return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) HasChassis() bool {
	if o != nil && !isNil(o.Chassis) {
		return true
	}

	return false
}

// SetChassis gets a reference to the given OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis and assigns it to the Chassis field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) SetChassis(v OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) {
	o.Chassis = &v
}

func (o OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Chassis) {
		toSerialize["chassis"] = o.Chassis
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed struct {
	value *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) Get() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) Set(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed {
	return &NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


