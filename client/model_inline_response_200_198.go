/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200198 struct for InlineResponse200198
type InlineResponse200198 struct {
	// The start time of the channel utilization interval.
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the channel utilization interval.
	EndTs *time.Time `json:"endTs,omitempty"`
	// The serial number for the device.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the device.
	Mac *string `json:"mac,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`
}

// NewInlineResponse200198 instantiates a new InlineResponse200198 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200198() *InlineResponse200198 {
	this := InlineResponse200198{}
	return &this
}

// NewInlineResponse200198WithDefaults instantiates a new InlineResponse200198 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200198WithDefaults() *InlineResponse200198 {
	this := InlineResponse200198{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200198) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200198) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200198) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200198) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200198) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200198) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200198) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200198) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200198) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200198) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200198) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200198) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200198) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200198) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork and assigns it to the Network field.
func (o *InlineResponse200198) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *InlineResponse200198) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	if o == nil || isNil(o.ByBand) {
		var ret []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetByBandOk() ([]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool) {
	if o == nil || isNil(o.ByBand) {
    return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *InlineResponse200198) HasByBand() bool {
	if o != nil && !isNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand and assigns it to the ByBand field.
func (o *InlineResponse200198) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) {
	o.ByBand = v
}

func (o InlineResponse200198) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200198 struct {
	value *InlineResponse200198
	isSet bool
}

func (v NullableInlineResponse200198) Get() *InlineResponse200198 {
	return v.value
}

func (v *NullableInlineResponse200198) Set(val *InlineResponse200198) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200198) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200198) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200198(val *InlineResponse200198) *NullableInlineResponse200198 {
	return &NullableInlineResponse200198{value: val, isSet: true}
}

func (v NullableInlineResponse200198) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200198) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


