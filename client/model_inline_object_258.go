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

// InlineObject258 struct for InlineObject258
type InlineObject258 struct {
	// A set of devices to import (or update)
	Devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices `json:"devices"`
	Options *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions `json:"options,omitempty"`
}

// NewInlineObject258 instantiates a new InlineObject258 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject258(devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) *InlineObject258 {
	this := InlineObject258{}
	this.Devices = devices
	return &this
}

// NewInlineObject258WithDefaults instantiates a new InlineObject258 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject258WithDefaults() *InlineObject258 {
	this := InlineObject258{}
	return &this
}

// GetDevices returns the Devices field value
func (o *InlineObject258) GetDevices() []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	if o == nil {
		var ret []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineObject258) GetDevicesOk() ([]OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices, bool) {
	if o == nil {
    return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *InlineObject258) SetDevices(v []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) {
	o.Devices = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InlineObject258) GetOptions() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions {
	if o == nil || isNil(o.Options) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject258) GetOptionsOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InlineObject258) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions and assigns it to the Options field.
func (o *InlineObject258) SetOptions(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareOptions) {
	o.Options = &v
}

func (o InlineObject258) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject258 struct {
	value *InlineObject258
	isSet bool
}

func (v NullableInlineObject258) Get() *InlineObject258 {
	return v.value
}

func (v *NullableInlineObject258) Set(val *InlineObject258) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject258) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject258) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject258(val *InlineObject258) *NullableInlineObject258 {
	return &NullableInlineObject258{value: val, isSet: true}
}

func (v NullableInlineObject258) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject258) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


