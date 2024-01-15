/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject192 struct for InlineObject192
type InlineObject192 struct {
	SourceGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup `json:"sourceGroup"`
	DestinationGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup `json:"destinationGroup"`
	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: [])
	Acls []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL (default: \"default\")
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
}

// NewInlineObject192 instantiates a new InlineObject192 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject192(sourceGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, destinationGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup) *InlineObject192 {
	this := InlineObject192{}
	this.SourceGroup = sourceGroup
	this.DestinationGroup = destinationGroup
	return &this
}

// NewInlineObject192WithDefaults instantiates a new InlineObject192 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject192WithDefaults() *InlineObject192 {
	this := InlineObject192{}
	return &this
}

// GetSourceGroup returns the SourceGroup field value
func (o *InlineObject192) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup {
	if o == nil {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup
		return ret
	}

	return o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject192) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceGroup, true
}

// SetSourceGroup sets field value
func (o *InlineObject192) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup) {
	o.SourceGroup = v
}

// GetDestinationGroup returns the DestinationGroup field value
func (o *InlineObject192) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup {
	if o == nil {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup
		return ret
	}

	return o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject192) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestinationGroup, true
}

// SetDestinationGroup sets field value
func (o *InlineObject192) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup) {
	o.DestinationGroup = v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *InlineObject192) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	if o == nil || isNil(o.Acls) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject192) GetAclsOk() ([]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool) {
	if o == nil || isNil(o.Acls) {
    return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *InlineObject192) HasAcls() bool {
	if o != nil && !isNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls and assigns it to the Acls field.
func (o *InlineObject192) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *InlineObject192) GetLastEntryRule() string {
	if o == nil || isNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject192) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || isNil(o.LastEntryRule) {
    return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *InlineObject192) HasLastEntryRule() bool {
	if o != nil && !isNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *InlineObject192) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

func (o InlineObject192) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if true {
		toSerialize["destinationGroup"] = o.DestinationGroup
	}
	if !isNil(o.Acls) {
		toSerialize["acls"] = o.Acls
	}
	if !isNil(o.LastEntryRule) {
		toSerialize["lastEntryRule"] = o.LastEntryRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject192 struct {
	value *InlineObject192
	isSet bool
}

func (v NullableInlineObject192) Get() *InlineObject192 {
	return v.value
}

func (v *NullableInlineObject192) Set(val *InlineObject192) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject192) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject192) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject192(val *InlineObject192) *NullableInlineObject192 {
	return &NullableInlineObject192{value: val, isSet: true}
}

func (v NullableInlineObject192) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject192) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


