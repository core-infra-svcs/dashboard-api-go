/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200337Items struct for InlineResponse200337Items
type InlineResponse200337Items struct {
	// Wireless LAN controller cloud ID
	Serial *string `json:"serial,omitempty"`
	// Wireless LAN controller redundancy SSO (stateful switchover)
	Mode *string `json:"mode,omitempty"`
	// Wireless LAN controller redundancy enablement
	Enabled *bool `json:"enabled,omitempty"`
	Failover *InlineResponse200337Failover `json:"failover,omitempty"`
	// Wireless LAN controller redundancy mobility mac 
	MobilityMac *string `json:"mobilityMac,omitempty"`
}

// NewInlineResponse200337Items instantiates a new InlineResponse200337Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337Items() *InlineResponse200337Items {
	this := InlineResponse200337Items{}
	return &this
}

// NewInlineResponse200337ItemsWithDefaults instantiates a new InlineResponse200337Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337ItemsWithDefaults() *InlineResponse200337Items {
	this := InlineResponse200337Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200337Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200337Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200337Items) SetSerial(v string) {
	o.Serial = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse200337Items) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Items) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse200337Items) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse200337Items) SetMode(v string) {
	o.Mode = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200337Items) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Items) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200337Items) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200337Items) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *InlineResponse200337Items) GetFailover() InlineResponse200337Failover {
	if o == nil || isNil(o.Failover) {
		var ret InlineResponse200337Failover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Items) GetFailoverOk() (*InlineResponse200337Failover, bool) {
	if o == nil || isNil(o.Failover) {
    return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *InlineResponse200337Items) HasFailover() bool {
	if o != nil && !isNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given InlineResponse200337Failover and assigns it to the Failover field.
func (o *InlineResponse200337Items) SetFailover(v InlineResponse200337Failover) {
	o.Failover = &v
}

// GetMobilityMac returns the MobilityMac field value if set, zero value otherwise.
func (o *InlineResponse200337Items) GetMobilityMac() string {
	if o == nil || isNil(o.MobilityMac) {
		var ret string
		return ret
	}
	return *o.MobilityMac
}

// GetMobilityMacOk returns a tuple with the MobilityMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Items) GetMobilityMacOk() (*string, bool) {
	if o == nil || isNil(o.MobilityMac) {
    return nil, false
	}
	return o.MobilityMac, true
}

// HasMobilityMac returns a boolean if a field has been set.
func (o *InlineResponse200337Items) HasMobilityMac() bool {
	if o != nil && !isNil(o.MobilityMac) {
		return true
	}

	return false
}

// SetMobilityMac gets a reference to the given string and assigns it to the MobilityMac field.
func (o *InlineResponse200337Items) SetMobilityMac(v string) {
	o.MobilityMac = &v
}

func (o InlineResponse200337Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Failover) {
		toSerialize["failover"] = o.Failover
	}
	if !isNil(o.MobilityMac) {
		toSerialize["mobilityMac"] = o.MobilityMac
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337Items struct {
	value *InlineResponse200337Items
	isSet bool
}

func (v NullableInlineResponse200337Items) Get() *InlineResponse200337Items {
	return v.value
}

func (v *NullableInlineResponse200337Items) Set(val *InlineResponse200337Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337Items(val *InlineResponse200337Items) *NullableInlineResponse200337Items {
	return &NullableInlineResponse200337Items{value: val, isSet: true}
}

func (v NullableInlineResponse200337Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


