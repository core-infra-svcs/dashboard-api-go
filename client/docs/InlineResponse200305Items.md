# InlineResponse200305Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroupId** | Pointer to **string** | Identifier for the rule group the counts belong to. | [optional] 
**Policy** | Pointer to [**InlineResponse200305Policy**](InlineResponse200305Policy.md) |  | [optional] 
**Counts** | Pointer to [**InlineResponse200305Counts**](InlineResponse200305Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200305Items

`func NewInlineResponse200305Items() *InlineResponse200305Items`

NewInlineResponse200305Items instantiates a new InlineResponse200305Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305ItemsWithDefaults

`func NewInlineResponse200305ItemsWithDefaults() *InlineResponse200305Items`

NewInlineResponse200305ItemsWithDefaults instantiates a new InlineResponse200305Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroupId

`func (o *InlineResponse200305Items) GetRuleGroupId() string`

GetRuleGroupId returns the RuleGroupId field if non-nil, zero value otherwise.

### GetRuleGroupIdOk

`func (o *InlineResponse200305Items) GetRuleGroupIdOk() (*string, bool)`

GetRuleGroupIdOk returns a tuple with the RuleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupId

`func (o *InlineResponse200305Items) SetRuleGroupId(v string)`

SetRuleGroupId sets RuleGroupId field to given value.

### HasRuleGroupId

`func (o *InlineResponse200305Items) HasRuleGroupId() bool`

HasRuleGroupId returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200305Items) GetPolicy() InlineResponse200305Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200305Items) GetPolicyOk() (*InlineResponse200305Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200305Items) SetPolicy(v InlineResponse200305Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200305Items) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200305Items) GetCounts() InlineResponse200305Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200305Items) GetCountsOk() (*InlineResponse200305Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200305Items) SetCounts(v InlineResponse200305Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200305Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


