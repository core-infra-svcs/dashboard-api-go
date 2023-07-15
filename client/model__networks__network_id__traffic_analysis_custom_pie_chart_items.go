/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdTrafficAnalysisCustomPieChartItems struct for NetworksNetworkIdTrafficAnalysisCustomPieChartItems
type NetworksNetworkIdTrafficAnalysisCustomPieChartItems struct {
	// The name of the custom pie chart item.
	Name string `json:"name"`
	//     The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'. 
	Type string `json:"type"`
	//     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details). 
	Value string `json:"value"`
}

// NewNetworksNetworkIdTrafficAnalysisCustomPieChartItems instantiates a new NetworksNetworkIdTrafficAnalysisCustomPieChartItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdTrafficAnalysisCustomPieChartItems(name string, type_ string, value string) *NetworksNetworkIdTrafficAnalysisCustomPieChartItems {
	this := NetworksNetworkIdTrafficAnalysisCustomPieChartItems{}
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewNetworksNetworkIdTrafficAnalysisCustomPieChartItemsWithDefaults instantiates a new NetworksNetworkIdTrafficAnalysisCustomPieChartItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdTrafficAnalysisCustomPieChartItemsWithDefaults() *NetworksNetworkIdTrafficAnalysisCustomPieChartItems {
	this := NetworksNetworkIdTrafficAnalysisCustomPieChartItems{}
	return &this
}

// GetName returns the Name field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) SetValue(v string) {
	o.Value = v
}

func (o NetworksNetworkIdTrafficAnalysisCustomPieChartItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems struct {
	value *NetworksNetworkIdTrafficAnalysisCustomPieChartItems
	isSet bool
}

func (v NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) Get() *NetworksNetworkIdTrafficAnalysisCustomPieChartItems {
	return v.value
}

func (v *NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) Set(val *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems(val *NetworksNetworkIdTrafficAnalysisCustomPieChartItems) *NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems {
	return &NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdTrafficAnalysisCustomPieChartItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


