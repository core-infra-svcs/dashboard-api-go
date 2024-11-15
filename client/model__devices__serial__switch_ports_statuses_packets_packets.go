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

// DevicesSerialSwitchPortsStatusesPacketsPackets struct for DevicesSerialSwitchPortsStatusesPacketsPackets
type DevicesSerialSwitchPortsStatusesPacketsPackets struct {
	// The type of packets being counted.
	Desc *string `json:"desc,omitempty"`
	// The total count of sent and received packets.
	Total *int32 `json:"total,omitempty"`
	// The total count of packets sent by the switch during the timespan.
	Sent *int32 `json:"sent,omitempty"`
	// The total count of packets received by the switch during the timespan.
	Recv *int32 `json:"recv,omitempty"`
	RatePerSec *DevicesSerialSwitchPortsStatusesPacketsRatePerSec `json:"ratePerSec,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesPacketsPackets instantiates a new DevicesSerialSwitchPortsStatusesPacketsPackets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesPacketsPackets() *DevicesSerialSwitchPortsStatusesPacketsPackets {
	this := DevicesSerialSwitchPortsStatusesPacketsPackets{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesPacketsPacketsWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesPacketsPackets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesPacketsPacketsWithDefaults() *DevicesSerialSwitchPortsStatusesPacketsPackets {
	this := DevicesSerialSwitchPortsStatusesPacketsPackets{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetDesc() string {
	if o == nil || isNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetDescOk() (*string, bool) {
	if o == nil || isNil(o.Desc) {
    return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) HasDesc() bool {
	if o != nil && !isNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) SetDesc(v string) {
	o.Desc = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) SetTotal(v int32) {
	o.Total = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) SetSent(v int32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetRecv() int32 {
	if o == nil || isNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetRecvOk() (*int32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) SetRecv(v int32) {
	o.Recv = &v
}

// GetRatePerSec returns the RatePerSec field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetRatePerSec() DevicesSerialSwitchPortsStatusesPacketsRatePerSec {
	if o == nil || isNil(o.RatePerSec) {
		var ret DevicesSerialSwitchPortsStatusesPacketsRatePerSec
		return ret
	}
	return *o.RatePerSec
}

// GetRatePerSecOk returns a tuple with the RatePerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) GetRatePerSecOk() (*DevicesSerialSwitchPortsStatusesPacketsRatePerSec, bool) {
	if o == nil || isNil(o.RatePerSec) {
    return nil, false
	}
	return o.RatePerSec, true
}

// HasRatePerSec returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) HasRatePerSec() bool {
	if o != nil && !isNil(o.RatePerSec) {
		return true
	}

	return false
}

// SetRatePerSec gets a reference to the given DevicesSerialSwitchPortsStatusesPacketsRatePerSec and assigns it to the RatePerSec field.
func (o *DevicesSerialSwitchPortsStatusesPacketsPackets) SetRatePerSec(v DevicesSerialSwitchPortsStatusesPacketsRatePerSec) {
	o.RatePerSec = &v
}

func (o DevicesSerialSwitchPortsStatusesPacketsPackets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !isNil(o.RatePerSec) {
		toSerialize["ratePerSec"] = o.RatePerSec
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesPacketsPackets struct {
	value *DevicesSerialSwitchPortsStatusesPacketsPackets
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesPacketsPackets) Get() *DevicesSerialSwitchPortsStatusesPacketsPackets {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesPacketsPackets) Set(val *DevicesSerialSwitchPortsStatusesPacketsPackets) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesPacketsPackets) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesPacketsPackets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesPacketsPackets(val *DevicesSerialSwitchPortsStatusesPacketsPackets) *NullableDevicesSerialSwitchPortsStatusesPacketsPackets {
	return &NullableDevicesSerialSwitchPortsStatusesPacketsPackets{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesPacketsPackets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesPacketsPackets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


