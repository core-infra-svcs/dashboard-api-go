# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The duration in seconds. Must be between 5 and 120. Default is 20 seconds | [optional] 
**Period** | Pointer to **int32** | The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds | [optional] 
**Duty** | Pointer to **int32** | The duty cycle as the percent active. Must be between 10 and 90. Default is 50. | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2() *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *InlineObject2) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject2) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject2) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject2) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPeriod

`func (o *InlineObject2) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InlineObject2) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InlineObject2) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *InlineObject2) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetDuty

`func (o *InlineObject2) GetDuty() int32`

GetDuty returns the Duty field if non-nil, zero value otherwise.

### GetDutyOk

`func (o *InlineObject2) GetDutyOk() (*int32, bool)`

GetDutyOk returns a tuple with the Duty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuty

`func (o *InlineObject2) SetDuty(v int32)`

SetDuty sets Duty field to given value.

### HasDuty

`func (o *InlineObject2) HasDuty() bool`

HasDuty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


