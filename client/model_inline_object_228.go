/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject228 struct for InlineObject228
type InlineObject228 struct {
	// The name of the new role. Must be unique.
	Name *string `json:"name,omitempty"`
	// Device tag on which this specified permission is applied.
	AppliedOnDevices []OrganizationsOrganizationIdCameraRolesAppliedOnDevices `json:"appliedOnDevices,omitempty"`
	// Network tag on which this specified permission is applied.
	AppliedOnNetworks []OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks `json:"appliedOnNetworks,omitempty"`
	// Permissions to be applied org wide.
	AppliedOrgWide []OrganizationsOrganizationIdCameraRolesAppliedOrgWide `json:"appliedOrgWide,omitempty"`
}

// NewInlineObject228 instantiates a new InlineObject228 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject228() *InlineObject228 {
	this := InlineObject228{}
	return &this
}

// NewInlineObject228WithDefaults instantiates a new InlineObject228 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject228WithDefaults() *InlineObject228 {
	this := InlineObject228{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject228) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject228) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject228) SetName(v string) {
	o.Name = &v
}

// GetAppliedOnDevices returns the AppliedOnDevices field value if set, zero value otherwise.
func (o *InlineObject228) GetAppliedOnDevices() []OrganizationsOrganizationIdCameraRolesAppliedOnDevices {
	if o == nil || isNil(o.AppliedOnDevices) {
		var ret []OrganizationsOrganizationIdCameraRolesAppliedOnDevices
		return ret
	}
	return o.AppliedOnDevices
}

// GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetAppliedOnDevicesOk() ([]OrganizationsOrganizationIdCameraRolesAppliedOnDevices, bool) {
	if o == nil || isNil(o.AppliedOnDevices) {
    return nil, false
	}
	return o.AppliedOnDevices, true
}

// HasAppliedOnDevices returns a boolean if a field has been set.
func (o *InlineObject228) HasAppliedOnDevices() bool {
	if o != nil && !isNil(o.AppliedOnDevices) {
		return true
	}

	return false
}

// SetAppliedOnDevices gets a reference to the given []OrganizationsOrganizationIdCameraRolesAppliedOnDevices and assigns it to the AppliedOnDevices field.
func (o *InlineObject228) SetAppliedOnDevices(v []OrganizationsOrganizationIdCameraRolesAppliedOnDevices) {
	o.AppliedOnDevices = v
}

// GetAppliedOnNetworks returns the AppliedOnNetworks field value if set, zero value otherwise.
func (o *InlineObject228) GetAppliedOnNetworks() []OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks {
	if o == nil || isNil(o.AppliedOnNetworks) {
		var ret []OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks
		return ret
	}
	return o.AppliedOnNetworks
}

// GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetAppliedOnNetworksOk() ([]OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks, bool) {
	if o == nil || isNil(o.AppliedOnNetworks) {
    return nil, false
	}
	return o.AppliedOnNetworks, true
}

// HasAppliedOnNetworks returns a boolean if a field has been set.
func (o *InlineObject228) HasAppliedOnNetworks() bool {
	if o != nil && !isNil(o.AppliedOnNetworks) {
		return true
	}

	return false
}

// SetAppliedOnNetworks gets a reference to the given []OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks and assigns it to the AppliedOnNetworks field.
func (o *InlineObject228) SetAppliedOnNetworks(v []OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) {
	o.AppliedOnNetworks = v
}

// GetAppliedOrgWide returns the AppliedOrgWide field value if set, zero value otherwise.
func (o *InlineObject228) GetAppliedOrgWide() []OrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	if o == nil || isNil(o.AppliedOrgWide) {
		var ret []OrganizationsOrganizationIdCameraRolesAppliedOrgWide
		return ret
	}
	return o.AppliedOrgWide
}

// GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject228) GetAppliedOrgWideOk() ([]OrganizationsOrganizationIdCameraRolesAppliedOrgWide, bool) {
	if o == nil || isNil(o.AppliedOrgWide) {
    return nil, false
	}
	return o.AppliedOrgWide, true
}

// HasAppliedOrgWide returns a boolean if a field has been set.
func (o *InlineObject228) HasAppliedOrgWide() bool {
	if o != nil && !isNil(o.AppliedOrgWide) {
		return true
	}

	return false
}

// SetAppliedOrgWide gets a reference to the given []OrganizationsOrganizationIdCameraRolesAppliedOrgWide and assigns it to the AppliedOrgWide field.
func (o *InlineObject228) SetAppliedOrgWide(v []OrganizationsOrganizationIdCameraRolesAppliedOrgWide) {
	o.AppliedOrgWide = v
}

func (o InlineObject228) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AppliedOnDevices) {
		toSerialize["appliedOnDevices"] = o.AppliedOnDevices
	}
	if !isNil(o.AppliedOnNetworks) {
		toSerialize["appliedOnNetworks"] = o.AppliedOnNetworks
	}
	if !isNil(o.AppliedOrgWide) {
		toSerialize["appliedOrgWide"] = o.AppliedOrgWide
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject228 struct {
	value *InlineObject228
	isSet bool
}

func (v NullableInlineObject228) Get() *InlineObject228 {
	return v.value
}

func (v *NullableInlineObject228) Set(val *InlineObject228) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject228) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject228) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject228(val *InlineObject228) *NullableInlineObject228 {
	return &NullableInlineObject228{value: val, isSet: true}
}

func (v NullableInlineObject228) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject228) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


