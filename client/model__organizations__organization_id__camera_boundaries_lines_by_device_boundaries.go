/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries Configured line boundaries of the camera
type OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries struct {
	// The line boundary id
	Id *string `json:"id,omitempty"`
	// The line boundary type
	Type *string `json:"type,omitempty"`
	// The line boundary name
	Name *string `json:"name,omitempty"`
	// The line boundary vertices
	Vertices []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices `json:"vertices,omitempty"`
	DirectionVertex *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex `json:"directionVertex,omitempty"`
}

// NewOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries instantiates a new OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries() *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries {
	this := OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries{}
	return &this
}

// NewOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesWithDefaults instantiates a new OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesWithDefaults() *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries {
	this := OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) SetName(v string) {
	o.Name = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetVertices() []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	if o == nil || isNil(o.Vertices) {
		var ret []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetVerticesOk() ([]OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices, bool) {
	if o == nil || isNil(o.Vertices) {
    return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) HasVertices() bool {
	if o != nil && !isNil(o.Vertices) {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices and assigns it to the Vertices field.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) SetVertices(v []OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) {
	o.Vertices = v
}

// GetDirectionVertex returns the DirectionVertex field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetDirectionVertex() OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex {
	if o == nil || isNil(o.DirectionVertex) {
		var ret OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex
		return ret
	}
	return *o.DirectionVertex
}

// GetDirectionVertexOk returns a tuple with the DirectionVertex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) GetDirectionVertexOk() (*OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex, bool) {
	if o == nil || isNil(o.DirectionVertex) {
    return nil, false
	}
	return o.DirectionVertex, true
}

// HasDirectionVertex returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) HasDirectionVertex() bool {
	if o != nil && !isNil(o.DirectionVertex) {
		return true
	}

	return false
}

// SetDirectionVertex gets a reference to the given OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex and assigns it to the DirectionVertex field.
func (o *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) SetDirectionVertex(v OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundariesDirectionVertex) {
	o.DirectionVertex = &v
}

func (o OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.DirectionVertex) {
		toSerialize["directionVertex"] = o.DirectionVertex
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries struct {
	value *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) Get() *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) Set(val *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries(val *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) *NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries {
	return &NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

