/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20066 struct for InlineResponse20066
type InlineResponse20066 struct {
	// ID of the network whose VPN exclusion rules are returned.
	NetworkId string `json:"networkId"`
	// Name of the network whose VPN exclusion rules are returned.
	NetworkName string `json:"networkName"`
	// Custom VPN exclusion rules.
	Custom []InlineResponse20066Custom `json:"custom"`
	// Major Application based VPN exclusion rules.
	MajorApplications []InlineResponse20066MajorApplications `json:"majorApplications"`
}

// NewInlineResponse20066 instantiates a new InlineResponse20066 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066(networkId string, networkName string, custom []InlineResponse20066Custom, majorApplications []InlineResponse20066MajorApplications) *InlineResponse20066 {
	this := InlineResponse20066{}
	this.NetworkId = networkId
	this.NetworkName = networkName
	this.Custom = custom
	this.MajorApplications = majorApplications
	return &this
}

// NewInlineResponse20066WithDefaults instantiates a new InlineResponse20066 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066WithDefaults() *InlineResponse20066 {
	this := InlineResponse20066{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *InlineResponse20066) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066) GetNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *InlineResponse20066) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetNetworkName returns the NetworkName field value
func (o *InlineResponse20066) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066) GetNetworkNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *InlineResponse20066) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetCustom returns the Custom field value
func (o *InlineResponse20066) GetCustom() []InlineResponse20066Custom {
	if o == nil {
		var ret []InlineResponse20066Custom
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066) GetCustomOk() ([]InlineResponse20066Custom, bool) {
	if o == nil {
    return nil, false
	}
	return o.Custom, true
}

// SetCustom sets field value
func (o *InlineResponse20066) SetCustom(v []InlineResponse20066Custom) {
	o.Custom = v
}

// GetMajorApplications returns the MajorApplications field value
func (o *InlineResponse20066) GetMajorApplications() []InlineResponse20066MajorApplications {
	if o == nil {
		var ret []InlineResponse20066MajorApplications
		return ret
	}

	return o.MajorApplications
}

// GetMajorApplicationsOk returns a tuple with the MajorApplications field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066) GetMajorApplicationsOk() ([]InlineResponse20066MajorApplications, bool) {
	if o == nil {
    return nil, false
	}
	return o.MajorApplications, true
}

// SetMajorApplications sets field value
func (o *InlineResponse20066) SetMajorApplications(v []InlineResponse20066MajorApplications) {
	o.MajorApplications = v
}

func (o InlineResponse20066) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if true {
		toSerialize["networkName"] = o.NetworkName
	}
	if true {
		toSerialize["custom"] = o.Custom
	}
	if true {
		toSerialize["majorApplications"] = o.MajorApplications
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066 struct {
	value *InlineResponse20066
	isSet bool
}

func (v NullableInlineResponse20066) Get() *InlineResponse20066 {
	return v.value
}

func (v *NullableInlineResponse20066) Set(val *InlineResponse20066) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066(val *InlineResponse20066) *NullableInlineResponse20066 {
	return &NullableInlineResponse20066{value: val, isSet: true}
}

func (v NullableInlineResponse20066) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


