# InlineResponse200305CountsRuleGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of child rule groups for this rule group. | [optional] 
**BySecurityLevel** | Pointer to [**InlineResponse200305CountsRuleGroupsBySecurityLevel**](InlineResponse200305CountsRuleGroupsBySecurityLevel.md) |  | [optional] 

## Methods

### NewInlineResponse200305CountsRuleGroups

`func NewInlineResponse200305CountsRuleGroups() *InlineResponse200305CountsRuleGroups`

NewInlineResponse200305CountsRuleGroups instantiates a new InlineResponse200305CountsRuleGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305CountsRuleGroupsWithDefaults

`func NewInlineResponse200305CountsRuleGroupsWithDefaults() *InlineResponse200305CountsRuleGroups`

NewInlineResponse200305CountsRuleGroupsWithDefaults instantiates a new InlineResponse200305CountsRuleGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200305CountsRuleGroups) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200305CountsRuleGroups) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200305CountsRuleGroups) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200305CountsRuleGroups) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBySecurityLevel

`func (o *InlineResponse200305CountsRuleGroups) GetBySecurityLevel() InlineResponse200305CountsRuleGroupsBySecurityLevel`

GetBySecurityLevel returns the BySecurityLevel field if non-nil, zero value otherwise.

### GetBySecurityLevelOk

`func (o *InlineResponse200305CountsRuleGroups) GetBySecurityLevelOk() (*InlineResponse200305CountsRuleGroupsBySecurityLevel, bool)`

GetBySecurityLevelOk returns a tuple with the BySecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySecurityLevel

`func (o *InlineResponse200305CountsRuleGroups) SetBySecurityLevel(v InlineResponse200305CountsRuleGroupsBySecurityLevel)`

SetBySecurityLevel sets BySecurityLevel field to given value.

### HasBySecurityLevel

`func (o *InlineResponse200305CountsRuleGroups) HasBySecurityLevel() bool`

HasBySecurityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


