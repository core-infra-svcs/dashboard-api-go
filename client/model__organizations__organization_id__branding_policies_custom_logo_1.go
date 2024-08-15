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

// OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 Properties describing the custom logo attached to the branding policy.
type OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 struct {
	// Whether or not there is a custom logo enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Image *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image `json:"image,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1 instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo1{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1WithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1WithDefaults() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo1{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) GetImage() OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image {
	if o == nil || isNil(o.Image) {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) GetImageOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image and assigns it to the Image field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) SetImage(v OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) {
	o.Image = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1 struct {
	value *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) Get() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) Set(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


