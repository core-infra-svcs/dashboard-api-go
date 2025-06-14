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

// DevicesSerialSwitchPortsStatusesLldp The Link Layer Discovery Protocol (LLDP) information of the connected device.
type DevicesSerialSwitchPortsStatusesLldp struct {
	// The device's system name.
	SystemName *string `json:"systemName,omitempty"`
	// The device's system description.
	SystemDescription *string `json:"systemDescription,omitempty"`
	// The device's chassis ID.
	ChassisId *string `json:"chassisId,omitempty"`
	// Identifies the port from which the LLDP packet was sent
	PortId *string `json:"portId,omitempty"`
	// The device's management VLAN.
	ManagementVlan *int32 `json:"managementVlan,omitempty"`
	// The port's VLAN.
	PortVlan *int32 `json:"portVlan,omitempty"`
	// The device's management IP.
	ManagementAddress *string `json:"managementAddress,omitempty"`
	// Description of the port from which the LLDP packet was sent.
	PortDescription *string `json:"portDescription,omitempty"`
	// Identifies the device type, which indicates the functional capabilities of the device.
	SystemCapabilities *string `json:"systemCapabilities,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesLldp instantiates a new DevicesSerialSwitchPortsStatusesLldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesLldp() *DevicesSerialSwitchPortsStatusesLldp {
	this := DevicesSerialSwitchPortsStatusesLldp{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesLldpWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesLldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesLldpWithDefaults() *DevicesSerialSwitchPortsStatusesLldp {
	this := DevicesSerialSwitchPortsStatusesLldp{}
	return &this
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemName() string {
	if o == nil || isNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemNameOk() (*string, bool) {
	if o == nil || isNil(o.SystemName) {
    return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemName() bool {
	if o != nil && !isNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemName(v string) {
	o.SystemName = &v
}

// GetSystemDescription returns the SystemDescription field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemDescription() string {
	if o == nil || isNil(o.SystemDescription) {
		var ret string
		return ret
	}
	return *o.SystemDescription
}

// GetSystemDescriptionOk returns a tuple with the SystemDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.SystemDescription) {
    return nil, false
	}
	return o.SystemDescription, true
}

// HasSystemDescription returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemDescription() bool {
	if o != nil && !isNil(o.SystemDescription) {
		return true
	}

	return false
}

// SetSystemDescription gets a reference to the given string and assigns it to the SystemDescription field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemDescription(v string) {
	o.SystemDescription = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetChassisId() string {
	if o == nil || isNil(o.ChassisId) {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetChassisIdOk() (*string, bool) {
	if o == nil || isNil(o.ChassisId) {
    return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasChassisId() bool {
	if o != nil && !isNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortId(v string) {
	o.PortId = &v
}

// GetManagementVlan returns the ManagementVlan field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementVlan() int32 {
	if o == nil || isNil(o.ManagementVlan) {
		var ret int32
		return ret
	}
	return *o.ManagementVlan
}

// GetManagementVlanOk returns a tuple with the ManagementVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementVlanOk() (*int32, bool) {
	if o == nil || isNil(o.ManagementVlan) {
    return nil, false
	}
	return o.ManagementVlan, true
}

// HasManagementVlan returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasManagementVlan() bool {
	if o != nil && !isNil(o.ManagementVlan) {
		return true
	}

	return false
}

// SetManagementVlan gets a reference to the given int32 and assigns it to the ManagementVlan field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetManagementVlan(v int32) {
	o.ManagementVlan = &v
}

// GetPortVlan returns the PortVlan field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortVlan() int32 {
	if o == nil || isNil(o.PortVlan) {
		var ret int32
		return ret
	}
	return *o.PortVlan
}

// GetPortVlanOk returns a tuple with the PortVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortVlanOk() (*int32, bool) {
	if o == nil || isNil(o.PortVlan) {
    return nil, false
	}
	return o.PortVlan, true
}

// HasPortVlan returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortVlan() bool {
	if o != nil && !isNil(o.PortVlan) {
		return true
	}

	return false
}

// SetPortVlan gets a reference to the given int32 and assigns it to the PortVlan field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortVlan(v int32) {
	o.PortVlan = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementAddress() string {
	if o == nil || isNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementAddressOk() (*string, bool) {
	if o == nil || isNil(o.ManagementAddress) {
    return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasManagementAddress() bool {
	if o != nil && !isNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPortDescription returns the PortDescription field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortDescription() string {
	if o == nil || isNil(o.PortDescription) {
		var ret string
		return ret
	}
	return *o.PortDescription
}

// GetPortDescriptionOk returns a tuple with the PortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.PortDescription) {
    return nil, false
	}
	return o.PortDescription, true
}

// HasPortDescription returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortDescription() bool {
	if o != nil && !isNil(o.PortDescription) {
		return true
	}

	return false
}

// SetPortDescription gets a reference to the given string and assigns it to the PortDescription field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortDescription(v string) {
	o.PortDescription = &v
}

// GetSystemCapabilities returns the SystemCapabilities field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemCapabilities() string {
	if o == nil || isNil(o.SystemCapabilities) {
		var ret string
		return ret
	}
	return *o.SystemCapabilities
}

// GetSystemCapabilitiesOk returns a tuple with the SystemCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemCapabilitiesOk() (*string, bool) {
	if o == nil || isNil(o.SystemCapabilities) {
    return nil, false
	}
	return o.SystemCapabilities, true
}

// HasSystemCapabilities returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemCapabilities() bool {
	if o != nil && !isNil(o.SystemCapabilities) {
		return true
	}

	return false
}

// SetSystemCapabilities gets a reference to the given string and assigns it to the SystemCapabilities field.
func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemCapabilities(v string) {
	o.SystemCapabilities = &v
}

func (o DevicesSerialSwitchPortsStatusesLldp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !isNil(o.SystemDescription) {
		toSerialize["systemDescription"] = o.SystemDescription
	}
	if !isNil(o.ChassisId) {
		toSerialize["chassisId"] = o.ChassisId
	}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.ManagementVlan) {
		toSerialize["managementVlan"] = o.ManagementVlan
	}
	if !isNil(o.PortVlan) {
		toSerialize["portVlan"] = o.PortVlan
	}
	if !isNil(o.ManagementAddress) {
		toSerialize["managementAddress"] = o.ManagementAddress
	}
	if !isNil(o.PortDescription) {
		toSerialize["portDescription"] = o.PortDescription
	}
	if !isNil(o.SystemCapabilities) {
		toSerialize["systemCapabilities"] = o.SystemCapabilities
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesLldp struct {
	value *DevicesSerialSwitchPortsStatusesLldp
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesLldp) Get() *DevicesSerialSwitchPortsStatusesLldp {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesLldp) Set(val *DevicesSerialSwitchPortsStatusesLldp) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesLldp) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesLldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesLldp(val *DevicesSerialSwitchPortsStatusesLldp) *NullableDevicesSerialSwitchPortsStatusesLldp {
	return &NullableDevicesSerialSwitchPortsStatusesLldp{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesLldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesLldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


