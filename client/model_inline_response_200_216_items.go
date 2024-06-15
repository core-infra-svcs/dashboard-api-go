/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200216Items struct for InlineResponse200216Items
type InlineResponse200216Items struct {
	// id
	NetworkId string `json:"networkId"`
	// Name
	NetworkName string `json:"networkName"`
	// Total Alerts
	AlertCount int32 `json:"alertCount"`
	// Alerts By Severity
	SeverityCounts []InlineResponse200216SeverityCounts `json:"severityCounts"`
}

// NewInlineResponse200216Items instantiates a new InlineResponse200216Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216Items(networkId string, networkName string, alertCount int32, severityCounts []InlineResponse200216SeverityCounts) *InlineResponse200216Items {
	this := InlineResponse200216Items{}
	this.NetworkId = networkId
	this.NetworkName = networkName
	this.AlertCount = alertCount
	this.SeverityCounts = severityCounts
	return &this
}

// NewInlineResponse200216ItemsWithDefaults instantiates a new InlineResponse200216Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216ItemsWithDefaults() *InlineResponse200216Items {
	this := InlineResponse200216Items{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *InlineResponse200216Items) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Items) GetNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *InlineResponse200216Items) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetNetworkName returns the NetworkName field value
func (o *InlineResponse200216Items) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Items) GetNetworkNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *InlineResponse200216Items) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetAlertCount returns the AlertCount field value
func (o *InlineResponse200216Items) GetAlertCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertCount
}

// GetAlertCountOk returns a tuple with the AlertCount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Items) GetAlertCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlertCount, true
}

// SetAlertCount sets field value
func (o *InlineResponse200216Items) SetAlertCount(v int32) {
	o.AlertCount = v
}

// GetSeverityCounts returns the SeverityCounts field value
func (o *InlineResponse200216Items) GetSeverityCounts() []InlineResponse200216SeverityCounts {
	if o == nil {
		var ret []InlineResponse200216SeverityCounts
		return ret
	}

	return o.SeverityCounts
}

// GetSeverityCountsOk returns a tuple with the SeverityCounts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Items) GetSeverityCountsOk() ([]InlineResponse200216SeverityCounts, bool) {
	if o == nil {
    return nil, false
	}
	return o.SeverityCounts, true
}

// SetSeverityCounts sets field value
func (o *InlineResponse200216Items) SetSeverityCounts(v []InlineResponse200216SeverityCounts) {
	o.SeverityCounts = v
}

func (o InlineResponse200216Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if true {
		toSerialize["networkName"] = o.NetworkName
	}
	if true {
		toSerialize["alertCount"] = o.AlertCount
	}
	if true {
		toSerialize["severityCounts"] = o.SeverityCounts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200216Items struct {
	value *InlineResponse200216Items
	isSet bool
}

func (v NullableInlineResponse200216Items) Get() *InlineResponse200216Items {
	return v.value
}

func (v *NullableInlineResponse200216Items) Set(val *InlineResponse200216Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200216Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200216Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200216Items(val *InlineResponse200216Items) *NullableInlineResponse200216Items {
	return &NullableInlineResponse200216Items{value: val, isSet: true}
}

func (v NullableInlineResponse200216Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200216Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


