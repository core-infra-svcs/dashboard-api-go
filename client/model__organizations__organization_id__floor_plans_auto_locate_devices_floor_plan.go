/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan The assigned floor plan for this device
type OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan struct {
	// Floor plan ID
	Id *string `json:"id,omitempty"`
	// Floor plan name
	Status *string `json:"status,omitempty"`
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan() *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan{}
	return &this
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlanWithDefaults instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlanWithDefaults() *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) SetStatus(v string) {
	o.Status = &v
}

func (o OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan struct {
	value *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) Get() *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) Set(val *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan(val *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) *NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan {
	return &NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


