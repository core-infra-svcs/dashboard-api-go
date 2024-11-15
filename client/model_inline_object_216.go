/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject216 struct for InlineObject216
type InlineObject216 struct {
	// Sets a list of specific SNORT signatures to allow
	AllowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules `json:"allowedRules"`
}

// NewInlineObject216 instantiates a new InlineObject216 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject216(allowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) *InlineObject216 {
	this := InlineObject216{}
	this.AllowedRules = allowedRules
	return &this
}

// NewInlineObject216WithDefaults instantiates a new InlineObject216 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject216WithDefaults() *InlineObject216 {
	this := InlineObject216{}
	return &this
}

// GetAllowedRules returns the AllowedRules field value
func (o *InlineObject216) GetAllowedRules() []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules
		return ret
	}

	return o.AllowedRules
}

// GetAllowedRulesOk returns a tuple with the AllowedRules field value
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetAllowedRulesOk() ([]OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedRules, true
}

// SetAllowedRules sets field value
func (o *InlineObject216) SetAllowedRules(v []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) {
	o.AllowedRules = v
}

func (o InlineObject216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedRules"] = o.AllowedRules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject216 struct {
	value *InlineObject216
	isSet bool
}

func (v NullableInlineObject216) Get() *InlineObject216 {
	return v.value
}

func (v *NullableInlineObject216) Set(val *InlineObject216) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject216) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject216) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject216(val *InlineObject216) *NullableInlineObject216 {
	return &NullableInlineObject216{value: val, isSet: true}
}

func (v NullableInlineObject216) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject216) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


