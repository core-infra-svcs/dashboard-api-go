/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications struct for NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications
type NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications instantiates a new NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications(id string) *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications {
	this := NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplicationsWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplicationsWithDefaults() *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications {
	this := NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications struct {
	value *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) Get() *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) Set(val *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications(val *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications {
	return &NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


