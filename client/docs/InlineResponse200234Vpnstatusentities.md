# InlineResponse200234Vpnstatusentities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device | [optional] 
**DeviceStatus** | Pointer to **string** | Device Status | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200234Uplinks**](InlineResponse200234Uplinks.md) | List of Uplink Information | [optional] 
**VpnMode** | Pointer to **string** | VPN Mode | [optional] 
**ExportedSubnets** | Pointer to [**[]InlineResponse200234ExportedSubnets**](InlineResponse200234ExportedSubnets.md) | List of Exported Subnets | [optional] 
**MerakiVpnPeers** | Pointer to [**[]InlineResponse200234MerakiVpnPeers**](InlineResponse200234MerakiVpnPeers.md) | Meraki VPN Peers | [optional] 
**ThirdPartyVpnPeers** | Pointer to [**[]InlineResponse200234ThirdPartyVpnPeers**](InlineResponse200234ThirdPartyVpnPeers.md) | Third Party VPN Peers | [optional] 

## Methods

### NewInlineResponse200234Vpnstatusentities

`func NewInlineResponse200234Vpnstatusentities() *InlineResponse200234Vpnstatusentities`

NewInlineResponse200234Vpnstatusentities instantiates a new InlineResponse200234Vpnstatusentities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200234VpnstatusentitiesWithDefaults

`func NewInlineResponse200234VpnstatusentitiesWithDefaults() *InlineResponse200234Vpnstatusentities`

NewInlineResponse200234VpnstatusentitiesWithDefaults instantiates a new InlineResponse200234Vpnstatusentities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200234Vpnstatusentities) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200234Vpnstatusentities) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200234Vpnstatusentities) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200234Vpnstatusentities) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200234Vpnstatusentities) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200234Vpnstatusentities) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200234Vpnstatusentities) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200234Vpnstatusentities) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200234Vpnstatusentities) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200234Vpnstatusentities) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200234Vpnstatusentities) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200234Vpnstatusentities) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200234Vpnstatusentities) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200234Vpnstatusentities) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200234Vpnstatusentities) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200234Vpnstatusentities) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200234Vpnstatusentities) GetUplinks() []InlineResponse200234Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200234Vpnstatusentities) GetUplinksOk() (*[]InlineResponse200234Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200234Vpnstatusentities) SetUplinks(v []InlineResponse200234Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200234Vpnstatusentities) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetVpnMode

`func (o *InlineResponse200234Vpnstatusentities) GetVpnMode() string`

GetVpnMode returns the VpnMode field if non-nil, zero value otherwise.

### GetVpnModeOk

`func (o *InlineResponse200234Vpnstatusentities) GetVpnModeOk() (*string, bool)`

GetVpnModeOk returns a tuple with the VpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMode

`func (o *InlineResponse200234Vpnstatusentities) SetVpnMode(v string)`

SetVpnMode sets VpnMode field to given value.

### HasVpnMode

`func (o *InlineResponse200234Vpnstatusentities) HasVpnMode() bool`

HasVpnMode returns a boolean if a field has been set.

### GetExportedSubnets

`func (o *InlineResponse200234Vpnstatusentities) GetExportedSubnets() []InlineResponse200234ExportedSubnets`

GetExportedSubnets returns the ExportedSubnets field if non-nil, zero value otherwise.

### GetExportedSubnetsOk

`func (o *InlineResponse200234Vpnstatusentities) GetExportedSubnetsOk() (*[]InlineResponse200234ExportedSubnets, bool)`

GetExportedSubnetsOk returns a tuple with the ExportedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedSubnets

`func (o *InlineResponse200234Vpnstatusentities) SetExportedSubnets(v []InlineResponse200234ExportedSubnets)`

SetExportedSubnets sets ExportedSubnets field to given value.

### HasExportedSubnets

`func (o *InlineResponse200234Vpnstatusentities) HasExportedSubnets() bool`

HasExportedSubnets returns a boolean if a field has been set.

### GetMerakiVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) GetMerakiVpnPeers() []InlineResponse200234MerakiVpnPeers`

GetMerakiVpnPeers returns the MerakiVpnPeers field if non-nil, zero value otherwise.

### GetMerakiVpnPeersOk

`func (o *InlineResponse200234Vpnstatusentities) GetMerakiVpnPeersOk() (*[]InlineResponse200234MerakiVpnPeers, bool)`

GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerakiVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) SetMerakiVpnPeers(v []InlineResponse200234MerakiVpnPeers)`

SetMerakiVpnPeers sets MerakiVpnPeers field to given value.

### HasMerakiVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) HasMerakiVpnPeers() bool`

HasMerakiVpnPeers returns a boolean if a field has been set.

### GetThirdPartyVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) GetThirdPartyVpnPeers() []InlineResponse200234ThirdPartyVpnPeers`

GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field if non-nil, zero value otherwise.

### GetThirdPartyVpnPeersOk

`func (o *InlineResponse200234Vpnstatusentities) GetThirdPartyVpnPeersOk() (*[]InlineResponse200234ThirdPartyVpnPeers, bool)`

GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) SetThirdPartyVpnPeers(v []InlineResponse200234ThirdPartyVpnPeers)`

SetThirdPartyVpnPeers sets ThirdPartyVpnPeers field to given value.

### HasThirdPartyVpnPeers

`func (o *InlineResponse200234Vpnstatusentities) HasThirdPartyVpnPeers() bool`

HasThirdPartyVpnPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


