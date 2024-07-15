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

// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	// list of all reserved IP ranges for a single MG
	ReservedIpRanges []DevicesSerialCellularGatewayLanReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// list of all fixed IP assignments for a single MG
	FixedIpAssignments []DevicesSerialCellularGatewayLanFixedIpAssignments `json:"fixedIpAssignments,omitempty"`
}

// NewInlineObject14 instantiates a new InlineObject14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// NewInlineObject14WithDefaults instantiates a new InlineObject14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14WithDefaults() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject14) GetReservedIpRanges() []DevicesSerialCellularGatewayLanReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []DevicesSerialCellularGatewayLanReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetReservedIpRangesOk() ([]DevicesSerialCellularGatewayLanReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject14) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []DevicesSerialCellularGatewayLanReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject14) SetReservedIpRanges(v []DevicesSerialCellularGatewayLanReservedIpRanges) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject14) GetFixedIpAssignments() []DevicesSerialCellularGatewayLanFixedIpAssignments {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret []DevicesSerialCellularGatewayLanFixedIpAssignments
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetFixedIpAssignmentsOk() ([]DevicesSerialCellularGatewayLanFixedIpAssignments, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject14) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []DevicesSerialCellularGatewayLanFixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineObject14) SetFixedIpAssignments(v []DevicesSerialCellularGatewayLanFixedIpAssignments) {
	o.FixedIpAssignments = v
}

func (o InlineObject14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject14 struct {
	value *InlineObject14
	isSet bool
}

func (v NullableInlineObject14) Get() *InlineObject14 {
	return v.value
}

func (v *NullableInlineObject14) Set(val *InlineObject14) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14(val *InlineObject14) *NullableInlineObject14 {
	return &NullableInlineObject14{value: val, isSet: true}
}

func (v NullableInlineObject14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


