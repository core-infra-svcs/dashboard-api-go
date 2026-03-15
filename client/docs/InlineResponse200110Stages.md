# InlineResponse200110Stages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**InlineResponse200110Group**](InlineResponse200110Group.md) |  | [optional] 
**Milestones** | Pointer to [**InlineResponse200110Milestones**](InlineResponse200110Milestones.md) |  | [optional] 
**Status** | Pointer to **string** | Current upgrade status of the group | [optional] 

## Methods

### NewInlineResponse200110Stages

`func NewInlineResponse200110Stages() *InlineResponse200110Stages`

NewInlineResponse200110Stages instantiates a new InlineResponse200110Stages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110StagesWithDefaults

`func NewInlineResponse200110StagesWithDefaults() *InlineResponse200110Stages`

NewInlineResponse200110StagesWithDefaults instantiates a new InlineResponse200110Stages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InlineResponse200110Stages) GetGroup() InlineResponse200110Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200110Stages) GetGroupOk() (*InlineResponse200110Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200110Stages) SetGroup(v InlineResponse200110Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200110Stages) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMilestones

`func (o *InlineResponse200110Stages) GetMilestones() InlineResponse200110Milestones`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *InlineResponse200110Stages) GetMilestonesOk() (*InlineResponse200110Milestones, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *InlineResponse200110Stages) SetMilestones(v InlineResponse200110Milestones)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *InlineResponse200110Stages) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200110Stages) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200110Stages) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200110Stages) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200110Stages) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


