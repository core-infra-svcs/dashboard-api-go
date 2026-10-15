# InlineResponse200530Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Description** | Pointer to **string** | The description of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Status** | Pointer to **string** | The status of the wireless LAN controller interface | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the interface is operating at. enum &#x3D; [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps] | [optional] 
**Addresses** | Pointer to [**[]InlineResponse200530Addresses**](InlineResponse200530Addresses.md) | Available addresses for the interface. | [optional] 
**Vrf** | Pointer to [**InlineResponse200530Vrf**](InlineResponse200530Vrf.md) |  | [optional] 
**IsUplink** | Pointer to **bool** | Indicate whether the interface is uplink | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**LinkNegotiation** | Pointer to **string** | The interface negotiation mode | [optional] 
**ChannelGroup** | Pointer to [**InlineResponse200527ChannelGroup**](InlineResponse200527ChannelGroup.md) |  | [optional] 
**Module** | Pointer to [**InlineResponse200527Module**](InlineResponse200527Module.md) |  | [optional] 

## Methods

### NewInlineResponse200530Interfaces

`func NewInlineResponse200530Interfaces() *InlineResponse200530Interfaces`

NewInlineResponse200530Interfaces instantiates a new InlineResponse200530Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200530InterfacesWithDefaults

`func NewInlineResponse200530InterfacesWithDefaults() *InlineResponse200530Interfaces`

NewInlineResponse200530InterfacesWithDefaults instantiates a new InlineResponse200530Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200530Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200530Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200530Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200530Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200530Interfaces) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200530Interfaces) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200530Interfaces) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200530Interfaces) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200530Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200530Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200530Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200530Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200530Interfaces) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200530Interfaces) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200530Interfaces) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200530Interfaces) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse200530Interfaces) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse200530Interfaces) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse200530Interfaces) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse200530Interfaces) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse200530Interfaces) GetAddresses() []InlineResponse200530Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse200530Interfaces) GetAddressesOk() (*[]InlineResponse200530Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse200530Interfaces) SetAddresses(v []InlineResponse200530Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse200530Interfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse200530Interfaces) GetVrf() InlineResponse200530Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200530Interfaces) GetVrfOk() (*InlineResponse200530Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200530Interfaces) SetVrf(v InlineResponse200530Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200530Interfaces) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse200530Interfaces) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse200530Interfaces) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse200530Interfaces) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse200530Interfaces) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200530Interfaces) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200530Interfaces) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200530Interfaces) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200530Interfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineResponse200530Interfaces) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineResponse200530Interfaces) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineResponse200530Interfaces) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineResponse200530Interfaces) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetChannelGroup

`func (o *InlineResponse200530Interfaces) GetChannelGroup() InlineResponse200527ChannelGroup`

GetChannelGroup returns the ChannelGroup field if non-nil, zero value otherwise.

### GetChannelGroupOk

`func (o *InlineResponse200530Interfaces) GetChannelGroupOk() (*InlineResponse200527ChannelGroup, bool)`

GetChannelGroupOk returns a tuple with the ChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGroup

`func (o *InlineResponse200530Interfaces) SetChannelGroup(v InlineResponse200527ChannelGroup)`

SetChannelGroup sets ChannelGroup field to given value.

### HasChannelGroup

`func (o *InlineResponse200530Interfaces) HasChannelGroup() bool`

HasChannelGroup returns a boolean if a field has been set.

### GetModule

`func (o *InlineResponse200530Interfaces) GetModule() InlineResponse200527Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *InlineResponse200530Interfaces) GetModuleOk() (*InlineResponse200527Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *InlineResponse200530Interfaces) SetModule(v InlineResponse200527Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *InlineResponse200530Interfaces) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


