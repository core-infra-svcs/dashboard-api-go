/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject206 struct for InlineObject206
type InlineObject206 struct {
	// Name of the adaptive policy ACL
	Name string `json:"name"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules.
	Rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 `json:"rules"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion string `json:"ipVersion"`
}

// NewInlineObject206 instantiates a new InlineObject206 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject206(name string, rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules1, ipVersion string) *InlineObject206 {
	this := InlineObject206{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Rules = rules
	this.IpVersion = ipVersion
	return &this
}

// NewInlineObject206WithDefaults instantiates a new InlineObject206 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject206WithDefaults() *InlineObject206 {
	this := InlineObject206{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *InlineObject206) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject206) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject206) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject206) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject206) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value
func (o *InlineObject206) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	if o == nil {
		var ret []OrganizationsOrganizationIdAdaptivePolicyAclsRules1
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetRulesOk() ([]OrganizationsOrganizationIdAdaptivePolicyAclsRules1, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject206) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules1) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value
func (o *InlineObject206) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetIpVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *InlineObject206) SetIpVersion(v string) {
	o.IpVersion = v
}

func (o InlineObject206) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["ipVersion"] = o.IpVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject206 struct {
	value *InlineObject206
	isSet bool
}

func (v NullableInlineObject206) Get() *InlineObject206 {
	return v.value
}

func (v *NullableInlineObject206) Set(val *InlineObject206) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject206) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject206) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject206(val *InlineObject206) *NullableInlineObject206 {
	return &NullableInlineObject206{value: val, isSet: true}
}

func (v NullableInlineObject206) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject206) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


