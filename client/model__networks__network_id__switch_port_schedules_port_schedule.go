/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchPortSchedulesPortSchedule     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day. 
type NetworksNetworkIdSwitchPortSchedulesPortSchedule struct {
	Monday *NetworksNetworkIdGroupPoliciesSchedulingMonday `json:"monday,omitempty"`
	Tuesday *NetworksNetworkIdGroupPoliciesSchedulingTuesday `json:"tuesday,omitempty"`
	Wednesday *NetworksNetworkIdGroupPoliciesSchedulingWednesday `json:"wednesday,omitempty"`
	Thursday *NetworksNetworkIdGroupPoliciesSchedulingThursday `json:"thursday,omitempty"`
	Friday *NetworksNetworkIdGroupPoliciesSchedulingFriday `json:"friday,omitempty"`
	Saturday *NetworksNetworkIdGroupPoliciesSchedulingSaturday `json:"saturday,omitempty"`
	Sunday *NetworksNetworkIdGroupPoliciesSchedulingSunday `json:"sunday,omitempty"`
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
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetMonday() NetworksNetworkIdGroupPoliciesSchedulingMonday {
	if o == nil || o.Monday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetMondayOk() (*NetworksNetworkIdGroupPoliciesSchedulingMonday, bool) {
	if o == nil || o.Monday == nil {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasMonday() bool {
	if o != nil && o.Monday != nil {
		return true
	}

	return false
}

// SetMonday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingMonday and assigns it to the Monday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetMonday(v NetworksNetworkIdGroupPoliciesSchedulingMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetTuesday() NetworksNetworkIdGroupPoliciesSchedulingTuesday {
	if o == nil || o.Tuesday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetTuesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingTuesday, bool) {
	if o == nil || o.Tuesday == nil {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasTuesday() bool {
	if o != nil && o.Tuesday != nil {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingTuesday and assigns it to the Tuesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetTuesday(v NetworksNetworkIdGroupPoliciesSchedulingTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetWednesday() NetworksNetworkIdGroupPoliciesSchedulingWednesday {
	if o == nil || o.Wednesday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetWednesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingWednesday, bool) {
	if o == nil || o.Wednesday == nil {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasWednesday() bool {
	if o != nil && o.Wednesday != nil {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingWednesday and assigns it to the Wednesday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetWednesday(v NetworksNetworkIdGroupPoliciesSchedulingWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetThursday() NetworksNetworkIdGroupPoliciesSchedulingThursday {
	if o == nil || o.Thursday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetThursdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingThursday, bool) {
	if o == nil || o.Thursday == nil {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasThursday() bool {
	if o != nil && o.Thursday != nil {
		return true
	}

	return false
}

// SetThursday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingThursday and assigns it to the Thursday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetThursday(v NetworksNetworkIdGroupPoliciesSchedulingThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetFriday() NetworksNetworkIdGroupPoliciesSchedulingFriday {
	if o == nil || o.Friday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetFridayOk() (*NetworksNetworkIdGroupPoliciesSchedulingFriday, bool) {
	if o == nil || o.Friday == nil {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasFriday() bool {
	if o != nil && o.Friday != nil {
		return true
	}

	return false
}

// SetFriday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingFriday and assigns it to the Friday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetFriday(v NetworksNetworkIdGroupPoliciesSchedulingFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSaturday() NetworksNetworkIdGroupPoliciesSchedulingSaturday {
	if o == nil || o.Saturday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSaturdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSaturday, bool) {
	if o == nil || o.Saturday == nil {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasSaturday() bool {
	if o != nil && o.Saturday != nil {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingSaturday and assigns it to the Saturday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetSaturday(v NetworksNetworkIdGroupPoliciesSchedulingSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSunday() NetworksNetworkIdGroupPoliciesSchedulingSunday {
	if o == nil || o.Sunday == nil {
		var ret NetworksNetworkIdGroupPoliciesSchedulingSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) GetSundayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSunday, bool) {
	if o == nil || o.Sunday == nil {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) HasSunday() bool {
	if o != nil && o.Sunday != nil {
		return true
	}

	return false
}

// SetSunday gets a reference to the given NetworksNetworkIdGroupPoliciesSchedulingSunday and assigns it to the Sunday field.
func (o *NetworksNetworkIdSwitchPortSchedulesPortSchedule) SetSunday(v NetworksNetworkIdGroupPoliciesSchedulingSunday) {
	o.Sunday = &v
}

func (o NetworksNetworkIdSwitchPortSchedulesPortSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Monday != nil {
		toSerialize["monday"] = o.Monday
	}
	if o.Tuesday != nil {
		toSerialize["tuesday"] = o.Tuesday
	}
	if o.Wednesday != nil {
		toSerialize["wednesday"] = o.Wednesday
	}
	if o.Thursday != nil {
		toSerialize["thursday"] = o.Thursday
	}
	if o.Friday != nil {
		toSerialize["friday"] = o.Friday
	}
	if o.Saturday != nil {
		toSerialize["saturday"] = o.Saturday
	}
	if o.Sunday != nil {
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


