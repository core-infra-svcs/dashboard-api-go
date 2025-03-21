# InlineResponse200353Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Description** | Pointer to **string** | The description of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Status** | Pointer to **string** | The status of the wireless LAN controller interface | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the interface is operating at. enum &#x3D; [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps] | [optional] 
**Addresses** | Pointer to [**[]InlineResponse200353Addresses**](InlineResponse200353Addresses.md) | Available addresses for the interface. | [optional] 
**Vrf** | Pointer to [**InlineResponse200353Vrf**](InlineResponse200353Vrf.md) |  | [optional] 
**IsUplink** | Pointer to **bool** | Indicate whether the interface is uplink | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**LinkNegotiation** | Pointer to **string** | The interface negotiation mode | [optional] 
**ChannelGroup** | Pointer to [**InlineResponse200350ChannelGroup**](InlineResponse200350ChannelGroup.md) |  | [optional] 
**Module** | Pointer to [**InlineResponse200350Module**](InlineResponse200350Module.md) |  | [optional] 

## Methods

### NewInlineResponse200353Interfaces

`func NewInlineResponse200353Interfaces() *InlineResponse200353Interfaces`

NewInlineResponse200353Interfaces instantiates a new InlineResponse200353Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200353InterfacesWithDefaults

`func NewInlineResponse200353InterfacesWithDefaults() *InlineResponse200353Interfaces`

NewInlineResponse200353InterfacesWithDefaults instantiates a new InlineResponse200353Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200353Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200353Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200353Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200353Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200353Interfaces) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200353Interfaces) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200353Interfaces) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200353Interfaces) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200353Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200353Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200353Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200353Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200353Interfaces) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200353Interfaces) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200353Interfaces) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200353Interfaces) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse200353Interfaces) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse200353Interfaces) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse200353Interfaces) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse200353Interfaces) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse200353Interfaces) GetAddresses() []InlineResponse200353Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse200353Interfaces) GetAddressesOk() (*[]InlineResponse200353Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse200353Interfaces) SetAddresses(v []InlineResponse200353Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse200353Interfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse200353Interfaces) GetVrf() InlineResponse200353Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200353Interfaces) GetVrfOk() (*InlineResponse200353Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200353Interfaces) SetVrf(v InlineResponse200353Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200353Interfaces) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse200353Interfaces) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse200353Interfaces) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse200353Interfaces) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse200353Interfaces) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200353Interfaces) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200353Interfaces) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200353Interfaces) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200353Interfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineResponse200353Interfaces) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineResponse200353Interfaces) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineResponse200353Interfaces) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineResponse200353Interfaces) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetChannelGroup

`func (o *InlineResponse200353Interfaces) GetChannelGroup() InlineResponse200350ChannelGroup`

GetChannelGroup returns the ChannelGroup field if non-nil, zero value otherwise.

### GetChannelGroupOk

`func (o *InlineResponse200353Interfaces) GetChannelGroupOk() (*InlineResponse200350ChannelGroup, bool)`

GetChannelGroupOk returns a tuple with the ChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGroup

`func (o *InlineResponse200353Interfaces) SetChannelGroup(v InlineResponse200350ChannelGroup)`

SetChannelGroup sets ChannelGroup field to given value.

### HasChannelGroup

`func (o *InlineResponse200353Interfaces) HasChannelGroup() bool`

HasChannelGroup returns a boolean if a field has been set.

### GetModule

`func (o *InlineResponse200353Interfaces) GetModule() InlineResponse200350Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *InlineResponse200353Interfaces) GetModuleOk() (*InlineResponse200350Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *InlineResponse200353Interfaces) SetModule(v InlineResponse200350Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *InlineResponse200353Interfaces) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


