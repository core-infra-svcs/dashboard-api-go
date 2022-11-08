/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20081 struct for InlineResponse20081
type InlineResponse20081 struct {
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

// NewInlineResponse20081 instantiates a new InlineResponse20081 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20081() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// NewInlineResponse20081WithDefaults instantiates a new InlineResponse20081 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20081WithDefaults() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20081) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20081) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20081) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20081) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20081) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20081) SetSerial(v string) {
	o.Serial = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *InlineResponse20081) GetUplink() string {
	if o == nil || isNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetUplinkOk() (*string, bool) {
	if o == nil || isNil(o.Uplink) {
    return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *InlineResponse20081) HasUplink() bool {
	if o != nil && !isNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *InlineResponse20081) SetUplink(v string) {
	o.Uplink = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20081) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20081) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20081) SetIp(v string) {
	o.Ip = &v
}

// GetTimeSeries returns the TimeSeries field value if set, zero value otherwise.
func (o *InlineResponse20081) GetTimeSeries() []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	if o == nil || isNil(o.TimeSeries) {
		var ret []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries
		return ret
	}
	return o.TimeSeries
}

// GetTimeSeriesOk returns a tuple with the TimeSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetTimeSeriesOk() ([]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries, bool) {
	if o == nil || isNil(o.TimeSeries) {
    return nil, false
	}
	return o.TimeSeries, true
}

// HasTimeSeries returns a boolean if a field has been set.
func (o *InlineResponse20081) HasTimeSeries() bool {
	if o != nil && !isNil(o.TimeSeries) {
		return true
	}

	return false
}

// SetTimeSeries gets a reference to the given []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries and assigns it to the TimeSeries field.
func (o *InlineResponse20081) SetTimeSeries(v []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) {
	o.TimeSeries = v
}

func (o InlineResponse20081) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20081 struct {
	value *InlineResponse20081
	isSet bool
}

func (v NullableInlineResponse20081) Get() *InlineResponse20081 {
	return v.value
}

func (v *NullableInlineResponse20081) Set(val *InlineResponse20081) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20081) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20081) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20081(val *InlineResponse20081) *NullableInlineResponse20081 {
	return &NullableInlineResponse20081{value: val, isSet: true}
}

func (v NullableInlineResponse20081) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20081) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


