/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject240 struct for InlineObject240
type InlineObject240 struct {
	// The name of the new role. Must be unique. This parameter is required.
	Name string `json:"name"`
	// Device tag on which this specified permission is applied.
	AppliedOnDevices []OrganizationsOrganizationIdCameraRolesAppliedOnDevices `json:"appliedOnDevices,omitempty"`
	// Network tag on which this specified permission is applied.
	AppliedOnNetworks []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks `json:"appliedOnNetworks,omitempty"`
	// Permissions to be applied org wide.
	AppliedOrgWide []OrganizationsOrganizationIdCameraRolesAppliedOrgWide `json:"appliedOrgWide,omitempty"`
}

// NewInlineObject240 instantiates a new InlineObject240 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject240(name string) *InlineObject240 {
	this := InlineObject240{}
	this.Name = name
	return &this
}

// NewInlineObject240WithDefaults instantiates a new InlineObject240 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject240WithDefaults() *InlineObject240 {
	this := InlineObject240{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject240) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject240) SetName(v string) {
	o.Name = v
}

// GetAppliedOnDevices returns the AppliedOnDevices field value if set, zero value otherwise.
func (o *InlineObject240) GetAppliedOnDevices() []OrganizationsOrganizationIdCameraRolesAppliedOnDevices {
	if o == nil || isNil(o.AppliedOnDevices) {
		var ret []OrganizationsOrganizationIdCameraRolesAppliedOnDevices
		return ret
	}
	return o.AppliedOnDevices
}

// GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetAppliedOnDevicesOk() ([]OrganizationsOrganizationIdCameraRolesAppliedOnDevices, bool) {
	if o == nil || isNil(o.AppliedOnDevices) {
    return nil, false
	}
	return o.AppliedOnDevices, true
}

// HasAppliedOnDevices returns a boolean if a field has been set.
func (o *InlineObject240) HasAppliedOnDevices() bool {
	if o != nil && !isNil(o.AppliedOnDevices) {
		return true
	}

	return false
}

// SetAppliedOnDevices gets a reference to the given []OrganizationsOrganizationIdCameraRolesAppliedOnDevices and assigns it to the AppliedOnDevices field.
func (o *InlineObject240) SetAppliedOnDevices(v []OrganizationsOrganizationIdCameraRolesAppliedOnDevices) {
	o.AppliedOnDevices = v
}

// GetAppliedOnNetworks returns the AppliedOnNetworks field value if set, zero value otherwise.
func (o *InlineObject240) GetAppliedOnNetworks() []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks {
	if o == nil || isNil(o.AppliedOnNetworks) {
		var ret []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks
		return ret
	}
	return o.AppliedOnNetworks
}

// GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetAppliedOnNetworksOk() ([]OrganizationsOrganizationIdCameraRolesAppliedOnNetworks, bool) {
	if o == nil || isNil(o.AppliedOnNetworks) {
    return nil, false
	}
	return o.AppliedOnNetworks, true
}

// HasAppliedOnNetworks returns a boolean if a field has been set.
func (o *InlineObject240) HasAppliedOnNetworks() bool {
	if o != nil && !isNil(o.AppliedOnNetworks) {
		return true
	}

	return false
}

// SetAppliedOnNetworks gets a reference to the given []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks and assigns it to the AppliedOnNetworks field.
func (o *InlineObject240) SetAppliedOnNetworks(v []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks) {
	o.AppliedOnNetworks = v
}

// GetAppliedOrgWide returns the AppliedOrgWide field value if set, zero value otherwise.
func (o *InlineObject240) GetAppliedOrgWide() []OrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	if o == nil || isNil(o.AppliedOrgWide) {
		var ret []OrganizationsOrganizationIdCameraRolesAppliedOrgWide
		return ret
	}
	return o.AppliedOrgWide
}

// GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetAppliedOrgWideOk() ([]OrganizationsOrganizationIdCameraRolesAppliedOrgWide, bool) {
	if o == nil || isNil(o.AppliedOrgWide) {
    return nil, false
	}
	return o.AppliedOrgWide, true
}

// HasAppliedOrgWide returns a boolean if a field has been set.
func (o *InlineObject240) HasAppliedOrgWide() bool {
	if o != nil && !isNil(o.AppliedOrgWide) {
		return true
	}

	return false
}

// SetAppliedOrgWide gets a reference to the given []OrganizationsOrganizationIdCameraRolesAppliedOrgWide and assigns it to the AppliedOrgWide field.
func (o *InlineObject240) SetAppliedOrgWide(v []OrganizationsOrganizationIdCameraRolesAppliedOrgWide) {
	o.AppliedOrgWide = v
}

func (o InlineObject240) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableInlineObject240 struct {
	value *InlineObject240
	isSet bool
}

func (v NullableInlineObject240) Get() *InlineObject240 {
	return v.value
}

func (v *NullableInlineObject240) Set(val *InlineObject240) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject240) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject240) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject240(val *InlineObject240) *NullableInlineObject240 {
	return &NullableInlineObject240{value: val, isSet: true}
}

func (v NullableInlineObject240) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject240) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


