# InlineResponse200209

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse200209PowerExceptions**](InlineResponse200209PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 
**UplinkClientSampling** | Pointer to [**InlineResponse200209UplinkClientSampling**](InlineResponse200209UplinkClientSampling.md) |  | [optional] 
**MacBlocklist** | Pointer to [**InlineResponse200209MacBlocklist**](InlineResponse200209MacBlocklist.md) |  | [optional] 
**ClientTracking** | Pointer to [**InlineResponse200209ClientTracking**](InlineResponse200209ClientTracking.md) |  | [optional] 
**PortChannelFallback** | Pointer to **bool** | Enable port channel fallback on Catalyst MS switches. | [optional] 
**UplinkSelection** | Pointer to [**InlineResponse200209UplinkSelection**](InlineResponse200209UplinkSelection.md) |  | [optional] 

## Methods

### NewInlineResponse200209

`func NewInlineResponse200209() *InlineResponse200209`

NewInlineResponse200209 instantiates a new InlineResponse200209 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200209WithDefaults

`func NewInlineResponse200209WithDefaults() *InlineResponse200209`

NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse200209) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200209) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200209) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200209) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse200209) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse200209) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse200209) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse200209) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse200209) GetPowerExceptions() []InlineResponse200209PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse200209) GetPowerExceptionsOk() (*[]InlineResponse200209PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse200209) SetPowerExceptions(v []InlineResponse200209PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse200209) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.

### GetUplinkClientSampling

`func (o *InlineResponse200209) GetUplinkClientSampling() InlineResponse200209UplinkClientSampling`

GetUplinkClientSampling returns the UplinkClientSampling field if non-nil, zero value otherwise.

### GetUplinkClientSamplingOk

`func (o *InlineResponse200209) GetUplinkClientSamplingOk() (*InlineResponse200209UplinkClientSampling, bool)`

GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkClientSampling

`func (o *InlineResponse200209) SetUplinkClientSampling(v InlineResponse200209UplinkClientSampling)`

SetUplinkClientSampling sets UplinkClientSampling field to given value.

### HasUplinkClientSampling

`func (o *InlineResponse200209) HasUplinkClientSampling() bool`

HasUplinkClientSampling returns a boolean if a field has been set.

### GetMacBlocklist

`func (o *InlineResponse200209) GetMacBlocklist() InlineResponse200209MacBlocklist`

GetMacBlocklist returns the MacBlocklist field if non-nil, zero value otherwise.

### GetMacBlocklistOk

`func (o *InlineResponse200209) GetMacBlocklistOk() (*InlineResponse200209MacBlocklist, bool)`

GetMacBlocklistOk returns a tuple with the MacBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocklist

`func (o *InlineResponse200209) SetMacBlocklist(v InlineResponse200209MacBlocklist)`

SetMacBlocklist sets MacBlocklist field to given value.

### HasMacBlocklist

`func (o *InlineResponse200209) HasMacBlocklist() bool`

HasMacBlocklist returns a boolean if a field has been set.

### GetClientTracking

`func (o *InlineResponse200209) GetClientTracking() InlineResponse200209ClientTracking`

GetClientTracking returns the ClientTracking field if non-nil, zero value otherwise.

### GetClientTrackingOk

`func (o *InlineResponse200209) GetClientTrackingOk() (*InlineResponse200209ClientTracking, bool)`

GetClientTrackingOk returns a tuple with the ClientTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTracking

`func (o *InlineResponse200209) SetClientTracking(v InlineResponse200209ClientTracking)`

SetClientTracking sets ClientTracking field to given value.

### HasClientTracking

`func (o *InlineResponse200209) HasClientTracking() bool`

HasClientTracking returns a boolean if a field has been set.

### GetPortChannelFallback

`func (o *InlineResponse200209) GetPortChannelFallback() bool`

GetPortChannelFallback returns the PortChannelFallback field if non-nil, zero value otherwise.

### GetPortChannelFallbackOk

`func (o *InlineResponse200209) GetPortChannelFallbackOk() (*bool, bool)`

GetPortChannelFallbackOk returns a tuple with the PortChannelFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelFallback

`func (o *InlineResponse200209) SetPortChannelFallback(v bool)`

SetPortChannelFallback sets PortChannelFallback field to given value.

### HasPortChannelFallback

`func (o *InlineResponse200209) HasPortChannelFallback() bool`

HasPortChannelFallback returns a boolean if a field has been set.

### GetUplinkSelection

`func (o *InlineResponse200209) GetUplinkSelection() InlineResponse200209UplinkSelection`

GetUplinkSelection returns the UplinkSelection field if non-nil, zero value otherwise.

### GetUplinkSelectionOk

`func (o *InlineResponse200209) GetUplinkSelectionOk() (*InlineResponse200209UplinkSelection, bool)`

GetUplinkSelectionOk returns a tuple with the UplinkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSelection

`func (o *InlineResponse200209) SetUplinkSelection(v InlineResponse200209UplinkSelection)`

SetUplinkSelection sets UplinkSelection field to given value.

### HasUplinkSelection

`func (o *InlineResponse200209) HasUplinkSelection() bool`

HasUplinkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


