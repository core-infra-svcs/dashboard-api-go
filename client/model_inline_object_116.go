/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject116 struct for InlineObject116
type InlineObject116 struct {
	// Name of the sensor alert profile.
	Name *string `json:"name,omitempty"`
	Schedule *NetworksNetworkIdSensorAlertsProfilesSchedule1 `json:"schedule,omitempty"`
	// List of conditions that will cause the profile to send an alert.
	Conditions []NetworksNetworkIdSensorAlertsProfilesConditions `json:"conditions,omitempty"`
	Recipients *NetworksNetworkIdSensorAlertsProfilesRecipients `json:"recipients,omitempty"`
	// List of device serials assigned to this sensor alert profile.
	Serials []string `json:"serials,omitempty"`
}

// NewInlineObject116 instantiates a new InlineObject116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject116() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// NewInlineObject116WithDefaults instantiates a new InlineObject116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject116WithDefaults() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject116) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject116) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject116) SetName(v string) {
	o.Name = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *InlineObject116) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule1 {
	if o == nil || isNil(o.Schedule) {
		var ret NetworksNetworkIdSensorAlertsProfilesSchedule1
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule1, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *InlineObject116) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NetworksNetworkIdSensorAlertsProfilesSchedule1 and assigns it to the Schedule field.
func (o *InlineObject116) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule1) {
	o.Schedule = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *InlineObject116) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions {
	if o == nil || isNil(o.Conditions) {
		var ret []NetworksNetworkIdSensorAlertsProfilesConditions
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetConditionsOk() ([]NetworksNetworkIdSensorAlertsProfilesConditions, bool) {
	if o == nil || isNil(o.Conditions) {
    return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *InlineObject116) HasConditions() bool {
	if o != nil && !isNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []NetworksNetworkIdSensorAlertsProfilesConditions and assigns it to the Conditions field.
func (o *InlineObject116) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions) {
	o.Conditions = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InlineObject116) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients {
	if o == nil || isNil(o.Recipients) {
		var ret NetworksNetworkIdSensorAlertsProfilesRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool) {
	if o == nil || isNil(o.Recipients) {
    return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InlineObject116) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given NetworksNetworkIdSensorAlertsProfilesRecipients and assigns it to the Recipients field.
func (o *InlineObject116) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients) {
	o.Recipients = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject116) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject116) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject116) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !isNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !isNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject116 struct {
	value *InlineObject116
	isSet bool
}

func (v NullableInlineObject116) Get() *InlineObject116 {
	return v.value
}

func (v *NullableInlineObject116) Set(val *InlineObject116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject116(val *InlineObject116) *NullableInlineObject116 {
	return &NullableInlineObject116{value: val, isSet: true}
}

func (v NullableInlineObject116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


