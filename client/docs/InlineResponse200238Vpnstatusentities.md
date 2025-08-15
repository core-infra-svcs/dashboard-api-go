# InlineResponse200238Vpnstatusentities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device | [optional] 
**DeviceStatus** | Pointer to **string** | Device Status | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200238Uplinks**](InlineResponse200238Uplinks.md) | List of Uplink Information | [optional] 
**VpnMode** | Pointer to **string** | VPN Mode | [optional] 
**ExportedSubnets** | Pointer to [**[]InlineResponse200238ExportedSubnets**](InlineResponse200238ExportedSubnets.md) | List of Exported Subnets | [optional] 
**MerakiVpnPeers** | Pointer to [**[]InlineResponse200238MerakiVpnPeers**](InlineResponse200238MerakiVpnPeers.md) | Meraki VPN Peers | [optional] 
**ThirdPartyVpnPeers** | Pointer to [**[]InlineResponse200238ThirdPartyVpnPeers**](InlineResponse200238ThirdPartyVpnPeers.md) | Third Party VPN Peers | [optional] 

## Methods

### NewInlineResponse200238Vpnstatusentities

`func NewInlineResponse200238Vpnstatusentities() *InlineResponse200238Vpnstatusentities`

NewInlineResponse200238Vpnstatusentities instantiates a new InlineResponse200238Vpnstatusentities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200238VpnstatusentitiesWithDefaults

`func NewInlineResponse200238VpnstatusentitiesWithDefaults() *InlineResponse200238Vpnstatusentities`

NewInlineResponse200238VpnstatusentitiesWithDefaults instantiates a new InlineResponse200238Vpnstatusentities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200238Vpnstatusentities) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200238Vpnstatusentities) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200238Vpnstatusentities) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200238Vpnstatusentities) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200238Vpnstatusentities) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200238Vpnstatusentities) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200238Vpnstatusentities) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200238Vpnstatusentities) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200238Vpnstatusentities) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200238Vpnstatusentities) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200238Vpnstatusentities) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200238Vpnstatusentities) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200238Vpnstatusentities) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200238Vpnstatusentities) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200238Vpnstatusentities) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200238Vpnstatusentities) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200238Vpnstatusentities) GetUplinks() []InlineResponse200238Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200238Vpnstatusentities) GetUplinksOk() (*[]InlineResponse200238Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200238Vpnstatusentities) SetUplinks(v []InlineResponse200238Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200238Vpnstatusentities) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetVpnMode

`func (o *InlineResponse200238Vpnstatusentities) GetVpnMode() string`

GetVpnMode returns the VpnMode field if non-nil, zero value otherwise.

### GetVpnModeOk

`func (o *InlineResponse200238Vpnstatusentities) GetVpnModeOk() (*string, bool)`

GetVpnModeOk returns a tuple with the VpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMode

`func (o *InlineResponse200238Vpnstatusentities) SetVpnMode(v string)`

SetVpnMode sets VpnMode field to given value.

### HasVpnMode

`func (o *InlineResponse200238Vpnstatusentities) HasVpnMode() bool`

HasVpnMode returns a boolean if a field has been set.

### GetExportedSubnets

`func (o *InlineResponse200238Vpnstatusentities) GetExportedSubnets() []InlineResponse200238ExportedSubnets`

GetExportedSubnets returns the ExportedSubnets field if non-nil, zero value otherwise.

### GetExportedSubnetsOk

`func (o *InlineResponse200238Vpnstatusentities) GetExportedSubnetsOk() (*[]InlineResponse200238ExportedSubnets, bool)`

GetExportedSubnetsOk returns a tuple with the ExportedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedSubnets

`func (o *InlineResponse200238Vpnstatusentities) SetExportedSubnets(v []InlineResponse200238ExportedSubnets)`

SetExportedSubnets sets ExportedSubnets field to given value.

### HasExportedSubnets

`func (o *InlineResponse200238Vpnstatusentities) HasExportedSubnets() bool`

HasExportedSubnets returns a boolean if a field has been set.

### GetMerakiVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) GetMerakiVpnPeers() []InlineResponse200238MerakiVpnPeers`

GetMerakiVpnPeers returns the MerakiVpnPeers field if non-nil, zero value otherwise.

### GetMerakiVpnPeersOk

`func (o *InlineResponse200238Vpnstatusentities) GetMerakiVpnPeersOk() (*[]InlineResponse200238MerakiVpnPeers, bool)`

GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerakiVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) SetMerakiVpnPeers(v []InlineResponse200238MerakiVpnPeers)`

SetMerakiVpnPeers sets MerakiVpnPeers field to given value.

### HasMerakiVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) HasMerakiVpnPeers() bool`

HasMerakiVpnPeers returns a boolean if a field has been set.

### GetThirdPartyVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) GetThirdPartyVpnPeers() []InlineResponse200238ThirdPartyVpnPeers`

GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field if non-nil, zero value otherwise.

### GetThirdPartyVpnPeersOk

`func (o *InlineResponse200238Vpnstatusentities) GetThirdPartyVpnPeersOk() (*[]InlineResponse200238ThirdPartyVpnPeers, bool)`

GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) SetThirdPartyVpnPeers(v []InlineResponse200238ThirdPartyVpnPeers)`

SetThirdPartyVpnPeers sets ThirdPartyVpnPeers field to given value.

### HasThirdPartyVpnPeers

`func (o *InlineResponse200238Vpnstatusentities) HasThirdPartyVpnPeers() bool`

HasThirdPartyVpnPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


