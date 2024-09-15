/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject206 struct for InlineObject206
type InlineObject206 struct {
	SourceGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 `json:"sourceGroup,omitempty"`
	DestinationGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 `json:"destinationGroup,omitempty"`
	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	Acls []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
}

// NewInlineObject206 instantiates a new InlineObject206 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject206() *InlineObject206 {
	this := InlineObject206{}
	return &this
}

// NewInlineObject206WithDefaults instantiates a new InlineObject206 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject206WithDefaults() *InlineObject206 {
	this := InlineObject206{}
	return &this
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *InlineObject206) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 {
	if o == nil || isNil(o.SourceGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1, bool) {
	if o == nil || isNil(o.SourceGroup) {
    return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *InlineObject206) HasSourceGroup() bool {
	if o != nil && !isNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1 and assigns it to the SourceGroup field.
func (o *InlineObject206) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *InlineObject206) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 {
	if o == nil || isNil(o.DestinationGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1, bool) {
	if o == nil || isNil(o.DestinationGroup) {
    return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *InlineObject206) HasDestinationGroup() bool {
	if o != nil && !isNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1 and assigns it to the DestinationGroup field.
func (o *InlineObject206) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *InlineObject206) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 {
	if o == nil || isNil(o.Acls) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetAclsOk() ([]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1, bool) {
	if o == nil || isNil(o.Acls) {
    return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *InlineObject206) HasAcls() bool {
	if o != nil && !isNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1 and assigns it to the Acls field.
func (o *InlineObject206) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *InlineObject206) GetLastEntryRule() string {
	if o == nil || isNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || isNil(o.LastEntryRule) {
    return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *InlineObject206) HasLastEntryRule() bool {
	if o != nil && !isNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *InlineObject206) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

func (o InlineObject206) MarshalJSON() ([]byte, error) {
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


