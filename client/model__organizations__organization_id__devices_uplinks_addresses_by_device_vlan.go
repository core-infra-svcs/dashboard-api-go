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

// OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan VLAN information of the uplink interface
type OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan struct {
	// VLAN ID of the uplink interface
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlanWithDefaults instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlanWithDefaults() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan struct {
	value *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) Get() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) Set(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan {
	return &NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


