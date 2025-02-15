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

// OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork 
type OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork struct {
	// Name of network
	Name *string `json:"name,omitempty"`
	// ID of network
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork instantiates a new OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork() *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	this := OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopClientsByUsageNetworkWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopClientsByUsageNetworkWithDefaults() *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	this := OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork struct {
	value *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) Get() *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) Set(val *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork(val *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	return &NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


