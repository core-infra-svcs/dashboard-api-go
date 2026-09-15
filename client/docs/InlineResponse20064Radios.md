# InlineResponse20064Radios

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

### NewInlineResponse20064Radios

`func NewInlineResponse20064Radios() *InlineResponse20064Radios`

NewInlineResponse20064Radios instantiates a new InlineResponse20064Radios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20064RadiosWithDefaults

`func NewInlineResponse20064RadiosWithDefaults() *InlineResponse20064Radios`

NewInlineResponse20064RadiosWithDefaults instantiates a new InlineResponse20064Radios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20064Radios) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20064Radios) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20064Radios) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20064Radios) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIndex

`func (o *InlineResponse20064Radios) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InlineResponse20064Radios) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InlineResponse20064Radios) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *InlineResponse20064Radios) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetBand

`func (o *InlineResponse20064Radios) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *InlineResponse20064Radios) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *InlineResponse20064Radios) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *InlineResponse20064Radios) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse20064Radios) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse20064Radios) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse20064Radios) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse20064Radios) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *InlineResponse20064Radios) GetChannelWidth() int32`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *InlineResponse20064Radios) GetChannelWidthOk() (*int32, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *InlineResponse20064Radios) SetChannelWidth(v int32)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *InlineResponse20064Radios) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetTargetPower

`func (o *InlineResponse20064Radios) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *InlineResponse20064Radios) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *InlineResponse20064Radios) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *InlineResponse20064Radios) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


