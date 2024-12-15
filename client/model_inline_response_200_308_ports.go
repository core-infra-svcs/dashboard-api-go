/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200308Ports struct for InlineResponse200308Ports
type InlineResponse200308Ports struct {
	// The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	PortId *string `json:"portId,omitempty"`
	// Whether the port is configured to be enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The current connection status of the port.
	Status *string `json:"status,omitempty"`
	// Whether the port is the switch's uplink.
	IsUplink *bool `json:"isUplink,omitempty"`
	// All errors present on the port.
	Errors []string `json:"errors,omitempty"`
	// All warnings present on the port.
	Warnings []string `json:"warnings,omitempty"`
	// The current data transfer rate which the port is operating at.
	Speed *string `json:"speed,omitempty"`
	// The current duplex of a connected port.
	Duplex *string `json:"duplex,omitempty"`
	SpanningTree *DevicesSerialSwitchPortsStatusesSpanningTree `json:"spanningTree,omitempty"`
	Poe *DevicesSerialSwitchPortsStatusesPoe `json:"poe,omitempty"`
	SecurePort *InlineResponse200308SecurePort `json:"securePort,omitempty"`
}

// NewInlineResponse200308Ports instantiates a new InlineResponse200308Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308Ports() *InlineResponse200308Ports {
	this := InlineResponse200308Ports{}
	return &this
}

// NewInlineResponse200308PortsWithDefaults instantiates a new InlineResponse200308Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308PortsWithDefaults() *InlineResponse200308Ports {
	this := InlineResponse200308Ports{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *InlineResponse200308Ports) SetPortId(v string) {
	o.PortId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200308Ports) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200308Ports) SetStatus(v string) {
	o.Status = &v
}

// GetIsUplink returns the IsUplink field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetIsUplink() bool {
	if o == nil || isNil(o.IsUplink) {
		var ret bool
		return ret
	}
	return *o.IsUplink
}

// GetIsUplinkOk returns a tuple with the IsUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetIsUplinkOk() (*bool, bool) {
	if o == nil || isNil(o.IsUplink) {
    return nil, false
	}
	return o.IsUplink, true
}

// HasIsUplink returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasIsUplink() bool {
	if o != nil && !isNil(o.IsUplink) {
		return true
	}

	return false
}

// SetIsUplink gets a reference to the given bool and assigns it to the IsUplink field.
func (o *InlineResponse200308Ports) SetIsUplink(v bool) {
	o.IsUplink = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse200308Ports) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
    return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *InlineResponse200308Ports) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetSpeed() string {
	if o == nil || isNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetSpeedOk() (*string, bool) {
	if o == nil || isNil(o.Speed) {
    return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasSpeed() bool {
	if o != nil && !isNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *InlineResponse200308Ports) SetSpeed(v string) {
	o.Speed = &v
}

// GetDuplex returns the Duplex field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetDuplex() string {
	if o == nil || isNil(o.Duplex) {
		var ret string
		return ret
	}
	return *o.Duplex
}

// GetDuplexOk returns a tuple with the Duplex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetDuplexOk() (*string, bool) {
	if o == nil || isNil(o.Duplex) {
    return nil, false
	}
	return o.Duplex, true
}

// HasDuplex returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasDuplex() bool {
	if o != nil && !isNil(o.Duplex) {
		return true
	}

	return false
}

// SetDuplex gets a reference to the given string and assigns it to the Duplex field.
func (o *InlineResponse200308Ports) SetDuplex(v string) {
	o.Duplex = &v
}

// GetSpanningTree returns the SpanningTree field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetSpanningTree() DevicesSerialSwitchPortsStatusesSpanningTree {
	if o == nil || isNil(o.SpanningTree) {
		var ret DevicesSerialSwitchPortsStatusesSpanningTree
		return ret
	}
	return *o.SpanningTree
}

// GetSpanningTreeOk returns a tuple with the SpanningTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetSpanningTreeOk() (*DevicesSerialSwitchPortsStatusesSpanningTree, bool) {
	if o == nil || isNil(o.SpanningTree) {
    return nil, false
	}
	return o.SpanningTree, true
}

// HasSpanningTree returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasSpanningTree() bool {
	if o != nil && !isNil(o.SpanningTree) {
		return true
	}

	return false
}

// SetSpanningTree gets a reference to the given DevicesSerialSwitchPortsStatusesSpanningTree and assigns it to the SpanningTree field.
func (o *InlineResponse200308Ports) SetSpanningTree(v DevicesSerialSwitchPortsStatusesSpanningTree) {
	o.SpanningTree = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetPoe() DevicesSerialSwitchPortsStatusesPoe {
	if o == nil || isNil(o.Poe) {
		var ret DevicesSerialSwitchPortsStatusesPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetPoeOk() (*DevicesSerialSwitchPortsStatusesPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given DevicesSerialSwitchPortsStatusesPoe and assigns it to the Poe field.
func (o *InlineResponse200308Ports) SetPoe(v DevicesSerialSwitchPortsStatusesPoe) {
	o.Poe = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *InlineResponse200308Ports) GetSecurePort() InlineResponse200308SecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret InlineResponse200308SecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308Ports) GetSecurePortOk() (*InlineResponse200308SecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *InlineResponse200308Ports) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given InlineResponse200308SecurePort and assigns it to the SecurePort field.
func (o *InlineResponse200308Ports) SetSecurePort(v InlineResponse200308SecurePort) {
	o.SecurePort = &v
}

func (o InlineResponse200308Ports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.IsUplink) {
		toSerialize["isUplink"] = o.IsUplink
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !isNil(o.Duplex) {
		toSerialize["duplex"] = o.Duplex
	}
	if !isNil(o.SpanningTree) {
		toSerialize["spanningTree"] = o.SpanningTree
	}
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	if !isNil(o.SecurePort) {
		toSerialize["securePort"] = o.SecurePort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308Ports struct {
	value *InlineResponse200308Ports
	isSet bool
}

func (v NullableInlineResponse200308Ports) Get() *InlineResponse200308Ports {
	return v.value
}

func (v *NullableInlineResponse200308Ports) Set(val *InlineResponse200308Ports) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308Ports) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308Ports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308Ports(val *InlineResponse200308Ports) *NullableInlineResponse200308Ports {
	return &NullableInlineResponse200308Ports{value: val, isSet: true}
}

func (v NullableInlineResponse200308Ports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308Ports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


