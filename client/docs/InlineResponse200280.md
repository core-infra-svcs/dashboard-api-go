# InlineResponse200280

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device | [optional] 
**DeviceStatus** | Pointer to **string** | Device Status | [optional] 
**Uplinks** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnStatusesUplinks**](OrganizationsOrganizationIdApplianceVpnStatusesUplinks.md) | List of Uplink Information | [optional] 
**VpnMode** | Pointer to **string** | VPN Mode | [optional] 
**ExportedSubnets** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnStatusesExportedSubnets**](OrganizationsOrganizationIdApplianceVpnStatusesExportedSubnets.md) | List of Exported Subnets | [optional] 
**MerakiVpnPeers** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnStatusesMerakiVpnPeers**](OrganizationsOrganizationIdApplianceVpnStatusesMerakiVpnPeers.md) | Meraki VPN Peers | [optional] 
**ThirdPartyVpnPeers** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnStatusesThirdPartyVpnPeers**](OrganizationsOrganizationIdApplianceVpnStatusesThirdPartyVpnPeers.md) | Third Party VPN Peers | [optional] 

## Methods

### NewInlineResponse200280

`func NewInlineResponse200280() *InlineResponse200280`

NewInlineResponse200280 instantiates a new InlineResponse200280 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200280WithDefaults

`func NewInlineResponse200280WithDefaults() *InlineResponse200280`

NewInlineResponse200280WithDefaults instantiates a new InlineResponse200280 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200280) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200280) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200280) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200280) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200280) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200280) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200280) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200280) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200280) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200280) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200280) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200280) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200280) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200280) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200280) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200280) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200280) GetUplinks() []OrganizationsOrganizationIdApplianceVpnStatusesUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200280) GetUplinksOk() (*[]OrganizationsOrganizationIdApplianceVpnStatusesUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200280) SetUplinks(v []OrganizationsOrganizationIdApplianceVpnStatusesUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200280) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetVpnMode

`func (o *InlineResponse200280) GetVpnMode() string`

GetVpnMode returns the VpnMode field if non-nil, zero value otherwise.

### GetVpnModeOk

`func (o *InlineResponse200280) GetVpnModeOk() (*string, bool)`

GetVpnModeOk returns a tuple with the VpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMode

`func (o *InlineResponse200280) SetVpnMode(v string)`

SetVpnMode sets VpnMode field to given value.

### HasVpnMode

`func (o *InlineResponse200280) HasVpnMode() bool`

HasVpnMode returns a boolean if a field has been set.

### GetExportedSubnets

`func (o *InlineResponse200280) GetExportedSubnets() []OrganizationsOrganizationIdApplianceVpnStatusesExportedSubnets`

GetExportedSubnets returns the ExportedSubnets field if non-nil, zero value otherwise.

### GetExportedSubnetsOk

`func (o *InlineResponse200280) GetExportedSubnetsOk() (*[]OrganizationsOrganizationIdApplianceVpnStatusesExportedSubnets, bool)`

GetExportedSubnetsOk returns a tuple with the ExportedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedSubnets

`func (o *InlineResponse200280) SetExportedSubnets(v []OrganizationsOrganizationIdApplianceVpnStatusesExportedSubnets)`

SetExportedSubnets sets ExportedSubnets field to given value.

### HasExportedSubnets

`func (o *InlineResponse200280) HasExportedSubnets() bool`

HasExportedSubnets returns a boolean if a field has been set.

### GetMerakiVpnPeers

`func (o *InlineResponse200280) GetMerakiVpnPeers() []OrganizationsOrganizationIdApplianceVpnStatusesMerakiVpnPeers`

GetMerakiVpnPeers returns the MerakiVpnPeers field if non-nil, zero value otherwise.

### GetMerakiVpnPeersOk

`func (o *InlineResponse200280) GetMerakiVpnPeersOk() (*[]OrganizationsOrganizationIdApplianceVpnStatusesMerakiVpnPeers, bool)`

GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerakiVpnPeers

`func (o *InlineResponse200280) SetMerakiVpnPeers(v []OrganizationsOrganizationIdApplianceVpnStatusesMerakiVpnPeers)`

SetMerakiVpnPeers sets MerakiVpnPeers field to given value.

### HasMerakiVpnPeers

`func (o *InlineResponse200280) HasMerakiVpnPeers() bool`

HasMerakiVpnPeers returns a boolean if a field has been set.

### GetThirdPartyVpnPeers

`func (o *InlineResponse200280) GetThirdPartyVpnPeers() []OrganizationsOrganizationIdApplianceVpnStatusesThirdPartyVpnPeers`

GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field if non-nil, zero value otherwise.

### GetThirdPartyVpnPeersOk

`func (o *InlineResponse200280) GetThirdPartyVpnPeersOk() (*[]OrganizationsOrganizationIdApplianceVpnStatusesThirdPartyVpnPeers, bool)`

GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyVpnPeers

`func (o *InlineResponse200280) SetThirdPartyVpnPeers(v []OrganizationsOrganizationIdApplianceVpnStatusesThirdPartyVpnPeers)`

SetThirdPartyVpnPeers sets ThirdPartyVpnPeers field to given value.

### HasThirdPartyVpnPeers

`func (o *InlineResponse200280) HasThirdPartyVpnPeers() bool`

HasThirdPartyVpnPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


