# InlineResponse200159

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac address of the server. | [optional] 
**Vlan** | Pointer to **int32** | Vlan id of the server. | [optional] 
**ClientId** | Pointer to **string** | Client id of the server if available. | [optional] 
**IsAllowed** | Pointer to **bool** | Whether the server is allowed or blocked. Always true for configured servers. | [optional] 
**LastSeenAt** | Pointer to **time.Time** | Last time the server was seen. | [optional] 
**SeenBy** | Pointer to [**[]NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy**](NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy.md) | Devices that saw the server. | [optional] 
**Type** | Pointer to **string** | server type. Can be a &#39;device&#39;, &#39;stack&#39;, or &#39;discovered&#39; (i.e client). | [optional] 
**Device** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenDevice**](NetworksNetworkIdSwitchDhcpV4ServersSeenDevice.md) |  | [optional] 
**Ipv4** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4**](NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4.md) |  | [optional] 
**IsConfigured** | Pointer to **bool** | Whether the server is configured. | [optional] 
**LastAck** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck**](NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck.md) |  | [optional] 
**LastPacket** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket**](NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket.md) |  | [optional] 

## Methods

### NewInlineResponse200159

`func NewInlineResponse200159() *InlineResponse200159`

NewInlineResponse200159 instantiates a new InlineResponse200159 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200159WithDefaults

`func NewInlineResponse200159WithDefaults() *InlineResponse200159`

NewInlineResponse200159WithDefaults instantiates a new InlineResponse200159 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200159) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200159) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200159) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200159) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200159) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200159) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200159) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200159) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetClientId

`func (o *InlineResponse200159) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200159) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200159) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200159) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIsAllowed

`func (o *InlineResponse200159) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *InlineResponse200159) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *InlineResponse200159) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *InlineResponse200159) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *InlineResponse200159) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *InlineResponse200159) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *InlineResponse200159) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *InlineResponse200159) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetSeenBy

`func (o *InlineResponse200159) GetSeenBy() []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy`

GetSeenBy returns the SeenBy field if non-nil, zero value otherwise.

### GetSeenByOk

`func (o *InlineResponse200159) GetSeenByOk() (*[]NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy, bool)`

GetSeenByOk returns a tuple with the SeenBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenBy

`func (o *InlineResponse200159) SetSeenBy(v []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy)`

SetSeenBy sets SeenBy field to given value.

### HasSeenBy

`func (o *InlineResponse200159) HasSeenBy() bool`

HasSeenBy returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200159) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200159) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200159) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200159) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200159) GetDevice() NetworksNetworkIdSwitchDhcpV4ServersSeenDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200159) GetDeviceOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200159) SetDevice(v NetworksNetworkIdSwitchDhcpV4ServersSeenDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200159) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineResponse200159) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineResponse200159) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineResponse200159) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineResponse200159) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIsConfigured

`func (o *InlineResponse200159) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *InlineResponse200159) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *InlineResponse200159) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *InlineResponse200159) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetLastAck

`func (o *InlineResponse200159) GetLastAck() NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck`

GetLastAck returns the LastAck field if non-nil, zero value otherwise.

### GetLastAckOk

`func (o *InlineResponse200159) GetLastAckOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck, bool)`

GetLastAckOk returns a tuple with the LastAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAck

`func (o *InlineResponse200159) SetLastAck(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck)`

SetLastAck sets LastAck field to given value.

### HasLastAck

`func (o *InlineResponse200159) HasLastAck() bool`

HasLastAck returns a boolean if a field has been set.

### GetLastPacket

`func (o *InlineResponse200159) GetLastPacket() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket`

GetLastPacket returns the LastPacket field if non-nil, zero value otherwise.

### GetLastPacketOk

`func (o *InlineResponse200159) GetLastPacketOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket, bool)`

GetLastPacketOk returns a tuple with the LastPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPacket

`func (o *InlineResponse200159) SetLastPacket(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket)`

SetLastPacket sets LastPacket field to given value.

### HasLastPacket

`func (o *InlineResponse200159) HasLastPacket() bool`

HasLastPacket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


