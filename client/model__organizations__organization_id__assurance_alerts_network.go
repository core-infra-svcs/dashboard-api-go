/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAssuranceAlertsNetwork Network details
type OrganizationsOrganizationIdAssuranceAlertsNetwork struct {
	// Name of the network where alert appears
	Name string `json:"name"`
	// ID of the network where alert appears
	Id string `json:"id"`
}

// NewOrganizationsOrganizationIdAssuranceAlertsNetwork instantiates a new OrganizationsOrganizationIdAssuranceAlertsNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAssuranceAlertsNetwork(name string, id string) *OrganizationsOrganizationIdAssuranceAlertsNetwork {
	this := OrganizationsOrganizationIdAssuranceAlertsNetwork{}
	this.Name = name
	this.Id = id
	return &this
}

// NewOrganizationsOrganizationIdAssuranceAlertsNetworkWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAssuranceAlertsNetworkWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsNetwork {
	this := OrganizationsOrganizationIdAssuranceAlertsNetwork{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationsOrganizationIdAssuranceAlertsNetwork) SetId(v string) {
	o.Id = v
}

func (o OrganizationsOrganizationIdAssuranceAlertsNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAssuranceAlertsNetwork struct {
	value *OrganizationsOrganizationIdAssuranceAlertsNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) Get() *OrganizationsOrganizationIdAssuranceAlertsNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) Set(val *OrganizationsOrganizationIdAssuranceAlertsNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAssuranceAlertsNetwork(val *OrganizationsOrganizationIdAssuranceAlertsNetwork) *NullableOrganizationsOrganizationIdAssuranceAlertsNetwork {
	return &NullableOrganizationsOrganizationIdAssuranceAlertsNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


