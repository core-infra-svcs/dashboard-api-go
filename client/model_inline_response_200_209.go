/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200209 struct for InlineResponse200209
type InlineResponse200209 struct {
	// ID of the adaptive policy ACL
	AclId *string `json:"aclId,omitempty"`
	// Name of the adaptive policy ACL
	Name *string `json:"name,omitempty"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// IP version of adpative policy ACL
	IpVersion *string `json:"ipVersion,omitempty"`
	// An ordered array of the adaptive policy ACL rules
	Rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules `json:"rules,omitempty"`
	// When the adaptive policy ACL was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// When the adaptive policy ACL was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewInlineResponse200209 instantiates a new InlineResponse200209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200209() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200209WithDefaults() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *InlineResponse200209) GetAclId() string {
	if o == nil || isNil(o.AclId) {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetAclIdOk() (*string, bool) {
	if o == nil || isNil(o.AclId) {
    return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *InlineResponse200209) HasAclId() bool {
	if o != nil && !isNil(o.AclId) {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *InlineResponse200209) SetAclId(v string) {
	o.AclId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200209) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200209) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200209) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200209) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200209) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200209) SetDescription(v string) {
	o.Description = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *InlineResponse200209) GetIpVersion() string {
	if o == nil || isNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetIpVersionOk() (*string, bool) {
	if o == nil || isNil(o.IpVersion) {
    return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *InlineResponse200209) HasIpVersion() bool {
	if o != nil && !isNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *InlineResponse200209) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200209) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules {
	if o == nil || isNil(o.Rules) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyAclsRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetRulesOk() ([]OrganizationsOrganizationIdAdaptivePolicyAclsRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200209) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyAclsRules and assigns it to the Rules field.
func (o *InlineResponse200209) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules) {
	o.Rules = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200209) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200209) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200209) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200209) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200209) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InlineResponse200209) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InlineResponse200209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AclId) {
		toSerialize["aclId"] = o.AclId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200209 struct {
	value *InlineResponse200209
	isSet bool
}

func (v NullableInlineResponse200209) Get() *InlineResponse200209 {
	return v.value
}

func (v *NullableInlineResponse200209) Set(val *InlineResponse200209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200209(val *InlineResponse200209) *NullableInlineResponse200209 {
	return &NullableInlineResponse200209{value: val, isSet: true}
}

func (v NullableInlineResponse200209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


