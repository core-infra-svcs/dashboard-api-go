/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSdwanInternetPoliciesValue Value of traffic filter
type NetworksNetworkIdApplianceSdwanInternetPoliciesValue struct {
	// Protocol of the traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource `json:"source"`
	Destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination `json:"destination"`
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesValue instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValue(source NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource, destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) *NetworksNetworkIdApplianceSdwanInternetPoliciesValue {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesValue{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesValue {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesValue{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetSource() NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource {
	if o == nil {
		var ret NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetSourceOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetSource(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetDestination() NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination {
	if o == nil {
		var ret NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetDestinationOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetDestination(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) {
	o.Destination = v
}

func (o NetworksNetworkIdApplianceSdwanInternetPoliciesValue) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue struct {
	value *NetworksNetworkIdApplianceSdwanInternetPoliciesValue
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) Get() *NetworksNetworkIdApplianceSdwanInternetPoliciesValue {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) Set(val *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue(val *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue {
	return &NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


