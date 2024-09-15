/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork Network info.
type OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork struct {
	// ID for the network that the device is associated with.
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork instantiates a new OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork {
	this := OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetworkWithDefaults instantiates a new OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetworkWithDefaults() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork {
	this := OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork struct {
	value *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) Get() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) Set(val *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork(val *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork {
	return &NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


