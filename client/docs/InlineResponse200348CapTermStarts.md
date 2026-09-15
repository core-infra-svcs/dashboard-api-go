# InlineResponse200348CapTermStarts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HourOfDay** | Pointer to **int32** | When the data usage term resets daily, the hour the accounted bandwidth usage should reset. This can be between 0 and 23 (where 0 is 0:00 GMT and 23 is 23:00 GMT). | [optional] 
**DayOfWeek** | Pointer to **string** | When the data usage term resets weekly, the day of week that the accounted bandwidth usage should reset. Uses the lowercase three-letter weekday abbreviation: &#39;mon&#39;, &#39;tue&#39;, &#39;wed&#39;, &#39;thu&#39;, &#39;fri&#39;, &#39;sat&#39;, or &#39;sun&#39;. Requests are normalized case-insensitively to this form. | [optional] 
**DayOfMonth** | Pointer to **int32** | When the data usage term resets monthly, the day of month that the accounted bandwidth usage should reset. This can be between 1 and 31 (months that have a number of days less than the value for dayOfMonth will begin on the last day of the month). | [optional] 

## Methods

### NewInlineResponse200348CapTermStarts

`func NewInlineResponse200348CapTermStarts() *InlineResponse200348CapTermStarts`

NewInlineResponse200348CapTermStarts instantiates a new InlineResponse200348CapTermStarts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200348CapTermStartsWithDefaults

`func NewInlineResponse200348CapTermStartsWithDefaults() *InlineResponse200348CapTermStarts`

NewInlineResponse200348CapTermStartsWithDefaults instantiates a new InlineResponse200348CapTermStarts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourOfDay

`func (o *InlineResponse200348CapTermStarts) GetHourOfDay() int32`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *InlineResponse200348CapTermStarts) GetHourOfDayOk() (*int32, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *InlineResponse200348CapTermStarts) SetHourOfDay(v int32)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *InlineResponse200348CapTermStarts) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *InlineResponse200348CapTermStarts) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *InlineResponse200348CapTermStarts) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *InlineResponse200348CapTermStarts) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *InlineResponse200348CapTermStarts) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *InlineResponse200348CapTermStarts) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *InlineResponse200348CapTermStarts) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *InlineResponse200348CapTermStarts) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *InlineResponse200348CapTermStarts) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


