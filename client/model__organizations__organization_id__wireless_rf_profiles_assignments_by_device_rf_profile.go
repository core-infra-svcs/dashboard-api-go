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

// OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile Information regarding the RF Profile of the device.
type OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile struct {
	// The ID of the RF Profile the device belongs to.
	Id *string `json:"id,omitempty"`
	// The name of the RF Profile the device belongs to.
	Name *string `json:"name,omitempty"`
	// Status to show if this profile is default indoor profile.
	IsIndoorDefault *bool `json:"isIndoorDefault,omitempty"`
	// Status to show if this profile is default outdoor profile.
	IsOutdoorDefault *bool `json:"isOutdoorDefault,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfileWithDefaults instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfileWithDefaults() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) SetName(v string) {
	o.Name = &v
}

// GetIsIndoorDefault returns the IsIndoorDefault field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetIsIndoorDefault() bool {
	if o == nil || isNil(o.IsIndoorDefault) {
		var ret bool
		return ret
	}
	return *o.IsIndoorDefault
}

// GetIsIndoorDefaultOk returns a tuple with the IsIndoorDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetIsIndoorDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsIndoorDefault) {
    return nil, false
	}
	return o.IsIndoorDefault, true
}

// HasIsIndoorDefault returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) HasIsIndoorDefault() bool {
	if o != nil && !isNil(o.IsIndoorDefault) {
		return true
	}

	return false
}

// SetIsIndoorDefault gets a reference to the given bool and assigns it to the IsIndoorDefault field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) SetIsIndoorDefault(v bool) {
	o.IsIndoorDefault = &v
}

// GetIsOutdoorDefault returns the IsOutdoorDefault field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetIsOutdoorDefault() bool {
	if o == nil || isNil(o.IsOutdoorDefault) {
		var ret bool
		return ret
	}
	return *o.IsOutdoorDefault
}

// GetIsOutdoorDefaultOk returns a tuple with the IsOutdoorDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) GetIsOutdoorDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsOutdoorDefault) {
    return nil, false
	}
	return o.IsOutdoorDefault, true
}

// HasIsOutdoorDefault returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) HasIsOutdoorDefault() bool {
	if o != nil && !isNil(o.IsOutdoorDefault) {
		return true
	}

	return false
}

// SetIsOutdoorDefault gets a reference to the given bool and assigns it to the IsOutdoorDefault field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) SetIsOutdoorDefault(v bool) {
	o.IsOutdoorDefault = &v
}

func (o OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsIndoorDefault) {
		toSerialize["isIndoorDefault"] = o.IsIndoorDefault
	}
	if !isNil(o.IsOutdoorDefault) {
		toSerialize["isOutdoorDefault"] = o.IsOutdoorDefault
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile struct {
	value *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) Get() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) Set(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile {
	return &NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


