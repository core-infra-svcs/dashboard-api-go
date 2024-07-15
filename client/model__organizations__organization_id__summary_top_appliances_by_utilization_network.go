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

// OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork Network info
type OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork struct {
	// Network name
	Name *string `json:"name,omitempty"`
	// Network id
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork instantiates a new OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	this := OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetworkWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetworkWithDefaults() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	this := OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork struct {
	value *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) Get() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) Set(val *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork(val *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	return &NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


