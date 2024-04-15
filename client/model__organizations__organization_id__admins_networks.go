/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdminsNetworks struct for OrganizationsOrganizationIdAdminsNetworks
type OrganizationsOrganizationIdAdminsNetworks struct {
	// Network ID
	Id *string `json:"id,omitempty"`
	// Admin's level of access to the network
	Access *string `json:"access,omitempty"`
}

// NewOrganizationsOrganizationIdAdminsNetworks instantiates a new OrganizationsOrganizationIdAdminsNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdminsNetworks() *OrganizationsOrganizationIdAdminsNetworks {
	this := OrganizationsOrganizationIdAdminsNetworks{}
	return &this
}

// NewOrganizationsOrganizationIdAdminsNetworksWithDefaults instantiates a new OrganizationsOrganizationIdAdminsNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdminsNetworksWithDefaults() *OrganizationsOrganizationIdAdminsNetworks {
	this := OrganizationsOrganizationIdAdminsNetworks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdminsNetworks) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdAdminsNetworks) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdminsNetworks) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *OrganizationsOrganizationIdAdminsNetworks) SetAccess(v string) {
	o.Access = &v
}

func (o OrganizationsOrganizationIdAdminsNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdminsNetworks struct {
	value *OrganizationsOrganizationIdAdminsNetworks
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks) Get() *OrganizationsOrganizationIdAdminsNetworks {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks) Set(val *OrganizationsOrganizationIdAdminsNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdminsNetworks(val *OrganizationsOrganizationIdAdminsNetworks) *NullableOrganizationsOrganizationIdAdminsNetworks {
	return &NullableOrganizationsOrganizationIdAdminsNetworks{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


