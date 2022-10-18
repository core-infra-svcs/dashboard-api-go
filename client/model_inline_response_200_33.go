/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	// The time the data was measured at.
	MeasuredAt *string `json:"measuredAt,omitempty"`
	// The user during connection.
	User *string `json:"user,omitempty"`
	// The network device for the device used for connection.
	NetworkDevice *string `json:"networkDevice,omitempty"`
	// The network driver for the device.
	NetworkDriver *string `json:"networkDriver,omitempty"`
	// Channel through which the connection is routing.
	WifiChannel *string `json:"wifiChannel,omitempty"`
	// The type of authentication used by the SSID.
	WifiAuth *string `json:"wifiAuth,omitempty"`
	// The MAC of the access point the device is connected to.
	WifiBssid *string `json:"wifiBssid,omitempty"`
	// The name of the network the device is connected to.
	WifiSsid *string `json:"wifiSsid,omitempty"`
	// The Received Signal Strength Indicator for the device.
	WifiRssi *string `json:"wifiRssi,omitempty"`
	// The wireless signal power level received by the device.
	WifiNoise *string `json:"wifiNoise,omitempty"`
	// The IP address of the DCHP Server.
	DhcpServer *string `json:"dhcpServer,omitempty"`
	// The IP of the device during connection.
	Ip *string `json:"ip,omitempty"`
	// The network max transmission unit.
	NetworkMTU *string `json:"networkMTU,omitempty"`
	// The subnet of the device connection.
	Subnet *string `json:"subnet,omitempty"`
	// The gateway IP the device was connected to.
	Gateway *string `json:"gateway,omitempty"`
	// The public IP address of the device.
	PublicIP *string `json:"publicIP,omitempty"`
	// The DNS Server during the connection.
	DnsServer *string `json:"dnsServer,omitempty"`
	// The time the connection was logged.
	Ts *string `json:"ts,omitempty"`
}

// NewInlineResponse20033 instantiates a new InlineResponse20033 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20033() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20033WithDefaults() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// GetMeasuredAt returns the MeasuredAt field value if set, zero value otherwise.
func (o *InlineResponse20033) GetMeasuredAt() string {
	if o == nil || o.MeasuredAt == nil {
		var ret string
		return ret
	}
	return *o.MeasuredAt
}

// GetMeasuredAtOk returns a tuple with the MeasuredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetMeasuredAtOk() (*string, bool) {
	if o == nil || o.MeasuredAt == nil {
		return nil, false
	}
	return o.MeasuredAt, true
}

// HasMeasuredAt returns a boolean if a field has been set.
func (o *InlineResponse20033) HasMeasuredAt() bool {
	if o != nil && o.MeasuredAt != nil {
		return true
	}

	return false
}

// SetMeasuredAt gets a reference to the given string and assigns it to the MeasuredAt field.
func (o *InlineResponse20033) SetMeasuredAt(v string) {
	o.MeasuredAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *InlineResponse20033) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *InlineResponse20033) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *InlineResponse20033) SetUser(v string) {
	o.User = &v
}

// GetNetworkDevice returns the NetworkDevice field value if set, zero value otherwise.
func (o *InlineResponse20033) GetNetworkDevice() string {
	if o == nil || o.NetworkDevice == nil {
		var ret string
		return ret
	}
	return *o.NetworkDevice
}

// GetNetworkDeviceOk returns a tuple with the NetworkDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetNetworkDeviceOk() (*string, bool) {
	if o == nil || o.NetworkDevice == nil {
		return nil, false
	}
	return o.NetworkDevice, true
}

// HasNetworkDevice returns a boolean if a field has been set.
func (o *InlineResponse20033) HasNetworkDevice() bool {
	if o != nil && o.NetworkDevice != nil {
		return true
	}

	return false
}

// SetNetworkDevice gets a reference to the given string and assigns it to the NetworkDevice field.
func (o *InlineResponse20033) SetNetworkDevice(v string) {
	o.NetworkDevice = &v
}

// GetNetworkDriver returns the NetworkDriver field value if set, zero value otherwise.
func (o *InlineResponse20033) GetNetworkDriver() string {
	if o == nil || o.NetworkDriver == nil {
		var ret string
		return ret
	}
	return *o.NetworkDriver
}

// GetNetworkDriverOk returns a tuple with the NetworkDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetNetworkDriverOk() (*string, bool) {
	if o == nil || o.NetworkDriver == nil {
		return nil, false
	}
	return o.NetworkDriver, true
}

// HasNetworkDriver returns a boolean if a field has been set.
func (o *InlineResponse20033) HasNetworkDriver() bool {
	if o != nil && o.NetworkDriver != nil {
		return true
	}

	return false
}

// SetNetworkDriver gets a reference to the given string and assigns it to the NetworkDriver field.
func (o *InlineResponse20033) SetNetworkDriver(v string) {
	o.NetworkDriver = &v
}

// GetWifiChannel returns the WifiChannel field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiChannel() string {
	if o == nil || o.WifiChannel == nil {
		var ret string
		return ret
	}
	return *o.WifiChannel
}

// GetWifiChannelOk returns a tuple with the WifiChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiChannelOk() (*string, bool) {
	if o == nil || o.WifiChannel == nil {
		return nil, false
	}
	return o.WifiChannel, true
}

// HasWifiChannel returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiChannel() bool {
	if o != nil && o.WifiChannel != nil {
		return true
	}

	return false
}

// SetWifiChannel gets a reference to the given string and assigns it to the WifiChannel field.
func (o *InlineResponse20033) SetWifiChannel(v string) {
	o.WifiChannel = &v
}

// GetWifiAuth returns the WifiAuth field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiAuth() string {
	if o == nil || o.WifiAuth == nil {
		var ret string
		return ret
	}
	return *o.WifiAuth
}

// GetWifiAuthOk returns a tuple with the WifiAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiAuthOk() (*string, bool) {
	if o == nil || o.WifiAuth == nil {
		return nil, false
	}
	return o.WifiAuth, true
}

// HasWifiAuth returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiAuth() bool {
	if o != nil && o.WifiAuth != nil {
		return true
	}

	return false
}

// SetWifiAuth gets a reference to the given string and assigns it to the WifiAuth field.
func (o *InlineResponse20033) SetWifiAuth(v string) {
	o.WifiAuth = &v
}

// GetWifiBssid returns the WifiBssid field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiBssid() string {
	if o == nil || o.WifiBssid == nil {
		var ret string
		return ret
	}
	return *o.WifiBssid
}

// GetWifiBssidOk returns a tuple with the WifiBssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiBssidOk() (*string, bool) {
	if o == nil || o.WifiBssid == nil {
		return nil, false
	}
	return o.WifiBssid, true
}

// HasWifiBssid returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiBssid() bool {
	if o != nil && o.WifiBssid != nil {
		return true
	}

	return false
}

// SetWifiBssid gets a reference to the given string and assigns it to the WifiBssid field.
func (o *InlineResponse20033) SetWifiBssid(v string) {
	o.WifiBssid = &v
}

// GetWifiSsid returns the WifiSsid field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiSsid() string {
	if o == nil || o.WifiSsid == nil {
		var ret string
		return ret
	}
	return *o.WifiSsid
}

// GetWifiSsidOk returns a tuple with the WifiSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiSsidOk() (*string, bool) {
	if o == nil || o.WifiSsid == nil {
		return nil, false
	}
	return o.WifiSsid, true
}

// HasWifiSsid returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiSsid() bool {
	if o != nil && o.WifiSsid != nil {
		return true
	}

	return false
}

// SetWifiSsid gets a reference to the given string and assigns it to the WifiSsid field.
func (o *InlineResponse20033) SetWifiSsid(v string) {
	o.WifiSsid = &v
}

// GetWifiRssi returns the WifiRssi field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiRssi() string {
	if o == nil || o.WifiRssi == nil {
		var ret string
		return ret
	}
	return *o.WifiRssi
}

// GetWifiRssiOk returns a tuple with the WifiRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiRssiOk() (*string, bool) {
	if o == nil || o.WifiRssi == nil {
		return nil, false
	}
	return o.WifiRssi, true
}

// HasWifiRssi returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiRssi() bool {
	if o != nil && o.WifiRssi != nil {
		return true
	}

	return false
}

// SetWifiRssi gets a reference to the given string and assigns it to the WifiRssi field.
func (o *InlineResponse20033) SetWifiRssi(v string) {
	o.WifiRssi = &v
}

// GetWifiNoise returns the WifiNoise field value if set, zero value otherwise.
func (o *InlineResponse20033) GetWifiNoise() string {
	if o == nil || o.WifiNoise == nil {
		var ret string
		return ret
	}
	return *o.WifiNoise
}

// GetWifiNoiseOk returns a tuple with the WifiNoise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWifiNoiseOk() (*string, bool) {
	if o == nil || o.WifiNoise == nil {
		return nil, false
	}
	return o.WifiNoise, true
}

// HasWifiNoise returns a boolean if a field has been set.
func (o *InlineResponse20033) HasWifiNoise() bool {
	if o != nil && o.WifiNoise != nil {
		return true
	}

	return false
}

// SetWifiNoise gets a reference to the given string and assigns it to the WifiNoise field.
func (o *InlineResponse20033) SetWifiNoise(v string) {
	o.WifiNoise = &v
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *InlineResponse20033) GetDhcpServer() string {
	if o == nil || o.DhcpServer == nil {
		var ret string
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetDhcpServerOk() (*string, bool) {
	if o == nil || o.DhcpServer == nil {
		return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *InlineResponse20033) HasDhcpServer() bool {
	if o != nil && o.DhcpServer != nil {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given string and assigns it to the DhcpServer field.
func (o *InlineResponse20033) SetDhcpServer(v string) {
	o.DhcpServer = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20033) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20033) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20033) SetIp(v string) {
	o.Ip = &v
}

// GetNetworkMTU returns the NetworkMTU field value if set, zero value otherwise.
func (o *InlineResponse20033) GetNetworkMTU() string {
	if o == nil || o.NetworkMTU == nil {
		var ret string
		return ret
	}
	return *o.NetworkMTU
}

// GetNetworkMTUOk returns a tuple with the NetworkMTU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetNetworkMTUOk() (*string, bool) {
	if o == nil || o.NetworkMTU == nil {
		return nil, false
	}
	return o.NetworkMTU, true
}

// HasNetworkMTU returns a boolean if a field has been set.
func (o *InlineResponse20033) HasNetworkMTU() bool {
	if o != nil && o.NetworkMTU != nil {
		return true
	}

	return false
}

// SetNetworkMTU gets a reference to the given string and assigns it to the NetworkMTU field.
func (o *InlineResponse20033) SetNetworkMTU(v string) {
	o.NetworkMTU = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20033) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20033) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20033) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse20033) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse20033) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse20033) SetGateway(v string) {
	o.Gateway = &v
}

// GetPublicIP returns the PublicIP field value if set, zero value otherwise.
func (o *InlineResponse20033) GetPublicIP() string {
	if o == nil || o.PublicIP == nil {
		var ret string
		return ret
	}
	return *o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetPublicIPOk() (*string, bool) {
	if o == nil || o.PublicIP == nil {
		return nil, false
	}
	return o.PublicIP, true
}

// HasPublicIP returns a boolean if a field has been set.
func (o *InlineResponse20033) HasPublicIP() bool {
	if o != nil && o.PublicIP != nil {
		return true
	}

	return false
}

// SetPublicIP gets a reference to the given string and assigns it to the PublicIP field.
func (o *InlineResponse20033) SetPublicIP(v string) {
	o.PublicIP = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *InlineResponse20033) GetDnsServer() string {
	if o == nil || o.DnsServer == nil {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetDnsServerOk() (*string, bool) {
	if o == nil || o.DnsServer == nil {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *InlineResponse20033) HasDnsServer() bool {
	if o != nil && o.DnsServer != nil {
		return true
	}

	return false
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *InlineResponse20033) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20033) GetTs() string {
	if o == nil || o.Ts == nil {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetTsOk() (*string, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20033) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse20033) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse20033) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MeasuredAt != nil {
		toSerialize["measuredAt"] = o.MeasuredAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.NetworkDevice != nil {
		toSerialize["networkDevice"] = o.NetworkDevice
	}
	if o.NetworkDriver != nil {
		toSerialize["networkDriver"] = o.NetworkDriver
	}
	if o.WifiChannel != nil {
		toSerialize["wifiChannel"] = o.WifiChannel
	}
	if o.WifiAuth != nil {
		toSerialize["wifiAuth"] = o.WifiAuth
	}
	if o.WifiBssid != nil {
		toSerialize["wifiBssid"] = o.WifiBssid
	}
	if o.WifiSsid != nil {
		toSerialize["wifiSsid"] = o.WifiSsid
	}
	if o.WifiRssi != nil {
		toSerialize["wifiRssi"] = o.WifiRssi
	}
	if o.WifiNoise != nil {
		toSerialize["wifiNoise"] = o.WifiNoise
	}
	if o.DhcpServer != nil {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.NetworkMTU != nil {
		toSerialize["networkMTU"] = o.NetworkMTU
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.PublicIP != nil {
		toSerialize["publicIP"] = o.PublicIP
	}
	if o.DnsServer != nil {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20033 struct {
	value *InlineResponse20033
	isSet bool
}

func (v NullableInlineResponse20033) Get() *InlineResponse20033 {
	return v.value
}

func (v *NullableInlineResponse20033) Set(val *InlineResponse20033) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20033) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20033) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20033(val *InlineResponse20033) *NullableInlineResponse20033 {
	return &NullableInlineResponse20033{value: val, isSet: true}
}

func (v NullableInlineResponse20033) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20033) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


