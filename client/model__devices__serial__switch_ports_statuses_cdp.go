/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesCdp The Cisco Discovery Protocol (CDP) information of the connected device.
type DevicesSerialSwitchPortsStatusesCdp struct {
	// The system name.
	SystemName *string `json:"systemName,omitempty"`
	// Identifies the hardware platform of the device.
	Platform *string `json:"platform,omitempty"`
	// Identifies the device name.
	DeviceId *string `json:"deviceId,omitempty"`
	// Identifies the port from which the CDP packet was sent.
	PortId *string `json:"portId,omitempty"`
	// Indicates, per interface, the assumed VLAN for untagged packets on the interface.
	NativeVlan *int32 `json:"nativeVlan,omitempty"`
	// Contains network addresses of both receiving and sending devices.
	Address *string `json:"address,omitempty"`
	// The device's management IP.
	ManagementAddress *string `json:"managementAddress,omitempty"`
	// Contains the device software release information.
	Version *string `json:"version,omitempty"`
	// Advertises the configured VLAN Trunking Protocl (VTP)-management-domain name of the system.
	VtpManagementDomain *string `json:"vtpManagementDomain,omitempty"`
	// Identifies the device type, which indicates the functional capabilities of the device.
	Capabilities *string `json:"capabilities,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesCdp instantiates a new DevicesSerialSwitchPortsStatusesCdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesCdp() *DevicesSerialSwitchPortsStatusesCdp {
	this := DevicesSerialSwitchPortsStatusesCdp{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesCdpWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesCdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesCdpWithDefaults() *DevicesSerialSwitchPortsStatusesCdp {
	this := DevicesSerialSwitchPortsStatusesCdp{}
	return &this
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetSystemName() string {
	if o == nil || isNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetSystemNameOk() (*string, bool) {
	if o == nil || isNil(o.SystemName) {
    return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasSystemName() bool {
	if o != nil && !isNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetSystemName(v string) {
	o.SystemName = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetPlatform() string {
	if o == nil || isNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetPlatformOk() (*string, bool) {
	if o == nil || isNil(o.Platform) {
    return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasPlatform() bool {
	if o != nil && !isNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetPlatform(v string) {
	o.Platform = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetPortId(v string) {
	o.PortId = &v
}

// GetNativeVlan returns the NativeVlan field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetNativeVlan() int32 {
	if o == nil || isNil(o.NativeVlan) {
		var ret int32
		return ret
	}
	return *o.NativeVlan
}

// GetNativeVlanOk returns a tuple with the NativeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetNativeVlanOk() (*int32, bool) {
	if o == nil || isNil(o.NativeVlan) {
    return nil, false
	}
	return o.NativeVlan, true
}

// HasNativeVlan returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasNativeVlan() bool {
	if o != nil && !isNil(o.NativeVlan) {
		return true
	}

	return false
}

// SetNativeVlan gets a reference to the given int32 and assigns it to the NativeVlan field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetNativeVlan(v int32) {
	o.NativeVlan = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetAddress(v string) {
	o.Address = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetManagementAddress() string {
	if o == nil || isNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetManagementAddressOk() (*string, bool) {
	if o == nil || isNil(o.ManagementAddress) {
    return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasManagementAddress() bool {
	if o != nil && !isNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetVersion(v string) {
	o.Version = &v
}

// GetVtpManagementDomain returns the VtpManagementDomain field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetVtpManagementDomain() string {
	if o == nil || isNil(o.VtpManagementDomain) {
		var ret string
		return ret
	}
	return *o.VtpManagementDomain
}

// GetVtpManagementDomainOk returns a tuple with the VtpManagementDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetVtpManagementDomainOk() (*string, bool) {
	if o == nil || isNil(o.VtpManagementDomain) {
    return nil, false
	}
	return o.VtpManagementDomain, true
}

// HasVtpManagementDomain returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasVtpManagementDomain() bool {
	if o != nil && !isNil(o.VtpManagementDomain) {
		return true
	}

	return false
}

// SetVtpManagementDomain gets a reference to the given string and assigns it to the VtpManagementDomain field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetVtpManagementDomain(v string) {
	o.VtpManagementDomain = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetCapabilities() string {
	if o == nil || isNil(o.Capabilities) {
		var ret string
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) GetCapabilitiesOk() (*string, bool) {
	if o == nil || isNil(o.Capabilities) {
    return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesCdp) HasCapabilities() bool {
	if o != nil && !isNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given string and assigns it to the Capabilities field.
func (o *DevicesSerialSwitchPortsStatusesCdp) SetCapabilities(v string) {
	o.Capabilities = &v
}

func (o DevicesSerialSwitchPortsStatusesCdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !isNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.NativeVlan) {
		toSerialize["nativeVlan"] = o.NativeVlan
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.ManagementAddress) {
		toSerialize["managementAddress"] = o.ManagementAddress
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.VtpManagementDomain) {
		toSerialize["vtpManagementDomain"] = o.VtpManagementDomain
	}
	if !isNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesCdp struct {
	value *DevicesSerialSwitchPortsStatusesCdp
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesCdp) Get() *DevicesSerialSwitchPortsStatusesCdp {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesCdp) Set(val *DevicesSerialSwitchPortsStatusesCdp) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesCdp) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesCdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesCdp(val *DevicesSerialSwitchPortsStatusesCdp) *NullableDevicesSerialSwitchPortsStatusesCdp {
	return &NullableDevicesSerialSwitchPortsStatusesCdp{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesCdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesCdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


