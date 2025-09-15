# InlineResponse200101Stages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**InlineResponse200101Group**](InlineResponse200101Group.md) |  | [optional] 
**Milestones** | Pointer to [**InlineResponse200101Milestones**](InlineResponse200101Milestones.md) |  | [optional] 
**Status** | Pointer to **string** | Current upgrade status of the group | [optional] 

## Methods

### NewInlineResponse200101Stages

`func NewInlineResponse200101Stages() *InlineResponse200101Stages`

NewInlineResponse200101Stages instantiates a new InlineResponse200101Stages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101StagesWithDefaults

`func NewInlineResponse200101StagesWithDefaults() *InlineResponse200101Stages`

NewInlineResponse200101StagesWithDefaults instantiates a new InlineResponse200101Stages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InlineResponse200101Stages) GetGroup() InlineResponse200101Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200101Stages) GetGroupOk() (*InlineResponse200101Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200101Stages) SetGroup(v InlineResponse200101Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200101Stages) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMilestones

`func (o *InlineResponse200101Stages) GetMilestones() InlineResponse200101Milestones`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *InlineResponse200101Stages) GetMilestonesOk() (*InlineResponse200101Milestones, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *InlineResponse200101Stages) SetMilestones(v InlineResponse200101Milestones)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *InlineResponse200101Stages) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200101Stages) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200101Stages) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200101Stages) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200101Stages) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


