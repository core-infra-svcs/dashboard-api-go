/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200275Device Device associated to the schedule
type InlineResponse200275Device struct {
	// The serial of the device
	Serial *string `json:"serial,omitempty"`
	// The switchports on which to take the packet capture
	Switchports *string `json:"switchports,omitempty"`
	// The interfaces on which to take the packet capture (applicable for Catalyst devices)
	Interface *string `json:"interface,omitempty"`
}

// NewInlineResponse200275Device instantiates a new InlineResponse200275Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200275Device() *InlineResponse200275Device {
	this := InlineResponse200275Device{}
	return &this
}

// NewInlineResponse200275DeviceWithDefaults instantiates a new InlineResponse200275Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200275DeviceWithDefaults() *InlineResponse200275Device {
	this := InlineResponse200275Device{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200275Device) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200275Device) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200275Device) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200275Device) SetSerial(v string) {
	o.Serial = &v
}

// GetSwitchports returns the Switchports field value if set, zero value otherwise.
func (o *InlineResponse200275Device) GetSwitchports() string {
	if o == nil || isNil(o.Switchports) {
		var ret string
		return ret
	}
	return *o.Switchports
}

// GetSwitchportsOk returns a tuple with the Switchports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200275Device) GetSwitchportsOk() (*string, bool) {
	if o == nil || isNil(o.Switchports) {
    return nil, false
	}
	return o.Switchports, true
}

// HasSwitchports returns a boolean if a field has been set.
func (o *InlineResponse200275Device) HasSwitchports() bool {
	if o != nil && !isNil(o.Switchports) {
		return true
	}

	return false
}

// SetSwitchports gets a reference to the given string and assigns it to the Switchports field.
func (o *InlineResponse200275Device) SetSwitchports(v string) {
	o.Switchports = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *InlineResponse200275Device) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200275Device) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *InlineResponse200275Device) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *InlineResponse200275Device) SetInterface(v string) {
	o.Interface = &v
}

func (o InlineResponse200275Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Switchports) {
		toSerialize["switchports"] = o.Switchports
	}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200275Device struct {
	value *InlineResponse200275Device
	isSet bool
}

func (v NullableInlineResponse200275Device) Get() *InlineResponse200275Device {
	return v.value
}

func (v *NullableInlineResponse200275Device) Set(val *InlineResponse200275Device) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200275Device) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200275Device) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200275Device(val *InlineResponse200275Device) *NullableInlineResponse200275Device {
	return &NullableInlineResponse200275Device{value: val, isSet: true}
}

func (v NullableInlineResponse200275Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200275Device) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


