/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200202 struct for InlineResponse200202
type InlineResponse200202 struct {
	// The ID for the adaptive policy
	AdaptivePolicyId *string `json:"adaptivePolicyId,omitempty"`
	SourceGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup `json:"sourceGroup,omitempty"`
	DestinationGroup *OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup `json:"destinationGroup,omitempty"`
	// The access control lists for the adaptive policy
	Acls []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
	// The created at timestamp for the adaptive policy
	CreatedAt *string `json:"createdAt,omitempty"`
	// The updated at timestamp for the adaptive policy
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewInlineResponse200202 instantiates a new InlineResponse200202 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200202() *InlineResponse200202 {
	this := InlineResponse200202{}
	return &this
}

// NewInlineResponse200202WithDefaults instantiates a new InlineResponse200202 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200202WithDefaults() *InlineResponse200202 {
	this := InlineResponse200202{}
	return &this
}

// GetAdaptivePolicyId returns the AdaptivePolicyId field value if set, zero value otherwise.
func (o *InlineResponse200202) GetAdaptivePolicyId() string {
	if o == nil || isNil(o.AdaptivePolicyId) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyId
}

// GetAdaptivePolicyIdOk returns a tuple with the AdaptivePolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetAdaptivePolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.AdaptivePolicyId) {
    return nil, false
	}
	return o.AdaptivePolicyId, true
}

// HasAdaptivePolicyId returns a boolean if a field has been set.
func (o *InlineResponse200202) HasAdaptivePolicyId() bool {
	if o != nil && !isNil(o.AdaptivePolicyId) {
		return true
	}

	return false
}

// SetAdaptivePolicyId gets a reference to the given string and assigns it to the AdaptivePolicyId field.
func (o *InlineResponse200202) SetAdaptivePolicyId(v string) {
	o.AdaptivePolicyId = &v
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *InlineResponse200202) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup {
	if o == nil || isNil(o.SourceGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool) {
	if o == nil || isNil(o.SourceGroup) {
    return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *InlineResponse200202) HasSourceGroup() bool {
	if o != nil && !isNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup and assigns it to the SourceGroup field.
func (o *InlineResponse200202) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *InlineResponse200202) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup {
	if o == nil || isNil(o.DestinationGroup) {
		var ret OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool) {
	if o == nil || isNil(o.DestinationGroup) {
    return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *InlineResponse200202) HasDestinationGroup() bool {
	if o != nil && !isNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup and assigns it to the DestinationGroup field.
func (o *InlineResponse200202) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *InlineResponse200202) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls {
	if o == nil || isNil(o.Acls) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetAclsOk() ([]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool) {
	if o == nil || isNil(o.Acls) {
    return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *InlineResponse200202) HasAcls() bool {
	if o != nil && !isNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls and assigns it to the Acls field.
func (o *InlineResponse200202) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *InlineResponse200202) GetLastEntryRule() string {
	if o == nil || isNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || isNil(o.LastEntryRule) {
    return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *InlineResponse200202) HasLastEntryRule() bool {
	if o != nil && !isNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *InlineResponse200202) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200202) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200202) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse200202) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200202) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200202) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200202) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *InlineResponse200202) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o InlineResponse200202) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdaptivePolicyId) {
		toSerialize["adaptivePolicyId"] = o.AdaptivePolicyId
	}
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200202 struct {
	value *InlineResponse200202
	isSet bool
}

func (v NullableInlineResponse200202) Get() *InlineResponse200202 {
	return v.value
}

func (v *NullableInlineResponse200202) Set(val *InlineResponse200202) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200202) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200202) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200202(val *InlineResponse200202) *NullableInlineResponse200202 {
	return &NullableInlineResponse200202{value: val, isSet: true}
}

func (v NullableInlineResponse200202) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200202) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


