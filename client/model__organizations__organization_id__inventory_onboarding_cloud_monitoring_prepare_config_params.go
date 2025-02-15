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

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams Params used in order to connect to the device
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams struct {
	Tunnel *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel `json:"tunnel,omitempty"`
	// Static IP Address used to connect to the device
	CloudStaticIp *string `json:"cloudStaticIp,omitempty"`
	User *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser `json:"user,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams{}
	return &this
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetTunnel() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel {
	if o == nil || isNil(o.Tunnel) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetTunnelOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel, bool) {
	if o == nil || isNil(o.Tunnel) {
    return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) HasTunnel() bool {
	if o != nil && !isNil(o.Tunnel) {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel and assigns it to the Tunnel field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) SetTunnel(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel) {
	o.Tunnel = &v
}

// GetCloudStaticIp returns the CloudStaticIp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetCloudStaticIp() string {
	if o == nil || isNil(o.CloudStaticIp) {
		var ret string
		return ret
	}
	return *o.CloudStaticIp
}

// GetCloudStaticIpOk returns a tuple with the CloudStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetCloudStaticIpOk() (*string, bool) {
	if o == nil || isNil(o.CloudStaticIp) {
    return nil, false
	}
	return o.CloudStaticIp, true
}

// HasCloudStaticIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) HasCloudStaticIp() bool {
	if o != nil && !isNil(o.CloudStaticIp) {
		return true
	}

	return false
}

// SetCloudStaticIp gets a reference to the given string and assigns it to the CloudStaticIp field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) SetCloudStaticIp(v string) {
	o.CloudStaticIp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetUser() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
	if o == nil || isNil(o.User) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) GetUserOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser and assigns it to the User field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) SetUser(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) {
	o.User = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tunnel) {
		toSerialize["tunnel"] = o.Tunnel
	}
	if !isNil(o.CloudStaticIp) {
		toSerialize["cloudStaticIp"] = o.CloudStaticIp
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


