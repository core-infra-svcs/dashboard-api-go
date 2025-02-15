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

// OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems struct for OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems
type OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems struct {
	Network *InlineResponse200219Network `json:"network,omitempty"`
	Profile *InlineResponse200219Profile `json:"profile,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems instantiates a new OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems() *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems {
	this := OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItemsWithDefaults instantiates a new OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItemsWithDefaults() *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems {
	this := OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) GetNetwork() InlineResponse200219Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200219Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) GetNetworkOk() (*InlineResponse200219Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200219Network and assigns it to the Network field.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) SetNetwork(v InlineResponse200219Network) {
	o.Network = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) GetProfile() InlineResponse200219Profile {
	if o == nil || isNil(o.Profile) {
		var ret InlineResponse200219Profile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) GetProfileOk() (*InlineResponse200219Profile, bool) {
	if o == nil || isNil(o.Profile) {
    return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) HasProfile() bool {
	if o != nil && !isNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given InlineResponse200219Profile and assigns it to the Profile field.
func (o *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) SetProfile(v InlineResponse200219Profile) {
	o.Profile = &v
}

func (o OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems struct {
	value *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) Get() *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) Set(val *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems(val *OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) *NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems {
	return &NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


