# InlineResponse200204

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse200204PowerExceptions**](InlineResponse200204PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**InlineResponse200204UplinkClientSampling**](InlineResponse200204UplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**InlineResponse200204MacBlocklist**](InlineResponse200204MacBlocklist.md) |  | [optional] 
**PortChannelFallback** | Pointer to **bool** | Enable port channel fallback on Catalyst MS switches. | [optional] 
**UplinkSelection** | Pointer to [**InlineResponse200204UplinkSelection**](InlineResponse200204UplinkSelection.md) |  | [optional] 

## Methods

### NewInlineResponse200204

`func NewInlineResponse200204() *InlineResponse200204`

NewInlineResponse200204 instantiates a new InlineResponse200204 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200204WithDefaults

`func NewInlineResponse200204WithDefaults() *InlineResponse200204`

NewInlineResponse200204WithDefaults instantiates a new InlineResponse200204 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse200204) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200204) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200204) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200204) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse200204) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse200204) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse200204) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse200204) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse200204) GetPowerExceptions() []InlineResponse200204PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse200204) GetPowerExceptionsOk() (*[]InlineResponse200204PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse200204) SetPowerExceptions(v []InlineResponse200204PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse200204) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineResponse200204) GetUplinkClientSampling() InlineResponse200204UplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineResponse200204) GetUplinkClientSamplingOk() (*InlineResponse200204UplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineResponse200204) SetUplinkClientSampling(v InlineResponse200204UplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineResponse200204) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineResponse200204) GetMacBlocklist() InlineResponse200204MacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineResponse200204) GetMacBlocklistOk() (*InlineResponse200204MacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineResponse200204) SetMacBlocklist(v InlineResponse200204MacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineResponse200204) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.

### GetPortChannelFallback

`func (o *InlineResponse200204) GetPortChannelFallback() bool`

GetPortChannelFallback returns the PortChannelFallback field if non-nil, zero value otherwise.

### GetPortChannelFallbackOk

`func (o *InlineResponse200204) GetPortChannelFallbackOk() (*bool, bool)`

GetPortChannelFallbackOk returns a tuple with the PortChannelFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelFallback

`func (o *InlineResponse200204) SetPortChannelFallback(v bool)`

SetPortChannelFallback sets PortChannelFallback field to given value.

### HasPortChannelFallback

`func (o *InlineResponse200204) HasPortChannelFallback() bool`

HasPortChannelFallback returns a boolean if a field has been set.

### GetUplinkSelection

`func (o *InlineResponse200204) GetUplinkSelection() InlineResponse200204UplinkSelection`

GetUplinkSelection returns the UplinkSelection field if non-nil, zero value otherwise.

### GetUplinkSelectionOk

`func (o *InlineResponse200204) GetUplinkSelectionOk() (*InlineResponse200204UplinkSelection, bool)`

GetUplinkSelectionOk returns a tuple with the UplinkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSelection

`func (o *InlineResponse200204) SetUplinkSelection(v InlineResponse200204UplinkSelection)`

SetUplinkSelection sets UplinkSelection field to given value.

### HasUplinkSelection

`func (o *InlineResponse200204) HasUplinkSelection() bool`

HasUplinkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


