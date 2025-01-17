/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200207Counts The current amount of various adaptive policy objects.
type InlineResponse200207Counts struct {
	// Number of adaptive policy groups currently in the organization.
	Groups *int32 `json:"groups,omitempty"`
	// Number of user-created adaptive policy groups currently in the organization.
	CustomGroups *int32 `json:"customGroups,omitempty"`
	// Number of user-created adaptive policy ACLs currently in the organization.
	CustomAcls *int32 `json:"customAcls,omitempty"`
	// Number of adaptive policies currently in the organization.
	Policies *int32 `json:"policies,omitempty"`
	// Number of adaptive policies currently in the organization that deny all traffic.
	DenyPolicies *int32 `json:"denyPolicies,omitempty"`
	// Number of adaptive policies currently in the organization that allow all traffic.
	AllowPolicies *int32 `json:"allowPolicies,omitempty"`
	// Number of policy objects (with the adaptive policy type) currently in the organization.
	PolicyObjects *int32 `json:"policyObjects,omitempty"`
}

// NewInlineResponse200207Counts instantiates a new InlineResponse200207Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200207Counts() *InlineResponse200207Counts {
	this := InlineResponse200207Counts{}
	return &this
}

// NewInlineResponse200207CountsWithDefaults instantiates a new InlineResponse200207Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200207CountsWithDefaults() *InlineResponse200207Counts {
	this := InlineResponse200207Counts{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetGroups() int32 {
	if o == nil || isNil(o.Groups) {
		var ret int32
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetGroupsOk() (*int32, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given int32 and assigns it to the Groups field.
func (o *InlineResponse200207Counts) SetGroups(v int32) {
	o.Groups = &v
}

// GetCustomGroups returns the CustomGroups field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetCustomGroups() int32 {
	if o == nil || isNil(o.CustomGroups) {
		var ret int32
		return ret
	}
	return *o.CustomGroups
}

// GetCustomGroupsOk returns a tuple with the CustomGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetCustomGroupsOk() (*int32, bool) {
	if o == nil || isNil(o.CustomGroups) {
    return nil, false
	}
	return o.CustomGroups, true
}

// HasCustomGroups returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasCustomGroups() bool {
	if o != nil && !isNil(o.CustomGroups) {
		return true
	}

	return false
}

// SetCustomGroups gets a reference to the given int32 and assigns it to the CustomGroups field.
func (o *InlineResponse200207Counts) SetCustomGroups(v int32) {
	o.CustomGroups = &v
}

// GetCustomAcls returns the CustomAcls field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetCustomAcls() int32 {
	if o == nil || isNil(o.CustomAcls) {
		var ret int32
		return ret
	}
	return *o.CustomAcls
}

// GetCustomAclsOk returns a tuple with the CustomAcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetCustomAclsOk() (*int32, bool) {
	if o == nil || isNil(o.CustomAcls) {
    return nil, false
	}
	return o.CustomAcls, true
}

// HasCustomAcls returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasCustomAcls() bool {
	if o != nil && !isNil(o.CustomAcls) {
		return true
	}

	return false
}

// SetCustomAcls gets a reference to the given int32 and assigns it to the CustomAcls field.
func (o *InlineResponse200207Counts) SetCustomAcls(v int32) {
	o.CustomAcls = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetPolicies() int32 {
	if o == nil || isNil(o.Policies) {
		var ret int32
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetPoliciesOk() (*int32, bool) {
	if o == nil || isNil(o.Policies) {
    return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given int32 and assigns it to the Policies field.
func (o *InlineResponse200207Counts) SetPolicies(v int32) {
	o.Policies = &v
}

// GetDenyPolicies returns the DenyPolicies field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetDenyPolicies() int32 {
	if o == nil || isNil(o.DenyPolicies) {
		var ret int32
		return ret
	}
	return *o.DenyPolicies
}

// GetDenyPoliciesOk returns a tuple with the DenyPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetDenyPoliciesOk() (*int32, bool) {
	if o == nil || isNil(o.DenyPolicies) {
    return nil, false
	}
	return o.DenyPolicies, true
}

// HasDenyPolicies returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasDenyPolicies() bool {
	if o != nil && !isNil(o.DenyPolicies) {
		return true
	}

	return false
}

// SetDenyPolicies gets a reference to the given int32 and assigns it to the DenyPolicies field.
func (o *InlineResponse200207Counts) SetDenyPolicies(v int32) {
	o.DenyPolicies = &v
}

// GetAllowPolicies returns the AllowPolicies field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetAllowPolicies() int32 {
	if o == nil || isNil(o.AllowPolicies) {
		var ret int32
		return ret
	}
	return *o.AllowPolicies
}

// GetAllowPoliciesOk returns a tuple with the AllowPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetAllowPoliciesOk() (*int32, bool) {
	if o == nil || isNil(o.AllowPolicies) {
    return nil, false
	}
	return o.AllowPolicies, true
}

// HasAllowPolicies returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasAllowPolicies() bool {
	if o != nil && !isNil(o.AllowPolicies) {
		return true
	}

	return false
}

// SetAllowPolicies gets a reference to the given int32 and assigns it to the AllowPolicies field.
func (o *InlineResponse200207Counts) SetAllowPolicies(v int32) {
	o.AllowPolicies = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineResponse200207Counts) GetPolicyObjects() int32 {
	if o == nil || isNil(o.PolicyObjects) {
		var ret int32
		return ret
	}
	return *o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200207Counts) GetPolicyObjectsOk() (*int32, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineResponse200207Counts) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given int32 and assigns it to the PolicyObjects field.
func (o *InlineResponse200207Counts) SetPolicyObjects(v int32) {
	o.PolicyObjects = &v
}

func (o InlineResponse200207Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.CustomGroups) {
		toSerialize["customGroups"] = o.CustomGroups
	}
	if !isNil(o.CustomAcls) {
		toSerialize["customAcls"] = o.CustomAcls
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !isNil(o.DenyPolicies) {
		toSerialize["denyPolicies"] = o.DenyPolicies
	}
	if !isNil(o.AllowPolicies) {
		toSerialize["allowPolicies"] = o.AllowPolicies
	}
	if !isNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200207Counts struct {
	value *InlineResponse200207Counts
	isSet bool
}

func (v NullableInlineResponse200207Counts) Get() *InlineResponse200207Counts {
	return v.value
}

func (v *NullableInlineResponse200207Counts) Set(val *InlineResponse200207Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200207Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200207Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200207Counts(val *InlineResponse200207Counts) *NullableInlineResponse200207Counts {
	return &NullableInlineResponse200207Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200207Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200207Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


