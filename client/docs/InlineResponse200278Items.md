# InlineResponse200278Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** | Id of scheduled packet capture | [optional] 
**Devices** | Pointer to [**[]InlineResponse200278Devices**](InlineResponse200278Devices.md) | Devices associated to the schedule | [optional] 
**Name** | Pointer to **string** | Name of scheduled packet capture | [optional] 
**Admin** | Pointer to [**InlineResponse200278Admin**](InlineResponse200278Admin.md) |  | [optional] 
**Notes** | Pointer to **string** | Reason of scheduled packet capture | [optional] 
**Duration** | Pointer to **int32** | Duration of scheduled packet capture | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the packet capture | [optional] 
**CreatedAt** | Pointer to **string** | Time of creation of scheduled packet capture | [optional] 
**UpdatedAt** | Pointer to **string** | Time of updation of scheduled packet capture | [optional] 
**CaptureCount** | Pointer to **int32** | The number of pcaps captured/performed | [optional] 
**LastCaptureId** | Pointer to **string** | Pcap log id of the latest pcap from this schedule | [optional] 
**Enabled** | Pointer to **bool** | Whether the packet capture schedule is enabled | [optional] 
**Priority** | Pointer to **int32** | Priority of the packet capture | [optional] 
**Schedule** | Pointer to [**InlineResponse200278Schedule**](InlineResponse200278Schedule.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Any warnings pertaining to the schedule and it&#39;s nodes | [optional] 

## Methods

### NewInlineResponse200278Items

`func NewInlineResponse200278Items() *InlineResponse200278Items`

NewInlineResponse200278Items instantiates a new InlineResponse200278Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200278ItemsWithDefaults

`func NewInlineResponse200278ItemsWithDefaults() *InlineResponse200278Items`

NewInlineResponse200278ItemsWithDefaults instantiates a new InlineResponse200278Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleId

`func (o *InlineResponse200278Items) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *InlineResponse200278Items) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *InlineResponse200278Items) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *InlineResponse200278Items) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200278Items) GetDevices() []InlineResponse200278Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200278Items) GetDevicesOk() (*[]InlineResponse200278Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200278Items) SetDevices(v []InlineResponse200278Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200278Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200278Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200278Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200278Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200278Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdmin

`func (o *InlineResponse200278Items) GetAdmin() InlineResponse200278Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *InlineResponse200278Items) GetAdminOk() (*InlineResponse200278Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *InlineResponse200278Items) SetAdmin(v InlineResponse200278Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *InlineResponse200278Items) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200278Items) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200278Items) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200278Items) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200278Items) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200278Items) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200278Items) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200278Items) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200278Items) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineResponse200278Items) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineResponse200278Items) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineResponse200278Items) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineResponse200278Items) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200278Items) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200278Items) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200278Items) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200278Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200278Items) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200278Items) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200278Items) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200278Items) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCaptureCount

`func (o *InlineResponse200278Items) GetCaptureCount() int32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *InlineResponse200278Items) GetCaptureCountOk() (*int32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *InlineResponse200278Items) SetCaptureCount(v int32)`

SetCaptureCount sets CaptureCount field to given value.

### HasCaptureCount

`func (o *InlineResponse200278Items) HasCaptureCount() bool`

HasCaptureCount returns a boolean if a field has been set.

### GetLastCaptureId

`func (o *InlineResponse200278Items) GetLastCaptureId() string`

GetLastCaptureId returns the LastCaptureId field if non-nil, zero value otherwise.

### GetLastCaptureIdOk

`func (o *InlineResponse200278Items) GetLastCaptureIdOk() (*string, bool)`

GetLastCaptureIdOk returns a tuple with the LastCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaptureId

`func (o *InlineResponse200278Items) SetLastCaptureId(v string)`

SetLastCaptureId sets LastCaptureId field to given value.

### HasLastCaptureId

`func (o *InlineResponse200278Items) HasLastCaptureId() bool`

HasLastCaptureId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200278Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200278Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200278Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200278Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200278Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200278Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200278Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200278Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineResponse200278Items) GetSchedule() InlineResponse200278Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineResponse200278Items) GetScheduleOk() (*InlineResponse200278Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineResponse200278Items) SetSchedule(v InlineResponse200278Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineResponse200278Items) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse200278Items) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse200278Items) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse200278Items) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse200278Items) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


