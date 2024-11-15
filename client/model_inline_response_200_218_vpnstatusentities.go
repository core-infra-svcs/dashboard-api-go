/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200218Vpnstatusentities struct for InlineResponse200218Vpnstatusentities
type InlineResponse200218Vpnstatusentities struct {
	// Network Id
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	NetworkName *string `json:"networkName,omitempty"`
	// Serial number of the device
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// Device Status
	DeviceStatus *string `json:"deviceStatus,omitempty"`
	// List of Uplink Information
	Uplinks []InlineResponse200218Uplinks `json:"uplinks,omitempty"`
	// VPN Mode
	VpnMode *string `json:"vpnMode,omitempty"`
	// List of Exported Subnets
	ExportedSubnets []InlineResponse200218ExportedSubnets `json:"exportedSubnets,omitempty"`
	// Meraki VPN Peers
	MerakiVpnPeers []InlineResponse200218MerakiVpnPeers `json:"merakiVpnPeers,omitempty"`
	// Third Party VPN Peers
	ThirdPartyVpnPeers []InlineResponse200218ThirdPartyVpnPeers `json:"thirdPartyVpnPeers,omitempty"`
}

// NewInlineResponse200218Vpnstatusentities instantiates a new InlineResponse200218Vpnstatusentities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200218Vpnstatusentities() *InlineResponse200218Vpnstatusentities {
	this := InlineResponse200218Vpnstatusentities{}
	return &this
}

// NewInlineResponse200218VpnstatusentitiesWithDefaults instantiates a new InlineResponse200218Vpnstatusentities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200218VpnstatusentitiesWithDefaults() *InlineResponse200218Vpnstatusentities {
	this := InlineResponse200218Vpnstatusentities{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200218Vpnstatusentities) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetNetworkName() string {
	if o == nil || isNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetNetworkNameOk() (*string, bool) {
	if o == nil || isNil(o.NetworkName) {
    return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasNetworkName() bool {
	if o != nil && !isNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *InlineResponse200218Vpnstatusentities) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineResponse200218Vpnstatusentities) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetDeviceStatus() string {
	if o == nil || isNil(o.DeviceStatus) {
		var ret string
		return ret
	}
	return *o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetDeviceStatusOk() (*string, bool) {
	if o == nil || isNil(o.DeviceStatus) {
    return nil, false
	}
	return o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasDeviceStatus() bool {
	if o != nil && !isNil(o.DeviceStatus) {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given string and assigns it to the DeviceStatus field.
func (o *InlineResponse200218Vpnstatusentities) SetDeviceStatus(v string) {
	o.DeviceStatus = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetUplinks() []InlineResponse200218Uplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []InlineResponse200218Uplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetUplinksOk() ([]InlineResponse200218Uplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []InlineResponse200218Uplinks and assigns it to the Uplinks field.
func (o *InlineResponse200218Vpnstatusentities) SetUplinks(v []InlineResponse200218Uplinks) {
	o.Uplinks = v
}

// GetVpnMode returns the VpnMode field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetVpnMode() string {
	if o == nil || isNil(o.VpnMode) {
		var ret string
		return ret
	}
	return *o.VpnMode
}

// GetVpnModeOk returns a tuple with the VpnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetVpnModeOk() (*string, bool) {
	if o == nil || isNil(o.VpnMode) {
    return nil, false
	}
	return o.VpnMode, true
}

// HasVpnMode returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasVpnMode() bool {
	if o != nil && !isNil(o.VpnMode) {
		return true
	}

	return false
}

// SetVpnMode gets a reference to the given string and assigns it to the VpnMode field.
func (o *InlineResponse200218Vpnstatusentities) SetVpnMode(v string) {
	o.VpnMode = &v
}

// GetExportedSubnets returns the ExportedSubnets field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetExportedSubnets() []InlineResponse200218ExportedSubnets {
	if o == nil || isNil(o.ExportedSubnets) {
		var ret []InlineResponse200218ExportedSubnets
		return ret
	}
	return o.ExportedSubnets
}

// GetExportedSubnetsOk returns a tuple with the ExportedSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetExportedSubnetsOk() ([]InlineResponse200218ExportedSubnets, bool) {
	if o == nil || isNil(o.ExportedSubnets) {
    return nil, false
	}
	return o.ExportedSubnets, true
}

// HasExportedSubnets returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasExportedSubnets() bool {
	if o != nil && !isNil(o.ExportedSubnets) {
		return true
	}

	return false
}

// SetExportedSubnets gets a reference to the given []InlineResponse200218ExportedSubnets and assigns it to the ExportedSubnets field.
func (o *InlineResponse200218Vpnstatusentities) SetExportedSubnets(v []InlineResponse200218ExportedSubnets) {
	o.ExportedSubnets = v
}

// GetMerakiVpnPeers returns the MerakiVpnPeers field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetMerakiVpnPeers() []InlineResponse200218MerakiVpnPeers {
	if o == nil || isNil(o.MerakiVpnPeers) {
		var ret []InlineResponse200218MerakiVpnPeers
		return ret
	}
	return o.MerakiVpnPeers
}

// GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetMerakiVpnPeersOk() ([]InlineResponse200218MerakiVpnPeers, bool) {
	if o == nil || isNil(o.MerakiVpnPeers) {
    return nil, false
	}
	return o.MerakiVpnPeers, true
}

// HasMerakiVpnPeers returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasMerakiVpnPeers() bool {
	if o != nil && !isNil(o.MerakiVpnPeers) {
		return true
	}

	return false
}

// SetMerakiVpnPeers gets a reference to the given []InlineResponse200218MerakiVpnPeers and assigns it to the MerakiVpnPeers field.
func (o *InlineResponse200218Vpnstatusentities) SetMerakiVpnPeers(v []InlineResponse200218MerakiVpnPeers) {
	o.MerakiVpnPeers = v
}

// GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field value if set, zero value otherwise.
func (o *InlineResponse200218Vpnstatusentities) GetThirdPartyVpnPeers() []InlineResponse200218ThirdPartyVpnPeers {
	if o == nil || isNil(o.ThirdPartyVpnPeers) {
		var ret []InlineResponse200218ThirdPartyVpnPeers
		return ret
	}
	return o.ThirdPartyVpnPeers
}

// GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Vpnstatusentities) GetThirdPartyVpnPeersOk() ([]InlineResponse200218ThirdPartyVpnPeers, bool) {
	if o == nil || isNil(o.ThirdPartyVpnPeers) {
    return nil, false
	}
	return o.ThirdPartyVpnPeers, true
}

// HasThirdPartyVpnPeers returns a boolean if a field has been set.
func (o *InlineResponse200218Vpnstatusentities) HasThirdPartyVpnPeers() bool {
	if o != nil && !isNil(o.ThirdPartyVpnPeers) {
		return true
	}

	return false
}

// SetThirdPartyVpnPeers gets a reference to the given []InlineResponse200218ThirdPartyVpnPeers and assigns it to the ThirdPartyVpnPeers field.
func (o *InlineResponse200218Vpnstatusentities) SetThirdPartyVpnPeers(v []InlineResponse200218ThirdPartyVpnPeers) {
	o.ThirdPartyVpnPeers = v
}

func (o InlineResponse200218Vpnstatusentities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !isNil(o.DeviceStatus) {
		toSerialize["deviceStatus"] = o.DeviceStatus
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	if !isNil(o.VpnMode) {
		toSerialize["vpnMode"] = o.VpnMode
	}
	if !isNil(o.ExportedSubnets) {
		toSerialize["exportedSubnets"] = o.ExportedSubnets
	}
	if !isNil(o.MerakiVpnPeers) {
		toSerialize["merakiVpnPeers"] = o.MerakiVpnPeers
	}
	if !isNil(o.ThirdPartyVpnPeers) {
		toSerialize["thirdPartyVpnPeers"] = o.ThirdPartyVpnPeers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200218Vpnstatusentities struct {
	value *InlineResponse200218Vpnstatusentities
	isSet bool
}

func (v NullableInlineResponse200218Vpnstatusentities) Get() *InlineResponse200218Vpnstatusentities {
	return v.value
}

func (v *NullableInlineResponse200218Vpnstatusentities) Set(val *InlineResponse200218Vpnstatusentities) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200218Vpnstatusentities) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200218Vpnstatusentities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200218Vpnstatusentities(val *InlineResponse200218Vpnstatusentities) *NullableInlineResponse200218Vpnstatusentities {
	return &NullableInlineResponse200218Vpnstatusentities{value: val, isSet: true}
}

func (v NullableInlineResponse200218Vpnstatusentities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200218Vpnstatusentities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

