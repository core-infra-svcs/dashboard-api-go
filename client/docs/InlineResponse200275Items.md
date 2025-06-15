# InlineResponse200275Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** | Id of scheduled packet capture | [optional] 
**Devices** | Pointer to [**[]InlineResponse200275Devices**](InlineResponse200275Devices.md) | Devices associated to the schedule | [optional] 
**Name** | Pointer to **string** | Name of scheduled packet capture | [optional] 
**Admin** | Pointer to [**InlineResponse200275Admin**](InlineResponse200275Admin.md) |  | [optional] 
**Notes** | Pointer to **string** | Reason of scheduled packet capture | [optional] 
**Duration** | Pointer to **int32** | Duration of scheduled packet capture | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the packet capture | [optional] 
**CreatedAt** | Pointer to **string** | Time of creation of scheduled packet capture | [optional] 
**UpdatedAt** | Pointer to **string** | Time of updation of scheduled packet capture | [optional] 
**CaptureCount** | Pointer to **int32** | The number of pcaps captured/performed | [optional] 
**LastCaptureId** | Pointer to **string** | Pcap log id of the latest pcap from this schedule | [optional] 
**Enabled** | Pointer to **bool** | Whether the packet capture schedule is enabled | [optional] 
**Priority** | Pointer to **int32** | Priority of the packet capture | [optional] 
**Schedule** | Pointer to [**InlineResponse200275Schedule**](InlineResponse200275Schedule.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Any warnings pertaining to the schedule and it&#39;s nodes | [optional] 

## Methods

### NewInlineResponse200275Items

`func NewInlineResponse200275Items() *InlineResponse200275Items`

NewInlineResponse200275Items instantiates a new InlineResponse200275Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275ItemsWithDefaults

`func NewInlineResponse200275ItemsWithDefaults() *InlineResponse200275Items`

NewInlineResponse200275ItemsWithDefaults instantiates a new InlineResponse200275Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleId

`func (o *InlineResponse200275Items) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *InlineResponse200275Items) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *InlineResponse200275Items) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *InlineResponse200275Items) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200275Items) GetDevices() []InlineResponse200275Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200275Items) GetDevicesOk() (*[]InlineResponse200275Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200275Items) SetDevices(v []InlineResponse200275Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200275Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200275Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200275Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200275Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200275Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdmin

`func (o *InlineResponse200275Items) GetAdmin() InlineResponse200275Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *InlineResponse200275Items) GetAdminOk() (*InlineResponse200275Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *InlineResponse200275Items) SetAdmin(v InlineResponse200275Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *InlineResponse200275Items) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200275Items) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200275Items) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200275Items) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200275Items) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200275Items) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200275Items) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200275Items) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200275Items) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineResponse200275Items) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineResponse200275Items) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineResponse200275Items) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineResponse200275Items) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200275Items) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200275Items) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200275Items) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200275Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200275Items) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200275Items) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200275Items) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200275Items) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCaptureCount

`func (o *InlineResponse200275Items) GetCaptureCount() int32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *InlineResponse200275Items) GetCaptureCountOk() (*int32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *InlineResponse200275Items) SetCaptureCount(v int32)`

SetCaptureCount sets CaptureCount field to given value.

### HasCaptureCount

`func (o *InlineResponse200275Items) HasCaptureCount() bool`

HasCaptureCount returns a boolean if a field has been set.

### GetLastCaptureId

`func (o *InlineResponse200275Items) GetLastCaptureId() string`

GetLastCaptureId returns the LastCaptureId field if non-nil, zero value otherwise.

### GetLastCaptureIdOk

`func (o *InlineResponse200275Items) GetLastCaptureIdOk() (*string, bool)`

GetLastCaptureIdOk returns a tuple with the LastCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCaptureId

`func (o *InlineResponse200275Items) SetLastCaptureId(v string)`

SetLastCaptureId sets LastCaptureId field to given value.

### HasLastCaptureId

`func (o *InlineResponse200275Items) HasLastCaptureId() bool`

HasLastCaptureId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200275Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200275Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200275Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200275Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200275Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200275Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200275Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200275Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineResponse200275Items) GetSchedule() InlineResponse200275Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineResponse200275Items) GetScheduleOk() (*InlineResponse200275Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineResponse200275Items) SetSchedule(v InlineResponse200275Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineResponse200275Items) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse200275Items) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse200275Items) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse200275Items) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse200275Items) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


