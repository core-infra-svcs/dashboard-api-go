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

// OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch Reading for the 'remoteLockoutSwitch' metric. This will only be present if the 'metric' property equals 'remoteLockoutSwitch'.
type OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch struct {
	// True if power controls are disabled via the MT40's physical remote lockout switch.
	Locked *bool `json:"locked,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch() *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch {
	this := OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitchWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitchWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch {
	this := OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch{}
	return &this
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
    return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) SetLocked(v bool) {
	o.Locked = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) Get() *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch(val *OrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) *NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryRemoteLockoutSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


