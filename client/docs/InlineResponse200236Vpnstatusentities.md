# InlineResponse200236Vpnstatusentities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device | [optional] 
**DeviceStatus** | Pointer to **string** | Device Status | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200236Uplinks**](InlineResponse200236Uplinks.md) | List of Uplink Information | [optional] 
**VpnMode** | Pointer to **string** | VPN Mode | [optional] 
**ExportedSubnets** | Pointer to [**[]InlineResponse200236ExportedSubnets**](InlineResponse200236ExportedSubnets.md) | List of Exported Subnets | [optional] 
**MerakiVpnPeers** | Pointer to [**[]InlineResponse200236MerakiVpnPeers**](InlineResponse200236MerakiVpnPeers.md) | Meraki VPN Peers | [optional] 
**ThirdPartyVpnPeers** | Pointer to [**[]InlineResponse200236ThirdPartyVpnPeers**](InlineResponse200236ThirdPartyVpnPeers.md) | Third Party VPN Peers | [optional] 

## Methods

### NewInlineResponse200236Vpnstatusentities

`func NewInlineResponse200236Vpnstatusentities() *InlineResponse200236Vpnstatusentities`

NewInlineResponse200236Vpnstatusentities instantiates a new InlineResponse200236Vpnstatusentities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200236VpnstatusentitiesWithDefaults

`func NewInlineResponse200236VpnstatusentitiesWithDefaults() *InlineResponse200236Vpnstatusentities`

NewInlineResponse200236VpnstatusentitiesWithDefaults instantiates a new InlineResponse200236Vpnstatusentities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200236Vpnstatusentities) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200236Vpnstatusentities) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200236Vpnstatusentities) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200236Vpnstatusentities) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200236Vpnstatusentities) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200236Vpnstatusentities) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200236Vpnstatusentities) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200236Vpnstatusentities) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200236Vpnstatusentities) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200236Vpnstatusentities) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200236Vpnstatusentities) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200236Vpnstatusentities) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200236Vpnstatusentities) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200236Vpnstatusentities) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200236Vpnstatusentities) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200236Vpnstatusentities) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200236Vpnstatusentities) GetUplinks() []InlineResponse200236Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200236Vpnstatusentities) GetUplinksOk() (*[]InlineResponse200236Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200236Vpnstatusentities) SetUplinks(v []InlineResponse200236Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200236Vpnstatusentities) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetVpnMode

`func (o *InlineResponse200236Vpnstatusentities) GetVpnMode() string`

GetVpnMode returns the VpnMode field if non-nil, zero value otherwise.

### GetVpnModeOk

`func (o *InlineResponse200236Vpnstatusentities) GetVpnModeOk() (*string, bool)`

GetVpnModeOk returns a tuple with the VpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMode

`func (o *InlineResponse200236Vpnstatusentities) SetVpnMode(v string)`

SetVpnMode sets VpnMode field to given value.

### HasVpnMode

`func (o *InlineResponse200236Vpnstatusentities) HasVpnMode() bool`

HasVpnMode returns a boolean if a field has been set.

### GetExportedSubnets

`func (o *InlineResponse200236Vpnstatusentities) GetExportedSubnets() []InlineResponse200236ExportedSubnets`

GetExportedSubnets returns the ExportedSubnets field if non-nil, zero value otherwise.

### GetExportedSubnetsOk

`func (o *InlineResponse200236Vpnstatusentities) GetExportedSubnetsOk() (*[]InlineResponse200236ExportedSubnets, bool)`

GetExportedSubnetsOk returns a tuple with the ExportedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedSubnets

`func (o *InlineResponse200236Vpnstatusentities) SetExportedSubnets(v []InlineResponse200236ExportedSubnets)`

SetExportedSubnets sets ExportedSubnets field to given value.

### HasExportedSubnets

`func (o *InlineResponse200236Vpnstatusentities) HasExportedSubnets() bool`

HasExportedSubnets returns a boolean if a field has been set.

### GetMerakiVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) GetMerakiVpnPeers() []InlineResponse200236MerakiVpnPeers`

GetMerakiVpnPeers returns the MerakiVpnPeers field if non-nil, zero value otherwise.

### GetMerakiVpnPeersOk

`func (o *InlineResponse200236Vpnstatusentities) GetMerakiVpnPeersOk() (*[]InlineResponse200236MerakiVpnPeers, bool)`

GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerakiVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) SetMerakiVpnPeers(v []InlineResponse200236MerakiVpnPeers)`

SetMerakiVpnPeers sets MerakiVpnPeers field to given value.

### HasMerakiVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) HasMerakiVpnPeers() bool`

HasMerakiVpnPeers returns a boolean if a field has been set.

### GetThirdPartyVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) GetThirdPartyVpnPeers() []InlineResponse200236ThirdPartyVpnPeers`

GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field if non-nil, zero value otherwise.

### GetThirdPartyVpnPeersOk

`func (o *InlineResponse200236Vpnstatusentities) GetThirdPartyVpnPeersOk() (*[]InlineResponse200236ThirdPartyVpnPeers, bool)`

GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) SetThirdPartyVpnPeers(v []InlineResponse200236ThirdPartyVpnPeers)`

SetThirdPartyVpnPeers sets ThirdPartyVpnPeers field to given value.

### HasThirdPartyVpnPeers

`func (o *InlineResponse200236Vpnstatusentities) HasThirdPartyVpnPeers() bool`

HasThirdPartyVpnPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


