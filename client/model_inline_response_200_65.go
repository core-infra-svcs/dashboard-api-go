/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20065 struct for InlineResponse20065
type InlineResponse20065 struct {
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

// NewInlineResponse20065 instantiates a new InlineResponse20065 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065WithDefaults() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *InlineResponse20065) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
    return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *InlineResponse20065) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *InlineResponse20065) SetAction(v string) {
	o.Action = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20065) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20065) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20065) SetName(v string) {
	o.Name = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse20065) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse20065) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *InlineResponse20065) SetDetails(v string) {
	o.Details = &v
}

// GetDashboardUser returns the DashboardUser field value if set, zero value otherwise.
func (o *InlineResponse20065) GetDashboardUser() string {
	if o == nil || isNil(o.DashboardUser) {
		var ret string
		return ret
	}
	return *o.DashboardUser
}

// GetDashboardUserOk returns a tuple with the DashboardUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetDashboardUserOk() (*string, bool) {
	if o == nil || isNil(o.DashboardUser) {
    return nil, false
	}
	return o.DashboardUser, true
}

// HasDashboardUser returns a boolean if a field has been set.
func (o *InlineResponse20065) HasDashboardUser() bool {
	if o != nil && !isNil(o.DashboardUser) {
		return true
	}

	return false
}

// SetDashboardUser gets a reference to the given string and assigns it to the DashboardUser field.
func (o *InlineResponse20065) SetDashboardUser(v string) {
	o.DashboardUser = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20065) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20065) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse20065) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse20065) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20065 struct {
	value *InlineResponse20065
	isSet bool
}

func (v NullableInlineResponse20065) Get() *InlineResponse20065 {
	return v.value
}

func (v *NullableInlineResponse20065) Set(val *InlineResponse20065) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065(val *InlineResponse20065) *NullableInlineResponse20065 {
	return &NullableInlineResponse20065{value: val, isSet: true}
}

func (v NullableInlineResponse20065) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


