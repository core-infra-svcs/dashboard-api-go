/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200322 struct for InlineResponse200322
type InlineResponse200322 struct {
	// The start time of the channel utilization interval.
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the channel utilization interval.
	EndTs *time.Time `json:"endTs,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`
}

// NewInlineResponse200322 instantiates a new InlineResponse200322 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200322() *InlineResponse200322 {
	this := InlineResponse200322{}
	return &this
}

// NewInlineResponse200322WithDefaults instantiates a new InlineResponse200322 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200322WithDefaults() *InlineResponse200322 {
	this := InlineResponse200322{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200322) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200322) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200322) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200322) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200322) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200322) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200322) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200322) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork and assigns it to the Network field.
func (o *InlineResponse200322) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *InlineResponse200322) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	if o == nil || isNil(o.ByBand) {
		var ret []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322) GetByBandOk() ([]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool) {
	if o == nil || isNil(o.ByBand) {
    return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *InlineResponse200322) HasByBand() bool {
	if o != nil && !isNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand and assigns it to the ByBand field.
func (o *InlineResponse200322) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) {
	o.ByBand = v
}

func (o InlineResponse200322) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200322 struct {
	value *InlineResponse200322
	isSet bool
}

func (v NullableInlineResponse200322) Get() *InlineResponse200322 {
	return v.value
}

func (v *NullableInlineResponse200322) Set(val *InlineResponse200322) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200322) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200322) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200322(val *InlineResponse200322) *NullableInlineResponse200322 {
	return &NullableInlineResponse200322{value: val, isSet: true}
}

func (v NullableInlineResponse200322) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200322) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


