/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject190 struct for InlineObject190
type InlineObject190 struct {
	SourceGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup `json:"sourceGroup,omitempty"`
	DestinationGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup `json:"destinationGroup,omitempty"`
	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	Acls []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
}

// NewInlineObject190 instantiates a new InlineObject190 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject190() *InlineObject190 {
	this := InlineObject190{}
	return &this
}

// NewInlineObject190WithDefaults instantiates a new InlineObject190 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject190WithDefaults() *InlineObject190 {
	this := InlineObject190{}
	return &this
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *InlineObject190) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup {
	if o == nil || isNil(o.SourceGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject190) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool) {
	if o == nil || isNil(o.SourceGroup) {
    return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *InlineObject190) HasSourceGroup() bool {
	if o != nil && !isNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup and assigns it to the SourceGroup field.
func (o *InlineObject190) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *InlineObject190) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup {
	if o == nil || isNil(o.DestinationGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject190) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool) {
	if o == nil || isNil(o.DestinationGroup) {
    return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *InlineObject190) HasDestinationGroup() bool {
	if o != nil && !isNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup and assigns it to the DestinationGroup field.
func (o *InlineObject190) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *InlineObject190) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	if o == nil || isNil(o.Acls) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject190) GetAclsOk() ([]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool) {
	if o == nil || isNil(o.Acls) {
    return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *InlineObject190) HasAcls() bool {
	if o != nil && !isNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls and assigns it to the Acls field.
func (o *InlineObject190) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *InlineObject190) GetLastEntryRule() string {
	if o == nil || isNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject190) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || isNil(o.LastEntryRule) {
    return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *InlineObject190) HasLastEntryRule() bool {
	if o != nil && !isNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *InlineObject190) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

func (o InlineObject190) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject190 struct {
	value *InlineObject190
	isSet bool
}

func (v NullableInlineObject190) Get() *InlineObject190 {
	return v.value
}

func (v *NullableInlineObject190) Set(val *InlineObject190) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject190) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject190) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject190(val *InlineObject190) *NullableInlineObject190 {
	return &NullableInlineObject190{value: val, isSet: true}
}

func (v NullableInlineObject190) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject190) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


