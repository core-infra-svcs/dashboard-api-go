# InlineResponse200305CountsRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of child rules for this rule group. | [optional] 
**ByAction** | Pointer to [**InlineResponse200305CountsRulesByAction**](InlineResponse200305CountsRulesByAction.md) |  | [optional] 

## Methods

### NewInlineResponse200305CountsRules

`func NewInlineResponse200305CountsRules() *InlineResponse200305CountsRules`

NewInlineResponse200305CountsRules instantiates a new InlineResponse200305CountsRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305CountsRulesWithDefaults

`func NewInlineResponse200305CountsRulesWithDefaults() *InlineResponse200305CountsRules`

NewInlineResponse200305CountsRulesWithDefaults instantiates a new InlineResponse200305CountsRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200305CountsRules) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200305CountsRules) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200305CountsRules) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200305CountsRules) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByAction

`func (o *InlineResponse200305CountsRules) GetByAction() InlineResponse200305CountsRulesByAction`

GetByAction returns the ByAction field if non-nil, zero value otherwise.

### GetByActionOk

`func (o *InlineResponse200305CountsRules) GetByActionOk() (*InlineResponse200305CountsRulesByAction, bool)`

GetByActionOk returns a tuple with the ByAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAction

`func (o *InlineResponse200305CountsRules) SetByAction(v InlineResponse200305CountsRulesByAction)`

SetByAction sets ByAction field to given value.

### HasByAction

`func (o *InlineResponse200305CountsRules) HasByAction() bool`

HasByAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


