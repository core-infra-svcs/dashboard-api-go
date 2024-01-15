/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday Sunday schedule
type NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday struct {
	// Whether the schedule is active or inactive
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'
	To *string `json:"to,omitempty"`
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday() *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday{}
	return &this
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleSundayWithDefaults instantiates a new NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleSundayWithDefaults() *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday {
	this := NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) SetTo(v string) {
	o.To = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday struct {
	value *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) Get() *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) Set(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday(val *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday {
	return &NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


