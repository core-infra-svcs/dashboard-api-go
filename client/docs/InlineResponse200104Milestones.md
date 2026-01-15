# InlineResponse200104Milestones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledFor** | Pointer to **time.Time** | Scheduled start time for the group | [optional] 
**StartedAt** | Pointer to **time.Time** | Start time for the group | [optional] 
**CompletedAt** | Pointer to **time.Time** | Finish time for the group | [optional] 
**CanceledAt** | Pointer to **time.Time** | Time that the group was canceled | [optional] 

## Methods

### NewInlineResponse200104Milestones

`func NewInlineResponse200104Milestones() *InlineResponse200104Milestones`

NewInlineResponse200104Milestones instantiates a new InlineResponse200104Milestones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104MilestonesWithDefaults

`func NewInlineResponse200104MilestonesWithDefaults() *InlineResponse200104Milestones`

NewInlineResponse200104MilestonesWithDefaults instantiates a new InlineResponse200104Milestones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledFor

`func (o *InlineResponse200104Milestones) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *InlineResponse200104Milestones) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *InlineResponse200104Milestones) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *InlineResponse200104Milestones) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetStartedAt

`func (o *InlineResponse200104Milestones) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InlineResponse200104Milestones) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InlineResponse200104Milestones) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *InlineResponse200104Milestones) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *InlineResponse200104Milestones) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InlineResponse200104Milestones) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InlineResponse200104Milestones) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InlineResponse200104Milestones) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCanceledAt

`func (o *InlineResponse200104Milestones) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *InlineResponse200104Milestones) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *InlineResponse200104Milestones) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *InlineResponse200104Milestones) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


