/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20027Lldp lldp information
type InlineResponse20027Lldp struct {
	// Device system name
	SystemName *string `json:"systemName,omitempty"`
	// ID for the port
	PortId *string `json:"portId,omitempty"`
	// Management IP address
	ManagementAddress *string `json:"managementAddress,omitempty"`
	// Source port
	SourcePort *string `json:"sourcePort,omitempty"`
}

// NewInlineResponse20027Lldp instantiates a new InlineResponse20027Lldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027Lldp() *InlineResponse20027Lldp {
	this := InlineResponse20027Lldp{}
	return &this
}

// NewInlineResponse20027LldpWithDefaults instantiates a new InlineResponse20027Lldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027LldpWithDefaults() *InlineResponse20027Lldp {
	this := InlineResponse20027Lldp{}
	return &this
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *InlineResponse20027Lldp) GetSystemName() string {
	if o == nil || isNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027Lldp) GetSystemNameOk() (*string, bool) {
	if o == nil || isNil(o.SystemName) {
    return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *InlineResponse20027Lldp) HasSystemName() bool {
	if o != nil && !isNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *InlineResponse20027Lldp) SetSystemName(v string) {
	o.SystemName = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *InlineResponse20027Lldp) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027Lldp) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *InlineResponse20027Lldp) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *InlineResponse20027Lldp) SetPortId(v string) {
	o.PortId = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *InlineResponse20027Lldp) GetManagementAddress() string {
	if o == nil || isNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027Lldp) GetManagementAddressOk() (*string, bool) {
	if o == nil || isNil(o.ManagementAddress) {
    return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *InlineResponse20027Lldp) HasManagementAddress() bool {
	if o != nil && !isNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *InlineResponse20027Lldp) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetSourcePort returns the SourcePort field value if set, zero value otherwise.
func (o *InlineResponse20027Lldp) GetSourcePort() string {
	if o == nil || isNil(o.SourcePort) {
		var ret string
		return ret
	}
	return *o.SourcePort
}

// GetSourcePortOk returns a tuple with the SourcePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027Lldp) GetSourcePortOk() (*string, bool) {
	if o == nil || isNil(o.SourcePort) {
    return nil, false
	}
	return o.SourcePort, true
}

// HasSourcePort returns a boolean if a field has been set.
func (o *InlineResponse20027Lldp) HasSourcePort() bool {
	if o != nil && !isNil(o.SourcePort) {
		return true
	}

	return false
}

// SetSourcePort gets a reference to the given string and assigns it to the SourcePort field.
func (o *InlineResponse20027Lldp) SetSourcePort(v string) {
	o.SourcePort = &v
}

func (o InlineResponse20027Lldp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.ManagementAddress) {
		toSerialize["managementAddress"] = o.ManagementAddress
	}
	if !isNil(o.SourcePort) {
		toSerialize["sourcePort"] = o.SourcePort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027Lldp struct {
	value *InlineResponse20027Lldp
	isSet bool
}

func (v NullableInlineResponse20027Lldp) Get() *InlineResponse20027Lldp {
	return v.value
}

func (v *NullableInlineResponse20027Lldp) Set(val *InlineResponse20027Lldp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027Lldp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027Lldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027Lldp(val *InlineResponse20027Lldp) *NullableInlineResponse20027Lldp {
	return &NullableInlineResponse20027Lldp{value: val, isSet: true}
}

func (v NullableInlineResponse20027Lldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027Lldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


