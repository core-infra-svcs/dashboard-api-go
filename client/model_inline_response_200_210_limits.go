/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200210Limits The current limits of various adaptive policy objects.
type InlineResponse200210Limits struct {
	// Maximum number of user-created adaptive policy groups allowed in the organization.
	CustomGroups *int32 `json:"customGroups,omitempty"`
	// Maximum number of rules allowed in an adaptive policy ACL in the organization.
	RulesInAnAcl *int32 `json:"rulesInAnAcl,omitempty"`
	// Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.
	AclsInAPolicy *int32 `json:"aclsInAPolicy,omitempty"`
	// Maximum number of policy objects (with the adaptive policy type) allowed in the organization.
	PolicyObjects *int32 `json:"policyObjects,omitempty"`
}

// NewInlineResponse200210Limits instantiates a new InlineResponse200210Limits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200210Limits() *InlineResponse200210Limits {
	this := InlineResponse200210Limits{}
	return &this
}

// NewInlineResponse200210LimitsWithDefaults instantiates a new InlineResponse200210Limits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200210LimitsWithDefaults() *InlineResponse200210Limits {
	this := InlineResponse200210Limits{}
	return &this
}

// GetCustomGroups returns the CustomGroups field value if set, zero value otherwise.
func (o *InlineResponse200210Limits) GetCustomGroups() int32 {
	if o == nil || isNil(o.CustomGroups) {
		var ret int32
		return ret
	}
	return *o.CustomGroups
}

// GetCustomGroupsOk returns a tuple with the CustomGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200210Limits) GetCustomGroupsOk() (*int32, bool) {
	if o == nil || isNil(o.CustomGroups) {
    return nil, false
	}
	return o.CustomGroups, true
}

// HasCustomGroups returns a boolean if a field has been set.
func (o *InlineResponse200210Limits) HasCustomGroups() bool {
	if o != nil && !isNil(o.CustomGroups) {
		return true
	}

	return false
}

// SetCustomGroups gets a reference to the given int32 and assigns it to the CustomGroups field.
func (o *InlineResponse200210Limits) SetCustomGroups(v int32) {
	o.CustomGroups = &v
}

// GetRulesInAnAcl returns the RulesInAnAcl field value if set, zero value otherwise.
func (o *InlineResponse200210Limits) GetRulesInAnAcl() int32 {
	if o == nil || isNil(o.RulesInAnAcl) {
		var ret int32
		return ret
	}
	return *o.RulesInAnAcl
}

// GetRulesInAnAclOk returns a tuple with the RulesInAnAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200210Limits) GetRulesInAnAclOk() (*int32, bool) {
	if o == nil || isNil(o.RulesInAnAcl) {
    return nil, false
	}
	return o.RulesInAnAcl, true
}

// HasRulesInAnAcl returns a boolean if a field has been set.
func (o *InlineResponse200210Limits) HasRulesInAnAcl() bool {
	if o != nil && !isNil(o.RulesInAnAcl) {
		return true
	}

	return false
}

// SetRulesInAnAcl gets a reference to the given int32 and assigns it to the RulesInAnAcl field.
func (o *InlineResponse200210Limits) SetRulesInAnAcl(v int32) {
	o.RulesInAnAcl = &v
}

// GetAclsInAPolicy returns the AclsInAPolicy field value if set, zero value otherwise.
func (o *InlineResponse200210Limits) GetAclsInAPolicy() int32 {
	if o == nil || isNil(o.AclsInAPolicy) {
		var ret int32
		return ret
	}
	return *o.AclsInAPolicy
}

// GetAclsInAPolicyOk returns a tuple with the AclsInAPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200210Limits) GetAclsInAPolicyOk() (*int32, bool) {
	if o == nil || isNil(o.AclsInAPolicy) {
    return nil, false
	}
	return o.AclsInAPolicy, true
}

// HasAclsInAPolicy returns a boolean if a field has been set.
func (o *InlineResponse200210Limits) HasAclsInAPolicy() bool {
	if o != nil && !isNil(o.AclsInAPolicy) {
		return true
	}

	return false
}

// SetAclsInAPolicy gets a reference to the given int32 and assigns it to the AclsInAPolicy field.
func (o *InlineResponse200210Limits) SetAclsInAPolicy(v int32) {
	o.AclsInAPolicy = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineResponse200210Limits) GetPolicyObjects() int32 {
	if o == nil || isNil(o.PolicyObjects) {
		var ret int32
		return ret
	}
	return *o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200210Limits) GetPolicyObjectsOk() (*int32, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineResponse200210Limits) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given int32 and assigns it to the PolicyObjects field.
func (o *InlineResponse200210Limits) SetPolicyObjects(v int32) {
	o.PolicyObjects = &v
}

func (o InlineResponse200210Limits) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200210Limits struct {
	value *InlineResponse200210Limits
	isSet bool
}

func (v NullableInlineResponse200210Limits) Get() *InlineResponse200210Limits {
	return v.value
}

func (v *NullableInlineResponse200210Limits) Set(val *InlineResponse200210Limits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200210Limits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200210Limits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200210Limits(val *InlineResponse200210Limits) *NullableInlineResponse200210Limits {
	return &NullableInlineResponse200210Limits{value: val, isSet: true}
}

func (v NullableInlineResponse200210Limits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200210Limits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


