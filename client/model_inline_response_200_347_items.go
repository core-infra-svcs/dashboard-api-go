/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200347Items struct for InlineResponse200347Items
type InlineResponse200347Items struct {
	// Access points cloud ID
	Serial *string `json:"serial,omitempty"`
	Controller *InlineResponse200347Controller `json:"controller,omitempty"`
	Network *InlineResponse200347Network `json:"network,omitempty"`
}

// NewInlineResponse200347Items instantiates a new InlineResponse200347Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200347Items() *InlineResponse200347Items {
	this := InlineResponse200347Items{}
	return &this
}

// NewInlineResponse200347ItemsWithDefaults instantiates a new InlineResponse200347Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200347ItemsWithDefaults() *InlineResponse200347Items {
	this := InlineResponse200347Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200347Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200347Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200347Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200347Items) SetSerial(v string) {
	o.Serial = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *InlineResponse200347Items) GetController() InlineResponse200347Controller {
	if o == nil || isNil(o.Controller) {
		var ret InlineResponse200347Controller
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200347Items) GetControllerOk() (*InlineResponse200347Controller, bool) {
	if o == nil || isNil(o.Controller) {
    return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *InlineResponse200347Items) HasController() bool {
	if o != nil && !isNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given InlineResponse200347Controller and assigns it to the Controller field.
func (o *InlineResponse200347Items) SetController(v InlineResponse200347Controller) {
	o.Controller = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200347Items) GetNetwork() InlineResponse200347Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200347Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200347Items) GetNetworkOk() (*InlineResponse200347Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200347Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200347Network and assigns it to the Network field.
func (o *InlineResponse200347Items) SetNetwork(v InlineResponse200347Network) {
	o.Network = &v
}

func (o InlineResponse200347Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200347Items struct {
	value *InlineResponse200347Items
	isSet bool
}

func (v NullableInlineResponse200347Items) Get() *InlineResponse200347Items {
	return v.value
}

func (v *NullableInlineResponse200347Items) Set(val *InlineResponse200347Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200347Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200347Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200347Items(val *InlineResponse200347Items) *NullableInlineResponse200347Items {
	return &NullableInlineResponse200347Items{value: val, isSet: true}
}

func (v NullableInlineResponse200347Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200347Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


