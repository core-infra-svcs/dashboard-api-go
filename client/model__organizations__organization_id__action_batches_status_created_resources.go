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

// OrganizationsOrganizationIdActionBatchesStatusCreatedResources struct for OrganizationsOrganizationIdActionBatchesStatusCreatedResources
type OrganizationsOrganizationIdActionBatchesStatusCreatedResources struct {
	// ID of the created resource
	Id *string `json:"id,omitempty"`
	// URI, not including base, of the created resource
	Uri *string `json:"uri,omitempty"`
}

// NewOrganizationsOrganizationIdActionBatchesStatusCreatedResources instantiates a new OrganizationsOrganizationIdActionBatchesStatusCreatedResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesStatusCreatedResources() *OrganizationsOrganizationIdActionBatchesStatusCreatedResources {
	this := OrganizationsOrganizationIdActionBatchesStatusCreatedResources{}
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesStatusCreatedResourcesWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesStatusCreatedResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesStatusCreatedResourcesWithDefaults() *OrganizationsOrganizationIdActionBatchesStatusCreatedResources {
	this := OrganizationsOrganizationIdActionBatchesStatusCreatedResources{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) SetUri(v string) {
	o.Uri = &v
}

func (o OrganizationsOrganizationIdActionBatchesStatusCreatedResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources struct {
	value *OrganizationsOrganizationIdActionBatchesStatusCreatedResources
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) Get() *OrganizationsOrganizationIdActionBatchesStatusCreatedResources {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) Set(val *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources(val *OrganizationsOrganizationIdActionBatchesStatusCreatedResources) *NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources {
	return &NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatusCreatedResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


