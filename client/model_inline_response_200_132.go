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

// InlineResponse200132 struct for InlineResponse200132
type InlineResponse200132 struct {
	// The type of command sent to the device.
	Action *string `json:"action,omitempty"`
	// The name of the device to which the command is sent.
	Name *string `json:"name,omitempty"`
	// A JSON string object containing command details.
	Details *string `json:"details,omitempty"`
	// The Meraki dashboard user who initiated the command.
	DashboardUser *string `json:"dashboardUser,omitempty"`
	// The time the command was sent to the device.
	Ts *string `json:"ts,omitempty"`
}

// NewInlineResponse200132 instantiates a new InlineResponse200132 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200132() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// NewInlineResponse200132WithDefaults instantiates a new InlineResponse200132 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200132WithDefaults() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *InlineResponse200132) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
    return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *InlineResponse200132) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *InlineResponse200132) SetAction(v string) {
	o.Action = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200132) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200132) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200132) SetName(v string) {
	o.Name = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse200132) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse200132) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *InlineResponse200132) SetDetails(v string) {
	o.Details = &v
}

// GetDashboardUser returns the DashboardUser field value if set, zero value otherwise.
func (o *InlineResponse200132) GetDashboardUser() string {
	if o == nil || isNil(o.DashboardUser) {
		var ret string
		return ret
	}
	return *o.DashboardUser
}

// GetDashboardUserOk returns a tuple with the DashboardUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetDashboardUserOk() (*string, bool) {
	if o == nil || isNil(o.DashboardUser) {
    return nil, false
	}
	return o.DashboardUser, true
}

// HasDashboardUser returns a boolean if a field has been set.
func (o *InlineResponse200132) HasDashboardUser() bool {
	if o != nil && !isNil(o.DashboardUser) {
		return true
	}

	return false
}

// SetDashboardUser gets a reference to the given string and assigns it to the DashboardUser field.
func (o *InlineResponse200132) SetDashboardUser(v string) {
	o.DashboardUser = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200132) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200132) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse200132) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse200132) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.DashboardUser) {
		toSerialize["dashboardUser"] = o.DashboardUser
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200132 struct {
	value *InlineResponse200132
	isSet bool
}

func (v NullableInlineResponse200132) Get() *InlineResponse200132 {
	return v.value
}

func (v *NullableInlineResponse200132) Set(val *InlineResponse200132) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200132) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200132) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200132(val *InlineResponse200132) *NullableInlineResponse200132 {
	return &NullableInlineResponse200132{value: val, isSet: true}
}

func (v NullableInlineResponse200132) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200132) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


