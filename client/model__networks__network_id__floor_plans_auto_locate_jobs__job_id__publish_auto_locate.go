/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate The auto locate position for this device
type NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate struct {
	// Whether or not this device's location should be saved as a user-defined anchor
	IsAnchor *bool `json:"isAnchor,omitempty"`
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate{}
	return &this
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocateWithDefaults instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocateWithDefaults() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate{}
	return &this
}

// GetIsAnchor returns the IsAnchor field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) GetIsAnchor() bool {
	if o == nil || isNil(o.IsAnchor) {
		var ret bool
		return ret
	}
	return *o.IsAnchor
}

// GetIsAnchorOk returns a tuple with the IsAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) GetIsAnchorOk() (*bool, bool) {
	if o == nil || isNil(o.IsAnchor) {
    return nil, false
	}
	return o.IsAnchor, true
}

// HasIsAnchor returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) HasIsAnchor() bool {
	if o != nil && !isNil(o.IsAnchor) {
		return true
	}

	return false
}

// SetIsAnchor gets a reference to the given bool and assigns it to the IsAnchor field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) SetIsAnchor(v bool) {
	o.IsAnchor = &v
}

func (o NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAnchor) {
		toSerialize["isAnchor"] = o.IsAnchor
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate struct {
	value *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) Get() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) Set(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate {
	return &NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


