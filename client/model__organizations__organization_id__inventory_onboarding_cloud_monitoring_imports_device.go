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

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice Represents the details of an imported device.
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice struct {
	// The url to the device details page within dashboard.
	Url *string `json:"url,omitempty"`
	// Whether or not the device was successfully created in dashboard.
	Created *bool `json:"created,omitempty"`
	// Represents the current state of importing the device.
	Status *string `json:"status,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDeviceWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDeviceWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) SetUrl(v string) {
	o.Url = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetCreated() bool {
	if o == nil || isNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetCreatedOk() (*bool, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) SetCreated(v bool) {
	o.Created = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) SetStatus(v string) {
	o.Status = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


