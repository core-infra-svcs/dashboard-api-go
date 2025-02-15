/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails struct for OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails
type OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails struct {
	// Name of device detail
	Name string `json:"name"`
	// Value of device detail
	Value *string `json:"value,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails instantiates a new OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails(name string) *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails {
	this := OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails{}
	this.Name = name
	return &this
}

// NewOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetailsWithDefaults instantiates a new OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetailsWithDefaults() *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails {
	this := OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) SetValue(v string) {
	o.Value = &v
}

func (o OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails struct {
	value *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) Get() *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) Set(val *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails(val *OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) *NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails {
	return &NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


