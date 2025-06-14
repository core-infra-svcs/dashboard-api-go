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

// OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems struct for OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems
type OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems struct {
	// The Id of the Network
	NetworkId string `json:"networkId"`
	// Array of Sentry Group Policies for the Network
	Policies []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies `json:"policies,omitempty"`
}

// NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems(networkId string) *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems {
	this := OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems{}
	this.NetworkId = networkId
	return &this
}

// NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItemsWithDefaults instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItemsWithDefaults() *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems {
	this := OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) GetNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) GetPolicies() []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies {
	if o == nil || isNil(o.Policies) {
		var ret []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) GetPoliciesOk() ([]OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies, bool) {
	if o == nil || isNil(o.Policies) {
    return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies and assigns it to the Policies field.
func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) SetPolicies(v []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) {
	o.Policies = v
}

func (o OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems struct {
	value *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) Get() *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) Set(val *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems(val *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems {
	return &NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


