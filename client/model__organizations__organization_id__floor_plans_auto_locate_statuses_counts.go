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

// OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts Counts for this floor plan
type OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts struct {
	Devices *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices `json:"devices,omitempty"`
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts{}
	return &this
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsWithDefaults instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsWithDefaults() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) GetDevices() OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices {
	if o == nil || isNil(o.Devices) {
		var ret OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) GetDevicesOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices and assigns it to the Devices field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) SetDevices(v OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCountsDevices) {
	o.Devices = &v
}

func (o OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts struct {
	value *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) Get() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) Set(val *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts(val *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts {
	return &NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


