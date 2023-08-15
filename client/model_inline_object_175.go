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

// InlineObject175 struct for InlineObject175
type InlineObject175 struct {
	// Name of the adaptive policy ACL
	Name *string `json:"name,omitempty"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules. An empty array will clear the rules.
	Rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 `json:"rules,omitempty"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion *string `json:"ipVersion,omitempty"`
}

// NewInlineObject175 instantiates a new InlineObject175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject175() *InlineObject175 {
	this := InlineObject175{}
	return &this
}

// NewInlineObject175WithDefaults instantiates a new InlineObject175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject175WithDefaults() *InlineObject175 {
	this := InlineObject175{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject175) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject175) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject175) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject175) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject175) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject175) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject175) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	if o == nil || isNil(o.Rules) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyAclsRules1
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetRulesOk() ([]OrganizationsOrganizationIdAdaptivePolicyAclsRules1, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject175) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 and assigns it to the Rules field.
func (o *InlineObject175) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules1) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *InlineObject175) GetIpVersion() string {
	if o == nil || isNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetIpVersionOk() (*string, bool) {
	if o == nil || isNil(o.IpVersion) {
    return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *InlineObject175) HasIpVersion() bool {
	if o != nil && !isNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *InlineObject175) SetIpVersion(v string) {
	o.IpVersion = &v
}

func (o InlineObject175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject175 struct {
	value *InlineObject175
	isSet bool
}

func (v NullableInlineObject175) Get() *InlineObject175 {
	return v.value
}

func (v *NullableInlineObject175) Set(val *InlineObject175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject175(val *InlineObject175) *NullableInlineObject175 {
	return &NullableInlineObject175{value: val, isSet: true}
}

func (v NullableInlineObject175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


