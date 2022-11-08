/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20027 struct for InlineResponse20027
type InlineResponse20027 struct {
	// ID of the sensor alert profile.
	ProfileId *string `json:"profileId,omitempty"`
	// Name of the sensor alert profile.
	Name *string `json:"name,omitempty"`
	Schedule *NetworksNetworkIdSensorAlertsProfilesSchedule `json:"schedule,omitempty"`
	// List of conditions that will cause the profile to send an alert.
	Conditions []NetworksNetworkIdSensorAlertsProfilesConditions `json:"conditions"`
	Recipients *NetworksNetworkIdSensorAlertsProfilesRecipients `json:"recipients,omitempty"`
	// List of device serials assigned to this sensor alert profile.
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse20027 instantiates a new InlineResponse20027 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027(conditions []NetworksNetworkIdSensorAlertsProfilesConditions) *InlineResponse20027 {
	this := InlineResponse20027{}
	this.Conditions = conditions
	return &this
}

// NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027WithDefaults() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *InlineResponse20027) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *InlineResponse20027) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *InlineResponse20027) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20027) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20027) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20027) SetName(v string) {
	o.Name = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule {
	if o == nil || isNil(o.Schedule) {
		var ret NetworksNetworkIdSensorAlertsProfilesSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NetworksNetworkIdSensorAlertsProfilesSchedule and assigns it to the Schedule field.
func (o *InlineResponse20027) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule) {
	o.Schedule = &v
}

// GetConditions returns the Conditions field value
func (o *InlineResponse20027) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions {
	if o == nil {
		var ret []NetworksNetworkIdSensorAlertsProfilesConditions
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetConditionsOk() ([]NetworksNetworkIdSensorAlertsProfilesConditions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *InlineResponse20027) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions) {
	o.Conditions = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InlineResponse20027) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients {
	if o == nil || isNil(o.Recipients) {
		var ret NetworksNetworkIdSensorAlertsProfilesRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool) {
	if o == nil || isNil(o.Recipients) {
    return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InlineResponse20027) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given NetworksNetworkIdSensorAlertsProfilesRecipients and assigns it to the Recipients field.
func (o *InlineResponse20027) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients) {
	o.Recipients = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse20027) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse20027) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
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

type NullableInlineResponse20027 struct {
	value *InlineResponse20027
	isSet bool
}

func (v NullableInlineResponse20027) Get() *InlineResponse20027 {
	return v.value
}

func (v *NullableInlineResponse20027) Set(val *InlineResponse20027) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027(val *InlineResponse20027) *NullableInlineResponse20027 {
	return &NullableInlineResponse20027{value: val, isSet: true}
}

func (v NullableInlineResponse20027) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


