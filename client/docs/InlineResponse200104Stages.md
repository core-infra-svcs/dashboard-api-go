# InlineResponse200104Stages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**InlineResponse200104Group**](InlineResponse200104Group.md) |  | [optional] 
**Milestones** | Pointer to [**InlineResponse200104Milestones**](InlineResponse200104Milestones.md) |  | [optional] 
**Status** | Pointer to **string** | Current upgrade status of the group | [optional] 

## Methods

### NewInlineResponse200104Stages

`func NewInlineResponse200104Stages() *InlineResponse200104Stages`

NewInlineResponse200104Stages instantiates a new InlineResponse200104Stages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104StagesWithDefaults

`func NewInlineResponse200104StagesWithDefaults() *InlineResponse200104Stages`

NewInlineResponse200104StagesWithDefaults instantiates a new InlineResponse200104Stages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InlineResponse200104Stages) GetGroup() InlineResponse200104Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200104Stages) GetGroupOk() (*InlineResponse200104Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200104Stages) SetGroup(v InlineResponse200104Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200104Stages) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMilestones

`func (o *InlineResponse200104Stages) GetMilestones() InlineResponse200104Milestones`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *InlineResponse200104Stages) GetMilestonesOk() (*InlineResponse200104Milestones, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *InlineResponse200104Stages) SetMilestones(v InlineResponse200104Milestones)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *InlineResponse200104Stages) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200104Stages) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200104Stages) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200104Stages) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200104Stages) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


