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

// InlineResponse200201Limits The current limits of various adaptive policy objects.
type InlineResponse200201Limits struct {
	// Maximum number of user-created adaptive policy groups allowed in the organization.
	CustomGroups *int32 `json:"customGroups,omitempty"`
	// Maximum number of rules allowed in an adaptive policy ACL in the organization.
	RulesInAnAcl *int32 `json:"rulesInAnAcl,omitempty"`
	// Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.
	AclsInAPolicy *int32 `json:"aclsInAPolicy,omitempty"`
	// Maximum number of policy objects (with the adaptive policy type) allowed in the organization.
	PolicyObjects *int32 `json:"policyObjects,omitempty"`
}

// NewInlineResponse200201Limits instantiates a new InlineResponse200201Limits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200201Limits() *InlineResponse200201Limits {
	this := InlineResponse200201Limits{}
	return &this
}

// NewInlineResponse200201LimitsWithDefaults instantiates a new InlineResponse200201Limits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200201LimitsWithDefaults() *InlineResponse200201Limits {
	this := InlineResponse200201Limits{}
	return &this
}

// GetCustomGroups returns the CustomGroups field value if set, zero value otherwise.
func (o *InlineResponse200201Limits) GetCustomGroups() int32 {
	if o == nil || isNil(o.CustomGroups) {
		var ret int32
		return ret
	}
	return *o.CustomGroups
}

// GetCustomGroupsOk returns a tuple with the CustomGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200201Limits) GetCustomGroupsOk() (*int32, bool) {
	if o == nil || isNil(o.CustomGroups) {
    return nil, false
	}
	return o.CustomGroups, true
}

// HasCustomGroups returns a boolean if a field has been set.
func (o *InlineResponse200201Limits) HasCustomGroups() bool {
	if o != nil && !isNil(o.CustomGroups) {
		return true
	}

	return false
}

// SetCustomGroups gets a reference to the given int32 and assigns it to the CustomGroups field.
func (o *InlineResponse200201Limits) SetCustomGroups(v int32) {
	o.CustomGroups = &v
}

// GetRulesInAnAcl returns the RulesInAnAcl field value if set, zero value otherwise.
func (o *InlineResponse200201Limits) GetRulesInAnAcl() int32 {
	if o == nil || isNil(o.RulesInAnAcl) {
		var ret int32
		return ret
	}
	return *o.RulesInAnAcl
}

// GetRulesInAnAclOk returns a tuple with the RulesInAnAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200201Limits) GetRulesInAnAclOk() (*int32, bool) {
	if o == nil || isNil(o.RulesInAnAcl) {
    return nil, false
	}
	return o.RulesInAnAcl, true
}

// HasRulesInAnAcl returns a boolean if a field has been set.
func (o *InlineResponse200201Limits) HasRulesInAnAcl() bool {
	if o != nil && !isNil(o.RulesInAnAcl) {
		return true
	}

	return false
}

// SetRulesInAnAcl gets a reference to the given int32 and assigns it to the RulesInAnAcl field.
func (o *InlineResponse200201Limits) SetRulesInAnAcl(v int32) {
	o.RulesInAnAcl = &v
}

// GetAclsInAPolicy returns the AclsInAPolicy field value if set, zero value otherwise.
func (o *InlineResponse200201Limits) GetAclsInAPolicy() int32 {
	if o == nil || isNil(o.AclsInAPolicy) {
		var ret int32
		return ret
	}
	return *o.AclsInAPolicy
}

// GetAclsInAPolicyOk returns a tuple with the AclsInAPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200201Limits) GetAclsInAPolicyOk() (*int32, bool) {
	if o == nil || isNil(o.AclsInAPolicy) {
    return nil, false
	}
	return o.AclsInAPolicy, true
}

// HasAclsInAPolicy returns a boolean if a field has been set.
func (o *InlineResponse200201Limits) HasAclsInAPolicy() bool {
	if o != nil && !isNil(o.AclsInAPolicy) {
		return true
	}

	return false
}

// SetAclsInAPolicy gets a reference to the given int32 and assigns it to the AclsInAPolicy field.
func (o *InlineResponse200201Limits) SetAclsInAPolicy(v int32) {
	o.AclsInAPolicy = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineResponse200201Limits) GetPolicyObjects() int32 {
	if o == nil || isNil(o.PolicyObjects) {
		var ret int32
		return ret
	}
	return *o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200201Limits) GetPolicyObjectsOk() (*int32, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineResponse200201Limits) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given int32 and assigns it to the PolicyObjects field.
func (o *InlineResponse200201Limits) SetPolicyObjects(v int32) {
	o.PolicyObjects = &v
}

func (o InlineResponse200201Limits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CustomGroups) {
		toSerialize["customGroups"] = o.CustomGroups
	}
	if !isNil(o.RulesInAnAcl) {
		toSerialize["rulesInAnAcl"] = o.RulesInAnAcl
	}
	if !isNil(o.AclsInAPolicy) {
		toSerialize["aclsInAPolicy"] = o.AclsInAPolicy
	}
	if !isNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200201Limits struct {
	value *InlineResponse200201Limits
	isSet bool
}

func (v NullableInlineResponse200201Limits) Get() *InlineResponse200201Limits {
	return v.value
}

func (v *NullableInlineResponse200201Limits) Set(val *InlineResponse200201Limits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200201Limits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200201Limits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200201Limits(val *InlineResponse200201Limits) *NullableInlineResponse200201Limits {
	return &NullableInlineResponse200201Limits{value: val, isSet: true}
}

func (v NullableInlineResponse200201Limits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200201Limits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


