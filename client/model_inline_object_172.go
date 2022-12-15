/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject172 struct for InlineObject172
type InlineObject172 struct {
	// Name of the group
	Name *string `json:"name,omitempty"`
	// SGT value of the group
	Sgt *int32 `json:"sgt,omitempty"`
	// Description of the group
	Description *string `json:"description,omitempty"`
	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute)
	PolicyObjects []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects `json:"policyObjects,omitempty"`
}

// NewInlineObject172 instantiates a new InlineObject172 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject172() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// NewInlineObject172WithDefaults instantiates a new InlineObject172 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject172WithDefaults() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject172) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject172) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject172) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *InlineObject172) GetSgt() int32 {
	if o == nil || isNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetSgtOk() (*int32, bool) {
	if o == nil || isNil(o.Sgt) {
    return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *InlineObject172) HasSgt() bool {
	if o != nil && !isNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *InlineObject172) SetSgt(v int32) {
	o.Sgt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject172) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject172) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject172) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineObject172) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	if o == nil || isNil(o.PolicyObjects) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetPolicyObjectsOk() ([]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineObject172) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects and assigns it to the PolicyObjects field.
func (o *InlineObject172) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) {
	o.PolicyObjects = v
}

func (o InlineObject172) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableInlineObject172 struct {
	value *InlineObject172
	isSet bool
}

func (v NullableInlineObject172) Get() *InlineObject172 {
	return v.value
}

func (v *NullableInlineObject172) Set(val *InlineObject172) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject172) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject172) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject172(val *InlineObject172) *NullableInlineObject172 {
	return &NullableInlineObject172{value: val, isSet: true}
}

func (v NullableInlineObject172) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject172) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


