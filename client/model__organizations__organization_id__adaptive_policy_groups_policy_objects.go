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

// OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects struct for OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
type OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects struct {
	// The ID of the policy object
	Id *string `json:"id,omitempty"`
	// The name of the policy object
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects instantiates a new OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects() *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	this := OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects{}
	return &this
}

// NewOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjectsWithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjectsWithDefaults() *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	this := OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects struct {
	value *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) Get() *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) Set(val *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects(val *OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) *NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	return &NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


