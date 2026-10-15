# InlineResponse200363Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | The unique ID of the rule. | [optional] 
**Slot** | Pointer to **string** | The logical SIM identifier that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;esim&#39;. | [optional] 
**Uplink** | Pointer to [**InlineResponse200362Uplink**](InlineResponse200362Uplink.md) |  | [optional] 
**Cap** | Pointer to [**InlineResponse200363Cap**](InlineResponse200363Cap.md) |  | [optional] 
**Actions** | Pointer to [**[]InlineResponse200362Actions**](InlineResponse200362Actions.md) | The defined actions that will occur when the profile rule is triggered. No more than two actions may be defined for a rule. | [optional] 

## Methods

### NewInlineResponse200363Rules

`func NewInlineResponse200363Rules() *InlineResponse200363Rules`

NewInlineResponse200363Rules instantiates a new InlineResponse200363Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200363RulesWithDefaults

`func NewInlineResponse200363RulesWithDefaults() *InlineResponse200363Rules`

NewInlineResponse200363RulesWithDefaults instantiates a new InlineResponse200363Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200363Rules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200363Rules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200363Rules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200363Rules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200363Rules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200363Rules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200363Rules) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200363Rules) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200363Rules) GetUplink() InlineResponse200362Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200363Rules) GetUplinkOk() (*InlineResponse200362Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200363Rules) SetUplink(v InlineResponse200362Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200363Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetCap

`func (o *InlineResponse200363Rules) GetCap() InlineResponse200363Cap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *InlineResponse200363Rules) GetCapOk() (*InlineResponse200363Cap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *InlineResponse200363Rules) SetCap(v InlineResponse200363Cap)`

SetCap sets Cap field to given value.

### HasCap

`func (o *InlineResponse200363Rules) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200363Rules) GetActions() []InlineResponse200362Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200363Rules) GetActionsOk() (*[]InlineResponse200362Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200363Rules) SetActions(v []InlineResponse200362Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InlineResponse200363Rules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


