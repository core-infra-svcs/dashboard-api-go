/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVlansReservedIpRanges struct for NetworksNetworkIdApplianceVlansReservedIpRanges
type NetworksNetworkIdApplianceVlansReservedIpRanges struct {
	// The first IP in the reserved range
	Start *string `json:"start,omitempty"`
	// The last IP in the reserved range
	End *string `json:"end,omitempty"`
	// A text comment for the reserved range
	Comment *string `json:"comment,omitempty"`
}

// NewNetworksNetworkIdApplianceVlansReservedIpRanges instantiates a new NetworksNetworkIdApplianceVlansReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVlansReservedIpRanges() *NetworksNetworkIdApplianceVlansReservedIpRanges {
	this := NetworksNetworkIdApplianceVlansReservedIpRanges{}
	return &this
}

// NewNetworksNetworkIdApplianceVlansReservedIpRangesWithDefaults instantiates a new NetworksNetworkIdApplianceVlansReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVlansReservedIpRangesWithDefaults() *NetworksNetworkIdApplianceVlansReservedIpRanges {
	this := NetworksNetworkIdApplianceVlansReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdApplianceVlansReservedIpRanges) SetComment(v string) {
	o.Comment = &v
}

func (o NetworksNetworkIdApplianceVlansReservedIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVlansReservedIpRanges struct {
	value *NetworksNetworkIdApplianceVlansReservedIpRanges
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVlansReservedIpRanges) Get() *NetworksNetworkIdApplianceVlansReservedIpRanges {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVlansReservedIpRanges) Set(val *NetworksNetworkIdApplianceVlansReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVlansReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVlansReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVlansReservedIpRanges(val *NetworksNetworkIdApplianceVlansReservedIpRanges) *NullableNetworksNetworkIdApplianceVlansReservedIpRanges {
	return &NullableNetworksNetworkIdApplianceVlansReservedIpRanges{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVlansReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVlansReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


