/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps struct for OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps
type OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps struct {
	Devices OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices `json:"devices"`
	// What action to perform on devices.old after the device cloning is complete. 'remove from network' will return the device to inventory, while 'release from organization inventory' will free up the license attached to the device.
	AfterAction string `json:"afterAction"`
}

// NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps instantiates a new OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps(devices OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices, afterAction string) *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps {
	this := OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps{}
	this.Devices = devices
	this.AfterAction = afterAction
	return &this
}

// NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwapsWithDefaults instantiates a new OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwapsWithDefaults() *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps {
	this := OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps{}
	return &this
}

// GetDevices returns the Devices field value
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetDevices() OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices {
	if o == nil {
		var ret OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetDevicesOk() (*OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) SetDevices(v OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices) {
	o.Devices = v
}

// GetAfterAction returns the AfterAction field value
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetAfterAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfterAction
}

// GetAfterActionOk returns a tuple with the AfterAction field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetAfterActionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfterAction, true
}

// SetAfterAction sets field value
func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) SetAfterAction(v string) {
	o.AfterAction = v
}

func (o OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	if true {
		toSerialize["afterAction"] = o.AfterAction
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps struct {
	value *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) Get() *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) Set(val *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps(val *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) *NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps {
	return &NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


