# InlineResponse200255StatesExpiringWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdInDays** | Pointer to **int32** | The number of days from now denoting the warning threshold for an expiring license | [optional] 
**ExpiringCount** | Pointer to **int32** | The number of licenses that will expire in this window | [optional] 

## Methods

### NewInlineResponse200255StatesExpiringWarning

`func NewInlineResponse200255StatesExpiringWarning() *InlineResponse200255StatesExpiringWarning`

NewInlineResponse200255StatesExpiringWarning instantiates a new InlineResponse200255StatesExpiringWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200255StatesExpiringWarningWithDefaults

`func NewInlineResponse200255StatesExpiringWarningWithDefaults() *InlineResponse200255StatesExpiringWarning`

NewInlineResponse200255StatesExpiringWarningWithDefaults instantiates a new InlineResponse200255StatesExpiringWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdInDays

`func (o *InlineResponse200255StatesExpiringWarning) GetThresholdInDays() int32`

GetThresholdInDays returns the ThresholdInDays field if non-nil, zero value otherwise.

### GetThresholdInDaysOk

`func (o *InlineResponse200255StatesExpiringWarning) GetThresholdInDaysOk() (*int32, bool)`

GetThresholdInDaysOk returns a tuple with the ThresholdInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInDays

`func (o *InlineResponse200255StatesExpiringWarning) SetThresholdInDays(v int32)`

SetThresholdInDays sets ThresholdInDays field to given value.

### HasThresholdInDays

`func (o *InlineResponse200255StatesExpiringWarning) HasThresholdInDays() bool`

HasThresholdInDays returns a boolean if a field has been set.

### GetExpiringCount

`func (o *InlineResponse200255StatesExpiringWarning) GetExpiringCount() int32`

GetExpiringCount returns the ExpiringCount field if non-nil, zero value otherwise.

### GetExpiringCountOk

`func (o *InlineResponse200255StatesExpiringWarning) GetExpiringCountOk() (*int32, bool)`

GetExpiringCountOk returns a tuple with the ExpiringCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringCount

`func (o *InlineResponse200255StatesExpiringWarning) SetExpiringCount(v int32)`

SetExpiringCount sets ExpiringCount field to given value.

### HasExpiringCount

`func (o *InlineResponse200255StatesExpiringWarning) HasExpiringCount() bool`

HasExpiringCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


