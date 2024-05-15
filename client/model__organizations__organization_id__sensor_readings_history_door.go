/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryDoor Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
type OrganizationsOrganizationIdSensorReadingsHistoryDoor struct {
	// True if the door is open.
	Open *bool `json:"open,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryDoor instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryDoor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryDoor() *OrganizationsOrganizationIdSensorReadingsHistoryDoor {
	this := OrganizationsOrganizationIdSensorReadingsHistoryDoor{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryDoorWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryDoor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryDoorWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryDoor {
	this := OrganizationsOrganizationIdSensorReadingsHistoryDoor{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDoor) GetOpen() bool {
	if o == nil || isNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDoor) GetOpenOk() (*bool, bool) {
	if o == nil || isNil(o.Open) {
    return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDoor) HasOpen() bool {
	if o != nil && !isNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryDoor) SetOpen(v bool) {
	o.Open = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryDoor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryDoor
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) Get() *OrganizationsOrganizationIdSensorReadingsHistoryDoor {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryDoor) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryDoor(val *OrganizationsOrganizationIdSensorReadingsHistoryDoor) *NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryDoor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


