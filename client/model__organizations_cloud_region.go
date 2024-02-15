/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsCloudRegion Region info
type OrganizationsCloudRegion struct {
	// Name of region
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsCloudRegion instantiates a new OrganizationsCloudRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsCloudRegion() *OrganizationsCloudRegion {
	this := OrganizationsCloudRegion{}
	return &this
}

// NewOrganizationsCloudRegionWithDefaults instantiates a new OrganizationsCloudRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsCloudRegionWithDefaults() *OrganizationsCloudRegion {
	this := OrganizationsCloudRegion{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsCloudRegion) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCloudRegion) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsCloudRegion) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsCloudRegion) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsCloudRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsCloudRegion struct {
	value *OrganizationsCloudRegion
	isSet bool
}

func (v NullableOrganizationsCloudRegion) Get() *OrganizationsCloudRegion {
	return v.value
}

func (v *NullableOrganizationsCloudRegion) Set(val *OrganizationsCloudRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsCloudRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsCloudRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsCloudRegion(val *OrganizationsCloudRegion) *NullableOrganizationsCloudRegion {
	return &NullableOrganizationsCloudRegion{value: val, isSet: true}
}

func (v NullableOrganizationsCloudRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsCloudRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


