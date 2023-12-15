/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesStatusesComponents Components
type OrganizationsOrganizationIdDevicesStatusesComponents struct {
	// Power Supplies
	PowerSupplies []OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies `json:"powerSupplies,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesStatusesComponents instantiates a new OrganizationsOrganizationIdDevicesStatusesComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesStatusesComponents() *OrganizationsOrganizationIdDevicesStatusesComponents {
	this := OrganizationsOrganizationIdDevicesStatusesComponents{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesStatusesComponentsWithDefaults instantiates a new OrganizationsOrganizationIdDevicesStatusesComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesStatusesComponentsWithDefaults() *OrganizationsOrganizationIdDevicesStatusesComponents {
	this := OrganizationsOrganizationIdDevicesStatusesComponents{}
	return &this
}

// GetPowerSupplies returns the PowerSupplies field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponents) GetPowerSupplies() []OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies {
	if o == nil || isNil(o.PowerSupplies) {
		var ret []OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies
		return ret
	}
	return o.PowerSupplies
}

// GetPowerSuppliesOk returns a tuple with the PowerSupplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponents) GetPowerSuppliesOk() ([]OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies, bool) {
	if o == nil || isNil(o.PowerSupplies) {
    return nil, false
	}
	return o.PowerSupplies, true
}

// HasPowerSupplies returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponents) HasPowerSupplies() bool {
	if o != nil && !isNil(o.PowerSupplies) {
		return true
	}

	return false
}

// SetPowerSupplies gets a reference to the given []OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies and assigns it to the PowerSupplies field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponents) SetPowerSupplies(v []OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) {
	o.PowerSupplies = v
}

func (o OrganizationsOrganizationIdDevicesStatusesComponents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PowerSupplies) {
		toSerialize["powerSupplies"] = o.PowerSupplies
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesStatusesComponents struct {
	value *OrganizationsOrganizationIdDevicesStatusesComponents
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponents) Get() *OrganizationsOrganizationIdDevicesStatusesComponents {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponents) Set(val *OrganizationsOrganizationIdDevicesStatusesComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesStatusesComponents(val *OrganizationsOrganizationIdDevicesStatusesComponents) *NullableOrganizationsOrganizationIdDevicesStatusesComponents {
	return &NullableOrganizationsOrganizationIdDevicesStatusesComponents{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


