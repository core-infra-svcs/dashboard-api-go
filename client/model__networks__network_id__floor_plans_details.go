/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFloorPlansDetails struct for NetworksNetworkIdFloorPlansDetails
type NetworksNetworkIdFloorPlansDetails struct {
	// Additional property name
	Name *string `json:"name,omitempty"`
	// Additional property value
	Value *string `json:"value,omitempty"`
}

// NewNetworksNetworkIdFloorPlansDetails instantiates a new NetworksNetworkIdFloorPlansDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansDetails() *NetworksNetworkIdFloorPlansDetails {
	this := NetworksNetworkIdFloorPlansDetails{}
	return &this
}

// NewNetworksNetworkIdFloorPlansDetailsWithDefaults instantiates a new NetworksNetworkIdFloorPlansDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansDetailsWithDefaults() *NetworksNetworkIdFloorPlansDetails {
	this := NetworksNetworkIdFloorPlansDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFloorPlansDetails) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansDetails) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansDetails) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansDetails) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworksNetworkIdFloorPlansDetails) SetValue(v string) {
	o.Value = &v
}

func (o NetworksNetworkIdFloorPlansDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansDetails struct {
	value *NetworksNetworkIdFloorPlansDetails
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansDetails) Get() *NetworksNetworkIdFloorPlansDetails {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansDetails) Set(val *NetworksNetworkIdFloorPlansDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansDetails(val *NetworksNetworkIdFloorPlansDetails) *NullableNetworksNetworkIdFloorPlansDetails {
	return &NullableNetworksNetworkIdFloorPlansDetails{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


