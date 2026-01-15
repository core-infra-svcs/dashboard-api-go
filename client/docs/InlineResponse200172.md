# InlineResponse200172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse200172PowerExceptions**](InlineResponse200172PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**InlineResponse200172UplinkClientSampling**](InlineResponse200172UplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**InlineResponse200172MacBlocklist**](InlineResponse200172MacBlocklist.md) |  | [optional] 
**UplinkSelection** | Pointer to [**InlineResponse200172UplinkSelection**](InlineResponse200172UplinkSelection.md) |  | [optional] 

## Methods

### NewInlineResponse200172

`func NewInlineResponse200172() *InlineResponse200172`

NewInlineResponse200172 instantiates a new InlineResponse200172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200172WithDefaults

`func NewInlineResponse200172WithDefaults() *InlineResponse200172`

NewInlineResponse200172WithDefaults instantiates a new InlineResponse200172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse200172) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200172) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200172) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200172) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse200172) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse200172) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse200172) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse200172) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse200172) GetPowerExceptions() []InlineResponse200172PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse200172) GetPowerExceptionsOk() (*[]InlineResponse200172PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse200172) SetPowerExceptions(v []InlineResponse200172PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse200172) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineResponse200172) GetUplinkClientSampling() InlineResponse200172UplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineResponse200172) GetUplinkClientSamplingOk() (*InlineResponse200172UplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineResponse200172) SetUplinkClientSampling(v InlineResponse200172UplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineResponse200172) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineResponse200172) GetMacBlocklist() InlineResponse200172MacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineResponse200172) GetMacBlocklistOk() (*InlineResponse200172MacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineResponse200172) SetMacBlocklist(v InlineResponse200172MacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineResponse200172) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.

### GetUplinkSelection

`func (o *InlineResponse200172) GetUplinkSelection() InlineResponse200172UplinkSelection`

GetUplinkSelection returns the UplinkSelection field if non-nil, zero value otherwise.

### GetUplinkSelectionOk

`func (o *InlineResponse200172) GetUplinkSelectionOk() (*InlineResponse200172UplinkSelection, bool)`

GetUplinkSelectionOk returns a tuple with the UplinkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSelection

`func (o *InlineResponse200172) SetUplinkSelection(v InlineResponse200172UplinkSelection)`

SetUplinkSelection sets UplinkSelection field to given value.

### HasUplinkSelection

`func (o *InlineResponse200172) HasUplinkSelection() bool`

HasUplinkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


