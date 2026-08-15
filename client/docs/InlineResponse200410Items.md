# InlineResponse200410Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentId** | Pointer to **string** | ID of the ruleset assignment | [optional] 
**RulesetId** | Pointer to **string** | ID of the ruleset being assigned | [optional] 
**PolicyId** | Pointer to **string** | ID of the policy the ruleset is assigned to | [optional] 
**Priority** | Pointer to **int32** | Priority of the ruleset assignment (higher numbers &#x3D; higher priority) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Assignment creation time | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Assignment last update time | [optional] 

## Methods

### NewInlineResponse200410Items

`func NewInlineResponse200410Items() *InlineResponse200410Items`

NewInlineResponse200410Items instantiates a new InlineResponse200410Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200410ItemsWithDefaults

`func NewInlineResponse200410ItemsWithDefaults() *InlineResponse200410Items`

NewInlineResponse200410ItemsWithDefaults instantiates a new InlineResponse200410Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentId

`func (o *InlineResponse200410Items) GetAssignmentId() string`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *InlineResponse200410Items) GetAssignmentIdOk() (*string, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *InlineResponse200410Items) SetAssignmentId(v string)`

SetAssignmentId sets AssignmentId field to given value.

### HasAssignmentId

`func (o *InlineResponse200410Items) HasAssignmentId() bool`

HasAssignmentId returns a boolean if a field has been set.

### GetRulesetId

`func (o *InlineResponse200410Items) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineResponse200410Items) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineResponse200410Items) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *InlineResponse200410Items) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetPolicyId

`func (o *InlineResponse200410Items) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *InlineResponse200410Items) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *InlineResponse200410Items) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *InlineResponse200410Items) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200410Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200410Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200410Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200410Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200410Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200410Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200410Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200410Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200410Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200410Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200410Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200410Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


