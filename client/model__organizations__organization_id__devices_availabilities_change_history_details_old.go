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

// OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld struct for OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld
type OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld struct {
	// Name of the detail
	Name *string `json:"name,omitempty"`
	// Value of the detail
	Value *string `json:"value,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOldWithDefaults instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOldWithDefaults() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) SetValue(v string) {
	o.Value = &v
}

func (o OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld struct {
	value *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) Get() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) Set(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	return &NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


