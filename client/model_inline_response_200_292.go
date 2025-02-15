/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200292 struct for InlineResponse200292
type InlineResponse200292 struct {
	// Policy object ID
	Id *string `json:"id,omitempty"`
	// Name of the Policy object group.
	Name *string `json:"name,omitempty"`
	// Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)
	Category *string `json:"category,omitempty"`
	// Time Stamp of policy object creation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time Stamp of policy object updation.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Policy objects associated with Network Object Group or Port Object Group
	ObjectIds []int32 `json:"objectIds,omitempty"`
	// Network ID's associated with the policy objects.
	NetworkIds []string `json:"networkIds,omitempty"`
}

// NewInlineResponse200292 instantiates a new InlineResponse200292 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200292() *InlineResponse200292 {
	this := InlineResponse200292{}
	return &this
}

// NewInlineResponse200292WithDefaults instantiates a new InlineResponse200292 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200292WithDefaults() *InlineResponse200292 {
	this := InlineResponse200292{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200292) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200292) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200292) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200292) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200292) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200292) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse200292) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse200292) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse200292) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200292) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200292) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200292) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200292) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200292) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InlineResponse200292) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetObjectIds returns the ObjectIds field value if set, zero value otherwise.
func (o *InlineResponse200292) GetObjectIds() []int32 {
	if o == nil || isNil(o.ObjectIds) {
		var ret []int32
		return ret
	}
	return o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetObjectIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.ObjectIds) {
    return nil, false
	}
	return o.ObjectIds, true
}

// HasObjectIds returns a boolean if a field has been set.
func (o *InlineResponse200292) HasObjectIds() bool {
	if o != nil && !isNil(o.ObjectIds) {
		return true
	}

	return false
}

// SetObjectIds gets a reference to the given []int32 and assigns it to the ObjectIds field.
func (o *InlineResponse200292) SetObjectIds(v []int32) {
	o.ObjectIds = v
}

// GetNetworkIds returns the NetworkIds field value if set, zero value otherwise.
func (o *InlineResponse200292) GetNetworkIds() []string {
	if o == nil || isNil(o.NetworkIds) {
		var ret []string
		return ret
	}
	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292) GetNetworkIdsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkIds) {
    return nil, false
	}
	return o.NetworkIds, true
}

// HasNetworkIds returns a boolean if a field has been set.
func (o *InlineResponse200292) HasNetworkIds() bool {
	if o != nil && !isNil(o.NetworkIds) {
		return true
	}

	return false
}

// SetNetworkIds gets a reference to the given []string and assigns it to the NetworkIds field.
func (o *InlineResponse200292) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

func (o InlineResponse200292) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.ObjectIds) {
		toSerialize["objectIds"] = o.ObjectIds
	}
	if !isNil(o.NetworkIds) {
		toSerialize["networkIds"] = o.NetworkIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200292 struct {
	value *InlineResponse200292
	isSet bool
}

func (v NullableInlineResponse200292) Get() *InlineResponse200292 {
	return v.value
}

func (v *NullableInlineResponse200292) Set(val *InlineResponse200292) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200292) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200292) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200292(val *InlineResponse200292) *NullableInlineResponse200292 {
	return &NullableInlineResponse200292{value: val, isSet: true}
}

func (v NullableInlineResponse200292) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200292) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


