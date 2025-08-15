# InlineResponse20116Defaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransmitPowerLevel** | Pointer to **int32** | Default transmit power level for the network | [optional] 
**Channel** | Pointer to **string** | Channel the zigbee node will communicate on either, auto or between 11-25 | [optional] 

## Methods

### NewInlineResponse20116Defaults

`func NewInlineResponse20116Defaults() *InlineResponse20116Defaults`

NewInlineResponse20116Defaults instantiates a new InlineResponse20116Defaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20116DefaultsWithDefaults

`func NewInlineResponse20116DefaultsWithDefaults() *InlineResponse20116Defaults`

NewInlineResponse20116DefaultsWithDefaults instantiates a new InlineResponse20116Defaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransmitPowerLevel

`func (o *InlineResponse20116Defaults) GetTransmitPowerLevel() int32`

GetTransmitPowerLevel returns the TransmitPowerLevel field if non-nil, zero value otherwise.

### GetTransmitPowerLevelOk

`func (o *InlineResponse20116Defaults) GetTransmitPowerLevelOk() (*int32, bool)`

GetTransmitPowerLevelOk returns a tuple with the TransmitPowerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPowerLevel

`func (o *InlineResponse20116Defaults) SetTransmitPowerLevel(v int32)`

SetTransmitPowerLevel sets TransmitPowerLevel field to given value.

### HasTransmitPowerLevel

`func (o *InlineResponse20116Defaults) HasTransmitPowerLevel() bool`

HasTransmitPowerLevel returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse20116Defaults) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse20116Defaults) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse20116Defaults) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse20116Defaults) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


