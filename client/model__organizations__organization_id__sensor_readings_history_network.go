/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryNetwork Network to which the sensor belongs.
type OrganizationsOrganizationIdSensorReadingsHistoryNetwork struct {
	// ID of the network.
	Id *string `json:"id,omitempty"`
	// Name of the network.
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNetwork instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryNetwork() *OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryNetworkWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryNetworkWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	this := OrganizationsOrganizationIdSensorReadingsHistoryNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) Get() *OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork(val *OrganizationsOrganizationIdSensorReadingsHistoryNetwork) *NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


