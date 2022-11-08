/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInsightApplicationsThresholds Thresholds defined by a user or Meraki models for each application
type OrganizationsOrganizationIdInsightApplicationsThresholds struct {
	// Threshold type (static or smart)
	Type *string `json:"type,omitempty"`
	// Threshold for each network
	ByNetwork []OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork `json:"byNetwork,omitempty"`
}

// NewOrganizationsOrganizationIdInsightApplicationsThresholds instantiates a new OrganizationsOrganizationIdInsightApplicationsThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInsightApplicationsThresholds() *OrganizationsOrganizationIdInsightApplicationsThresholds {
	this := OrganizationsOrganizationIdInsightApplicationsThresholds{}
	return &this
}

// NewOrganizationsOrganizationIdInsightApplicationsThresholdsWithDefaults instantiates a new OrganizationsOrganizationIdInsightApplicationsThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInsightApplicationsThresholdsWithDefaults() *OrganizationsOrganizationIdInsightApplicationsThresholds {
	this := OrganizationsOrganizationIdInsightApplicationsThresholds{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) SetType(v string) {
	o.Type = &v
}

// GetByNetwork returns the ByNetwork field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) GetByNetwork() []OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork {
	if o == nil || isNil(o.ByNetwork) {
		var ret []OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork
		return ret
	}
	return o.ByNetwork
}

// GetByNetworkOk returns a tuple with the ByNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) GetByNetworkOk() ([]OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork, bool) {
	if o == nil || isNil(o.ByNetwork) {
    return nil, false
	}
	return o.ByNetwork, true
}

// HasByNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) HasByNetwork() bool {
	if o != nil && !isNil(o.ByNetwork) {
		return true
	}

	return false
}

// SetByNetwork gets a reference to the given []OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork and assigns it to the ByNetwork field.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholds) SetByNetwork(v []OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) {
	o.ByNetwork = v
}

func (o OrganizationsOrganizationIdInsightApplicationsThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ByNetwork) {
		toSerialize["byNetwork"] = o.ByNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInsightApplicationsThresholds struct {
	value *OrganizationsOrganizationIdInsightApplicationsThresholds
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholds) Get() *OrganizationsOrganizationIdInsightApplicationsThresholds {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholds) Set(val *OrganizationsOrganizationIdInsightApplicationsThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInsightApplicationsThresholds(val *OrganizationsOrganizationIdInsightApplicationsThresholds) *NullableOrganizationsOrganizationIdInsightApplicationsThresholds {
	return &NullableOrganizationsOrganizationIdInsightApplicationsThresholds{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


