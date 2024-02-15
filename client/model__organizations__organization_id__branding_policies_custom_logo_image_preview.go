/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview Preview of the image
type OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview struct {
	// Url of the preview image
	Url *string `json:"url,omitempty"`
	// Timestamp of the preview image
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreviewWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreviewWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) SetUrl(v string) {
	o.Url = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview struct {
	value *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) Get() *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) Set(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogoImagePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


