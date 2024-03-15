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

// OrganizationsOrganizationIdDevicesStatusesComponentsPoe PoE info of the power supply
type OrganizationsOrganizationIdDevicesStatusesComponentsPoe struct {
	// Unit of the PoE maximum
	Unit *string `json:"unit,omitempty"`
	// Maximum PoE this power supply can provide when connected to the current switch model
	Maximum *int32 `json:"maximum,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesStatusesComponentsPoe instantiates a new OrganizationsOrganizationIdDevicesStatusesComponentsPoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesStatusesComponentsPoe() *OrganizationsOrganizationIdDevicesStatusesComponentsPoe {
	this := OrganizationsOrganizationIdDevicesStatusesComponentsPoe{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesStatusesComponentsPoeWithDefaults instantiates a new OrganizationsOrganizationIdDevicesStatusesComponentsPoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesStatusesComponentsPoeWithDefaults() *OrganizationsOrganizationIdDevicesStatusesComponentsPoe {
	this := OrganizationsOrganizationIdDevicesStatusesComponentsPoe{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) SetUnit(v string) {
	o.Unit = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) GetMaximum() int32 {
	if o == nil || isNil(o.Maximum) {
		var ret int32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) GetMaximumOk() (*int32, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given int32 and assigns it to the Maximum field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) SetMaximum(v int32) {
	o.Maximum = &v
}

func (o OrganizationsOrganizationIdDevicesStatusesComponentsPoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe struct {
	value *OrganizationsOrganizationIdDevicesStatusesComponentsPoe
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) Get() *OrganizationsOrganizationIdDevicesStatusesComponentsPoe {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) Set(val *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe(val *OrganizationsOrganizationIdDevicesStatusesComponentsPoe) *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe {
	return &NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


