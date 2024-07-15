/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200199 struct for InlineResponse200199
type InlineResponse200199 struct {
	// The ID of the adaptive policy group
	GroupId *string `json:"groupId,omitempty"`
	// The name of the adaptive policy group
	Name *string `json:"name,omitempty"`
	// The security group tag for the adaptive policy group
	Sgt *int32 `json:"sgt,omitempty"`
	// The description for the adaptive policy group
	Description *string `json:"description,omitempty"`
	// The policy objects for the adaptive policy group
	PolicyObjects []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects `json:"policyObjects,omitempty"`
	// Whether the adaptive policy group is the default group
	IsDefaultGroup *bool `json:"isDefaultGroup,omitempty"`
	// List of required IP mappings for the adaptive policy group
	RequiredIpMappings []string `json:"requiredIpMappings,omitempty"`
	// Created at timestamp for the adaptive policy group
	CreatedAt *string `json:"createdAt,omitempty"`
	// Updated at timestamp for the adaptive policy group
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewInlineResponse200199 instantiates a new InlineResponse200199 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200199() *InlineResponse200199 {
	this := InlineResponse200199{}
	return &this
}

// NewInlineResponse200199WithDefaults instantiates a new InlineResponse200199 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200199WithDefaults() *InlineResponse200199 {
	this := InlineResponse200199{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *InlineResponse200199) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineResponse200199) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *InlineResponse200199) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200199) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200199) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200199) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *InlineResponse200199) GetSgt() int32 {
	if o == nil || isNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetSgtOk() (*int32, bool) {
	if o == nil || isNil(o.Sgt) {
    return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *InlineResponse200199) HasSgt() bool {
	if o != nil && !isNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *InlineResponse200199) SetSgt(v int32) {
	o.Sgt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200199) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200199) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200199) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineResponse200199) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	if o == nil || isNil(o.PolicyObjects) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetPolicyObjectsOk() ([]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineResponse200199) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects and assigns it to the PolicyObjects field.
func (o *InlineResponse200199) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) {
	o.PolicyObjects = v
}

// GetIsDefaultGroup returns the IsDefaultGroup field value if set, zero value otherwise.
func (o *InlineResponse200199) GetIsDefaultGroup() bool {
	if o == nil || isNil(o.IsDefaultGroup) {
		var ret bool
		return ret
	}
	return *o.IsDefaultGroup
}

// GetIsDefaultGroupOk returns a tuple with the IsDefaultGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetIsDefaultGroupOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefaultGroup) {
    return nil, false
	}
	return o.IsDefaultGroup, true
}

// HasIsDefaultGroup returns a boolean if a field has been set.
func (o *InlineResponse200199) HasIsDefaultGroup() bool {
	if o != nil && !isNil(o.IsDefaultGroup) {
		return true
	}

	return false
}

// SetIsDefaultGroup gets a reference to the given bool and assigns it to the IsDefaultGroup field.
func (o *InlineResponse200199) SetIsDefaultGroup(v bool) {
	o.IsDefaultGroup = &v
}

// GetRequiredIpMappings returns the RequiredIpMappings field value if set, zero value otherwise.
func (o *InlineResponse200199) GetRequiredIpMappings() []string {
	if o == nil || isNil(o.RequiredIpMappings) {
		var ret []string
		return ret
	}
	return o.RequiredIpMappings
}

// GetRequiredIpMappingsOk returns a tuple with the RequiredIpMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetRequiredIpMappingsOk() ([]string, bool) {
	if o == nil || isNil(o.RequiredIpMappings) {
    return nil, false
	}
	return o.RequiredIpMappings, true
}

// HasRequiredIpMappings returns a boolean if a field has been set.
func (o *InlineResponse200199) HasRequiredIpMappings() bool {
	if o != nil && !isNil(o.RequiredIpMappings) {
		return true
	}

	return false
}

// SetRequiredIpMappings gets a reference to the given []string and assigns it to the RequiredIpMappings field.
func (o *InlineResponse200199) SetRequiredIpMappings(v []string) {
	o.RequiredIpMappings = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200199) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200199) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse200199) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200199) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200199) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200199) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *InlineResponse200199) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o InlineResponse200199) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Sgt) {
		toSerialize["sgt"] = o.Sgt
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	if !isNil(o.IsDefaultGroup) {
		toSerialize["isDefaultGroup"] = o.IsDefaultGroup
	}
	if !isNil(o.RequiredIpMappings) {
		toSerialize["requiredIpMappings"] = o.RequiredIpMappings
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200199 struct {
	value *InlineResponse200199
	isSet bool
}

func (v NullableInlineResponse200199) Get() *InlineResponse200199 {
	return v.value
}

func (v *NullableInlineResponse200199) Set(val *InlineResponse200199) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200199) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200199) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200199(val *InlineResponse200199) *NullableInlineResponse200199 {
	return &NullableInlineResponse200199{value: val, isSet: true}
}

func (v NullableInlineResponse200199) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200199) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


