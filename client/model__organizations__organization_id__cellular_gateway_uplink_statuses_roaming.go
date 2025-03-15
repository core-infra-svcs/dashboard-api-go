/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming Roaming Status
type OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming struct {
	// Roaming Status
	Status *string `json:"status,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoamingWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoamingWithDefaults() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) SetStatus(v string) {
	o.Status = &v
}

func (o OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming struct {
	value *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) Get() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) Set(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming {
	return &NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


