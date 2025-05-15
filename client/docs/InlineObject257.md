# InlineObject257

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices**](OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices.md) | device details | 
**Name** | Pointer to **string** | Name of the packet capture file | [optional] 
**Notes** | Pointer to **string** | Reason for capture | [optional] 
**Duration** | Pointer to **int32** | Duration of the capture in seconds | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the capture | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the schedule | [optional] 
**Schedule** | Pointer to [**OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule**](OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule.md) |  | [optional] 

## Methods

### NewInlineObject257

`func NewInlineObject257(devices []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices, ) *InlineObject257`

NewInlineObject257 instantiates a new InlineObject257 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject257WithDefaults

`func NewInlineObject257WithDefaults() *InlineObject257`

NewInlineObject257WithDefaults instantiates a new InlineObject257 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *InlineObject257) GetDevices() []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject257) GetDevicesOk() (*[]OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject257) SetDevices(v []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices)`

SetDevices sets Devices field to given value.


### GetName

`func (o *InlineObject257) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject257) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject257) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject257) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject257) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject257) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject257) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject257) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject257) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject257) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject257) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject257) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject257) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject257) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject257) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject257) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject257) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject257) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject257) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject257) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineObject257) GetSchedule() OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineObject257) GetScheduleOk() (*OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineObject257) SetSchedule(v OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineObject257) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


