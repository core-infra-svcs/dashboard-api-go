/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 The destination adaptive policy group (requires one unique attribute)
type OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 struct {
	// The ID of the destination adaptive policy group
	Id *string `json:"id,omitempty"`
	// The name of the destination adaptive policy group
	Name *string `json:"name,omitempty"`
	// The SGT of the destination adaptive policy group
	Sgt *int32 `json:"sgt,omitempty"`
}

// NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 instantiates a new OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1() *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	this := OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1{}
	return &this
}

// NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1WithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1WithDefaults() *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	this := OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetSgt() int32 {
	if o == nil || isNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) GetSgtOk() (*int32, bool) {
	if o == nil || isNil(o.Sgt) {
    return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) HasSgt() bool {
	if o != nil && !isNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) SetSgt(v int32) {
	o.Sgt = &v
}

func (o OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Sgt) {
		toSerialize["sgt"] = o.Sgt
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 struct {
	value *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) Get() *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) Set(val *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1(val *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	return &NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


