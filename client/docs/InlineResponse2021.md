# InlineResponse2021

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The duration in seconds. Will be between 5 and 120. Default is 20 seconds | [optional] 
**Period** | Pointer to **int32** | The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds | [optional] 
**Duty** | Pointer to **int32** | The duty cycle as the percent active. Will be between 10 and 90. Default is 50 | [optional] 

## Methods

### NewInlineResponse2021

`func NewInlineResponse2021() *InlineResponse2021`

NewInlineResponse2021 instantiates a new InlineResponse2021 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2021WithDefaults

`func NewInlineResponse2021WithDefaults() *InlineResponse2021`

NewInlineResponse2021WithDefaults instantiates a new InlineResponse2021 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *InlineResponse2021) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse2021) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse2021) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse2021) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPeriod

`func (o *InlineResponse2021) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InlineResponse2021) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InlineResponse2021) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *InlineResponse2021) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetDuty

`func (o *InlineResponse2021) GetDuty() int32`

GetDuty returns the Duty field if non-nil, zero value otherwise.

### GetDutyOk

`func (o *InlineResponse2021) GetDutyOk() (*int32, bool)`

GetDutyOk returns a tuple with the Duty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuty

`func (o *InlineResponse2021) SetDuty(v int32)`

SetDuty sets Duty field to given value.

### HasDuty

`func (o *InlineResponse2021) HasDuty() bool`

HasDuty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


