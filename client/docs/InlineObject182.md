# InlineObject182

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]NetworksNetworkIdSwitchSettingsPowerExceptions**](NetworksNetworkIdSwitchSettingsPowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**NetworksNetworkIdSwitchSettingsUplinkClientSampling**](NetworksNetworkIdSwitchSettingsUplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**NetworksNetworkIdSwitchSettingsMacBlocklist**](NetworksNetworkIdSwitchSettingsMacBlocklist.md) |  | [optional] 
**ClientTracking** | Pointer to [**InlineResponse200209ClientTracking**](InlineResponse200209ClientTracking.md) |  | [optional] 
**PortChannelFallback** | Pointer to **bool** | Port channel fallback | [optional] 
**UplinkSelection** | Pointer to [**NetworksNetworkIdSwitchSettingsUplinkSelection**](NetworksNetworkIdSwitchSettingsUplinkSelection.md) |  | [optional] 

## Methods

### NewInlineObject182

`func NewInlineObject182() *InlineObject182`

NewInlineObject182 instantiates a new InlineObject182 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject182WithDefaults

`func NewInlineObject182WithDefaults() *InlineObject182`

NewInlineObject182WithDefaults instantiates a new InlineObject182 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineObject182) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject182) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject182) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject182) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineObject182) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineObject182) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineObject182) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineObject182) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineObject182) GetPowerExceptions() []NetworksNetworkIdSwitchSettingsPowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineObject182) GetPowerExceptionsOk() (*[]NetworksNetworkIdSwitchSettingsPowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineObject182) SetPowerExceptions(v []NetworksNetworkIdSwitchSettingsPowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineObject182) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineObject182) GetUplinkClientSampling() NetworksNetworkIdSwitchSettingsUplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineObject182) GetUplinkClientSamplingOk() (*NetworksNetworkIdSwitchSettingsUplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineObject182) SetUplinkClientSampling(v NetworksNetworkIdSwitchSettingsUplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineObject182) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineObject182) GetMacBlocklist() NetworksNetworkIdSwitchSettingsMacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineObject182) GetMacBlocklistOk() (*NetworksNetworkIdSwitchSettingsMacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineObject182) SetMacBlocklist(v NetworksNetworkIdSwitchSettingsMacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineObject182) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.

### GetClientTracking

`func (o *InlineObject182) GetClientTracking() InlineResponse200209ClientTracking`

GetClientTracking returns the ClientTracking field if non-nil, zero value otherwise.

### GetClientTrackingOk

`func (o *InlineObject182) GetClientTrackingOk() (*InlineResponse200209ClientTracking, bool)`

GetClientTrackingOk returns a tuple with the ClientTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTracking

`func (o *InlineObject182) SetClientTracking(v InlineResponse200209ClientTracking)`

SetClientTracking sets ClientTracking field to given value.

### HasClientTracking

`func (o *InlineObject182) HasClientTracking() bool`

HasClientTracking returns a boolean if a field has been set.

### GetPortChannelFallback

`func (o *InlineObject182) GetPortChannelFallback() bool`

GetPortChannelFallback returns the PortChannelFallback field if non-nil, zero value otherwise.

### GetPortChannelFallbackOk

`func (o *InlineObject182) GetPortChannelFallbackOk() (*bool, bool)`

GetPortChannelFallbackOk returns a tuple with the PortChannelFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelFallback

`func (o *InlineObject182) SetPortChannelFallback(v bool)`

SetPortChannelFallback sets PortChannelFallback field to given value.

### HasPortChannelFallback

`func (o *InlineObject182) HasPortChannelFallback() bool`

HasPortChannelFallback returns a boolean if a field has been set.

### GetUplinkSelection

`func (o *InlineObject182) GetUplinkSelection() NetworksNetworkIdSwitchSettingsUplinkSelection`

GetUplinkSelection returns the UplinkSelection field if non-nil, zero value otherwise.

### GetUplinkSelectionOk

`func (o *InlineObject182) GetUplinkSelectionOk() (*NetworksNetworkIdSwitchSettingsUplinkSelection, bool)`

GetUplinkSelectionOk returns a tuple with the UplinkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSelection

`func (o *InlineObject182) SetUplinkSelection(v NetworksNetworkIdSwitchSettingsUplinkSelection)`

SetUplinkSelection sets UplinkSelection field to given value.

### HasUplinkSelection

`func (o *InlineObject182) HasUplinkSelection() bool`

HasUplinkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


