/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject204 struct for InlineObject204
type InlineObject204 struct {
	SourceGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 `json:"sourceGroup,omitempty"`
	DestinationGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 `json:"destinationGroup,omitempty"`
	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	Acls []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
}

// NewInlineObject204 instantiates a new InlineObject204 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject204() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// NewInlineObject204WithDefaults instantiates a new InlineObject204 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject204WithDefaults() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *InlineObject204) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 {
	if o == nil || isNil(o.SourceGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1, bool) {
	if o == nil || isNil(o.SourceGroup) {
    return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *InlineObject204) HasSourceGroup() bool {
	if o != nil && !isNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 and assigns it to the SourceGroup field.
func (o *InlineObject204) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *InlineObject204) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	if o == nil || isNil(o.DestinationGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1, bool) {
	if o == nil || isNil(o.DestinationGroup) {
    return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *InlineObject204) HasDestinationGroup() bool {
	if o != nil && !isNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 and assigns it to the DestinationGroup field.
func (o *InlineObject204) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *InlineObject204) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 {
	if o == nil || isNil(o.Acls) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetAclsOk() ([]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1, bool) {
	if o == nil || isNil(o.Acls) {
    return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *InlineObject204) HasAcls() bool {
	if o != nil && !isNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 and assigns it to the Acls field.
func (o *InlineObject204) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *InlineObject204) GetLastEntryRule() string {
	if o == nil || isNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || isNil(o.LastEntryRule) {
    return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *InlineObject204) HasLastEntryRule() bool {
	if o != nil && !isNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *InlineObject204) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

func (o InlineObject204) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceGroup) {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if !isNil(o.DestinationGroup) {
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

type NullableInlineObject204 struct {
	value *InlineObject204
	isSet bool
}

func (v NullableInlineObject204) Get() *InlineObject204 {
	return v.value
}

func (v *NullableInlineObject204) Set(val *InlineObject204) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject204) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject204) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject204(val *InlineObject204) *NullableInlineObject204 {
	return &NullableInlineObject204{value: val, isSet: true}
}

func (v NullableInlineObject204) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject204) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


