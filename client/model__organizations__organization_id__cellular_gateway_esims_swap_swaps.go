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

// OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps struct for OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps
type OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps struct {
	// eSIM EID
	Eid string `json:"eid"`
	Target *OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget `json:"target,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps(eid string) *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps {
	this := OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{}
	this.Eid = eid
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwapsWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwapsWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps {
	this := OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{}
	return &this
}

// GetEid returns the Eid field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) GetEid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Eid
}

// GetEidOk returns a tuple with the Eid field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) GetEidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Eid, true
}

// SetEid sets field value
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) SetEid(v string) {
	o.Eid = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) GetTarget() OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget {
	if o == nil || isNil(o.Target) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) GetTargetOk() (*OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget, bool) {
	if o == nil || isNil(o.Target) {
    return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget and assigns it to the Target field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) SetTarget(v OrganizationsOrganizationIdCellularGatewayEsimsSwapTarget) {
	o.Target = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eid"] = o.Eid
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) Get() *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps(val *OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

