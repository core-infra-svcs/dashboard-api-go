/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue Value object of this traffic filter
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue struct {
	// Protocol of this custom type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource `json:"source"`
	Destination NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination `json:"destination"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue(source NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource, destination NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetSource() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource {
	if o == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetSourceOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) SetSource(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetDestination() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination {
	if o == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) GetDestinationOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) SetDestination(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) {
	o.Destination = v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


