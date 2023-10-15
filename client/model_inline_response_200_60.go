/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20060 struct for InlineResponse20060
type InlineResponse20060 struct {
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

// NewInlineResponse20060 instantiates a new InlineResponse20060 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060() *InlineResponse20060 {
	this := InlineResponse20060{}
	return &this
}

// NewInlineResponse20060WithDefaults instantiates a new InlineResponse20060 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060WithDefaults() *InlineResponse20060 {
	this := InlineResponse20060{}
	return &this
}

// GetMeasuredAt returns the MeasuredAt field value if set, zero value otherwise.
func (o *InlineResponse20060) GetMeasuredAt() string {
	if o == nil || isNil(o.MeasuredAt) {
		var ret string
		return ret
	}
	return *o.MeasuredAt
}

// GetMeasuredAtOk returns a tuple with the MeasuredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetMeasuredAtOk() (*string, bool) {
	if o == nil || isNil(o.MeasuredAt) {
    return nil, false
	}
	return o.MeasuredAt, true
}

// HasMeasuredAt returns a boolean if a field has been set.
func (o *InlineResponse20060) HasMeasuredAt() bool {
	if o != nil && !isNil(o.MeasuredAt) {
		return true
	}

	return false
}

// SetMeasuredAt gets a reference to the given string and assigns it to the MeasuredAt field.
func (o *InlineResponse20060) SetMeasuredAt(v string) {
	o.MeasuredAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *InlineResponse20060) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *InlineResponse20060) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *InlineResponse20060) SetUser(v string) {
	o.User = &v
}

// GetNetworkDevice returns the NetworkDevice field value if set, zero value otherwise.
func (o *InlineResponse20060) GetNetworkDevice() string {
	if o == nil || isNil(o.NetworkDevice) {
		var ret string
		return ret
	}
	return *o.NetworkDevice
}

// GetNetworkDeviceOk returns a tuple with the NetworkDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetNetworkDeviceOk() (*string, bool) {
	if o == nil || isNil(o.NetworkDevice) {
    return nil, false
	}
	return o.NetworkDevice, true
}

// HasNetworkDevice returns a boolean if a field has been set.
func (o *InlineResponse20060) HasNetworkDevice() bool {
	if o != nil && !isNil(o.NetworkDevice) {
		return true
	}

	return false
}

// SetNetworkDevice gets a reference to the given string and assigns it to the NetworkDevice field.
func (o *InlineResponse20060) SetNetworkDevice(v string) {
	o.NetworkDevice = &v
}

// GetNetworkDriver returns the NetworkDriver field value if set, zero value otherwise.
func (o *InlineResponse20060) GetNetworkDriver() string {
	if o == nil || isNil(o.NetworkDriver) {
		var ret string
		return ret
	}
	return *o.NetworkDriver
}

// GetNetworkDriverOk returns a tuple with the NetworkDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetNetworkDriverOk() (*string, bool) {
	if o == nil || isNil(o.NetworkDriver) {
    return nil, false
	}
	return o.NetworkDriver, true
}

// HasNetworkDriver returns a boolean if a field has been set.
func (o *InlineResponse20060) HasNetworkDriver() bool {
	if o != nil && !isNil(o.NetworkDriver) {
		return true
	}

	return false
}

// SetNetworkDriver gets a reference to the given string and assigns it to the NetworkDriver field.
func (o *InlineResponse20060) SetNetworkDriver(v string) {
	o.NetworkDriver = &v
}

// GetWifiChannel returns the WifiChannel field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiChannel() string {
	if o == nil || isNil(o.WifiChannel) {
		var ret string
		return ret
	}
	return *o.WifiChannel
}

// GetWifiChannelOk returns a tuple with the WifiChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiChannelOk() (*string, bool) {
	if o == nil || isNil(o.WifiChannel) {
    return nil, false
	}
	return o.WifiChannel, true
}

// HasWifiChannel returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiChannel() bool {
	if o != nil && !isNil(o.WifiChannel) {
		return true
	}

	return false
}

// SetWifiChannel gets a reference to the given string and assigns it to the WifiChannel field.
func (o *InlineResponse20060) SetWifiChannel(v string) {
	o.WifiChannel = &v
}

// GetWifiAuth returns the WifiAuth field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiAuth() string {
	if o == nil || isNil(o.WifiAuth) {
		var ret string
		return ret
	}
	return *o.WifiAuth
}

// GetWifiAuthOk returns a tuple with the WifiAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiAuthOk() (*string, bool) {
	if o == nil || isNil(o.WifiAuth) {
    return nil, false
	}
	return o.WifiAuth, true
}

// HasWifiAuth returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiAuth() bool {
	if o != nil && !isNil(o.WifiAuth) {
		return true
	}

	return false
}

// SetWifiAuth gets a reference to the given string and assigns it to the WifiAuth field.
func (o *InlineResponse20060) SetWifiAuth(v string) {
	o.WifiAuth = &v
}

// GetWifiBssid returns the WifiBssid field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiBssid() string {
	if o == nil || isNil(o.WifiBssid) {
		var ret string
		return ret
	}
	return *o.WifiBssid
}

// GetWifiBssidOk returns a tuple with the WifiBssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiBssidOk() (*string, bool) {
	if o == nil || isNil(o.WifiBssid) {
    return nil, false
	}
	return o.WifiBssid, true
}

// HasWifiBssid returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiBssid() bool {
	if o != nil && !isNil(o.WifiBssid) {
		return true
	}

	return false
}

// SetWifiBssid gets a reference to the given string and assigns it to the WifiBssid field.
func (o *InlineResponse20060) SetWifiBssid(v string) {
	o.WifiBssid = &v
}

// GetWifiSsid returns the WifiSsid field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiSsid() string {
	if o == nil || isNil(o.WifiSsid) {
		var ret string
		return ret
	}
	return *o.WifiSsid
}

// GetWifiSsidOk returns a tuple with the WifiSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiSsidOk() (*string, bool) {
	if o == nil || isNil(o.WifiSsid) {
    return nil, false
	}
	return o.WifiSsid, true
}

// HasWifiSsid returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiSsid() bool {
	if o != nil && !isNil(o.WifiSsid) {
		return true
	}

	return false
}

// SetWifiSsid gets a reference to the given string and assigns it to the WifiSsid field.
func (o *InlineResponse20060) SetWifiSsid(v string) {
	o.WifiSsid = &v
}

// GetWifiRssi returns the WifiRssi field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiRssi() string {
	if o == nil || isNil(o.WifiRssi) {
		var ret string
		return ret
	}
	return *o.WifiRssi
}

// GetWifiRssiOk returns a tuple with the WifiRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiRssiOk() (*string, bool) {
	if o == nil || isNil(o.WifiRssi) {
    return nil, false
	}
	return o.WifiRssi, true
}

// HasWifiRssi returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiRssi() bool {
	if o != nil && !isNil(o.WifiRssi) {
		return true
	}

	return false
}

// SetWifiRssi gets a reference to the given string and assigns it to the WifiRssi field.
func (o *InlineResponse20060) SetWifiRssi(v string) {
	o.WifiRssi = &v
}

// GetWifiNoise returns the WifiNoise field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWifiNoise() string {
	if o == nil || isNil(o.WifiNoise) {
		var ret string
		return ret
	}
	return *o.WifiNoise
}

// GetWifiNoiseOk returns a tuple with the WifiNoise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWifiNoiseOk() (*string, bool) {
	if o == nil || isNil(o.WifiNoise) {
    return nil, false
	}
	return o.WifiNoise, true
}

// HasWifiNoise returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWifiNoise() bool {
	if o != nil && !isNil(o.WifiNoise) {
		return true
	}

	return false
}

// SetWifiNoise gets a reference to the given string and assigns it to the WifiNoise field.
func (o *InlineResponse20060) SetWifiNoise(v string) {
	o.WifiNoise = &v
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *InlineResponse20060) GetDhcpServer() string {
	if o == nil || isNil(o.DhcpServer) {
		var ret string
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetDhcpServerOk() (*string, bool) {
	if o == nil || isNil(o.DhcpServer) {
    return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *InlineResponse20060) HasDhcpServer() bool {
	if o != nil && !isNil(o.DhcpServer) {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given string and assigns it to the DhcpServer field.
func (o *InlineResponse20060) SetDhcpServer(v string) {
	o.DhcpServer = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20060) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20060) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20060) SetIp(v string) {
	o.Ip = &v
}

// GetNetworkMTU returns the NetworkMTU field value if set, zero value otherwise.
func (o *InlineResponse20060) GetNetworkMTU() string {
	if o == nil || isNil(o.NetworkMTU) {
		var ret string
		return ret
	}
	return *o.NetworkMTU
}

// GetNetworkMTUOk returns a tuple with the NetworkMTU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetNetworkMTUOk() (*string, bool) {
	if o == nil || isNil(o.NetworkMTU) {
    return nil, false
	}
	return o.NetworkMTU, true
}

// HasNetworkMTU returns a boolean if a field has been set.
func (o *InlineResponse20060) HasNetworkMTU() bool {
	if o != nil && !isNil(o.NetworkMTU) {
		return true
	}

	return false
}

// SetNetworkMTU gets a reference to the given string and assigns it to the NetworkMTU field.
func (o *InlineResponse20060) SetNetworkMTU(v string) {
	o.NetworkMTU = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20060) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20060) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20060) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse20060) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse20060) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse20060) SetGateway(v string) {
	o.Gateway = &v
}

// GetPublicIP returns the PublicIP field value if set, zero value otherwise.
func (o *InlineResponse20060) GetPublicIP() string {
	if o == nil || isNil(o.PublicIP) {
		var ret string
		return ret
	}
	return *o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetPublicIPOk() (*string, bool) {
	if o == nil || isNil(o.PublicIP) {
    return nil, false
	}
	return o.PublicIP, true
}

// HasPublicIP returns a boolean if a field has been set.
func (o *InlineResponse20060) HasPublicIP() bool {
	if o != nil && !isNil(o.PublicIP) {
		return true
	}

	return false
}

// SetPublicIP gets a reference to the given string and assigns it to the PublicIP field.
func (o *InlineResponse20060) SetPublicIP(v string) {
	o.PublicIP = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *InlineResponse20060) GetDnsServer() string {
	if o == nil || isNil(o.DnsServer) {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetDnsServerOk() (*string, bool) {
	if o == nil || isNil(o.DnsServer) {
    return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *InlineResponse20060) HasDnsServer() bool {
	if o != nil && !isNil(o.DnsServer) {
		return true
	}

	return false
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *InlineResponse20060) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20060) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20060) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse20060) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse20060) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MeasuredAt) {
		toSerialize["measuredAt"] = o.MeasuredAt
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.NetworkDevice) {
		toSerialize["networkDevice"] = o.NetworkDevice
	}
	if !isNil(o.NetworkDriver) {
		toSerialize["networkDriver"] = o.NetworkDriver
	}
	if !isNil(o.WifiChannel) {
		toSerialize["wifiChannel"] = o.WifiChannel
	}
	if !isNil(o.WifiAuth) {
		toSerialize["wifiAuth"] = o.WifiAuth
	}
	if !isNil(o.WifiBssid) {
		toSerialize["wifiBssid"] = o.WifiBssid
	}
	if !isNil(o.WifiSsid) {
		toSerialize["wifiSsid"] = o.WifiSsid
	}
	if !isNil(o.WifiRssi) {
		toSerialize["wifiRssi"] = o.WifiRssi
	}
	if !isNil(o.WifiNoise) {
		toSerialize["wifiNoise"] = o.WifiNoise
	}
	if !isNil(o.DhcpServer) {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.NetworkMTU) {
		toSerialize["networkMTU"] = o.NetworkMTU
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.PublicIP) {
		toSerialize["publicIP"] = o.PublicIP
	}
	if !isNil(o.DnsServer) {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060 struct {
	value *InlineResponse20060
	isSet bool
}

func (v NullableInlineResponse20060) Get() *InlineResponse20060 {
	return v.value
}

func (v *NullableInlineResponse20060) Set(val *InlineResponse20060) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060(val *InlineResponse20060) *NullableInlineResponse20060 {
	return &NullableInlineResponse20060{value: val, isSet: true}
}

func (v NullableInlineResponse20060) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


