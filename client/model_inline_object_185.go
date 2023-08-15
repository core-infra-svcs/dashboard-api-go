/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject185 struct for InlineObject185
type InlineObject185 struct {
	// Sets a list of specific SNORT signatures to allow
	AllowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules `json:"allowedRules"`
}

// NewInlineObject185 instantiates a new InlineObject185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject185(allowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) *InlineObject185 {
	this := InlineObject185{}
	this.AllowedRules = allowedRules
	return &this
}

// NewInlineObject185WithDefaults instantiates a new InlineObject185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject185WithDefaults() *InlineObject185 {
	this := InlineObject185{}
	return &this
}

// GetAllowedRules returns the AllowedRules field value
func (o *InlineObject185) GetAllowedRules() []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules
		return ret
	}

	return o.AllowedRules
}

// GetAllowedRulesOk returns a tuple with the AllowedRules field value
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetAllowedRulesOk() ([]OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedRules, true
}

// SetAllowedRules sets field value
func (o *InlineObject185) SetAllowedRules(v []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) {
	o.AllowedRules = v
}

func (o InlineObject185) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedRules"] = o.AllowedRules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject185 struct {
	value *InlineObject185
	isSet bool
}

func (v NullableInlineObject185) Get() *InlineObject185 {
	return v.value
}

func (v *NullableInlineObject185) Set(val *InlineObject185) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject185) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject185(val *InlineObject185) *NullableInlineObject185 {
	return &NullableInlineObject185{value: val, isSet: true}
}

func (v NullableInlineObject185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


