# InlineResponse200357Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Description** | Pointer to **string** | The description of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Status** | Pointer to **string** | The status of the wireless LAN controller interface | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the interface is operating at. enum &#x3D; [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps] | [optional] 
**Addresses** | Pointer to [**[]InlineResponse200357Addresses**](InlineResponse200357Addresses.md) | Available addresses for the interface. | [optional] 
**Vrf** | Pointer to [**InlineResponse200357Vrf**](InlineResponse200357Vrf.md) |  | [optional] 
**IsUplink** | Pointer to **bool** | Indicate whether the interface is uplink | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**LinkNegotiation** | Pointer to **string** | The interface negotiation mode | [optional] 
**ChannelGroup** | Pointer to [**InlineResponse200354ChannelGroup**](InlineResponse200354ChannelGroup.md) |  | [optional] 
**Module** | Pointer to [**InlineResponse200354Module**](InlineResponse200354Module.md) |  | [optional] 

## Methods

### NewInlineResponse200357Interfaces

`func NewInlineResponse200357Interfaces() *InlineResponse200357Interfaces`

NewInlineResponse200357Interfaces instantiates a new InlineResponse200357Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200357InterfacesWithDefaults

`func NewInlineResponse200357InterfacesWithDefaults() *InlineResponse200357Interfaces`

NewInlineResponse200357InterfacesWithDefaults instantiates a new InlineResponse200357Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200357Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200357Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200357Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200357Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200357Interfaces) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200357Interfaces) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200357Interfaces) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200357Interfaces) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200357Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200357Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200357Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200357Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200357Interfaces) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200357Interfaces) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200357Interfaces) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200357Interfaces) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse200357Interfaces) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse200357Interfaces) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse200357Interfaces) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse200357Interfaces) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse200357Interfaces) GetAddresses() []InlineResponse200357Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse200357Interfaces) GetAddressesOk() (*[]InlineResponse200357Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse200357Interfaces) SetAddresses(v []InlineResponse200357Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse200357Interfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse200357Interfaces) GetVrf() InlineResponse200357Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200357Interfaces) GetVrfOk() (*InlineResponse200357Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200357Interfaces) SetVrf(v InlineResponse200357Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200357Interfaces) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse200357Interfaces) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse200357Interfaces) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse200357Interfaces) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse200357Interfaces) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200357Interfaces) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200357Interfaces) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200357Interfaces) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200357Interfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineResponse200357Interfaces) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineResponse200357Interfaces) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineResponse200357Interfaces) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineResponse200357Interfaces) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetChannelGroup

`func (o *InlineResponse200357Interfaces) GetChannelGroup() InlineResponse200354ChannelGroup`

GetChannelGroup returns the ChannelGroup field if non-nil, zero value otherwise.

### GetChannelGroupOk

`func (o *InlineResponse200357Interfaces) GetChannelGroupOk() (*InlineResponse200354ChannelGroup, bool)`

GetChannelGroupOk returns a tuple with the ChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGroup

`func (o *InlineResponse200357Interfaces) SetChannelGroup(v InlineResponse200354ChannelGroup)`

SetChannelGroup sets ChannelGroup field to given value.

### HasChannelGroup

`func (o *InlineResponse200357Interfaces) HasChannelGroup() bool`

HasChannelGroup returns a boolean if a field has been set.

### GetModule

`func (o *InlineResponse200357Interfaces) GetModule() InlineResponse200354Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *InlineResponse200357Interfaces) GetModuleOk() (*InlineResponse200354Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *InlineResponse200357Interfaces) SetModule(v InlineResponse200354Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *InlineResponse200357Interfaces) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


