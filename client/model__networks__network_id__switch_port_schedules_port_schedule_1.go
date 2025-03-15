/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchPortSchedulesPortSchedule1     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day. 
type NetworksNetworkIdSwitchPortSchedulesPortSchedule1 struct {
	Monday *NetworksNetworkIdGroupPoliciesSchedulingMonday `json:"monday,omitempty"`
	Tuesday *NetworksNetworkIdGroupPoliciesSchedulingTuesday `json:"tuesday,omitempty"`
	Wednesday *NetworksNetworkIdGroupPoliciesSchedulingWednesday `json:"wednesday,omitempty"`
	Thursday *NetworksNetworkIdGroupPoliciesSchedulingThursday `json:"thursday,omitempty"`
	Friday *NetworksNetworkIdGroupPoliciesSchedulingFriday `json:"friday,omitempty"`
	Saturday *NetworksNetworkIdGroupPoliciesSchedulingSaturday `json:"saturday,omitempty"`
	Sunday *NetworksNetworkIdGroupPoliciesSchedulingSunday `json:"sunday,omitempty"`
}

// NewNetworksNetworkIdSwitchPortSchedulesPortSchedule1 instantiates a new NetworksNetworkIdSwitchPortSchedulesPortSchedule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchPortSchedulesPortSchedule1() *NetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	this := NetworksNetworkIdSwitchPortSchedulesPortSchedule1{}
	return &this
}

// NewNetworksNetworkIdSwitchPortSchedulesPortSchedule1WithDefaults instantiates a new NetworksNetworkIdSwitchPortSchedulesPortSchedule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchPortSchedulesPortSchedule1WithDefaults() *NetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	this := NetworksNetworkIdSwitchPortSchedulesPortSchedule1{}
	return &this
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetMonday() NetworksNetworkIdGroupPoliciesSchedulingMonday {
	if o == nil || isNil(o.Monday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetMondayOk() (*NetworksNetworkIdGroupPoliciesSchedulingMonday, bool) {
	if o == nil || isNil(o.Monday) {
    return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasMonday() bool {
	if o != nil && !isNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingMonday and assigns it to the Monday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetMonday(v NetworksNetworkIdGroupPoliciesSchedulingMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetTuesday() NetworksNetworkIdGroupPoliciesSchedulingTuesday {
	if o == nil || isNil(o.Tuesday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetTuesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingTuesday, bool) {
	if o == nil || isNil(o.Tuesday) {
    return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasTuesday() bool {
	if o != nil && !isNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingTuesday and assigns it to the Tuesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetTuesday(v NetworksNetworkIdGroupPoliciesSchedulingTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetWednesday() NetworksNetworkIdGroupPoliciesSchedulingWednesday {
	if o == nil || isNil(o.Wednesday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetWednesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingWednesday, bool) {
	if o == nil || isNil(o.Wednesday) {
    return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasWednesday() bool {
	if o != nil && !isNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingWednesday and assigns it to the Wednesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetWednesday(v NetworksNetworkIdGroupPoliciesSchedulingWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetThursday() NetworksNetworkIdGroupPoliciesSchedulingThursday {
	if o == nil || isNil(o.Thursday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetThursdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingThursday, bool) {
	if o == nil || isNil(o.Thursday) {
    return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasThursday() bool {
	if o != nil && !isNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingThursday and assigns it to the Thursday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetThursday(v NetworksNetworkIdGroupPoliciesSchedulingThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetFriday() NetworksNetworkIdGroupPoliciesSchedulingFriday {
	if o == nil || isNil(o.Friday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetFridayOk() (*NetworksNetworkIdGroupPoliciesSchedulingFriday, bool) {
	if o == nil || isNil(o.Friday) {
    return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasFriday() bool {
	if o != nil && !isNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingFriday and assigns it to the Friday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetFriday(v NetworksNetworkIdGroupPoliciesSchedulingFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetSaturday() NetworksNetworkIdGroupPoliciesSchedulingSaturday {
	if o == nil || isNil(o.Saturday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetSaturdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSaturday, bool) {
	if o == nil || isNil(o.Saturday) {
    return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasSaturday() bool {
	if o != nil && !isNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingSaturday and assigns it to the Saturday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetSaturday(v NetworksNetworkIdGroupPoliciesSchedulingSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetSunday() NetworksNetworkIdGroupPoliciesSchedulingSunday {
	if o == nil || isNil(o.Sunday) {
		var ret NetworksNetworkIdGroupPoliciesSchedulingSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) GetSundayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSunday, bool) {
	if o == nil || isNil(o.Sunday) {
    return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) HasSunday() bool {
	if o != nil && !isNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingSunday and assigns it to the Sunday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) SetSunday(v NetworksNetworkIdGroupPoliciesSchedulingSunday) {
	o.Sunday = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortSchedule1) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1 struct {
	value *NetworksNetworkIdSwitchPortSchedulesPortSchedule1
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) Get() *NetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) Set(val *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1(val *NetworksNetworkIdSwitchPortSchedulesPortSchedule1) *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	return &NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchPortSchedulesPortSchedule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


