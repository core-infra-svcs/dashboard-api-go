# InlineResponse200156

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse200156PowerExceptions**](InlineResponse200156PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**InlineResponse200156UplinkClientSampling**](InlineResponse200156UplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**InlineResponse200156MacBlocklist**](InlineResponse200156MacBlocklist.md) |  | [optional] 

## Methods

### NewInlineResponse200156

`func NewInlineResponse200156() *InlineResponse200156`

NewInlineResponse200156 instantiates a new InlineResponse200156 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200156WithDefaults

`func NewInlineResponse200156WithDefaults() *InlineResponse200156`

NewInlineResponse200156WithDefaults instantiates a new InlineResponse200156 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse200156) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200156) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200156) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200156) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse200156) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse200156) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse200156) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse200156) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse200156) GetPowerExceptions() []InlineResponse200156PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse200156) GetPowerExceptionsOk() (*[]InlineResponse200156PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse200156) SetPowerExceptions(v []InlineResponse200156PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse200156) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineResponse200156) GetUplinkClientSampling() InlineResponse200156UplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineResponse200156) GetUplinkClientSamplingOk() (*InlineResponse200156UplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineResponse200156) SetUplinkClientSampling(v InlineResponse200156UplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineResponse200156) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineResponse200156) GetMacBlocklist() InlineResponse200156MacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineResponse200156) GetMacBlocklistOk() (*InlineResponse200156MacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineResponse200156) SetMacBlocklist(v InlineResponse200156MacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineResponse200156) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


