/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesTrafficInKbps A breakdown of the average speed of data that has passed through this port during the timespan.
type DevicesSerialSwitchPortsStatusesTrafficInKbps struct {
	// The average speed of the data sent and received (in kilobits-per-second).
	Total *float32 `json:"total,omitempty"`
	// The average speed of the data sent (in kilobits-per-second).
	Sent *float32 `json:"sent,omitempty"`
	// The average speed of the data received (in kilobits-per-second).
	Recv *float32 `json:"recv,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesTrafficInKbps instantiates a new DevicesSerialSwitchPortsStatusesTrafficInKbps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesTrafficInKbps() *DevicesSerialSwitchPortsStatusesTrafficInKbps {
	this := DevicesSerialSwitchPortsStatusesTrafficInKbps{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesTrafficInKbpsWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesTrafficInKbps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesTrafficInKbpsWithDefaults() *DevicesSerialSwitchPortsStatusesTrafficInKbps {
	this := DevicesSerialSwitchPortsStatusesTrafficInKbps{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) SetTotal(v float32) {
	o.Total = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) SetSent(v float32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetRecv() float32 {
	if o == nil || isNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) GetRecvOk() (*float32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *DevicesSerialSwitchPortsStatusesTrafficInKbps) SetRecv(v float32) {
	o.Recv = &v
}

func (o DevicesSerialSwitchPortsStatusesTrafficInKbps) MarshalJSON() ([]byte, error) {
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

type NullableDevicesSerialSwitchPortsStatusesTrafficInKbps struct {
	value *DevicesSerialSwitchPortsStatusesTrafficInKbps
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) Get() *DevicesSerialSwitchPortsStatusesTrafficInKbps {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) Set(val *DevicesSerialSwitchPortsStatusesTrafficInKbps) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesTrafficInKbps(val *DevicesSerialSwitchPortsStatusesTrafficInKbps) *NullableDevicesSerialSwitchPortsStatusesTrafficInKbps {
	return &NullableDevicesSerialSwitchPortsStatusesTrafficInKbps{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesTrafficInKbps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


