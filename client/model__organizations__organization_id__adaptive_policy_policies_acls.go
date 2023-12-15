/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls struct for OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls
type OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls struct {
	// The ID of the adaptive policy ACL
	Id *string `json:"id,omitempty"`
	// The name of the adaptive policy ACL
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls instantiates a new OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls() *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	this := OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls{}
	return &this
}

// NewOrganizationsOrganizationIdAdaptivePolicyPoliciesAclsWithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdaptivePolicyPoliciesAclsWithDefaults() *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	this := OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls struct {
	value *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) Get() *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) Set(val *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls(val *OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	return &NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


