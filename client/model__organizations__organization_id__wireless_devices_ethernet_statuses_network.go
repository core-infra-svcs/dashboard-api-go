/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork Network details object
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork struct {
	// The network ID the AP is associated to
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetworkWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetworkWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


