/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks struct for OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks
type OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks struct {
	// Interface for the device uplink. Available options are: cellular, man1, man2, wan1, wan2
	Interface *string `json:"interface,omitempty"`
	// Available addresses for the interface.
	Addresses []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses `json:"addresses,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinksWithDefaults instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinksWithDefaults() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) SetInterface(v string) {
	o.Interface = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) GetAddresses() []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses {
	if o == nil || isNil(o.Addresses) {
		var ret []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) GetAddressesOk() ([]OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses and assigns it to the Addresses field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) SetAddresses(v []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) {
	o.Addresses = v
}

func (o OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks struct {
	value *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) Get() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) Set(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks {
	return &NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


