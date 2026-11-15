# InlineResponse200373Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | The unique ID of the rule. | [optional] 
**Slot** | Pointer to **string** | The logical SIM identifier that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;esim&#39;. | [optional] 
**Uplink** | Pointer to [**InlineResponse200373Uplink**](InlineResponse200373Uplink.md) |  | [optional] 
**Cap** | Pointer to [**InlineResponse200373Cap**](InlineResponse200373Cap.md) |  | [optional] 
**Actions** | Pointer to [**[]InlineResponse200373Actions**](InlineResponse200373Actions.md) | The defined actions that will occur when the profile rule is triggered. At least one action and no more than two actions may be defined for a rule. | [optional] 

## Methods

### NewInlineResponse200373Rules

`func NewInlineResponse200373Rules() *InlineResponse200373Rules`

NewInlineResponse200373Rules instantiates a new InlineResponse200373Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373RulesWithDefaults

`func NewInlineResponse200373RulesWithDefaults() *InlineResponse200373Rules`

NewInlineResponse200373RulesWithDefaults instantiates a new InlineResponse200373Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200373Rules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200373Rules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200373Rules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200373Rules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200373Rules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200373Rules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200373Rules) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200373Rules) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200373Rules) GetUplink() InlineResponse200373Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200373Rules) GetUplinkOk() (*InlineResponse200373Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200373Rules) SetUplink(v InlineResponse200373Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200373Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetCap

`func (o *InlineResponse200373Rules) GetCap() InlineResponse200373Cap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *InlineResponse200373Rules) GetCapOk() (*InlineResponse200373Cap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *InlineResponse200373Rules) SetCap(v InlineResponse200373Cap)`

SetCap sets Cap field to given value.

### HasCap

`func (o *InlineResponse200373Rules) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200373Rules) GetActions() []InlineResponse200373Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200373Rules) GetActionsOk() (*[]InlineResponse200373Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200373Rules) SetActions(v []InlineResponse200373Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InlineResponse200373Rules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


