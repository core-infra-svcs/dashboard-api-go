/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20025Cdp cdp information
type InlineResponse20025Cdp struct {
	// ID for the device
	DeviceId *string `json:"deviceId,omitempty"`
	// ID for the port
	PortId *string `json:"portId,omitempty"`
	// MAC address
	Address *string `json:"address,omitempty"`
	// Source port
	SourcePort *string `json:"sourcePort,omitempty"`
}

// NewInlineResponse20025Cdp instantiates a new InlineResponse20025Cdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025Cdp() *InlineResponse20025Cdp {
	this := InlineResponse20025Cdp{}
	return &this
}

// NewInlineResponse20025CdpWithDefaults instantiates a new InlineResponse20025Cdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025CdpWithDefaults() *InlineResponse20025Cdp {
	this := InlineResponse20025Cdp{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20025Cdp) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Cdp) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20025Cdp) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20025Cdp) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *InlineResponse20025Cdp) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Cdp) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *InlineResponse20025Cdp) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *InlineResponse20025Cdp) SetPortId(v string) {
	o.PortId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse20025Cdp) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Cdp) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse20025Cdp) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse20025Cdp) SetAddress(v string) {
	o.Address = &v
}

// GetSourcePort returns the SourcePort field value if set, zero value otherwise.
func (o *InlineResponse20025Cdp) GetSourcePort() string {
	if o == nil || isNil(o.SourcePort) {
		var ret string
		return ret
	}
	return *o.SourcePort
}

// GetSourcePortOk returns a tuple with the SourcePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Cdp) GetSourcePortOk() (*string, bool) {
	if o == nil || isNil(o.SourcePort) {
    return nil, false
	}
	return o.SourcePort, true
}

// HasSourcePort returns a boolean if a field has been set.
func (o *InlineResponse20025Cdp) HasSourcePort() bool {
	if o != nil && !isNil(o.SourcePort) {
		return true
	}

	return false
}

// SetSourcePort gets a reference to the given string and assigns it to the SourcePort field.
func (o *InlineResponse20025Cdp) SetSourcePort(v string) {
	o.SourcePort = &v
}

func (o InlineResponse20025Cdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.SourcePort) {
		toSerialize["sourcePort"] = o.SourcePort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025Cdp struct {
	value *InlineResponse20025Cdp
	isSet bool
}

func (v NullableInlineResponse20025Cdp) Get() *InlineResponse20025Cdp {
	return v.value
}

func (v *NullableInlineResponse20025Cdp) Set(val *InlineResponse20025Cdp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025Cdp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025Cdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025Cdp(val *InlineResponse20025Cdp) *NullableInlineResponse20025Cdp {
	return &NullableInlineResponse20025Cdp{value: val, isSet: true}
}

func (v NullableInlineResponse20025Cdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025Cdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


