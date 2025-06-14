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

// InlineResponse200369Items struct for InlineResponse200369Items
type InlineResponse200369Items struct {
	// Wireless LAN controller cloud ID
	Serial *string `json:"serial,omitempty"`
	Network *InlineResponse200369Network `json:"network,omitempty"`
	// Overview history of a wireless LAN controller
	Readings []InlineResponse200369Readings `json:"readings,omitempty"`
}

// NewInlineResponse200369Items instantiates a new InlineResponse200369Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200369Items() *InlineResponse200369Items {
	this := InlineResponse200369Items{}
	return &this
}

// NewInlineResponse200369ItemsWithDefaults instantiates a new InlineResponse200369Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200369ItemsWithDefaults() *InlineResponse200369Items {
	this := InlineResponse200369Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200369Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200369Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200369Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200369Items) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200369Items) GetNetwork() InlineResponse200369Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200369Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200369Items) GetNetworkOk() (*InlineResponse200369Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200369Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200369Network and assigns it to the Network field.
func (o *InlineResponse200369Items) SetNetwork(v InlineResponse200369Network) {
	o.Network = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *InlineResponse200369Items) GetReadings() []InlineResponse200369Readings {
	if o == nil || isNil(o.Readings) {
		var ret []InlineResponse200369Readings
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200369Items) GetReadingsOk() ([]InlineResponse200369Readings, bool) {
	if o == nil || isNil(o.Readings) {
    return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *InlineResponse200369Items) HasReadings() bool {
	if o != nil && !isNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []InlineResponse200369Readings and assigns it to the Readings field.
func (o *InlineResponse200369Items) SetReadings(v []InlineResponse200369Readings) {
	o.Readings = v
}

func (o InlineResponse200369Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200369Items struct {
	value *InlineResponse200369Items
	isSet bool
}

func (v NullableInlineResponse200369Items) Get() *InlineResponse200369Items {
	return v.value
}

func (v *NullableInlineResponse200369Items) Set(val *InlineResponse200369Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200369Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200369Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200369Items(val *InlineResponse200369Items) *NullableInlineResponse200369Items {
	return &NullableInlineResponse200369Items{value: val, isSet: true}
}

func (v NullableInlineResponse200369Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200369Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


