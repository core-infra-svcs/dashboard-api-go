# InlineResponse200280Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the schedule | [optional] 
**StartTs** | Pointer to **string** | Start date of the recurring schedule entry | [optional] 
**EndTs** | Pointer to **string** | End date of the recurring schedule entry | [optional] 
**Frequency** | Pointer to **string** | Frequency of the recurring schedule entry ex. hour|week|month|day|minute | [optional] 
**Weekdays** | Pointer to **[]string** | The days of the week for the recurring schedule in string form. Multiple days can be combined. | [optional] 
**Recurrence** | Pointer to **int32** | The number of frequency units between each occurrence. For example, 1 means &#39;every [frequency]&#39;, 2 means &#39;every other [frequency]&#39;, etc. Used in conjunction with the &#39;frequency&#39; field. | [optional] 
**NextCaptureTs** | Pointer to **string** | The datetime at which next capture will occur | [optional] 

## Methods

### NewInlineResponse200280Schedule

`func NewInlineResponse200280Schedule() *InlineResponse200280Schedule`

NewInlineResponse200280Schedule instantiates a new InlineResponse200280Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200280ScheduleWithDefaults

`func NewInlineResponse200280ScheduleWithDefaults() *InlineResponse200280Schedule`

NewInlineResponse200280ScheduleWithDefaults instantiates a new InlineResponse200280Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200280Schedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200280Schedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200280Schedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200280Schedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTs

`func (o *InlineResponse200280Schedule) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200280Schedule) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200280Schedule) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200280Schedule) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200280Schedule) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200280Schedule) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200280Schedule) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200280Schedule) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetFrequency

`func (o *InlineResponse200280Schedule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InlineResponse200280Schedule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InlineResponse200280Schedule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InlineResponse200280Schedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetWeekdays

`func (o *InlineResponse200280Schedule) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *InlineResponse200280Schedule) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *InlineResponse200280Schedule) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *InlineResponse200280Schedule) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetRecurrence

`func (o *InlineResponse200280Schedule) GetRecurrence() int32`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *InlineResponse200280Schedule) GetRecurrenceOk() (*int32, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *InlineResponse200280Schedule) SetRecurrence(v int32)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *InlineResponse200280Schedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetNextCaptureTs

`func (o *InlineResponse200280Schedule) GetNextCaptureTs() string`

GetNextCaptureTs returns the NextCaptureTs field if non-nil, zero value otherwise.

### GetNextCaptureTsOk

`func (o *InlineResponse200280Schedule) GetNextCaptureTsOk() (*string, bool)`

GetNextCaptureTsOk returns a tuple with the NextCaptureTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCaptureTs

`func (o *InlineResponse200280Schedule) SetNextCaptureTs(v string)`

SetNextCaptureTs sets NextCaptureTs field to given value.

### HasNextCaptureTs

`func (o *InlineResponse200280Schedule) HasNextCaptureTs() bool`

HasNextCaptureTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


