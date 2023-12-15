/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject187 struct for InlineObject187
type InlineObject187 struct {
	// Name of the group
	Name string `json:"name"`
	// SGT value of the group
	Sgt int32 `json:"sgt"`
	// Description of the group (default: \"\")
	Description *string `json:"description,omitempty"`
	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute) (default: [])
	PolicyObjects []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects `json:"policyObjects,omitempty"`
}

// NewInlineObject187 instantiates a new InlineObject187 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject187(name string, sgt int32) *InlineObject187 {
	this := InlineObject187{}
	this.Name = name
	this.Sgt = sgt
	return &this
}

// NewInlineObject187WithDefaults instantiates a new InlineObject187 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject187WithDefaults() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject187) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject187) SetName(v string) {
	o.Name = v
}

// GetSgt returns the Sgt field value
func (o *InlineObject187) GetSgt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetSgtOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sgt, true
}

// SetSgt sets field value
func (o *InlineObject187) SetSgt(v int32) {
	o.Sgt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject187) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject187) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject187) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *InlineObject187) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects {
	if o == nil || isNil(o.PolicyObjects) {
		var ret []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetPolicyObjectsOk() ([]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool) {
	if o == nil || isNil(o.PolicyObjects) {
    return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *InlineObject187) HasPolicyObjects() bool {
	if o != nil && !isNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects and assigns it to the PolicyObjects field.
func (o *InlineObject187) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects) {
	o.PolicyObjects = v
}

func (o InlineObject187) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
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

type NullableInlineObject187 struct {
	value *InlineObject187
	isSet bool
}

func (v NullableInlineObject187) Get() *InlineObject187 {
	return v.value
}

func (v *NullableInlineObject187) Set(val *InlineObject187) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject187) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject187) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject187(val *InlineObject187) *NullableInlineObject187 {
	return &NullableInlineObject187{value: val, isSet: true}
}

func (v NullableInlineObject187) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject187) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


