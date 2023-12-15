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

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe PoE power details object
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe struct {
	// PoE power connected
	IsConnected *bool `json:"isConnected,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoeWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoeWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe{}
	return &this
}

// GetIsConnected returns the IsConnected field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) GetIsConnected() bool {
	if o == nil || isNil(o.IsConnected) {
		var ret bool
		return ret
	}
	return *o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) GetIsConnectedOk() (*bool, bool) {
	if o == nil || isNil(o.IsConnected) {
    return nil, false
	}
	return o.IsConnected, true
}

// HasIsConnected returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) HasIsConnected() bool {
	if o != nil && !isNil(o.IsConnected) {
		return true
	}

	return false
}

// SetIsConnected gets a reference to the given bool and assigns it to the IsConnected field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) SetIsConnected(v bool) {
	o.IsConnected = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsConnected) {
		toSerialize["isConnected"] = o.IsConnected
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


