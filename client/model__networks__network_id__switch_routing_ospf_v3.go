/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchRoutingOspfV3 OSPF v3 configuration
type NetworksNetworkIdSwitchRoutingOspfV3 struct {
	// Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	HelloTimerInSeconds *int32 `json:"helloTimerInSeconds,omitempty"`
	// Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	DeadTimerInSeconds *int32 `json:"deadTimerInSeconds,omitempty"`
	// OSPF v3 areas
	Areas []NetworksNetworkIdSwitchRoutingOspfAreas `json:"areas,omitempty"`
}

// NewNetworksNetworkIdSwitchRoutingOspfV3 instantiates a new NetworksNetworkIdSwitchRoutingOspfV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchRoutingOspfV3() *NetworksNetworkIdSwitchRoutingOspfV3 {
	this := NetworksNetworkIdSwitchRoutingOspfV3{}
	return &this
}

// NewNetworksNetworkIdSwitchRoutingOspfV3WithDefaults instantiates a new NetworksNetworkIdSwitchRoutingOspfV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchRoutingOspfV3WithDefaults() *NetworksNetworkIdSwitchRoutingOspfV3 {
	this := NetworksNetworkIdSwitchRoutingOspfV3{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelloTimerInSeconds returns the HelloTimerInSeconds field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetHelloTimerInSeconds() int32 {
	if o == nil || isNil(o.HelloTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.HelloTimerInSeconds
}

// GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetHelloTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.HelloTimerInSeconds) {
    return nil, false
	}
	return o.HelloTimerInSeconds, true
}

// HasHelloTimerInSeconds returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasHelloTimerInSeconds() bool {
	if o != nil && !isNil(o.HelloTimerInSeconds) {
		return true
	}

	return false
}

// SetHelloTimerInSeconds gets a reference to the given int32 and assigns it to the HelloTimerInSeconds field.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetHelloTimerInSeconds(v int32) {
	o.HelloTimerInSeconds = &v
}

// GetDeadTimerInSeconds returns the DeadTimerInSeconds field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetDeadTimerInSeconds() int32 {
	if o == nil || isNil(o.DeadTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.DeadTimerInSeconds
}

// GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetDeadTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.DeadTimerInSeconds) {
    return nil, false
	}
	return o.DeadTimerInSeconds, true
}

// HasDeadTimerInSeconds returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasDeadTimerInSeconds() bool {
	if o != nil && !isNil(o.DeadTimerInSeconds) {
		return true
	}

	return false
}

// SetDeadTimerInSeconds gets a reference to the given int32 and assigns it to the DeadTimerInSeconds field.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetDeadTimerInSeconds(v int32) {
	o.DeadTimerInSeconds = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetAreas() []NetworksNetworkIdSwitchRoutingOspfAreas {
	if o == nil || isNil(o.Areas) {
		var ret []NetworksNetworkIdSwitchRoutingOspfAreas
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetAreasOk() ([]NetworksNetworkIdSwitchRoutingOspfAreas, bool) {
	if o == nil || isNil(o.Areas) {
    return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasAreas() bool {
	if o != nil && !isNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []NetworksNetworkIdSwitchRoutingOspfAreas and assigns it to the Areas field.
func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetAreas(v []NetworksNetworkIdSwitchRoutingOspfAreas) {
	o.Areas = v
}

func (o NetworksNetworkIdSwitchRoutingOspfV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.HelloTimerInSeconds) {
		toSerialize["helloTimerInSeconds"] = o.HelloTimerInSeconds
	}
	if !isNil(o.DeadTimerInSeconds) {
		toSerialize["deadTimerInSeconds"] = o.DeadTimerInSeconds
	}
	if !isNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchRoutingOspfV3 struct {
	value *NetworksNetworkIdSwitchRoutingOspfV3
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchRoutingOspfV3) Get() *NetworksNetworkIdSwitchRoutingOspfV3 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchRoutingOspfV3) Set(val *NetworksNetworkIdSwitchRoutingOspfV3) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchRoutingOspfV3) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchRoutingOspfV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchRoutingOspfV3(val *NetworksNetworkIdSwitchRoutingOspfV3) *NullableNetworksNetworkIdSwitchRoutingOspfV3 {
	return &NullableNetworksNetworkIdSwitchRoutingOspfV3{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchRoutingOspfV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchRoutingOspfV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


