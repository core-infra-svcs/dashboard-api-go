/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsCloud Data for this organization
type OrganizationsCloud struct {
	Region *OrganizationsCloudRegion `json:"region,omitempty"`
}

// NewOrganizationsCloud instantiates a new OrganizationsCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsCloud() *OrganizationsCloud {
	this := OrganizationsCloud{}
	return &this
}

// NewOrganizationsCloudWithDefaults instantiates a new OrganizationsCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsCloudWithDefaults() *OrganizationsCloud {
	this := OrganizationsCloud{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OrganizationsCloud) GetRegion() OrganizationsCloudRegion {
	if o == nil || isNil(o.Region) {
		var ret OrganizationsCloudRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCloud) GetRegionOk() (*OrganizationsCloudRegion, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OrganizationsCloud) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given OrganizationsCloudRegion and assigns it to the Region field.
func (o *OrganizationsCloud) SetRegion(v OrganizationsCloudRegion) {
	o.Region = &v
}

func (o OrganizationsCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsCloud struct {
	value *OrganizationsCloud
	isSet bool
}

func (v NullableOrganizationsCloud) Get() *OrganizationsCloud {
	return v.value
}

func (v *NullableOrganizationsCloud) Set(val *OrganizationsCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsCloud(val *OrganizationsCloud) *NullableOrganizationsCloud {
	return &NullableOrganizationsCloud{value: val, isSet: true}
}

func (v NullableOrganizationsCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


