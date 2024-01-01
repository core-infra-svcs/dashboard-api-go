/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image Properties for setting the image.
type OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image struct {
	// The file contents (a base 64 encoded string) of your new logo.
	Contents *string `json:"contents,omitempty"`
	// The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1ImageWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo1ImageWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) GetContents() string {
	if o == nil || isNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) GetContentsOk() (*string, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) SetContents(v string) {
	o.Contents = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) SetFormat(v string) {
	o.Format = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image struct {
	value *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) Get() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) Set(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo1Image) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


