# InlineResponse200354Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Description** | Pointer to **string** | The description of the wireless LAN controller interface | [optional] 
**Enabled** | Pointer to **bool** | The status of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Status** | Pointer to **string** | The status of the wireless LAN controller interface | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the interface is operating at. enum &#x3D; [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps] | [optional] 
**IsUplink** | Pointer to **bool** | Indicate whether the interface is uplink | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**IsRedundancyPort** | Pointer to **bool** | Indicate whether the interface is a redundancy port used to perform HA role negotiation | [optional] 
**LinkNegotiation** | Pointer to **string** | The interface negotiation mode | [optional] 
**ChannelGroup** | Pointer to [**InlineResponse200354ChannelGroup**](InlineResponse200354ChannelGroup.md) |  | [optional] 
**Module** | Pointer to [**InlineResponse200354Module**](InlineResponse200354Module.md) |  | [optional] 

## Methods

### NewInlineResponse200354Interfaces

`func NewInlineResponse200354Interfaces() *InlineResponse200354Interfaces`

NewInlineResponse200354Interfaces instantiates a new InlineResponse200354Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200354InterfacesWithDefaults

`func NewInlineResponse200354InterfacesWithDefaults() *InlineResponse200354Interfaces`

NewInlineResponse200354InterfacesWithDefaults instantiates a new InlineResponse200354Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200354Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200354Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200354Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200354Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200354Interfaces) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200354Interfaces) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200354Interfaces) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200354Interfaces) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200354Interfaces) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200354Interfaces) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200354Interfaces) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200354Interfaces) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200354Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200354Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200354Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200354Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200354Interfaces) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200354Interfaces) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200354Interfaces) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200354Interfaces) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse200354Interfaces) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse200354Interfaces) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse200354Interfaces) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse200354Interfaces) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse200354Interfaces) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse200354Interfaces) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse200354Interfaces) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse200354Interfaces) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200354Interfaces) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200354Interfaces) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200354Interfaces) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200354Interfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetIsRedundancyPort

`func (o *InlineResponse200354Interfaces) GetIsRedundancyPort() bool`

GetIsRedundancyPort returns the IsRedundancyPort field if non-nil, zero value otherwise.

### GetIsRedundancyPortOk

`func (o *InlineResponse200354Interfaces) GetIsRedundancyPortOk() (*bool, bool)`

GetIsRedundancyPortOk returns a tuple with the IsRedundancyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedundancyPort

`func (o *InlineResponse200354Interfaces) SetIsRedundancyPort(v bool)`

SetIsRedundancyPort sets IsRedundancyPort field to given value.

### HasIsRedundancyPort

`func (o *InlineResponse200354Interfaces) HasIsRedundancyPort() bool`

HasIsRedundancyPort returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineResponse200354Interfaces) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineResponse200354Interfaces) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineResponse200354Interfaces) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineResponse200354Interfaces) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetChannelGroup

`func (o *InlineResponse200354Interfaces) GetChannelGroup() InlineResponse200354ChannelGroup`

GetChannelGroup returns the ChannelGroup field if non-nil, zero value otherwise.

### GetChannelGroupOk

`func (o *InlineResponse200354Interfaces) GetChannelGroupOk() (*InlineResponse200354ChannelGroup, bool)`

GetChannelGroupOk returns a tuple with the ChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGroup

`func (o *InlineResponse200354Interfaces) SetChannelGroup(v InlineResponse200354ChannelGroup)`

SetChannelGroup sets ChannelGroup field to given value.

### HasChannelGroup

`func (o *InlineResponse200354Interfaces) HasChannelGroup() bool`

HasChannelGroup returns a boolean if a field has been set.

### GetModule

`func (o *InlineResponse200354Interfaces) GetModule() InlineResponse200354Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *InlineResponse200354Interfaces) GetModuleOk() (*InlineResponse200354Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *InlineResponse200354Interfaces) SetModule(v InlineResponse200354Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *InlineResponse200354Interfaces) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


