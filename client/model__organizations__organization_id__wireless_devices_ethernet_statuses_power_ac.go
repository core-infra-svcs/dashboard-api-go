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

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc AC power details object
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc struct {
	// AC power connected
	IsConnected *bool `json:"isConnected,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAcWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAcWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc{}
	return &this
}

// GetIsConnected returns the IsConnected field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) GetIsConnected() bool {
	if o == nil || isNil(o.IsConnected) {
		var ret bool
		return ret
	}
	return *o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) GetIsConnectedOk() (*bool, bool) {
	if o == nil || isNil(o.IsConnected) {
    return nil, false
	}
	return o.IsConnected, true
}

// HasIsConnected returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) HasIsConnected() bool {
	if o != nil && !isNil(o.IsConnected) {
		return true
	}

	return false
}

// SetIsConnected gets a reference to the given bool and assigns it to the IsConnected field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) SetIsConnected(v bool) {
	o.IsConnected = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsConnected) {
		toSerialize["isConnected"] = o.IsConnected
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


