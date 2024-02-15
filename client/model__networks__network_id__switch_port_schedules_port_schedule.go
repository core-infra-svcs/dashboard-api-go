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

// NetworksNetworkIdSwitchPortSchedulesPortSchedule Port schedule
type NetworksNetworkIdSwitchPortSchedulesPortSchedule struct {
	Monday *NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday `json:"monday,omitempty"`
	Tuesday *NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday `json:"tuesday,omitempty"`
	Wednesday *NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday `json:"wednesday,omitempty"`
	Thursday *NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday `json:"thursday,omitempty"`
	Friday *NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday `json:"friday,omitempty"`
	Saturday *NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday `json:"saturday,omitempty"`
	Sunday *NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday `json:"sunday,omitempty"`
}

// NewNetworksNetworkIdSwitchPortSchedulesPortSchedule instantiates a new NetworksNetworkIdSwitchPortSchedulesPortSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchPortSchedulesPortSchedule() *NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	this := NetworksNetworkIdSwitchPortSchedulesPortSchedule{}
	return &this
}

// NewNetworksNetworkIdSwitchPortSchedulesPortScheduleWithDefaults instantiates a new NetworksNetworkIdSwitchPortSchedulesPortSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchPortSchedulesPortScheduleWithDefaults() *NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	this := NetworksNetworkIdSwitchPortSchedulesPortSchedule{}
	return &this
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetMonday() NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday {
	if o == nil || isNil(o.Monday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetMondayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday, bool) {
	if o == nil || isNil(o.Monday) {
    return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasMonday() bool {
	if o != nil && !isNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday and assigns it to the Monday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetMonday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetTuesday() NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday {
	if o == nil || isNil(o.Tuesday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetTuesdayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday, bool) {
	if o == nil || isNil(o.Tuesday) {
    return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasTuesday() bool {
	if o != nil && !isNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday and assigns it to the Tuesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetTuesday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetWednesday() NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday {
	if o == nil || isNil(o.Wednesday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetWednesdayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday, bool) {
	if o == nil || isNil(o.Wednesday) {
    return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasWednesday() bool {
	if o != nil && !isNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday and assigns it to the Wednesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetWednesday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetThursday() NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday {
	if o == nil || isNil(o.Thursday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetThursdayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday, bool) {
	if o == nil || isNil(o.Thursday) {
    return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasThursday() bool {
	if o != nil && !isNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday and assigns it to the Thursday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetThursday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetFriday() NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday {
	if o == nil || isNil(o.Friday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetFridayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday, bool) {
	if o == nil || isNil(o.Friday) {
    return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasFriday() bool {
	if o != nil && !isNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday and assigns it to the Friday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetFriday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSaturday() NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday {
	if o == nil || isNil(o.Saturday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSaturdayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday, bool) {
	if o == nil || isNil(o.Saturday) {
    return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasSaturday() bool {
	if o != nil && !isNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday and assigns it to the Saturday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetSaturday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSunday() NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday {
	if o == nil || isNil(o.Sunday) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSundayOk() (*NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday, bool) {
	if o == nil || isNil(o.Sunday) {
    return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasSunday() bool {
	if o != nil && !isNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday and assigns it to the Sunday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetSunday(v NetworksNetworkIdSwitchPortSchedulesPortScheduleSunday) {
	o.Sunday = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !isNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !isNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !isNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !isNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !isNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	if !isNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule struct {
	value *NetworksNetworkIdSwitchPortSchedulesPortSchedule
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) Get() *NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) Set(val *NetworksNetworkIdSwitchPortSchedulesPortSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchPortSchedulesPortSchedule(val *NetworksNetworkIdSwitchPortSchedulesPortSchedule) *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule {
	return &NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


