/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices struct for OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices struct {
	// Device SUDI certificate
	Sudi string `json:"sudi"`
	Tunnel *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel `json:"tunnel,omitempty"`
	User *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser `json:"user,omitempty"`
	Vty *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty `json:"vty,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices(sudi string) *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices{}
	this.Sudi = sudi
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevicesWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevicesWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices{}
	return &this
}

// GetSudi returns the Sudi field value
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetSudi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sudi
}

// GetSudiOk returns a tuple with the Sudi field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetSudiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sudi, true
}

// SetSudi sets field value
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) SetSudi(v string) {
	o.Sudi = v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetTunnel() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel {
	if o == nil || isNil(o.Tunnel) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetTunnelOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel, bool) {
	if o == nil || isNil(o.Tunnel) {
    return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) HasTunnel() bool {
	if o != nil && !isNil(o.Tunnel) {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel and assigns it to the Tunnel field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) SetTunnel(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) {
	o.Tunnel = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetUser() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser {
	if o == nil || isNil(o.User) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetUserOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser and assigns it to the User field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) SetUser(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareUser) {
	o.User = &v
}

// GetVty returns the Vty field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetVty() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty {
	if o == nil || isNil(o.Vty) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty
		return ret
	}
	return *o.Vty
}

// GetVtyOk returns a tuple with the Vty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) GetVtyOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty, bool) {
	if o == nil || isNil(o.Vty) {
    return nil, false
	}
	return o.Vty, true
}

// HasVty returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) HasVty() bool {
	if o != nil && !isNil(o.Vty) {
		return true
	}

	return false
}

// SetVty gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty and assigns it to the Vty field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) SetVty(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) {
	o.Vty = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sudi"] = o.Sudi
	}
	if !isNil(o.Tunnel) {
		toSerialize["tunnel"] = o.Tunnel
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Vty) {
		toSerialize["vty"] = o.Vty
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


