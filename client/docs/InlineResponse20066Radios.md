# InlineResponse20066Radios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Radio enabled | [optional] 
**Index** | Pointer to **string** | Radio identifier for the device | [optional] 
**Band** | Pointer to **string** | Radio band: one of &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39; | [optional] 
**Channel** | Pointer to **int32** | Radio channel: Valid channels depend on the radio band. | [optional] 
**ChannelWidth** | Pointer to **int32** | Radio channel width: Valid widths depend on the radio band. | [optional] 
**TargetPower** | Pointer to **int32** | Radio target power: Valid target power depends on the radio band. | [optional] 

## Methods

### NewInlineResponse20066Radios

`func NewInlineResponse20066Radios() *InlineResponse20066Radios`

NewInlineResponse20066Radios instantiates a new InlineResponse20066Radios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20066RadiosWithDefaults

`func NewInlineResponse20066RadiosWithDefaults() *InlineResponse20066Radios`

NewInlineResponse20066RadiosWithDefaults instantiates a new InlineResponse20066Radios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20066Radios) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20066Radios) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20066Radios) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20066Radios) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIndex

`func (o *InlineResponse20066Radios) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InlineResponse20066Radios) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InlineResponse20066Radios) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *InlineResponse20066Radios) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetBand

`func (o *InlineResponse20066Radios) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *InlineResponse20066Radios) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *InlineResponse20066Radios) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *InlineResponse20066Radios) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse20066Radios) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse20066Radios) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse20066Radios) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse20066Radios) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *InlineResponse20066Radios) GetChannelWidth() int32`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *InlineResponse20066Radios) GetChannelWidthOk() (*int32, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *InlineResponse20066Radios) SetChannelWidth(v int32)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *InlineResponse20066Radios) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetTargetPower

`func (o *InlineResponse20066Radios) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *InlineResponse20066Radios) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *InlineResponse20066Radios) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *InlineResponse20066Radios) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


