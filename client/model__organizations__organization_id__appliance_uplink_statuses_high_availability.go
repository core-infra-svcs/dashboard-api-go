/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability Device High Availability Capabilities
type OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability struct {
	// Indicates whether High Availability is enabled for the device. For devices that do not support HA, this will be 'false'
	Enabled *bool `json:"enabled,omitempty"`
	// The HA role of the device on the network. For devices that do not support HA, this will be 'primary'
	Role *string `json:"role,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability instantiates a new OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability() *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	this := OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailabilityWithDefaults instantiates a new OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailabilityWithDefaults() *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	this := OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) SetRole(v string) {
	o.Role = &v
}

func (o OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability struct {
	value *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) Get() *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) Set(val *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability(val *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) *NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	return &NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


