/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesSchedule The sensor schedule to use with the alert profile.
type NetworksNetworkIdSensorAlertsProfilesSchedule struct {
	// ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Id *string `json:"id,omitempty"`
	// Name of the sensor schedule to use with the alert profile.
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesSchedule instantiates a new NetworksNetworkIdSensorAlertsProfilesSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesSchedule() *NetworksNetworkIdSensorAlertsProfilesSchedule {
	this := NetworksNetworkIdSensorAlertsProfilesSchedule{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesScheduleWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesScheduleWithDefaults() *NetworksNetworkIdSensorAlertsProfilesSchedule {
	this := NetworksNetworkIdSensorAlertsProfilesSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSensorAlertsProfilesSchedule) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesSchedule struct {
	value *NetworksNetworkIdSensorAlertsProfilesSchedule
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesSchedule) Get() *NetworksNetworkIdSensorAlertsProfilesSchedule {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesSchedule) Set(val *NetworksNetworkIdSensorAlertsProfilesSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesSchedule(val *NetworksNetworkIdSensorAlertsProfilesSchedule) *NullableNetworksNetworkIdSensorAlertsProfilesSchedule {
	return &NullableNetworksNetworkIdSensorAlertsProfilesSchedule{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


