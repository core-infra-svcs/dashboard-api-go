/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200270 struct for InlineResponse200270
type InlineResponse200270 struct {
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Serial of MX device
	Serial *string `json:"serial,omitempty"`
	// Uplink interface (wan1, wan2, or cellular)
	Uplink *string `json:"uplink,omitempty"`
	// IP address of uplink
	Ip *string `json:"ip,omitempty"`
	// Loss and latency timeseries data
	TimeSeries []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries `json:"timeSeries,omitempty"`
}

// NewInlineResponse200270 instantiates a new InlineResponse200270 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200270() *InlineResponse200270 {
	this := InlineResponse200270{}
	return &this
}

// NewInlineResponse200270WithDefaults instantiates a new InlineResponse200270 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200270WithDefaults() *InlineResponse200270 {
	this := InlineResponse200270{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200270) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200270) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200270) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200270) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200270) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200270) SetSerial(v string) {
	o.Serial = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *InlineResponse200270) GetUplink() string {
	if o == nil || isNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270) GetUplinkOk() (*string, bool) {
	if o == nil || isNil(o.Uplink) {
    return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *InlineResponse200270) HasUplink() bool {
	if o != nil && !isNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *InlineResponse200270) SetUplink(v string) {
	o.Uplink = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse200270) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse200270) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse200270) SetIp(v string) {
	o.Ip = &v
}

// GetTimeSeries returns the TimeSeries field value if set, zero value otherwise.
func (o *InlineResponse200270) GetTimeSeries() []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	if o == nil || isNil(o.TimeSeries) {
		var ret []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries
		return ret
	}
	return o.TimeSeries
}

// GetTimeSeriesOk returns a tuple with the TimeSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270) GetTimeSeriesOk() ([]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries, bool) {
	if o == nil || isNil(o.TimeSeries) {
    return nil, false
	}
	return o.TimeSeries, true
}

// HasTimeSeries returns a boolean if a field has been set.
func (o *InlineResponse200270) HasTimeSeries() bool {
	if o != nil && !isNil(o.TimeSeries) {
		return true
	}

	return false
}

// SetTimeSeries gets a reference to the given []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries and assigns it to the TimeSeries field.
func (o *InlineResponse200270) SetTimeSeries(v []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) {
	o.TimeSeries = v
}

func (o InlineResponse200270) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Uplink) {
		toSerialize["uplink"] = o.Uplink
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.TimeSeries) {
		toSerialize["timeSeries"] = o.TimeSeries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200270 struct {
	value *InlineResponse200270
	isSet bool
}

func (v NullableInlineResponse200270) Get() *InlineResponse200270 {
	return v.value
}

func (v *NullableInlineResponse200270) Set(val *InlineResponse200270) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200270) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200270) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200270(val *InlineResponse200270) *NullableInlineResponse200270 {
	return &NullableInlineResponse200270{value: val, isSet: true}
}

func (v NullableInlineResponse200270) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200270) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


