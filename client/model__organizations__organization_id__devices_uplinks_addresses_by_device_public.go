/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic Public interface information.
type OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic struct {
	// The device uplink public IP address.
	Address *string `json:"address,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic() *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublicWithDefaults instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublicWithDefaults() *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) SetAddress(v string) {
	o.Address = &v
}

func (o OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic struct {
	value *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) Get() *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) Set(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic {
	return &NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


