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

// OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies struct for OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies
type OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies struct {
	// The Sentry Policy Id, if updating an existing Sentry Policy
	PolicyId *string `json:"policyId,omitempty"`
	// The Id of the Systems Manager Network
	SmNetworkId string `json:"smNetworkId"`
	// The scope of the Sentry Policy
	Scope string `json:"scope"`
	// The tags for the Sentry Policy
	Tags []string `json:"tags"`
	// The Group Policy Id
	GroupPolicyId string `json:"groupPolicyId"`
}

// NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies(smNetworkId string, scope string, tags []string, groupPolicyId string) *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies {
	this := OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies{}
	this.SmNetworkId = smNetworkId
	this.Scope = scope
	this.Tags = tags
	this.GroupPolicyId = groupPolicyId
	return &this
}

// NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPoliciesWithDefaults instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPoliciesWithDefaults() *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies {
	this := OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetPolicyId() string {
	if o == nil || isNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.PolicyId) {
    return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) HasPolicyId() bool {
	if o != nil && !isNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetSmNetworkId returns the SmNetworkId field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetSmNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmNetworkId
}

// GetSmNetworkIdOk returns a tuple with the SmNetworkId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetSmNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmNetworkId, true
}

// SetSmNetworkId sets field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetSmNetworkId(v string) {
	o.SmNetworkId = v
}

// GetScope returns the Scope field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetScopeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetScope(v string) {
	o.Scope = v
}

// GetTags returns the Tags field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetTags(v []string) {
	o.Tags = v
}

// GetGroupPolicyId returns the GroupPolicyId field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetGroupPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupPolicyId, true
}

// SetGroupPolicyId sets field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetGroupPolicyId(v string) {
	o.GroupPolicyId = v
}

func (o OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if true {
		toSerialize["smNetworkId"] = o.SmNetworkId
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies struct {
	value *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) Get() *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) Set(val *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies(val *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies {
	return &NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


