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

// InlineResponse200339Ports struct for InlineResponse200339Ports
type InlineResponse200339Ports struct {
	// The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	PortId *string `json:"portId,omitempty"`
	// Timestamp for most recent discovery info on this port.
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
	// The Cisco Discovery Protocol (CDP) information of the connected device.
	Cdp []InlineResponse200339Cdp `json:"cdp,omitempty"`
	// The Link Layer Discovery Protocol (LLDP) information of the connected device.
	Lldp []InlineResponse200339Lldp `json:"lldp,omitempty"`
}

// NewInlineResponse200339Ports instantiates a new InlineResponse200339Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339Ports() *InlineResponse200339Ports {
	this := InlineResponse200339Ports{}
	return &this
}

// NewInlineResponse200339PortsWithDefaults instantiates a new InlineResponse200339Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339PortsWithDefaults() *InlineResponse200339Ports {
	this := InlineResponse200339Ports{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *InlineResponse200339Ports) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Ports) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *InlineResponse200339Ports) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *InlineResponse200339Ports) SetPortId(v string) {
	o.PortId = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200339Ports) GetLastUpdatedAt() string {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Ports) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200339Ports) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *InlineResponse200339Ports) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse200339Ports) GetCdp() []InlineResponse200339Cdp {
	if o == nil || isNil(o.Cdp) {
		var ret []InlineResponse200339Cdp
		return ret
	}
	return o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Ports) GetCdpOk() ([]InlineResponse200339Cdp, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse200339Ports) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given []InlineResponse200339Cdp and assigns it to the Cdp field.
func (o *InlineResponse200339Ports) SetCdp(v []InlineResponse200339Cdp) {
	o.Cdp = v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse200339Ports) GetLldp() []InlineResponse200339Lldp {
	if o == nil || isNil(o.Lldp) {
		var ret []InlineResponse200339Lldp
		return ret
	}
	return o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Ports) GetLldpOk() ([]InlineResponse200339Lldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse200339Ports) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given []InlineResponse200339Lldp and assigns it to the Lldp field.
func (o *InlineResponse200339Ports) SetLldp(v []InlineResponse200339Lldp) {
	o.Lldp = v
}

func (o InlineResponse200339Ports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !isNil(o.Cdp) {
		toSerialize["cdp"] = o.Cdp
	}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339Ports struct {
	value *InlineResponse200339Ports
	isSet bool
}

func (v NullableInlineResponse200339Ports) Get() *InlineResponse200339Ports {
	return v.value
}

func (v *NullableInlineResponse200339Ports) Set(val *InlineResponse200339Ports) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339Ports) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339Ports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339Ports(val *InlineResponse200339Ports) *NullableInlineResponse200339Ports {
	return &NullableInlineResponse200339Ports{value: val, isSet: true}
}

func (v NullableInlineResponse200339Ports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339Ports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


