/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage Properties of the image.
type OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage struct {
	Preview *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview `json:"preview,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImageWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImageWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage{}
	return &this
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) GetPreview() OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview {
	if o == nil || isNil(o.Preview) {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) GetPreviewOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview, bool) {
	if o == nil || isNil(o.Preview) {
    return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) HasPreview() bool {
	if o != nil && !isNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview and assigns it to the Preview field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) SetPreview(v OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) {
	o.Preview = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage struct {
	value *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) Get() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) Set(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


