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

// NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday Friday schedule
type NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday struct {
	// Whether the schedule is active or inactive
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'
	To *string `json:"to,omitempty"`
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday() *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday{}
	return &this
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleFridayWithDefaults instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleFridayWithDefaults() *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) SetTo(v string) {
	o.To = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday struct {
	value *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) Get() *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) Set(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday {
	return &NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


