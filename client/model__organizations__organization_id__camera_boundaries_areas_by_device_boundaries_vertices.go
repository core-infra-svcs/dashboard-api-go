/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices struct for OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices
type OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices struct {
	// The vertex x coordinate
	X *float32 `json:"x,omitempty"`
	// The vertex y coordinate
	Y *float32 `json:"y,omitempty"`
}

// NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices instantiates a new OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	this := OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices{}
	return &this
}

// NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVerticesWithDefaults instantiates a new OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVerticesWithDefaults() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	this := OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) GetX() float32 {
	if o == nil || isNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) GetXOk() (*float32, bool) {
	if o == nil || isNil(o.X) {
    return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) GetY() float32 {
	if o == nil || isNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) GetYOk() (*float32, bool) {
	if o == nil || isNil(o.Y) {
    return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) SetY(v float32) {
	o.Y = &v
}

func (o OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices struct {
	value *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) Get() *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) Set(val *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices(val *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices {
	return &NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundariesVertices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


