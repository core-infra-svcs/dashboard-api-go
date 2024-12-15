/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis Details about the failed unit chassis
type OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis struct {
	// The name of the failed chassis unit
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassisWithDefaults instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassisWithDefaults() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis struct {
	value *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) Get() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) Set(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis {
	return &NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailedChassis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


