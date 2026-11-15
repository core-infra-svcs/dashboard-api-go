# InlineResponse200300CountsRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of rules evaluated by the policy. | [optional] 
**Overridden** | Pointer to [**InlineResponse200300CountsRulesOverridden**](InlineResponse200300CountsRulesOverridden.md) |  | [optional] 
**ByAction** | Pointer to [**InlineResponse200300CountsRulesByAction**](InlineResponse200300CountsRulesByAction.md) |  | [optional] 

## Methods

### NewInlineResponse200300CountsRules

`func NewInlineResponse200300CountsRules() *InlineResponse200300CountsRules`

NewInlineResponse200300CountsRules instantiates a new InlineResponse200300CountsRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300CountsRulesWithDefaults

`func NewInlineResponse200300CountsRulesWithDefaults() *InlineResponse200300CountsRules`

NewInlineResponse200300CountsRulesWithDefaults instantiates a new InlineResponse200300CountsRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200300CountsRules) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200300CountsRules) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200300CountsRules) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200300CountsRules) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOverridden

`func (o *InlineResponse200300CountsRules) GetOverridden() InlineResponse200300CountsRulesOverridden`

GetOverridden returns the Overridden field if non-nil, zero value otherwise.

### GetOverriddenOk

`func (o *InlineResponse200300CountsRules) GetOverriddenOk() (*InlineResponse200300CountsRulesOverridden, bool)`

GetOverriddenOk returns a tuple with the Overridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridden

`func (o *InlineResponse200300CountsRules) SetOverridden(v InlineResponse200300CountsRulesOverridden)`

SetOverridden sets Overridden field to given value.

### HasOverridden

`func (o *InlineResponse200300CountsRules) HasOverridden() bool`

HasOverridden returns a boolean if a field has been set.

### GetByAction

`func (o *InlineResponse200300CountsRules) GetByAction() InlineResponse200300CountsRulesByAction`

GetByAction returns the ByAction field if non-nil, zero value otherwise.

### GetByActionOk

`func (o *InlineResponse200300CountsRules) GetByActionOk() (*InlineResponse200300CountsRulesByAction, bool)`

GetByActionOk returns a tuple with the ByAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAction

`func (o *InlineResponse200300CountsRules) SetByAction(v InlineResponse200300CountsRulesByAction)`

SetByAction sets ByAction field to given value.

### HasByAction

`func (o *InlineResponse200300CountsRules) HasByAction() bool`

HasByAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


