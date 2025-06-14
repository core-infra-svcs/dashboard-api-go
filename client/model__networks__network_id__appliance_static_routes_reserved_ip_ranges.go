/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceStaticRoutesReservedIpRanges struct for NetworksNetworkIdApplianceStaticRoutesReservedIpRanges
type NetworksNetworkIdApplianceStaticRoutesReservedIpRanges struct {
	// First address in the reserved range
	Start *string `json:"start,omitempty"`
	// Last address in the reserved range
	End *string `json:"end,omitempty"`
	// Description of the range
	Comment *string `json:"comment,omitempty"`
}

// NewNetworksNetworkIdApplianceStaticRoutesReservedIpRanges instantiates a new NetworksNetworkIdApplianceStaticRoutesReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceStaticRoutesReservedIpRanges() *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges {
	this := NetworksNetworkIdApplianceStaticRoutesReservedIpRanges{}
	return &this
}

// NewNetworksNetworkIdApplianceStaticRoutesReservedIpRangesWithDefaults instantiates a new NetworksNetworkIdApplianceStaticRoutesReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceStaticRoutesReservedIpRangesWithDefaults() *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges {
	this := NetworksNetworkIdApplianceStaticRoutesReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) SetComment(v string) {
	o.Comment = &v
}

func (o NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges struct {
	value *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) Get() *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) Set(val *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges(val *NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) *NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges {
	return &NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


