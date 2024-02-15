/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream Packets sent from an AP to a client.
type OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream struct {
	// Total packets received by a client.
	Total *int32 `json:"total,omitempty"`
	// Total packets sent by an AP that did not reach the client.
	Lost *int32 `json:"lost,omitempty"`
	// Percentage of lost packets.
	LossPercentage *float32 `json:"lossPercentage,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstreamWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstreamWithDefaults() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) SetTotal(v int32) {
	o.Total = &v
}

// GetLost returns the Lost field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetLost() int32 {
	if o == nil || isNil(o.Lost) {
		var ret int32
		return ret
	}
	return *o.Lost
}

// GetLostOk returns a tuple with the Lost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetLostOk() (*int32, bool) {
	if o == nil || isNil(o.Lost) {
    return nil, false
	}
	return o.Lost, true
}

// HasLost returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) HasLost() bool {
	if o != nil && !isNil(o.Lost) {
		return true
	}

	return false
}

// SetLost gets a reference to the given int32 and assigns it to the Lost field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) SetLost(v int32) {
	o.Lost = &v
}

// GetLossPercentage returns the LossPercentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetLossPercentage() float32 {
	if o == nil || isNil(o.LossPercentage) {
		var ret float32
		return ret
	}
	return *o.LossPercentage
}

// GetLossPercentageOk returns a tuple with the LossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) GetLossPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.LossPercentage) {
    return nil, false
	}
	return o.LossPercentage, true
}

// HasLossPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) HasLossPercentage() bool {
	if o != nil && !isNil(o.LossPercentage) {
		return true
	}

	return false
}

// SetLossPercentage gets a reference to the given float32 and assigns it to the LossPercentage field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) SetLossPercentage(v float32) {
	o.LossPercentage = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Lost) {
		toSerialize["lost"] = o.Lost
	}
	if !isNil(o.LossPercentage) {
		toSerialize["lossPercentage"] = o.LossPercentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream struct {
	value *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) Get() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) Set(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	return &NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


