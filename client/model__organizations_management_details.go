/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsManagementDetails struct for OrganizationsManagementDetails
type OrganizationsManagementDetails struct {
	// Name of management data
	Name *string `json:"name,omitempty"`
	// Value of management data
	Value *string `json:"value,omitempty"`
}

// NewOrganizationsManagementDetails instantiates a new OrganizationsManagementDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsManagementDetails() *OrganizationsManagementDetails {
	this := OrganizationsManagementDetails{}
	return &this
}

// NewOrganizationsManagementDetailsWithDefaults instantiates a new OrganizationsManagementDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsManagementDetailsWithDefaults() *OrganizationsManagementDetails {
	this := OrganizationsManagementDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsManagementDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsManagementDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsManagementDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsManagementDetails) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrganizationsManagementDetails) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsManagementDetails) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrganizationsManagementDetails) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OrganizationsManagementDetails) SetValue(v string) {
	o.Value = &v
}

func (o OrganizationsManagementDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsManagementDetails struct {
	value *OrganizationsManagementDetails
	isSet bool
}

func (v NullableOrganizationsManagementDetails) Get() *OrganizationsManagementDetails {
	return v.value
}

func (v *NullableOrganizationsManagementDetails) Set(val *OrganizationsManagementDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsManagementDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsManagementDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsManagementDetails(val *OrganizationsManagementDetails) *NullableOrganizationsManagementDetails {
	return &NullableOrganizationsManagementDetails{value: val, isSet: true}
}

func (v NullableOrganizationsManagementDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsManagementDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


