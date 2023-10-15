/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject182 struct for InlineObject182
type InlineObject182 struct {
	// The name of the organization
	Name *string `json:"name,omitempty"`
	Management *OrganizationsManagement `json:"management,omitempty"`
	Api *OrganizationsOrganizationIdApi `json:"api,omitempty"`
}

// NewInlineObject182 instantiates a new InlineObject182 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject182() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// NewInlineObject182WithDefaults instantiates a new InlineObject182 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject182WithDefaults() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject182) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject182) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject182) SetName(v string) {
	o.Name = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *InlineObject182) GetManagement() OrganizationsManagement {
	if o == nil || isNil(o.Management) {
		var ret OrganizationsManagement
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetManagementOk() (*OrganizationsManagement, bool) {
	if o == nil || isNil(o.Management) {
    return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *InlineObject182) HasManagement() bool {
	if o != nil && !isNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given OrganizationsManagement and assigns it to the Management field.
func (o *InlineObject182) SetManagement(v OrganizationsManagement) {
	o.Management = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *InlineObject182) GetApi() OrganizationsOrganizationIdApi {
	if o == nil || isNil(o.Api) {
		var ret OrganizationsOrganizationIdApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetApiOk() (*OrganizationsOrganizationIdApi, bool) {
	if o == nil || isNil(o.Api) {
    return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *InlineObject182) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given OrganizationsOrganizationIdApi and assigns it to the Api field.
func (o *InlineObject182) SetApi(v OrganizationsOrganizationIdApi) {
	o.Api = &v
}

func (o InlineObject182) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject182 struct {
	value *InlineObject182
	isSet bool
}

func (v NullableInlineObject182) Get() *InlineObject182 {
	return v.value
}

func (v *NullableInlineObject182) Set(val *InlineObject182) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject182) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject182) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject182(val *InlineObject182) *NullableInlineObject182 {
	return &NullableInlineObject182{value: val, isSet: true}
}

func (v NullableInlineObject182) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject182) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


