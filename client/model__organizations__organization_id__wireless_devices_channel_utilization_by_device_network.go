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

// OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork Network for the given utilization metrics.
type OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	// Network ID of the given utilization metrics.
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetworkWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetworkWithDefaults() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	value *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) Get() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) Set(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	return &NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


