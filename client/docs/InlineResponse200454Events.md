# InlineResponse200454Events

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp of the event | [optional] 
**PowerMode** | Pointer to **string** | The power mode of the device | [optional] 

## Methods

### NewInlineResponse200454Events

`func NewInlineResponse200454Events() *InlineResponse200454Events`

NewInlineResponse200454Events instantiates a new InlineResponse200454Events object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200454EventsWithDefaults

`func NewInlineResponse200454EventsWithDefaults() *InlineResponse200454Events`

NewInlineResponse200454EventsWithDefaults instantiates a new InlineResponse200454Events object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200454Events) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200454Events) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200454Events) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200454Events) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetPowerMode

`func (o *InlineResponse200454Events) GetPowerMode() string`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *InlineResponse200454Events) GetPowerModeOk() (*string, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *InlineResponse200454Events) SetPowerMode(v string)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *InlineResponse200454Events) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


