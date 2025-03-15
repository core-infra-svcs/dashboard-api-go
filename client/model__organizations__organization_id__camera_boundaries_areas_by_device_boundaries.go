/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries Configured area boundaries of the camera
type OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries struct {
	// The area boundary id
	Id *string `json:"id,omitempty"`
	// The area boundary type
	Type *string `json:"type,omitempty"`
	// The area boundary name
	Name *string `json:"name,omitempty"`
	// The area boundary vertices
	Vertices []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices `json:"vertices,omitempty"`
}

// NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries instantiates a new OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries {
	this := OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries{}
	return &this
}

// NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesWithDefaults instantiates a new OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesWithDefaults() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries {
	this := OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) SetName(v string) {
	o.Name = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetVertices() []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	if o == nil || isNil(o.Vertices) {
		var ret []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) GetVerticesOk() ([]OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices, bool) {
	if o == nil || isNil(o.Vertices) {
    return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) HasVertices() bool {
	if o != nil && !isNil(o.Vertices) {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices and assigns it to the Vertices field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) SetVertices(v []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) {
	o.Vertices = v
}

func (o OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Vertices) {
		toSerialize["vertices"] = o.Vertices
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries struct {
	value *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) Get() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) Set(val *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries(val *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries {
	return &NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


