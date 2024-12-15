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

// OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork Network.
type OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork struct {
	// Network ID.
	Id *string `json:"id,omitempty"`
	// Name of the network.
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetworkWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetworkWithDefaults() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork struct {
	value *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) Get() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) Set(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	return &NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


