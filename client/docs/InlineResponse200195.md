# InlineResponse200195

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse200195PowerExceptions**](InlineResponse200195PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**InlineResponse200195UplinkClientSampling**](InlineResponse200195UplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**InlineResponse200195MacBlocklist**](InlineResponse200195MacBlocklist.md) |  | [optional] 
**PortChannelFallback** | Pointer to **bool** | Enable port channel fallback on Catalyst MS switches. | [optional] 
**UplinkSelection** | Pointer to [**InlineResponse200195UplinkSelection**](InlineResponse200195UplinkSelection.md) |  | [optional] 

## Methods

### NewInlineResponse200195

`func NewInlineResponse200195() *InlineResponse200195`

NewInlineResponse200195 instantiates a new InlineResponse200195 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200195WithDefaults

`func NewInlineResponse200195WithDefaults() *InlineResponse200195`

NewInlineResponse200195WithDefaults instantiates a new InlineResponse200195 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse200195) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200195) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200195) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200195) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse200195) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse200195) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse200195) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse200195) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse200195) GetPowerExceptions() []InlineResponse200195PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse200195) GetPowerExceptionsOk() (*[]InlineResponse200195PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse200195) SetPowerExceptions(v []InlineResponse200195PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse200195) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineResponse200195) GetUplinkClientSampling() InlineResponse200195UplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineResponse200195) GetUplinkClientSamplingOk() (*InlineResponse200195UplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineResponse200195) SetUplinkClientSampling(v InlineResponse200195UplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineResponse200195) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineResponse200195) GetMacBlocklist() InlineResponse200195MacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineResponse200195) GetMacBlocklistOk() (*InlineResponse200195MacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineResponse200195) SetMacBlocklist(v InlineResponse200195MacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineResponse200195) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.

### GetPortChannelFallback

`func (o *InlineResponse200195) GetPortChannelFallback() bool`

GetPortChannelFallback returns the PortChannelFallback field if non-nil, zero value otherwise.

### GetPortChannelFallbackOk

`func (o *InlineResponse200195) GetPortChannelFallbackOk() (*bool, bool)`

GetPortChannelFallbackOk returns a tuple with the PortChannelFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelFallback

`func (o *InlineResponse200195) SetPortChannelFallback(v bool)`

SetPortChannelFallback sets PortChannelFallback field to given value.

### HasPortChannelFallback

`func (o *InlineResponse200195) HasPortChannelFallback() bool`

HasPortChannelFallback returns a boolean if a field has been set.

### GetUplinkSelection

`func (o *InlineResponse200195) GetUplinkSelection() InlineResponse200195UplinkSelection`

GetUplinkSelection returns the UplinkSelection field if non-nil, zero value otherwise.

### GetUplinkSelectionOk

`func (o *InlineResponse200195) GetUplinkSelectionOk() (*InlineResponse200195UplinkSelection, bool)`

GetUplinkSelectionOk returns a tuple with the UplinkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSelection

`func (o *InlineResponse200195) SetUplinkSelection(v InlineResponse200195UplinkSelection)`

SetUplinkSelection sets UplinkSelection field to given value.

### HasUplinkSelection

`func (o *InlineResponse200195) HasUplinkSelection() bool`

HasUplinkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


