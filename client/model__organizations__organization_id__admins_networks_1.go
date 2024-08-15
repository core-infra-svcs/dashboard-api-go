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

// OrganizationsOrganizationIdAdminsNetworks1 struct for OrganizationsOrganizationIdAdminsNetworks1
type OrganizationsOrganizationIdAdminsNetworks1 struct {
	// The network ID
	Id string `json:"id"`
	// The privilege of the dashboard administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewOrganizationsOrganizationIdAdminsNetworks1 instantiates a new OrganizationsOrganizationIdAdminsNetworks1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdminsNetworks1(id string, access string) *OrganizationsOrganizationIdAdminsNetworks1 {
	this := OrganizationsOrganizationIdAdminsNetworks1{}
	this.Id = id
	this.Access = access
	return &this
}

// NewOrganizationsOrganizationIdAdminsNetworks1WithDefaults instantiates a new OrganizationsOrganizationIdAdminsNetworks1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdminsNetworks1WithDefaults() *OrganizationsOrganizationIdAdminsNetworks1 {
	this := OrganizationsOrganizationIdAdminsNetworks1{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationsOrganizationIdAdminsNetworks1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks1) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationsOrganizationIdAdminsNetworks1) SetId(v string) {
	o.Id = v
}

// GetAccess returns the Access field value
func (o *OrganizationsOrganizationIdAdminsNetworks1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsNetworks1) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OrganizationsOrganizationIdAdminsNetworks1) SetAccess(v string) {
	o.Access = v
}

func (o OrganizationsOrganizationIdAdminsNetworks1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdminsNetworks1 struct {
	value *OrganizationsOrganizationIdAdminsNetworks1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks1) Get() *OrganizationsOrganizationIdAdminsNetworks1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks1) Set(val *OrganizationsOrganizationIdAdminsNetworks1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdminsNetworks1(val *OrganizationsOrganizationIdAdminsNetworks1) *NullableOrganizationsOrganizationIdAdminsNetworks1 {
	return &NullableOrganizationsOrganizationIdAdminsNetworks1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdminsNetworks1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdminsNetworks1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


