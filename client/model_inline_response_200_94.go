/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20094 struct for InlineResponse20094
type InlineResponse20094 struct {
	// SSID Number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// LAN
	Vlan *int32 `json:"vlan,omitempty"`
	// Client Mac
	ClientMac *string `json:"clientMac,omitempty"`
	// Serial Number
	Serial *string `json:"serial,omitempty"`
	// The failed onboarding step. One of: assoc, auth, dhcp, dns.
	FailureStep *string `json:"failureStep,omitempty"`
	// The failure type in the onboarding step
	Type *string `json:"type,omitempty"`
	// The timestamp when the client mac failed
	Ts *time.Time `json:"ts,omitempty"`
}

// NewInlineResponse20094 instantiates a new InlineResponse20094 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20094() *InlineResponse20094 {
	this := InlineResponse20094{}
	return &this
}

// NewInlineResponse20094WithDefaults instantiates a new InlineResponse20094 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20094WithDefaults() *InlineResponse20094 {
	this := InlineResponse20094{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse20094) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse20094) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse20094) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20094) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20094) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse20094) SetVlan(v int32) {
	o.Vlan = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *InlineResponse20094) GetClientMac() string {
	if o == nil || isNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetClientMacOk() (*string, bool) {
	if o == nil || isNil(o.ClientMac) {
    return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *InlineResponse20094) HasClientMac() bool {
	if o != nil && !isNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *InlineResponse20094) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20094) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20094) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20094) SetSerial(v string) {
	o.Serial = &v
}

// GetFailureStep returns the FailureStep field value if set, zero value otherwise.
func (o *InlineResponse20094) GetFailureStep() string {
	if o == nil || isNil(o.FailureStep) {
		var ret string
		return ret
	}
	return *o.FailureStep
}

// GetFailureStepOk returns a tuple with the FailureStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetFailureStepOk() (*string, bool) {
	if o == nil || isNil(o.FailureStep) {
    return nil, false
	}
	return o.FailureStep, true
}

// HasFailureStep returns a boolean if a field has been set.
func (o *InlineResponse20094) HasFailureStep() bool {
	if o != nil && !isNil(o.FailureStep) {
		return true
	}

	return false
}

// SetFailureStep gets a reference to the given string and assigns it to the FailureStep field.
func (o *InlineResponse20094) SetFailureStep(v string) {
	o.FailureStep = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20094) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20094) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20094) SetType(v string) {
	o.Type = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20094) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20094) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20094) SetTs(v time.Time) {
	o.Ts = &v
}

func (o InlineResponse20094) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.ClientMac) {
		toSerialize["clientMac"] = o.ClientMac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.FailureStep) {
		toSerialize["failureStep"] = o.FailureStep
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20094 struct {
	value *InlineResponse20094
	isSet bool
}

func (v NullableInlineResponse20094) Get() *InlineResponse20094 {
	return v.value
}

func (v *NullableInlineResponse20094) Set(val *InlineResponse20094) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20094) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20094) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20094(val *InlineResponse20094) *NullableInlineResponse20094 {
	return &NullableInlineResponse20094{value: val, isSet: true}
}

func (v NullableInlineResponse20094) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20094) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


