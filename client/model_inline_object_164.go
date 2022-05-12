/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject164 struct for InlineObject164
type InlineObject164 struct {
	// Name of the Dashboard branding policy.
	Name string `json:"name"`
	// Boolean indicating whether this policy is enabled.
	Enabled bool `json:"enabled"`
	AdminSettings OrganizationsOrganizationIdBrandingPoliciesAdminSettings `json:"adminSettings"`
	HelpSettings *OrganizationsOrganizationIdBrandingPoliciesHelpSettings `json:"helpSettings,omitempty"`
	CustomLogo *OrganizationsOrganizationIdBrandingPoliciesCustomLogo `json:"customLogo,omitempty"`
}

// NewInlineObject164 instantiates a new InlineObject164 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject164(name string, enabled bool, adminSettings OrganizationsOrganizationIdBrandingPoliciesAdminSettings) *InlineObject164 {
	this := InlineObject164{}
	this.Name = name
	this.Enabled = enabled
	this.AdminSettings = adminSettings
	return &this
}

// NewInlineObject164WithDefaults instantiates a new InlineObject164 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject164WithDefaults() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject164) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject164) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *InlineObject164) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject164) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAdminSettings returns the AdminSettings field value
func (o *InlineObject164) GetAdminSettings() OrganizationsOrganizationIdBrandingPoliciesAdminSettings {
	if o == nil {
		var ret OrganizationsOrganizationIdBrandingPoliciesAdminSettings
		return ret
	}

	return o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetAdminSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesAdminSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdminSettings, true
}

// SetAdminSettings sets field value
func (o *InlineObject164) SetAdminSettings(v OrganizationsOrganizationIdBrandingPoliciesAdminSettings) {
	o.AdminSettings = v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *InlineObject164) GetHelpSettings() OrganizationsOrganizationIdBrandingPoliciesHelpSettings {
	if o == nil || o.HelpSettings == nil {
		var ret OrganizationsOrganizationIdBrandingPoliciesHelpSettings
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetHelpSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesHelpSettings, bool) {
	if o == nil || o.HelpSettings == nil {
		return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *InlineObject164) HasHelpSettings() bool {
	if o != nil && o.HelpSettings != nil {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesHelpSettings and assigns it to the HelpSettings field.
func (o *InlineObject164) SetHelpSettings(v OrganizationsOrganizationIdBrandingPoliciesHelpSettings) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *InlineObject164) GetCustomLogo() OrganizationsOrganizationIdBrandingPoliciesCustomLogo {
	if o == nil || o.CustomLogo == nil {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogo
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetCustomLogoOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogo, bool) {
	if o == nil || o.CustomLogo == nil {
		return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *InlineObject164) HasCustomLogo() bool {
	if o != nil && o.CustomLogo != nil {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogo and assigns it to the CustomLogo field.
func (o *InlineObject164) SetCustomLogo(v OrganizationsOrganizationIdBrandingPoliciesCustomLogo) {
	o.CustomLogo = &v
}

func (o InlineObject164) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if o.HelpSettings != nil {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if o.CustomLogo != nil {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject164 struct {
	value *InlineObject164
	isSet bool
}

func (v NullableInlineObject164) Get() *InlineObject164 {
	return v.value
}

func (v *NullableInlineObject164) Set(val *InlineObject164) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject164) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject164) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject164(val *InlineObject164) *NullableInlineObject164 {
	return &NullableInlineObject164{value: val, isSet: true}
}

func (v NullableInlineObject164) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject164) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

