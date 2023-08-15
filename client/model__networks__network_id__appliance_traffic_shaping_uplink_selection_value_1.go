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

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 Value object of this traffic filter
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 struct {
	// ID of this applicationCategory or application type traffic filter. E.g.: \"meraki:layer7/category/1\", \"meraki:layer7/application/4\"
	Id *string `json:"id,omitempty"`
	// Protocol of this custom type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source `json:"source,omitempty"`
	Destination *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination `json:"destination,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1WithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1WithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetId(v string) {
	o.Id = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetSource() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source {
	if o == nil || isNil(o.Source) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetSourceOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source and assigns it to the Source field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetSource(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetDestination() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination {
	if o == nil || isNil(o.Destination) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetDestinationOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination and assigns it to the Destination field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetDestination(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) {
	o.Destination = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


