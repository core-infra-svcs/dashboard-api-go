# OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the schedule | [optional] 
**StartTs** | Pointer to **string** | Start date and time of the recurring schedule entry | [optional] 
**EndTs** | Pointer to **string** | End date and time of the recurring schedule entry | [optional] 
**Frequency** | Pointer to **string** | Frequency of the recurring schedule entry (hour, week, month, day, minute) | [optional] 
**Weekdays** | Pointer to **[]string** | Weekdays for the schedule: Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday | [optional] 
**Recurrence** | Pointer to **int32** | Cardinality of the schedule frequency, ex. 1 &#x3D; every day, 2 &#x3D; every other day (when frequency &#x3D; day) | [optional] 

## Methods

### NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule

`func NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule() *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule`

NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesScheduleWithDefaults

`func NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesScheduleWithDefaults() *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule`

NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesScheduleWithDefaults instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetFrequency

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetWeekdays

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetRecurrence

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetRecurrence() int32`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) GetRecurrenceOk() (*int32, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) SetRecurrence(v int32)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


