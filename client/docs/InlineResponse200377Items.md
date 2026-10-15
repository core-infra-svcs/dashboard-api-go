# InlineResponse200377Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** | Id of scheduled packet capture | [optional] 
**Devices** | Pointer to [**[]InlineResponse200377Devices**](InlineResponse200377Devices.md) | Devices associated to the schedule | [optional] 
**Name** | Pointer to **string** | Name of scheduled packet capture | [optional] 
**Admin** | Pointer to [**InlineResponse200377Admin**](InlineResponse200377Admin.md) |  | [optional] 
**Notes** | Pointer to **string** | Reason of scheduled packet capture | [optional] 
**Duration** | Pointer to **int32** | Duration of scheduled packet capture | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the packet capture | [optional] 
**CreatedAt** | Pointer to **string** | Time of creation of scheduled packet capture | [optional] 
**UpdatedAt** | Pointer to **string** | Time of updation of scheduled packet capture | [optional] 
**CaptureCount** | Pointer to **int32** | The number of pcaps captured/performed | [optional] 
**LastCaptureId** | Pointer to **string** | Pcap log id of the latest pcap from this schedule | [optional] 
**Enabled** | Pointer to **bool** | Whether the packet capture schedule is enabled | [optional] 
**Priority** | Pointer to **int32** | Priority of the packet capture | [optional] 
**Schedule** | Pointer to [**InlineResponse200377Schedule**](InlineResponse200377Schedule.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Any warnings pertaining to the schedule and it&#39;s nodes | [optional] 

## Methods

### NewInlineResponse200377Items

`func NewInlineResponse200377Items() *InlineResponse200377Items`

NewInlineResponse200377Items instantiates a new InlineResponse200377Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200377ItemsWithDefaults

`func NewInlineResponse200377ItemsWithDefaults() *InlineResponse200377Items`

NewInlineResponse200377ItemsWithDefaults instantiates a new InlineResponse200377Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleId

`func (o *InlineResponse200377Items) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *InlineResponse200377Items) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *InlineResponse200377Items) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *InlineResponse200377Items) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200377Items) GetDevices() []InlineResponse200377Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200377Items) GetDevicesOk() (*[]InlineResponse200377Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200377Items) SetDevices(v []InlineResponse200377Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200377Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200377Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200377Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200377Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200377Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdmin

`func (o *InlineResponse200377Items) GetAdmin() InlineResponse200377Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *InlineResponse200377Items) GetAdminOk() (*InlineResponse200377Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *InlineResponse200377Items) SetAdmin(v InlineResponse200377Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *InlineResponse200377Items) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200377Items) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200377Items) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200377Items) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200377Items) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200377Items) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200377Items) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200377Items) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200377Items) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineResponse200377Items) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineResponse200377Items) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineResponse200377Items) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineResponse200377Items) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200377Items) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200377Items) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200377Items) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200377Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200377Items) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200377Items) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200377Items) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200377Items) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCaptureCount

`func (o *InlineResponse200377Items) GetCaptureCount() int32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *InlineResponse200377Items) GetCaptureCountOk() (*int32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *InlineResponse200377Items) SetCaptureCount(v int32)`

SetCaptureCount sets CaptureCount field to given value.

### HasCaptureCount

`func (o *InlineResponse200377Items) HasCaptureCount() bool`

HasCaptureCount returns a boolean if a field has been set.

### GetLastCaptureId

`func (o *InlineResponse200377Items) GetLastCaptureId() string`

GetLastCaptureId returns the LastCaptureId field if non-nil, zero value otherwise.

### GetLastCaptureIdOk

`func (o *InlineResponse200377Items) GetLastCaptureIdOk() (*string, bool)`

GetLastCaptureIdOk returns a tuple with the LastCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaptureId

`func (o *InlineResponse200377Items) SetLastCaptureId(v string)`

SetLastCaptureId sets LastCaptureId field to given value.

### HasLastCaptureId

`func (o *InlineResponse200377Items) HasLastCaptureId() bool`

HasLastCaptureId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200377Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200377Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200377Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200377Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200377Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200377Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200377Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200377Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineResponse200377Items) GetSchedule() InlineResponse200377Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineResponse200377Items) GetScheduleOk() (*InlineResponse200377Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineResponse200377Items) SetSchedule(v InlineResponse200377Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineResponse200377Items) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse200377Items) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse200377Items) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse200377Items) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse200377Items) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


