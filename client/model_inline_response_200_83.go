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

// InlineResponse20083 struct for InlineResponse20083
type InlineResponse20083 struct {
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
	// Client's most recent connection type
	RecentDeviceConnection *string `json:"recentDeviceConnection,omitempty"`
	// VPN connections associated with the client
	ClientVpnConnections []InlineResponse20083ClientVpnConnections `json:"clientVpnConnections,omitempty"`
	// The link layer discover protocol settings for the client
	Lldp [][]string `json:"lldp,omitempty"`
	// The Cisco discover protocol settings for the client
	Cdp [][]string `json:"cdp,omitempty"`
	// The connection status of the client
	Status *string `json:"status,omitempty"`
	// The notes associated with the client
	Notes *string `json:"notes,omitempty"`
}

// NewInlineResponse20083 instantiates a new InlineResponse20083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20083() *InlineResponse20083 {
	this := InlineResponse20083{}
	return &this
}

// NewInlineResponse20083WithDefaults instantiates a new InlineResponse20083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20083WithDefaults() *InlineResponse20083 {
	this := InlineResponse20083{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20083) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20083) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20083) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20083) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20083) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20083) SetMac(v string) {
	o.Mac = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20083) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20083) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20083) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *InlineResponse20083) GetIp6() string {
	if o == nil || isNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetIp6Ok() (*string, bool) {
	if o == nil || isNil(o.Ip6) {
    return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *InlineResponse20083) HasIp6() bool {
	if o != nil && !isNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *InlineResponse20083) SetIp6(v string) {
	o.Ip6 = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20083) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20083) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20083) SetDescription(v string) {
	o.Description = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *InlineResponse20083) GetFirstSeen() int32 {
	if o == nil || isNil(o.FirstSeen) {
		var ret int32
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetFirstSeenOk() (*int32, bool) {
	if o == nil || isNil(o.FirstSeen) {
    return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *InlineResponse20083) HasFirstSeen() bool {
	if o != nil && !isNil(o.FirstSeen) {
		return true
	}

	return false
}

// SetFirstSeen gets a reference to the given int32 and assigns it to the FirstSeen field.
func (o *InlineResponse20083) SetFirstSeen(v int32) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *InlineResponse20083) GetLastSeen() int32 {
	if o == nil || isNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetLastSeenOk() (*int32, bool) {
	if o == nil || isNil(o.LastSeen) {
    return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *InlineResponse20083) HasLastSeen() bool {
	if o != nil && !isNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *InlineResponse20083) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *InlineResponse20083) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.Manufacturer) {
    return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InlineResponse20083) HasManufacturer() bool {
	if o != nil && !isNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *InlineResponse20083) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *InlineResponse20083) GetOs() string {
	if o == nil || isNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetOsOk() (*string, bool) {
	if o == nil || isNil(o.Os) {
    return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *InlineResponse20083) HasOs() bool {
	if o != nil && !isNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *InlineResponse20083) SetOs(v string) {
	o.Os = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *InlineResponse20083) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *InlineResponse20083) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *InlineResponse20083) SetUser(v string) {
	o.User = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20083) GetVlan() string {
	if o == nil || isNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetVlanOk() (*string, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20083) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *InlineResponse20083) SetVlan(v string) {
	o.Vlan = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse20083) GetSsid() string {
	if o == nil || isNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetSsidOk() (*string, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse20083) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *InlineResponse20083) SetSsid(v string) {
	o.Ssid = &v
}

// GetSwitchport returns the Switchport field value if set, zero value otherwise.
func (o *InlineResponse20083) GetSwitchport() string {
	if o == nil || isNil(o.Switchport) {
		var ret string
		return ret
	}
	return *o.Switchport
}

// GetSwitchportOk returns a tuple with the Switchport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetSwitchportOk() (*string, bool) {
	if o == nil || isNil(o.Switchport) {
    return nil, false
	}
	return o.Switchport, true
}

// HasSwitchport returns a boolean if a field has been set.
func (o *InlineResponse20083) HasSwitchport() bool {
	if o != nil && !isNil(o.Switchport) {
		return true
	}

	return false
}

// SetSwitchport gets a reference to the given string and assigns it to the Switchport field.
func (o *InlineResponse20083) SetSwitchport(v string) {
	o.Switchport = &v
}

// GetWirelessCapabilities returns the WirelessCapabilities field value if set, zero value otherwise.
func (o *InlineResponse20083) GetWirelessCapabilities() string {
	if o == nil || isNil(o.WirelessCapabilities) {
		var ret string
		return ret
	}
	return *o.WirelessCapabilities
}

// GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetWirelessCapabilitiesOk() (*string, bool) {
	if o == nil || isNil(o.WirelessCapabilities) {
    return nil, false
	}
	return o.WirelessCapabilities, true
}

// HasWirelessCapabilities returns a boolean if a field has been set.
func (o *InlineResponse20083) HasWirelessCapabilities() bool {
	if o != nil && !isNil(o.WirelessCapabilities) {
		return true
	}

	return false
}

// SetWirelessCapabilities gets a reference to the given string and assigns it to the WirelessCapabilities field.
func (o *InlineResponse20083) SetWirelessCapabilities(v string) {
	o.WirelessCapabilities = &v
}

// GetSmInstalled returns the SmInstalled field value if set, zero value otherwise.
func (o *InlineResponse20083) GetSmInstalled() bool {
	if o == nil || isNil(o.SmInstalled) {
		var ret bool
		return ret
	}
	return *o.SmInstalled
}

// GetSmInstalledOk returns a tuple with the SmInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetSmInstalledOk() (*bool, bool) {
	if o == nil || isNil(o.SmInstalled) {
    return nil, false
	}
	return o.SmInstalled, true
}

// HasSmInstalled returns a boolean if a field has been set.
func (o *InlineResponse20083) HasSmInstalled() bool {
	if o != nil && !isNil(o.SmInstalled) {
		return true
	}

	return false
}

// SetSmInstalled gets a reference to the given bool and assigns it to the SmInstalled field.
func (o *InlineResponse20083) SetSmInstalled(v bool) {
	o.SmInstalled = &v
}

// GetRecentDeviceMac returns the RecentDeviceMac field value if set, zero value otherwise.
func (o *InlineResponse20083) GetRecentDeviceMac() string {
	if o == nil || isNil(o.RecentDeviceMac) {
		var ret string
		return ret
	}
	return *o.RecentDeviceMac
}

// GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetRecentDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceMac) {
    return nil, false
	}
	return o.RecentDeviceMac, true
}

// HasRecentDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse20083) HasRecentDeviceMac() bool {
	if o != nil && !isNil(o.RecentDeviceMac) {
		return true
	}

	return false
}

// SetRecentDeviceMac gets a reference to the given string and assigns it to the RecentDeviceMac field.
func (o *InlineResponse20083) SetRecentDeviceMac(v string) {
	o.RecentDeviceMac = &v
}

// GetRecentDeviceConnection returns the RecentDeviceConnection field value if set, zero value otherwise.
func (o *InlineResponse20083) GetRecentDeviceConnection() string {
	if o == nil || isNil(o.RecentDeviceConnection) {
		var ret string
		return ret
	}
	return *o.RecentDeviceConnection
}

// GetRecentDeviceConnectionOk returns a tuple with the RecentDeviceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetRecentDeviceConnectionOk() (*string, bool) {
	if o == nil || isNil(o.RecentDeviceConnection) {
    return nil, false
	}
	return o.RecentDeviceConnection, true
}

// HasRecentDeviceConnection returns a boolean if a field has been set.
func (o *InlineResponse20083) HasRecentDeviceConnection() bool {
	if o != nil && !isNil(o.RecentDeviceConnection) {
		return true
	}

	return false
}

// SetRecentDeviceConnection gets a reference to the given string and assigns it to the RecentDeviceConnection field.
func (o *InlineResponse20083) SetRecentDeviceConnection(v string) {
	o.RecentDeviceConnection = &v
}

// GetClientVpnConnections returns the ClientVpnConnections field value if set, zero value otherwise.
func (o *InlineResponse20083) GetClientVpnConnections() []InlineResponse20083ClientVpnConnections {
	if o == nil || isNil(o.ClientVpnConnections) {
		var ret []InlineResponse20083ClientVpnConnections
		return ret
	}
	return o.ClientVpnConnections
}

// GetClientVpnConnectionsOk returns a tuple with the ClientVpnConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetClientVpnConnectionsOk() ([]InlineResponse20083ClientVpnConnections, bool) {
	if o == nil || isNil(o.ClientVpnConnections) {
    return nil, false
	}
	return o.ClientVpnConnections, true
}

// HasClientVpnConnections returns a boolean if a field has been set.
func (o *InlineResponse20083) HasClientVpnConnections() bool {
	if o != nil && !isNil(o.ClientVpnConnections) {
		return true
	}

	return false
}

// SetClientVpnConnections gets a reference to the given []InlineResponse20083ClientVpnConnections and assigns it to the ClientVpnConnections field.
func (o *InlineResponse20083) SetClientVpnConnections(v []InlineResponse20083ClientVpnConnections) {
	o.ClientVpnConnections = v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse20083) GetLldp() [][]string {
	if o == nil || isNil(o.Lldp) {
		var ret [][]string
		return ret
	}
	return o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetLldpOk() ([][]string, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse20083) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given [][]string and assigns it to the Lldp field.
func (o *InlineResponse20083) SetLldp(v [][]string) {
	o.Lldp = v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse20083) GetCdp() [][]string {
	if o == nil || isNil(o.Cdp) {
		var ret [][]string
		return ret
	}
	return o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetCdpOk() ([][]string, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse20083) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given [][]string and assigns it to the Cdp field.
func (o *InlineResponse20083) SetCdp(v [][]string) {
	o.Cdp = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20083) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20083) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20083) SetStatus(v string) {
	o.Status = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse20083) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse20083) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse20083) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineResponse20083) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RecentDeviceConnection) {
		toSerialize["recentDeviceConnection"] = o.RecentDeviceConnection
	}
	if !isNil(o.ClientVpnConnections) {
		toSerialize["clientVpnConnections"] = o.ClientVpnConnections
	}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !isNil(o.Cdp) {
		toSerialize["cdp"] = o.Cdp
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20083 struct {
	value *InlineResponse20083
	isSet bool
}

func (v NullableInlineResponse20083) Get() *InlineResponse20083 {
	return v.value
}

func (v *NullableInlineResponse20083) Set(val *InlineResponse20083) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20083) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20083(val *InlineResponse20083) *NullableInlineResponse20083 {
	return &NullableInlineResponse20083{value: val, isSet: true}
}

func (v NullableInlineResponse20083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


