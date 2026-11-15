# InlineResponse200302Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideId** | Pointer to **string** | Identifier for the rule override resource. | [optional] 
**Policy** | Pointer to [**InlineResponse200301Policy**](InlineResponse200301Policy.md) |  | [optional] 
**Rule** | Pointer to [**InlineResponse200302Rule**](InlineResponse200302Rule.md) |  | [optional] 
**Action** | Pointer to **NullableString** | Action override applied to the rule. | [optional] 

## Methods

### NewInlineResponse200302Items

`func NewInlineResponse200302Items() *InlineResponse200302Items`

NewInlineResponse200302Items instantiates a new InlineResponse200302Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302ItemsWithDefaults

`func NewInlineResponse200302ItemsWithDefaults() *InlineResponse200302Items`

NewInlineResponse200302ItemsWithDefaults instantiates a new InlineResponse200302Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideId

`func (o *InlineResponse200302Items) GetOverrideId() string`

GetOverrideId returns the OverrideId field if non-nil, zero value otherwise.

### GetOverrideIdOk

`func (o *InlineResponse200302Items) GetOverrideIdOk() (*string, bool)`

GetOverrideIdOk returns a tuple with the OverrideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideId

`func (o *InlineResponse200302Items) SetOverrideId(v string)`

SetOverrideId sets OverrideId field to given value.

### HasOverrideId

`func (o *InlineResponse200302Items) HasOverrideId() bool`

HasOverrideId returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200302Items) GetPolicy() InlineResponse200301Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200302Items) GetPolicyOk() (*InlineResponse200301Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200302Items) SetPolicy(v InlineResponse200301Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200302Items) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRule

`func (o *InlineResponse200302Items) GetRule() InlineResponse200302Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *InlineResponse200302Items) GetRuleOk() (*InlineResponse200302Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *InlineResponse200302Items) SetRule(v InlineResponse200302Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *InlineResponse200302Items) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetAction

`func (o *InlineResponse200302Items) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InlineResponse200302Items) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InlineResponse200302Items) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InlineResponse200302Items) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *InlineResponse200302Items) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *InlineResponse200302Items) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


