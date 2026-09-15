# InlineResponse200348Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | The unique ID of the rule. | [optional] 
**Slot** | Pointer to **string** | The logical SIM identifier that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;esim&#39;. | [optional] 
**Uplink** | Pointer to [**InlineResponse200348Uplink**](InlineResponse200348Uplink.md) |  | [optional] 
**Cap** | Pointer to [**InlineResponse200348Cap**](InlineResponse200348Cap.md) |  | [optional] 
**Actions** | Pointer to [**[]InlineResponse200348Actions**](InlineResponse200348Actions.md) | The defined actions that will occur when the profile rule is triggered. At least one action and no more than two actions may be defined for a rule. | [optional] 

## Methods

### NewInlineResponse200348Rules

`func NewInlineResponse200348Rules() *InlineResponse200348Rules`

NewInlineResponse200348Rules instantiates a new InlineResponse200348Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200348RulesWithDefaults

`func NewInlineResponse200348RulesWithDefaults() *InlineResponse200348Rules`

NewInlineResponse200348RulesWithDefaults instantiates a new InlineResponse200348Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200348Rules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200348Rules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200348Rules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200348Rules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200348Rules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200348Rules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200348Rules) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200348Rules) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200348Rules) GetUplink() InlineResponse200348Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200348Rules) GetUplinkOk() (*InlineResponse200348Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200348Rules) SetUplink(v InlineResponse200348Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200348Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetCap

`func (o *InlineResponse200348Rules) GetCap() InlineResponse200348Cap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *InlineResponse200348Rules) GetCapOk() (*InlineResponse200348Cap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *InlineResponse200348Rules) SetCap(v InlineResponse200348Cap)`

SetCap sets Cap field to given value.

### HasCap

`func (o *InlineResponse200348Rules) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200348Rules) GetActions() []InlineResponse200348Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200348Rules) GetActionsOk() (*[]InlineResponse200348Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200348Rules) SetActions(v []InlineResponse200348Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InlineResponse200348Rules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


