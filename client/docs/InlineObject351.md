# InlineObject351

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetId** | **string** | ID of the ruleset to assign | 
**PolicyId** | **string** | ID of the policy to assign the ruleset to | 
**Priority** | Pointer to **int32** | Priority of the ruleset assignment (lower numbers &#x3D; higher priority) | [optional] 
**Staged** | Pointer to **bool** | Stage an assignment without applying it immediately to the policy | [optional] 

## Methods

### NewInlineObject351

`func NewInlineObject351(rulesetId string, policyId string, ) *InlineObject351`

NewInlineObject351 instantiates a new InlineObject351 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject351WithDefaults

`func NewInlineObject351WithDefaults() *InlineObject351`

NewInlineObject351WithDefaults instantiates a new InlineObject351 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetId

`func (o *InlineObject351) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineObject351) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineObject351) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.


### GetPolicyId

`func (o *InlineObject351) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *InlineObject351) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *InlineObject351) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetPriority

`func (o *InlineObject351) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineObject351) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineObject351) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineObject351) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStaged

`func (o *InlineObject351) GetStaged() bool`

GetStaged returns the Staged field if non-nil, zero value otherwise.

### GetStagedOk

`func (o *InlineObject351) GetStagedOk() (*bool, bool)`

GetStagedOk returns a tuple with the Staged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaged

`func (o *InlineObject351) SetStaged(v bool)`

SetStaged sets Staged field to given value.

### HasStaged

`func (o *InlineObject351) HasStaged() bool`

HasStaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


