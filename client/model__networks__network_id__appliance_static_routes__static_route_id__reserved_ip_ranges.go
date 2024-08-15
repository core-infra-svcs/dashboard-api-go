/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges struct for NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges
type NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges struct {
	// First address in the reserved range
	Start string `json:"start"`
	// Last address in the reserved range
	End string `json:"end"`
	// Description of the range
	Comment string `json:"comment"`
}

// NewNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges instantiates a new NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges(start string, end string, comment string) *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	this := NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges{}
	this.Start = start
	this.End = end
	this.Comment = comment
	return &this
}

// NewNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRangesWithDefaults instantiates a new NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRangesWithDefaults() *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	this := NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) SetComment(v string) {
	o.Comment = v
}

func (o NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges struct {
	value *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) Get() *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) Set(val *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges(val *NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) *NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	return &NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


