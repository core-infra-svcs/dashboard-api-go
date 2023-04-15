/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesUsageInKb A breakdown of how many kilobytes have passed through this port during the timespan.
type DevicesSerialSwitchPortsStatusesUsageInKb struct {
	// The total amount of data sent and received (in kilobytes).
	Total *int32 `json:"total,omitempty"`
	// The amount of data sent (in kilobytes).
	Sent *int32 `json:"sent,omitempty"`
	// The amount of data received (in kilobytes).
	Recv *int32 `json:"recv,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesUsageInKb instantiates a new DevicesSerialSwitchPortsStatusesUsageInKb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesUsageInKb() *DevicesSerialSwitchPortsStatusesUsageInKb {
	this := DevicesSerialSwitchPortsStatusesUsageInKb{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesUsageInKbWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesUsageInKb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesUsageInKbWithDefaults() *DevicesSerialSwitchPortsStatusesUsageInKb {
	this := DevicesSerialSwitchPortsStatusesUsageInKb{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) SetTotal(v int32) {
	o.Total = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) SetSent(v int32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetRecv() int32 {
	if o == nil || isNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) GetRecvOk() (*int32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *DevicesSerialSwitchPortsStatusesUsageInKb) SetRecv(v int32) {
	o.Recv = &v
}

func (o DevicesSerialSwitchPortsStatusesUsageInKb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesUsageInKb struct {
	value *DevicesSerialSwitchPortsStatusesUsageInKb
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesUsageInKb) Get() *DevicesSerialSwitchPortsStatusesUsageInKb {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesUsageInKb) Set(val *DevicesSerialSwitchPortsStatusesUsageInKb) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesUsageInKb) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesUsageInKb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesUsageInKb(val *DevicesSerialSwitchPortsStatusesUsageInKb) *NullableDevicesSerialSwitchPortsStatusesUsageInKb {
	return &NullableDevicesSerialSwitchPortsStatusesUsageInKb{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesUsageInKb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesUsageInKb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


