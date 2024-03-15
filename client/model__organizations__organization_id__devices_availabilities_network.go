/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesAvailabilitiesNetwork Network info.
type OrganizationsOrganizationIdDevicesAvailabilitiesNetwork struct {
	// ID for the network containing the device.
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesNetwork instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesAvailabilitiesNetwork() *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesNetworkWithDefaults instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesAvailabilitiesNetworkWithDefaults() *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork struct {
	value *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) Get() *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) Set(val *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork(val *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) *NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork {
	return &NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


