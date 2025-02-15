# InlineResponse200130

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasuredAt** | Pointer to **string** | The time the data was measured at. | [optional] 
**User** | Pointer to **string** | The user during connection. | [optional] 
**NetworkDevice** | Pointer to **string** | The network device for the device used for connection. | [optional] 
**NetworkDriver** | Pointer to **string** | The network driver for the device. | [optional] 
**WifiChannel** | Pointer to **string** | Channel through which the connection is routing. | [optional] 
**WifiAuth** | Pointer to **string** | The type of authentication used by the SSID. | [optional] 
**WifiBssid** | Pointer to **string** | The MAC of the access point the device is connected to. | [optional] 
**WifiSsid** | Pointer to **string** | The name of the network the device is connected to. | [optional] 
**WifiRssi** | Pointer to **string** | The Received Signal Strength Indicator for the device. | [optional] 
**WifiNoise** | Pointer to **string** | The wireless signal power level received by the device. | [optional] 
**DhcpServer** | Pointer to **string** | The IP address of the DCHP Server. | [optional] 
**Ip** | Pointer to **string** | The IP of the device during connection. | [optional] 
**NetworkMTU** | Pointer to **string** | The network max transmission unit. | [optional] 
**Subnet** | Pointer to **string** | The subnet of the device connection. | [optional] 
**Gateway** | Pointer to **string** | The gateway IP the device was connected to. | [optional] 
**PublicIP** | Pointer to **string** | The public IP address of the device. | [optional] 
**DnsServer** | Pointer to **string** | The DNS Server during the connection. | [optional] 
**Ts** | Pointer to **string** | The time the connection was logged. | [optional] 

## Methods

### NewInlineResponse200130

`func NewInlineResponse200130() *InlineResponse200130`

NewInlineResponse200130 instantiates a new InlineResponse200130 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200130WithDefaults

`func NewInlineResponse200130WithDefaults() *InlineResponse200130`

NewInlineResponse200130WithDefaults instantiates a new InlineResponse200130 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasuredAt

`func (o *InlineResponse200130) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *InlineResponse200130) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *InlineResponse200130) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *InlineResponse200130) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse200130) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse200130) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse200130) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse200130) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNetworkDevice

`func (o *InlineResponse200130) GetNetworkDevice() string`

GetNetworkDevice returns the NetworkDevice field if non-nil, zero value otherwise.

### GetNetworkDeviceOk

`func (o *InlineResponse200130) GetNetworkDeviceOk() (*string, bool)`

GetNetworkDeviceOk returns a tuple with the NetworkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevice

`func (o *InlineResponse200130) SetNetworkDevice(v string)`

SetNetworkDevice sets NetworkDevice field to given value.

### HasNetworkDevice

`func (o *InlineResponse200130) HasNetworkDevice() bool`

HasNetworkDevice returns a boolean if a field has been set.

### GetNetworkDriver

`func (o *InlineResponse200130) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *InlineResponse200130) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *InlineResponse200130) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.

### HasNetworkDriver

`func (o *InlineResponse200130) HasNetworkDriver() bool`

HasNetworkDriver returns a boolean if a field has been set.

### GetWifiChannel

`func (o *InlineResponse200130) GetWifiChannel() string`

GetWifiChannel returns the WifiChannel field if non-nil, zero value otherwise.

### GetWifiChannelOk

`func (o *InlineResponse200130) GetWifiChannelOk() (*string, bool)`

GetWifiChannelOk returns a tuple with the WifiChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiChannel

`func (o *InlineResponse200130) SetWifiChannel(v string)`

SetWifiChannel sets WifiChannel field to given value.

### HasWifiChannel

`func (o *InlineResponse200130) HasWifiChannel() bool`

HasWifiChannel returns a boolean if a field has been set.

### GetWifiAuth

`func (o *InlineResponse200130) GetWifiAuth() string`

GetWifiAuth returns the WifiAuth field if non-nil, zero value otherwise.

### GetWifiAuthOk

`func (o *InlineResponse200130) GetWifiAuthOk() (*string, bool)`

GetWifiAuthOk returns a tuple with the WifiAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiAuth

`func (o *InlineResponse200130) SetWifiAuth(v string)`

SetWifiAuth sets WifiAuth field to given value.

### HasWifiAuth

`func (o *InlineResponse200130) HasWifiAuth() bool`

HasWifiAuth returns a boolean if a field has been set.

### GetWifiBssid

`func (o *InlineResponse200130) GetWifiBssid() string`

GetWifiBssid returns the WifiBssid field if non-nil, zero value otherwise.

### GetWifiBssidOk

`func (o *InlineResponse200130) GetWifiBssidOk() (*string, bool)`

GetWifiBssidOk returns a tuple with the WifiBssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBssid

`func (o *InlineResponse200130) SetWifiBssid(v string)`

SetWifiBssid sets WifiBssid field to given value.

### HasWifiBssid

`func (o *InlineResponse200130) HasWifiBssid() bool`

HasWifiBssid returns a boolean if a field has been set.

### GetWifiSsid

`func (o *InlineResponse200130) GetWifiSsid() string`

GetWifiSsid returns the WifiSsid field if non-nil, zero value otherwise.

### GetWifiSsidOk

`func (o *InlineResponse200130) GetWifiSsidOk() (*string, bool)`

GetWifiSsidOk returns a tuple with the WifiSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiSsid

`func (o *InlineResponse200130) SetWifiSsid(v string)`

SetWifiSsid sets WifiSsid field to given value.

### HasWifiSsid

`func (o *InlineResponse200130) HasWifiSsid() bool`

HasWifiSsid returns a boolean if a field has been set.

### GetWifiRssi

`func (o *InlineResponse200130) GetWifiRssi() string`

GetWifiRssi returns the WifiRssi field if non-nil, zero value otherwise.

### GetWifiRssiOk

`func (o *InlineResponse200130) GetWifiRssiOk() (*string, bool)`

GetWifiRssiOk returns a tuple with the WifiRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiRssi

`func (o *InlineResponse200130) SetWifiRssi(v string)`

SetWifiRssi sets WifiRssi field to given value.

### HasWifiRssi

`func (o *InlineResponse200130) HasWifiRssi() bool`

HasWifiRssi returns a boolean if a field has been set.

### GetWifiNoise

`func (o *InlineResponse200130) GetWifiNoise() string`

GetWifiNoise returns the WifiNoise field if non-nil, zero value otherwise.

### GetWifiNoiseOk

`func (o *InlineResponse200130) GetWifiNoiseOk() (*string, bool)`

GetWifiNoiseOk returns a tuple with the WifiNoise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiNoise

`func (o *InlineResponse200130) SetWifiNoise(v string)`

SetWifiNoise sets WifiNoise field to given value.

### HasWifiNoise

`func (o *InlineResponse200130) HasWifiNoise() bool`

HasWifiNoise returns a boolean if a field has been set.

### GetDhcpServer

`func (o *InlineResponse200130) GetDhcpServer() string`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *InlineResponse200130) GetDhcpServerOk() (*string, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *InlineResponse200130) SetDhcpServer(v string)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *InlineResponse200130) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse200130) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse200130) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse200130) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse200130) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetworkMTU

`func (o *InlineResponse200130) GetNetworkMTU() string`

GetNetworkMTU returns the NetworkMTU field if non-nil, zero value otherwise.

### GetNetworkMTUOk

`func (o *InlineResponse200130) GetNetworkMTUOk() (*string, bool)`

GetNetworkMTUOk returns a tuple with the NetworkMTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMTU

`func (o *InlineResponse200130) SetNetworkMTU(v string)`

SetNetworkMTU sets NetworkMTU field to given value.

### HasNetworkMTU

`func (o *InlineResponse200130) HasNetworkMTU() bool`

HasNetworkMTU returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200130) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200130) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200130) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200130) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse200130) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200130) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200130) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse200130) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPublicIP

`func (o *InlineResponse200130) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *InlineResponse200130) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *InlineResponse200130) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *InlineResponse200130) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.

### GetDnsServer

`func (o *InlineResponse200130) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *InlineResponse200130) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *InlineResponse200130) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *InlineResponse200130) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200130) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200130) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200130) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200130) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


