/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200263 struct for InlineResponse200263
type InlineResponse200263 struct {
	// Policy object ID
	Id *string `json:"id,omitempty"`
	// Name of policy object (alphanumeric, space, dash, or underscore characters only).
	Name *string `json:"name,omitempty"`
	// Category of a policy object (one of: adaptivePolicy, network)
	Category *string `json:"category,omitempty"`
	// Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)
	Type *string `json:"type,omitempty"`
	// CIDR Value of a policy object
	Cidr *string `json:"cidr,omitempty"`
	// Time Stamp of policy object creation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time Stamp of policy object updation.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The IDs of policy object groups the policy object belongs to.
	GroupIds []string `json:"groupIds,omitempty"`
	// The IDs of the networks that use the policy object.
	NetworkIds []string `json:"networkIds,omitempty"`
}

// NewInlineResponse200263 instantiates a new InlineResponse200263 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200263() *InlineResponse200263 {
	this := InlineResponse200263{}
	return &this
}

// NewInlineResponse200263WithDefaults instantiates a new InlineResponse200263 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200263WithDefaults() *InlineResponse200263 {
	this := InlineResponse200263{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200263) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200263) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200263) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200263) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200263) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200263) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse200263) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse200263) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse200263) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200263) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200263) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200263) SetType(v string) {
	o.Type = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse200263) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse200263) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse200263) SetCidr(v string) {
	o.Cidr = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200263) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200263) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200263) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200263) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200263) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InlineResponse200263) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineResponse200263) GetGroupIds() []string {
	if o == nil || isNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineResponse200263) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineResponse200263) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetNetworkIds returns the NetworkIds field value if set, zero value otherwise.
func (o *InlineResponse200263) GetNetworkIds() []string {
	if o == nil || isNil(o.NetworkIds) {
		var ret []string
		return ret
	}
	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetNetworkIdsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkIds) {
    return nil, false
	}
	return o.NetworkIds, true
}

// HasNetworkIds returns a boolean if a field has been set.
func (o *InlineResponse200263) HasNetworkIds() bool {
	if o != nil && !isNil(o.NetworkIds) {
		return true
	}

	return false
}

// SetNetworkIds gets a reference to the given []string and assigns it to the NetworkIds field.
func (o *InlineResponse200263) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

func (o InlineResponse200263) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	if !isNil(o.NetworkIds) {
		toSerialize["networkIds"] = o.NetworkIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200263 struct {
	value *InlineResponse200263
	isSet bool
}

func (v NullableInlineResponse200263) Get() *InlineResponse200263 {
	return v.value
}

func (v *NullableInlineResponse200263) Set(val *InlineResponse200263) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200263) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200263) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200263(val *InlineResponse200263) *NullableInlineResponse200263 {
	return &NullableInlineResponse200263{value: val, isSet: true}
}

func (v NullableInlineResponse200263) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200263) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

