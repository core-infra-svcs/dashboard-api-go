/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject184 struct for InlineObject184
type InlineObject184 struct {
	// Name of the Dashboard branding policy.
	Name *string `json:"name,omitempty"`
	// Boolean indicating whether this policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AdminSettings *OrganizationsOrganizationIdBrandingPoliciesAdminSettings `json:"adminSettings,omitempty"`
	HelpSettings *OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 `json:"helpSettings,omitempty"`
	CustomLogo *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 `json:"customLogo,omitempty"`
}

// NewInlineObject184 instantiates a new InlineObject184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject184() *InlineObject184 {
	this := InlineObject184{}
	return &this
}

// NewInlineObject184WithDefaults instantiates a new InlineObject184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject184WithDefaults() *InlineObject184 {
	this := InlineObject184{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject184) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject184) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject184) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject184) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject184) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject184) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *InlineObject184) GetAdminSettings() OrganizationsOrganizationIdBrandingPoliciesAdminSettings {
	if o == nil || isNil(o.AdminSettings) {
		var ret OrganizationsOrganizationIdBrandingPoliciesAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetAdminSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesAdminSettings, bool) {
	if o == nil || isNil(o.AdminSettings) {
    return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *InlineObject184) HasAdminSettings() bool {
	if o != nil && !isNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesAdminSettings and assigns it to the AdminSettings field.
func (o *InlineObject184) SetAdminSettings(v OrganizationsOrganizationIdBrandingPoliciesAdminSettings) {
	o.AdminSettings = &v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *InlineObject184) GetHelpSettings() OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 {
	if o == nil || isNil(o.HelpSettings) {
		var ret OrganizationsOrganizationIdBrandingPoliciesHelpSettings1
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetHelpSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesHelpSettings1, bool) {
	if o == nil || isNil(o.HelpSettings) {
    return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *InlineObject184) HasHelpSettings() bool {
	if o != nil && !isNil(o.HelpSettings) {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 and assigns it to the HelpSettings field.
func (o *InlineObject184) SetHelpSettings(v OrganizationsOrganizationIdBrandingPoliciesHelpSettings1) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *InlineObject184) GetCustomLogo() OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	if o == nil || isNil(o.CustomLogo) {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogo1
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject184) GetCustomLogoOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogo1, bool) {
	if o == nil || isNil(o.CustomLogo) {
    return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *InlineObject184) HasCustomLogo() bool {
	if o != nil && !isNil(o.CustomLogo) {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 and assigns it to the CustomLogo field.
func (o *InlineObject184) SetCustomLogo(v OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) {
	o.CustomLogo = &v
}

func (o InlineObject184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AdminSettings) {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if !isNil(o.HelpSettings) {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if !isNil(o.CustomLogo) {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject184 struct {
	value *InlineObject184
	isSet bool
}

func (v NullableInlineObject184) Get() *InlineObject184 {
	return v.value
}

func (v *NullableInlineObject184) Set(val *InlineObject184) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject184) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject184(val *InlineObject184) *NullableInlineObject184 {
	return &NullableInlineObject184{value: val, isSet: true}
}

func (v NullableInlineObject184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


