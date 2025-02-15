/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSensorReadingsHistoryButton Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
type OrganizationsOrganizationIdSensorReadingsHistoryButton struct {
	// Type of button press that occurred.
	PressType *string `json:"pressType,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryButton instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryButton() *OrganizationsOrganizationIdSensorReadingsHistoryButton {
	this := OrganizationsOrganizationIdSensorReadingsHistoryButton{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryButtonWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryButtonWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryButton {
	this := OrganizationsOrganizationIdSensorReadingsHistoryButton{}
	return &this
}

// GetPressType returns the PressType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryButton) GetPressType() string {
	if o == nil || isNil(o.PressType) {
		var ret string
		return ret
	}
	return *o.PressType
}

// GetPressTypeOk returns a tuple with the PressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryButton) GetPressTypeOk() (*string, bool) {
	if o == nil || isNil(o.PressType) {
    return nil, false
	}
	return o.PressType, true
}

// HasPressType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryButton) HasPressType() bool {
	if o != nil && !isNil(o.PressType) {
		return true
	}

	return false
}

// SetPressType gets a reference to the given string and assigns it to the PressType field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryButton) SetPressType(v string) {
	o.PressType = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PressType) {
		toSerialize["pressType"] = o.PressType
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryButton struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryButton
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) Get() *OrganizationsOrganizationIdSensorReadingsHistoryButton {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryButton) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryButton(val *OrganizationsOrganizationIdSensorReadingsHistoryButton) *NullableOrganizationsOrganizationIdSensorReadingsHistoryButton {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryButton{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


