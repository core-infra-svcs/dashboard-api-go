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

// InlineObject204 struct for InlineObject204
type InlineObject204 struct {
	// Name of the group
	Name *string `json:"name,omitempty"`
	// SGT value of the group
	Sgt *int32 `json:"sgt,omitempty"`
	// Description of the group
	Description *string `json:"description,omitempty"`
	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute)
	PolicyObjects []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects `json:"policyObjects,omitempty"`
}

// NewInlineObject204 instantiates a new InlineObject204 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject204() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// NewInlineObject204WithDefaults instantiates a new InlineObject204 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject204WithDefaults() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject204) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject204) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject204) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *InlineObject204) GetSgt() int32 {
	if o == nil || isNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetSgtOk() (*int32, bool) {
	if o == nil || isNil(o.Sgt) {
    return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *InlineObject204) HasSgt() bool {
	if o != nil && !isNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *InlineObject204) SetSgt(v int32) {
	o.Sgt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject204) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject204) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject204) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineObject204) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	if o == nil || isNil(o.PolicyObjects) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetPolicyObjectsOk() ([]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineObject204) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects and assigns it to the PolicyObjects field.
func (o *InlineObject204) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) {
	o.PolicyObjects = v
}

func (o InlineObject204) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject204 struct {
	value *InlineObject204
	isSet bool
}

func (v NullableInlineObject204) Get() *InlineObject204 {
	return v.value
}

func (v *NullableInlineObject204) Set(val *InlineObject204) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject204) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject204) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject204(val *InlineObject204) *NullableInlineObject204 {
	return &NullableInlineObject204{value: val, isSet: true}
}

func (v NullableInlineObject204) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject204) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


