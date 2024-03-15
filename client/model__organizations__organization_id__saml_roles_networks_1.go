/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSamlRolesNetworks1 struct for OrganizationsOrganizationIdSamlRolesNetworks1
type OrganizationsOrganizationIdSamlRolesNetworks1 struct {
	// The network ID
	Id string `json:"id"`
	// The privilege of the SAML administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador', 'monitor-only' or 'ssid-admin'
	Access string `json:"access"`
}

// NewOrganizationsOrganizationIdSamlRolesNetworks1 instantiates a new OrganizationsOrganizationIdSamlRolesNetworks1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSamlRolesNetworks1(id string, access string) *OrganizationsOrganizationIdSamlRolesNetworks1 {
	this := OrganizationsOrganizationIdSamlRolesNetworks1{}
	this.Id = id
	this.Access = access
	return &this
}

// NewOrganizationsOrganizationIdSamlRolesNetworks1WithDefaults instantiates a new OrganizationsOrganizationIdSamlRolesNetworks1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSamlRolesNetworks1WithDefaults() *OrganizationsOrganizationIdSamlRolesNetworks1 {
	this := OrganizationsOrganizationIdSamlRolesNetworks1{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) SetId(v string) {
	o.Id = v
}

// GetAccess returns the Access field value
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OrganizationsOrganizationIdSamlRolesNetworks1) SetAccess(v string) {
	o.Access = v
}

func (o OrganizationsOrganizationIdSamlRolesNetworks1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSamlRolesNetworks1 struct {
	value *OrganizationsOrganizationIdSamlRolesNetworks1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSamlRolesNetworks1) Get() *OrganizationsOrganizationIdSamlRolesNetworks1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSamlRolesNetworks1) Set(val *OrganizationsOrganizationIdSamlRolesNetworks1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSamlRolesNetworks1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSamlRolesNetworks1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSamlRolesNetworks1(val *OrganizationsOrganizationIdSamlRolesNetworks1) *NullableOrganizationsOrganizationIdSamlRolesNetworks1 {
	return &NullableOrganizationsOrganizationIdSamlRolesNetworks1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSamlRolesNetworks1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSamlRolesNetworks1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


