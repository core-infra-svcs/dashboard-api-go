/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026 struct for InlineResponse20026
type InlineResponse20026 struct {
	// ID of the network whose VPN exclusion rules are returned.
	NetworkId string `json:"networkId"`
	// Name of the network whose VPN exclusion rules are returned.
	NetworkName string `json:"networkName"`
	// Custom VPN exclusion rules.
	Custom []InlineResponse20026Custom `json:"custom"`
	// Major Application based VPN exclusion rules.
	MajorApplications []InlineResponse20026MajorApplications `json:"majorApplications"`
}

// NewInlineResponse20026 instantiates a new InlineResponse20026 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026(networkId string, networkName string, custom []InlineResponse20026Custom, majorApplications []InlineResponse20026MajorApplications) *InlineResponse20026 {
	this := InlineResponse20026{}
	this.NetworkId = networkId
	this.NetworkName = networkName
	this.Custom = custom
	this.MajorApplications = majorApplications
	return &this
}

// NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026WithDefaults() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *InlineResponse20026) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetNetworkIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *InlineResponse20026) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetNetworkName returns the NetworkName field value
func (o *InlineResponse20026) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetNetworkNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *InlineResponse20026) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetCustom returns the Custom field value
func (o *InlineResponse20026) GetCustom() []InlineResponse20026Custom {
	if o == nil {
		var ret []InlineResponse20026Custom
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetCustomOk() ([]InlineResponse20026Custom, bool) {
	if o == nil {
    return nil, false
	}
	return o.Custom, true
}

// SetCustom sets field value
func (o *InlineResponse20026) SetCustom(v []InlineResponse20026Custom) {
	o.Custom = v
}

// GetMajorApplications returns the MajorApplications field value
func (o *InlineResponse20026) GetMajorApplications() []InlineResponse20026MajorApplications {
	if o == nil {
		var ret []InlineResponse20026MajorApplications
		return ret
	}

	return o.MajorApplications
}

// GetMajorApplicationsOk returns a tuple with the MajorApplications field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetMajorApplicationsOk() ([]InlineResponse20026MajorApplications, bool) {
	if o == nil {
    return nil, false
	}
	return o.MajorApplications, true
}

// SetMajorApplications sets field value
func (o *InlineResponse20026) SetMajorApplications(v []InlineResponse20026MajorApplications) {
	o.MajorApplications = v
}

func (o InlineResponse20026) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20026 struct {
	value *InlineResponse20026
	isSet bool
}

func (v NullableInlineResponse20026) Get() *InlineResponse20026 {
	return v.value
}

func (v *NullableInlineResponse20026) Set(val *InlineResponse20026) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026(val *InlineResponse20026) *NullableInlineResponse20026 {
	return &NullableInlineResponse20026{value: val, isSet: true}
}

func (v NullableInlineResponse20026) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


