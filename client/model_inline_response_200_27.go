/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20027 struct for InlineResponse20027
type InlineResponse20027 struct {
	// The ID of the client
	Id *string `json:"id,omitempty"`
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// The IP address of the client
	Ip *string `json:"ip,omitempty"`
	// The IPv6 address of the client
	Ip6 *string `json:"ip6,omitempty"`
	// Short description of the client
	Description *string `json:"description,omitempty"`
	// Timestamp client was first seen in the network
	FirstSeen *int32 `json:"firstSeen,omitempty"`
	// Timestamp client was last seen in the network
	LastSeen *int32 `json:"lastSeen,omitempty"`
	// Manufacturer of the client
	Manufacturer *string `json:"manufacturer,omitempty"`
	// The operating system of the client
	Os *string `json:"os,omitempty"`
	// The username of the user of the client
	User *string `json:"user,omitempty"`
	// The name of the VLAN that the client is connected to
	Vlan *string `json:"vlan,omitempty"`
	// The name of the SSID that the client is connected to
	Ssid *string `json:"ssid,omitempty"`
	// The switch port that the client is connected to
	Switchport *string `json:"switchport,omitempty"`
	// Wireless capabilities of the client
	WirelessCapabilities *string `json:"wirelessCapabilities,omitempty"`
	// Status of SM for the client
	SmInstalled *bool `json:"smInstalled,omitempty"`
	// The MAC address of the node that the device was last connected to
	RecentDeviceMac *string `json:"recentDeviceMac,omitempty"`
	// The connection status of the client
	Status *string `json:"status,omitempty"`
	Usage *InlineResponse20027Usage `json:"usage,omitempty"`
	// Named VLAN of the client
	NamedVlan *string `json:"namedVlan,omitempty"`
	// The adaptive policy group of the client
	AdaptivePolicyGroup *string `json:"adaptivePolicyGroup,omitempty"`
	// Prediction of the client's device type
	DeviceTypePrediction *string `json:"deviceTypePrediction,omitempty"`
	// The serial of the node the device was last connected to
	RecentDeviceSerial *string `json:"recentDeviceSerial,omitempty"`
	// The name of the node the device was last connected to
	RecentDeviceName *string `json:"recentDeviceName,omitempty"`
	// Client's most recent connection type
	RecentDeviceConnection *string `json:"recentDeviceConnection,omitempty"`
	// Notes on the client
	Notes *string `json:"notes,omitempty"`
	// Local IPv6 address of the client
	Ip6Local *string `json:"ip6Local,omitempty"`
	// 802.1x group policy of the client
	GroupPolicy8021x *string `json:"groupPolicy8021x,omitempty"`
	// iPSK name of the client
	PskGroup *string `json:"pskGroup,omitempty"`
}

// NewInlineResponse20027 instantiates a new InlineResponse20027 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027WithDefaults() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20027) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20027) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20027) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20027) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20027) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20027) SetMac(v string) {
	o.Mac = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20027) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20027) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20027) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *InlineResponse20027) GetIp6() string {
	if o == nil || isNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetIp6Ok() (*string, bool) {
	if o == nil || isNil(o.Ip6) {
    return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *InlineResponse20027) HasIp6() bool {
	if o != nil && !isNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *InlineResponse20027) SetIp6(v string) {
	o.Ip6 = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20027) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20027) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20027) SetDescription(v string) {
	o.Description = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *InlineResponse20027) GetFirstSeen() int32 {
	if o == nil || isNil(o.FirstSeen) {
		var ret int32
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetFirstSeenOk() (*int32, bool) {
	if o == nil || isNil(o.FirstSeen) {
    return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *InlineResponse20027) HasFirstSeen() bool {
	if o != nil && !isNil(o.FirstSeen) {
		return true
	}

	return false
}

// SetFirstSeen gets a reference to the given int32 and assigns it to the FirstSeen field.
func (o *InlineResponse20027) SetFirstSeen(v int32) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *InlineResponse20027) GetLastSeen() int32 {
	if o == nil || isNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetLastSeenOk() (*int32, bool) {
	if o == nil || isNil(o.LastSeen) {
    return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *InlineResponse20027) HasLastSeen() bool {
	if o != nil && !isNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *InlineResponse20027) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *InlineResponse20027) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.Manufacturer) {
    return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InlineResponse20027) HasManufacturer() bool {
	if o != nil && !isNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *InlineResponse20027) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *InlineResponse20027) GetOs() string {
	if o == nil || isNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetOsOk() (*string, bool) {
	if o == nil || isNil(o.Os) {
    return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *InlineResponse20027) HasOs() bool {
	if o != nil && !isNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *InlineResponse20027) SetOs(v string) {
	o.Os = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *InlineResponse20027) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *InlineResponse20027) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *InlineResponse20027) SetUser(v string) {
	o.User = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20027) GetVlan() string {
	if o == nil || isNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetVlanOk() (*string, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20027) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *InlineResponse20027) SetVlan(v string) {
	o.Vlan = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSsid() string {
	if o == nil || isNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetSsidOk() (*string, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *InlineResponse20027) SetSsid(v string) {
	o.Ssid = &v
}

// GetSwitchport returns the Switchport field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSwitchport() string {
	if o == nil || isNil(o.Switchport) {
		var ret string
		return ret
	}
	return *o.Switchport
}

// GetSwitchportOk returns a tuple with the Switchport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetSwitchportOk() (*string, bool) {
	if o == nil || isNil(o.Switchport) {
    return nil, false
	}
	return o.Switchport, true
}

// HasSwitchport returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSwitchport() bool {
	if o != nil && !isNil(o.Switchport) {
		return true
	}

	return false
}

// SetSwitchport gets a reference to the given string and assigns it to the Switchport field.
func (o *InlineResponse20027) SetSwitchport(v string) {
	o.Switchport = &v
}

// GetWirelessCapabilities returns the WirelessCapabilities field value if set, zero value otherwise.
func (o *InlineResponse20027) GetWirelessCapabilities() string {
	if o == nil || isNil(o.WirelessCapabilities) {
		var ret string
		return ret
	}
	return *o.WirelessCapabilities
}

// GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetWirelessCapabilitiesOk() (*string, bool) {
	if o == nil || isNil(o.WirelessCapabilities) {
    return nil, false
	}
	return o.WirelessCapabilities, true
}

// HasWirelessCapabilities returns a boolean if a field has been set.
func (o *InlineResponse20027) HasWirelessCapabilities() bool {
	if o != nil && !isNil(o.WirelessCapabilities) {
		return true
	}

	return false
}

// SetWirelessCapabilities gets a reference to the given string and assigns it to the WirelessCapabilities field.
func (o *InlineResponse20027) SetWirelessCapabilities(v string) {
	o.WirelessCapabilities = &v
}

// GetSmInstalled returns the SmInstalled field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSmInstalled() bool {
	if o == nil || isNil(o.SmInstalled) {
		var ret bool
		return ret
	}
	return *o.SmInstalled
}

// GetSmInstalledOk returns a tuple with the SmInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetSmInstalledOk() (*bool, bool) {
	if o == nil || isNil(o.SmInstalled) {
    return nil, false
	}
	return o.SmInstalled, true
}

// HasSmInstalled returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSmInstalled() bool {
	if o != nil && !isNil(o.SmInstalled) {
		return true
	}

	return false
}

// SetSmInstalled gets a reference to the given bool and assigns it to the SmInstalled field.
func (o *InlineResponse20027) SetSmInstalled(v bool) {
	o.SmInstalled = &v
}

// GetRecentDeviceMac returns the RecentDeviceMac field value if set, zero value otherwise.
func (o *InlineResponse20027) GetRecentDeviceMac() string {
	if o == nil || isNil(o.RecentDeviceMac) {
		var ret string
		return ret
	}
	return *o.RecentDeviceMac
}

// GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetRecentDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceMac) {
    return nil, false
	}
	return o.RecentDeviceMac, true
}

// HasRecentDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse20027) HasRecentDeviceMac() bool {
	if o != nil && !isNil(o.RecentDeviceMac) {
		return true
	}

	return false
}

// SetRecentDeviceMac gets a reference to the given string and assigns it to the RecentDeviceMac field.
func (o *InlineResponse20027) SetRecentDeviceMac(v string) {
	o.RecentDeviceMac = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20027) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20027) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20027) SetStatus(v string) {
	o.Status = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse20027) GetUsage() InlineResponse20027Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse20027Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetUsageOk() (*InlineResponse20027Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse20027) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse20027Usage and assigns it to the Usage field.
func (o *InlineResponse20027) SetUsage(v InlineResponse20027Usage) {
	o.Usage = &v
}

// GetNamedVlan returns the NamedVlan field value if set, zero value otherwise.
func (o *InlineResponse20027) GetNamedVlan() string {
	if o == nil || isNil(o.NamedVlan) {
		var ret string
		return ret
	}
	return *o.NamedVlan
}

// GetNamedVlanOk returns a tuple with the NamedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetNamedVlanOk() (*string, bool) {
	if o == nil || isNil(o.NamedVlan) {
    return nil, false
	}
	return o.NamedVlan, true
}

// HasNamedVlan returns a boolean if a field has been set.
func (o *InlineResponse20027) HasNamedVlan() bool {
	if o != nil && !isNil(o.NamedVlan) {
		return true
	}

	return false
}

// SetNamedVlan gets a reference to the given string and assigns it to the NamedVlan field.
func (o *InlineResponse20027) SetNamedVlan(v string) {
	o.NamedVlan = &v
}

// GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field value if set, zero value otherwise.
func (o *InlineResponse20027) GetAdaptivePolicyGroup() string {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyGroup
}

// GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetAdaptivePolicyGroupOk() (*string, bool) {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
    return nil, false
	}
	return o.AdaptivePolicyGroup, true
}

// HasAdaptivePolicyGroup returns a boolean if a field has been set.
func (o *InlineResponse20027) HasAdaptivePolicyGroup() bool {
	if o != nil && !isNil(o.AdaptivePolicyGroup) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroup gets a reference to the given string and assigns it to the AdaptivePolicyGroup field.
func (o *InlineResponse20027) SetAdaptivePolicyGroup(v string) {
	o.AdaptivePolicyGroup = &v
}

// GetDeviceTypePrediction returns the DeviceTypePrediction field value if set, zero value otherwise.
func (o *InlineResponse20027) GetDeviceTypePrediction() string {
	if o == nil || isNil(o.DeviceTypePrediction) {
		var ret string
		return ret
	}
	return *o.DeviceTypePrediction
}

// GetDeviceTypePredictionOk returns a tuple with the DeviceTypePrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetDeviceTypePredictionOk() (*string, bool) {
	if o == nil || isNil(o.DeviceTypePrediction) {
    return nil, false
	}
	return o.DeviceTypePrediction, true
}

// HasDeviceTypePrediction returns a boolean if a field has been set.
func (o *InlineResponse20027) HasDeviceTypePrediction() bool {
	if o != nil && !isNil(o.DeviceTypePrediction) {
		return true
	}

	return false
}

// SetDeviceTypePrediction gets a reference to the given string and assigns it to the DeviceTypePrediction field.
func (o *InlineResponse20027) SetDeviceTypePrediction(v string) {
	o.DeviceTypePrediction = &v
}

// GetRecentDeviceSerial returns the RecentDeviceSerial field value if set, zero value otherwise.
func (o *InlineResponse20027) GetRecentDeviceSerial() string {
	if o == nil || isNil(o.RecentDeviceSerial) {
		var ret string
		return ret
	}
	return *o.RecentDeviceSerial
}

// GetRecentDeviceSerialOk returns a tuple with the RecentDeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetRecentDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceSerial) {
    return nil, false
	}
	return o.RecentDeviceSerial, true
}

// HasRecentDeviceSerial returns a boolean if a field has been set.
func (o *InlineResponse20027) HasRecentDeviceSerial() bool {
	if o != nil && !isNil(o.RecentDeviceSerial) {
		return true
	}

	return false
}

// SetRecentDeviceSerial gets a reference to the given string and assigns it to the RecentDeviceSerial field.
func (o *InlineResponse20027) SetRecentDeviceSerial(v string) {
	o.RecentDeviceSerial = &v
}

// GetRecentDeviceName returns the RecentDeviceName field value if set, zero value otherwise.
func (o *InlineResponse20027) GetRecentDeviceName() string {
	if o == nil || isNil(o.RecentDeviceName) {
		var ret string
		return ret
	}
	return *o.RecentDeviceName
}

// GetRecentDeviceNameOk returns a tuple with the RecentDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetRecentDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceName) {
    return nil, false
	}
	return o.RecentDeviceName, true
}

// HasRecentDeviceName returns a boolean if a field has been set.
func (o *InlineResponse20027) HasRecentDeviceName() bool {
	if o != nil && !isNil(o.RecentDeviceName) {
		return true
	}

	return false
}

// SetRecentDeviceName gets a reference to the given string and assigns it to the RecentDeviceName field.
func (o *InlineResponse20027) SetRecentDeviceName(v string) {
	o.RecentDeviceName = &v
}

// GetRecentDeviceConnection returns the RecentDeviceConnection field value if set, zero value otherwise.
func (o *InlineResponse20027) GetRecentDeviceConnection() string {
	if o == nil || isNil(o.RecentDeviceConnection) {
		var ret string
		return ret
	}
	return *o.RecentDeviceConnection
}

// GetRecentDeviceConnectionOk returns a tuple with the RecentDeviceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetRecentDeviceConnectionOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceConnection) {
    return nil, false
	}
	return o.RecentDeviceConnection, true
}

// HasRecentDeviceConnection returns a boolean if a field has been set.
func (o *InlineResponse20027) HasRecentDeviceConnection() bool {
	if o != nil && !isNil(o.RecentDeviceConnection) {
		return true
	}

	return false
}

// SetRecentDeviceConnection gets a reference to the given string and assigns it to the RecentDeviceConnection field.
func (o *InlineResponse20027) SetRecentDeviceConnection(v string) {
	o.RecentDeviceConnection = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse20027) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse20027) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse20027) SetNotes(v string) {
	o.Notes = &v
}

// GetIp6Local returns the Ip6Local field value if set, zero value otherwise.
func (o *InlineResponse20027) GetIp6Local() string {
	if o == nil || isNil(o.Ip6Local) {
		var ret string
		return ret
	}
	return *o.Ip6Local
}

// GetIp6LocalOk returns a tuple with the Ip6Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetIp6LocalOk() (*string, bool) {
	if o == nil || isNil(o.Ip6Local) {
    return nil, false
	}
	return o.Ip6Local, true
}

// HasIp6Local returns a boolean if a field has been set.
func (o *InlineResponse20027) HasIp6Local() bool {
	if o != nil && !isNil(o.Ip6Local) {
		return true
	}

	return false
}

// SetIp6Local gets a reference to the given string and assigns it to the Ip6Local field.
func (o *InlineResponse20027) SetIp6Local(v string) {
	o.Ip6Local = &v
}

// GetGroupPolicy8021x returns the GroupPolicy8021x field value if set, zero value otherwise.
func (o *InlineResponse20027) GetGroupPolicy8021x() string {
	if o == nil || isNil(o.GroupPolicy8021x) {
		var ret string
		return ret
	}
	return *o.GroupPolicy8021x
}

// GetGroupPolicy8021xOk returns a tuple with the GroupPolicy8021x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetGroupPolicy8021xOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicy8021x) {
    return nil, false
	}
	return o.GroupPolicy8021x, true
}

// HasGroupPolicy8021x returns a boolean if a field has been set.
func (o *InlineResponse20027) HasGroupPolicy8021x() bool {
	if o != nil && !isNil(o.GroupPolicy8021x) {
		return true
	}

	return false
}

// SetGroupPolicy8021x gets a reference to the given string and assigns it to the GroupPolicy8021x field.
func (o *InlineResponse20027) SetGroupPolicy8021x(v string) {
	o.GroupPolicy8021x = &v
}

// GetPskGroup returns the PskGroup field value if set, zero value otherwise.
func (o *InlineResponse20027) GetPskGroup() string {
	if o == nil || isNil(o.PskGroup) {
		var ret string
		return ret
	}
	return *o.PskGroup
}

// GetPskGroupOk returns a tuple with the PskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetPskGroupOk() (*string, bool) {
	if o == nil || isNil(o.PskGroup) {
    return nil, false
	}
	return o.PskGroup, true
}

// HasPskGroup returns a boolean if a field has been set.
func (o *InlineResponse20027) HasPskGroup() bool {
	if o != nil && !isNil(o.PskGroup) {
		return true
	}

	return false
}

// SetPskGroup gets a reference to the given string and assigns it to the PskGroup field.
func (o *InlineResponse20027) SetPskGroup(v string) {
	o.PskGroup = &v
}

func (o InlineResponse20027) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Ip6) {
		toSerialize["ip6"] = o.Ip6
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.FirstSeen) {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	if !isNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !isNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !isNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Switchport) {
		toSerialize["switchport"] = o.Switchport
	}
	if !isNil(o.WirelessCapabilities) {
		toSerialize["wirelessCapabilities"] = o.WirelessCapabilities
	}
	if !isNil(o.SmInstalled) {
		toSerialize["smInstalled"] = o.SmInstalled
	}
	if !isNil(o.RecentDeviceMac) {
		toSerialize["recentDeviceMac"] = o.RecentDeviceMac
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.NamedVlan) {
		toSerialize["namedVlan"] = o.NamedVlan
	}
	if !isNil(o.AdaptivePolicyGroup) {
		toSerialize["adaptivePolicyGroup"] = o.AdaptivePolicyGroup
	}
	if !isNil(o.DeviceTypePrediction) {
		toSerialize["deviceTypePrediction"] = o.DeviceTypePrediction
	}
	if !isNil(o.RecentDeviceSerial) {
		toSerialize["recentDeviceSerial"] = o.RecentDeviceSerial
	}
	if !isNil(o.RecentDeviceName) {
		toSerialize["recentDeviceName"] = o.RecentDeviceName
	}
	if !isNil(o.RecentDeviceConnection) {
		toSerialize["recentDeviceConnection"] = o.RecentDeviceConnection
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Ip6Local) {
		toSerialize["ip6Local"] = o.Ip6Local
	}
	if !isNil(o.GroupPolicy8021x) {
		toSerialize["groupPolicy8021x"] = o.GroupPolicy8021x
	}
	if !isNil(o.PskGroup) {
		toSerialize["pskGroup"] = o.PskGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027 struct {
	value *InlineResponse20027
	isSet bool
}

func (v NullableInlineResponse20027) Get() *InlineResponse20027 {
	return v.value
}

func (v *NullableInlineResponse20027) Set(val *InlineResponse20027) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027(val *InlineResponse20027) *NullableInlineResponse20027 {
	return &NullableInlineResponse20027{value: val, isSet: true}
}

func (v NullableInlineResponse20027) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


