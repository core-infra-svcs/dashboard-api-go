/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject233 struct for InlineObject233
type InlineObject233 struct {
	// Name of the Dashboard branding policy.
	Name *string `json:"name,omitempty"`
	// Boolean indicating whether this policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AdminSettings *OrganizationsOrganizationIdBrandingPoliciesAdminSettings `json:"adminSettings,omitempty"`
	HelpSettings *OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 `json:"helpSettings,omitempty"`
	CustomLogo *OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 `json:"customLogo,omitempty"`
}

// NewInlineObject233 instantiates a new InlineObject233 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject233() *InlineObject233 {
	this := InlineObject233{}
	return &this
}

// NewInlineObject233WithDefaults instantiates a new InlineObject233 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject233WithDefaults() *InlineObject233 {
	this := InlineObject233{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject233) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject233) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject233) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject233) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject233) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject233) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *InlineObject233) GetAdminSettings() OrganizationsOrganizationIdBrandingPoliciesAdminSettings {
	if o == nil || isNil(o.AdminSettings) {
		var ret OrganizationsOrganizationIdBrandingPoliciesAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetAdminSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesAdminSettings, bool) {
	if o == nil || isNil(o.AdminSettings) {
    return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *InlineObject233) HasAdminSettings() bool {
	if o != nil && !isNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesAdminSettings and assigns it to the AdminSettings field.
func (o *InlineObject233) SetAdminSettings(v OrganizationsOrganizationIdBrandingPoliciesAdminSettings) {
	o.AdminSettings = &v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *InlineObject233) GetHelpSettings() OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 {
	if o == nil || isNil(o.HelpSettings) {
		var ret OrganizationsOrganizationIdBrandingPoliciesHelpSettings1
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetHelpSettingsOk() (*OrganizationsOrganizationIdBrandingPoliciesHelpSettings1, bool) {
	if o == nil || isNil(o.HelpSettings) {
    return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *InlineObject233) HasHelpSettings() bool {
	if o != nil && !isNil(o.HelpSettings) {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesHelpSettings1 and assigns it to the HelpSettings field.
func (o *InlineObject233) SetHelpSettings(v OrganizationsOrganizationIdBrandingPoliciesHelpSettings1) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *InlineObject233) GetCustomLogo() OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 {
	if o == nil || isNil(o.CustomLogo) {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogo1
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetCustomLogoOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogo1, bool) {
	if o == nil || isNil(o.CustomLogo) {
    return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *InlineObject233) HasCustomLogo() bool {
	if o != nil && !isNil(o.CustomLogo) {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogo1 and assigns it to the CustomLogo field.
func (o *InlineObject233) SetCustomLogo(v OrganizationsOrganizationIdBrandingPoliciesCustomLogo1) {
	o.CustomLogo = &v
}

func (o InlineObject233) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject233 struct {
	value *InlineObject233
	isSet bool
}

func (v NullableInlineObject233) Get() *InlineObject233 {
	return v.value
}

func (v *NullableInlineObject233) Set(val *InlineObject233) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject233) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject233) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject233(val *InlineObject233) *NullableInlineObject233 {
	return &NullableInlineObject233{value: val, isSet: true}
}

func (v NullableInlineObject233) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject233) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


