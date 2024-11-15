/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget Target Profile attributes
type OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget struct {
	// ID of the target account; can be the account currently tied to the eSIM
	AccountId string `json:"accountId"`
	// Name of the target communication plan
	CommunicationPlan string `json:"communicationPlan"`
	// Name of the target rate plan
	RatePlan string `json:"ratePlan"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget(accountId string, communicationPlan string, ratePlan string) *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget {
	this := OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget{}
	this.AccountId = accountId
	this.CommunicationPlan = communicationPlan
	this.RatePlan = ratePlan
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsSwapTargetWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsSwapTargetWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget {
	this := OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetAccountIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) SetAccountId(v string) {
	o.AccountId = v
}

// GetCommunicationPlan returns the CommunicationPlan field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetCommunicationPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommunicationPlan
}

// GetCommunicationPlanOk returns a tuple with the CommunicationPlan field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetCommunicationPlanOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CommunicationPlan, true
}

// SetCommunicationPlan sets field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) SetCommunicationPlan(v string) {
	o.CommunicationPlan = v
}

// GetRatePlan returns the RatePlan field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetRatePlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RatePlan
}

// GetRatePlanOk returns a tuple with the RatePlan field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) GetRatePlanOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RatePlan, true
}

// SetRatePlan sets field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) SetRatePlan(v string) {
	o.RatePlan = v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["communicationPlan"] = o.CommunicationPlan
	}
	if true {
		toSerialize["ratePlan"] = o.RatePlan
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) Get() *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget(val *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

