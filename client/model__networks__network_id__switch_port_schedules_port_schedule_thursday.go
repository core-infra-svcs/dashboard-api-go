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

// NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday Thursday schedule
type NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday struct {
	// Whether the schedule is active or inactive
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'
	To *string `json:"to,omitempty"`
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday() *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday{}
	return &this
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleThursdayWithDefaults instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleThursdayWithDefaults() *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) SetTo(v string) {
	o.To = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday struct {
	value *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) Get() *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) Set(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday {
	return &NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


