/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
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
	UsageInKb *DevicesSerialSwitchPortsStatusesUsageInKb `json:"usageInKb,omitempty"`
	Cdp *DevicesSerialSwitchPortsStatusesCdp `json:"cdp,omitempty"`
	Lldp *DevicesSerialSwitchPortsStatusesLldp `json:"lldp,omitempty"`
	// The number of clients connected through this port.
	ClientCount *int32 `json:"clientCount,omitempty"`
	// How much power (in watt-hours) has been delivered by this port during the timespan.
	PowerUsageInWh *float32 `json:"powerUsageInWh,omitempty"`
	TrafficInKbps *DevicesSerialSwitchPortsStatusesTrafficInKbps `json:"trafficInKbps,omitempty"`
	SecurePort *DevicesSerialSwitchPortsStatusesSecurePort `json:"securePort,omitempty"`
}

// NewInlineResponse20036 instantiates a new InlineResponse20036 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036WithDefaults() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *InlineResponse20036) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *InlineResponse20036) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *InlineResponse20036) SetPortId(v string) {
	o.PortId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20036) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20036) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20036) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20036) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20036) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20036) SetStatus(v string) {
	o.Status = &v
}

// GetIsUplink returns the IsUplink field value if set, zero value otherwise.
func (o *InlineResponse20036) GetIsUplink() bool {
	if o == nil || isNil(o.IsUplink) {
		var ret bool
		return ret
	}
	return *o.IsUplink
}

// GetIsUplinkOk returns a tuple with the IsUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetIsUplinkOk() (*bool, bool) {
	if o == nil || isNil(o.IsUplink) {
    return nil, false
	}
	return o.IsUplink, true
}

// HasIsUplink returns a boolean if a field has been set.
func (o *InlineResponse20036) HasIsUplink() bool {
	if o != nil && !isNil(o.IsUplink) {
		return true
	}

	return false
}

// SetIsUplink gets a reference to the given bool and assigns it to the IsUplink field.
func (o *InlineResponse20036) SetIsUplink(v bool) {
	o.IsUplink = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse20036) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse20036) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse20036) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *InlineResponse20036) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
    return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *InlineResponse20036) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *InlineResponse20036) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *InlineResponse20036) GetSpeed() string {
	if o == nil || isNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetSpeedOk() (*string, bool) {
	if o == nil || isNil(o.Speed) {
    return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *InlineResponse20036) HasSpeed() bool {
	if o != nil && !isNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *InlineResponse20036) SetSpeed(v string) {
	o.Speed = &v
}

// GetDuplex returns the Duplex field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDuplex() string {
	if o == nil || isNil(o.Duplex) {
		var ret string
		return ret
	}
	return *o.Duplex
}

// GetDuplexOk returns a tuple with the Duplex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDuplexOk() (*string, bool) {
	if o == nil || isNil(o.Duplex) {
    return nil, false
	}
	return o.Duplex, true
}

// HasDuplex returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDuplex() bool {
	if o != nil && !isNil(o.Duplex) {
		return true
	}

	return false
}

// SetDuplex gets a reference to the given string and assigns it to the Duplex field.
func (o *InlineResponse20036) SetDuplex(v string) {
	o.Duplex = &v
}

// GetSpanningTree returns the SpanningTree field value if set, zero value otherwise.
func (o *InlineResponse20036) GetSpanningTree() DevicesSerialSwitchPortsStatusesSpanningTree {
	if o == nil || isNil(o.SpanningTree) {
		var ret DevicesSerialSwitchPortsStatusesSpanningTree
		return ret
	}
	return *o.SpanningTree
}

// GetSpanningTreeOk returns a tuple with the SpanningTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetSpanningTreeOk() (*DevicesSerialSwitchPortsStatusesSpanningTree, bool) {
	if o == nil || isNil(o.SpanningTree) {
    return nil, false
	}
	return o.SpanningTree, true
}

// HasSpanningTree returns a boolean if a field has been set.
func (o *InlineResponse20036) HasSpanningTree() bool {
	if o != nil && !isNil(o.SpanningTree) {
		return true
	}

	return false
}

// SetSpanningTree gets a reference to the given DevicesSerialSwitchPortsStatusesSpanningTree and assigns it to the SpanningTree field.
func (o *InlineResponse20036) SetSpanningTree(v DevicesSerialSwitchPortsStatusesSpanningTree) {
	o.SpanningTree = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *InlineResponse20036) GetPoe() DevicesSerialSwitchPortsStatusesPoe {
	if o == nil || isNil(o.Poe) {
		var ret DevicesSerialSwitchPortsStatusesPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetPoeOk() (*DevicesSerialSwitchPortsStatusesPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *InlineResponse20036) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given DevicesSerialSwitchPortsStatusesPoe and assigns it to the Poe field.
func (o *InlineResponse20036) SetPoe(v DevicesSerialSwitchPortsStatusesPoe) {
	o.Poe = &v
}

// GetUsageInKb returns the UsageInKb field value if set, zero value otherwise.
func (o *InlineResponse20036) GetUsageInKb() DevicesSerialSwitchPortsStatusesUsageInKb {
	if o == nil || isNil(o.UsageInKb) {
		var ret DevicesSerialSwitchPortsStatusesUsageInKb
		return ret
	}
	return *o.UsageInKb
}

// GetUsageInKbOk returns a tuple with the UsageInKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetUsageInKbOk() (*DevicesSerialSwitchPortsStatusesUsageInKb, bool) {
	if o == nil || isNil(o.UsageInKb) {
    return nil, false
	}
	return o.UsageInKb, true
}

// HasUsageInKb returns a boolean if a field has been set.
func (o *InlineResponse20036) HasUsageInKb() bool {
	if o != nil && !isNil(o.UsageInKb) {
		return true
	}

	return false
}

// SetUsageInKb gets a reference to the given DevicesSerialSwitchPortsStatusesUsageInKb and assigns it to the UsageInKb field.
func (o *InlineResponse20036) SetUsageInKb(v DevicesSerialSwitchPortsStatusesUsageInKb) {
	o.UsageInKb = &v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse20036) GetCdp() DevicesSerialSwitchPortsStatusesCdp {
	if o == nil || isNil(o.Cdp) {
		var ret DevicesSerialSwitchPortsStatusesCdp
		return ret
	}
	return *o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetCdpOk() (*DevicesSerialSwitchPortsStatusesCdp, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse20036) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given DevicesSerialSwitchPortsStatusesCdp and assigns it to the Cdp field.
func (o *InlineResponse20036) SetCdp(v DevicesSerialSwitchPortsStatusesCdp) {
	o.Cdp = &v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse20036) GetLldp() DevicesSerialSwitchPortsStatusesLldp {
	if o == nil || isNil(o.Lldp) {
		var ret DevicesSerialSwitchPortsStatusesLldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetLldpOk() (*DevicesSerialSwitchPortsStatusesLldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse20036) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given DevicesSerialSwitchPortsStatusesLldp and assigns it to the Lldp field.
func (o *InlineResponse20036) SetLldp(v DevicesSerialSwitchPortsStatusesLldp) {
	o.Lldp = &v
}

// GetClientCount returns the ClientCount field value if set, zero value otherwise.
func (o *InlineResponse20036) GetClientCount() int32 {
	if o == nil || isNil(o.ClientCount) {
		var ret int32
		return ret
	}
	return *o.ClientCount
}

// GetClientCountOk returns a tuple with the ClientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetClientCountOk() (*int32, bool) {
	if o == nil || isNil(o.ClientCount) {
    return nil, false
	}
	return o.ClientCount, true
}

// HasClientCount returns a boolean if a field has been set.
func (o *InlineResponse20036) HasClientCount() bool {
	if o != nil && !isNil(o.ClientCount) {
		return true
	}

	return false
}

// SetClientCount gets a reference to the given int32 and assigns it to the ClientCount field.
func (o *InlineResponse20036) SetClientCount(v int32) {
	o.ClientCount = &v
}

// GetPowerUsageInWh returns the PowerUsageInWh field value if set, zero value otherwise.
func (o *InlineResponse20036) GetPowerUsageInWh() float32 {
	if o == nil || isNil(o.PowerUsageInWh) {
		var ret float32
		return ret
	}
	return *o.PowerUsageInWh
}

// GetPowerUsageInWhOk returns a tuple with the PowerUsageInWh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetPowerUsageInWhOk() (*float32, bool) {
	if o == nil || isNil(o.PowerUsageInWh) {
    return nil, false
	}
	return o.PowerUsageInWh, true
}

// HasPowerUsageInWh returns a boolean if a field has been set.
func (o *InlineResponse20036) HasPowerUsageInWh() bool {
	if o != nil && !isNil(o.PowerUsageInWh) {
		return true
	}

	return false
}

// SetPowerUsageInWh gets a reference to the given float32 and assigns it to the PowerUsageInWh field.
func (o *InlineResponse20036) SetPowerUsageInWh(v float32) {
	o.PowerUsageInWh = &v
}

// GetTrafficInKbps returns the TrafficInKbps field value if set, zero value otherwise.
func (o *InlineResponse20036) GetTrafficInKbps() DevicesSerialSwitchPortsStatusesTrafficInKbps {
	if o == nil || isNil(o.TrafficInKbps) {
		var ret DevicesSerialSwitchPortsStatusesTrafficInKbps
		return ret
	}
	return *o.TrafficInKbps
}

// GetTrafficInKbpsOk returns a tuple with the TrafficInKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetTrafficInKbpsOk() (*DevicesSerialSwitchPortsStatusesTrafficInKbps, bool) {
	if o == nil || isNil(o.TrafficInKbps) {
    return nil, false
	}
	return o.TrafficInKbps, true
}

// HasTrafficInKbps returns a boolean if a field has been set.
func (o *InlineResponse20036) HasTrafficInKbps() bool {
	if o != nil && !isNil(o.TrafficInKbps) {
		return true
	}

	return false
}

// SetTrafficInKbps gets a reference to the given DevicesSerialSwitchPortsStatusesTrafficInKbps and assigns it to the TrafficInKbps field.
func (o *InlineResponse20036) SetTrafficInKbps(v DevicesSerialSwitchPortsStatusesTrafficInKbps) {
	o.TrafficInKbps = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *InlineResponse20036) GetSecurePort() DevicesSerialSwitchPortsStatusesSecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret DevicesSerialSwitchPortsStatusesSecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetSecurePortOk() (*DevicesSerialSwitchPortsStatusesSecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *InlineResponse20036) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given DevicesSerialSwitchPortsStatusesSecurePort and assigns it to the SecurePort field.
func (o *InlineResponse20036) SetSecurePort(v DevicesSerialSwitchPortsStatusesSecurePort) {
	o.SecurePort = &v
}

func (o InlineResponse20036) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.UsageInKb) {
		toSerialize["usageInKb"] = o.UsageInKb
	}
	if !isNil(o.Cdp) {
		toSerialize["cdp"] = o.Cdp
	}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !isNil(o.ClientCount) {
		toSerialize["clientCount"] = o.ClientCount
	}
	if !isNil(o.PowerUsageInWh) {
		toSerialize["powerUsageInWh"] = o.PowerUsageInWh
	}
	if !isNil(o.TrafficInKbps) {
		toSerialize["trafficInKbps"] = o.TrafficInKbps
	}
	if !isNil(o.SecurePort) {
		toSerialize["securePort"] = o.SecurePort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036 struct {
	value *InlineResponse20036
	isSet bool
}

func (v NullableInlineResponse20036) Get() *InlineResponse20036 {
	return v.value
}

func (v *NullableInlineResponse20036) Set(val *InlineResponse20036) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036(val *InlineResponse20036) *NullableInlineResponse20036 {
	return &NullableInlineResponse20036{value: val, isSet: true}
}

func (v NullableInlineResponse20036) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


