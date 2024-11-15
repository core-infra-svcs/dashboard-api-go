/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges struct for NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges
type NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges struct {
	// The first IP in the reserved range
	Start string `json:"start"`
	// The last IP in the reserved range
	End string `json:"end"`
	// A text comment for the reserved range
	Comment string `json:"comment"`
}

// NewNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges instantiates a new NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges(start string, end string, comment string) *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges {
	this := NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges{}
	this.Start = start
	this.End = end
	this.Comment = comment
	return &this
}

// NewNetworksNetworkIdApplianceVlansVlanIdReservedIpRangesWithDefaults instantiates a new NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVlansVlanIdReservedIpRangesWithDefaults() *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges {
	this := NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) SetComment(v string) {
	o.Comment = v
}

func (o NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges struct {
	value *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) Get() *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) Set(val *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges(val *NetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) *NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges {
	return &NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVlansVlanIdReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


