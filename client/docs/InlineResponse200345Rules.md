# InlineResponse200345Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | The unique ID of the rule. | [optional] 
**Slot** | Pointer to **string** | The logical SIM identifier that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;esim&#39;. | [optional] 
**Uplink** | Pointer to [**InlineResponse200344Uplink**](InlineResponse200344Uplink.md) |  | [optional] 
**Cap** | Pointer to [**InlineResponse200345Cap**](InlineResponse200345Cap.md) |  | [optional] 
**Actions** | Pointer to [**[]InlineResponse200344Actions**](InlineResponse200344Actions.md) | The defined actions that will occur when the profile rule is triggered. No more than two actions may be defined for a rule. | [optional] 

## Methods

### NewInlineResponse200345Rules

`func NewInlineResponse200345Rules() *InlineResponse200345Rules`

NewInlineResponse200345Rules instantiates a new InlineResponse200345Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200345RulesWithDefaults

`func NewInlineResponse200345RulesWithDefaults() *InlineResponse200345Rules`

NewInlineResponse200345RulesWithDefaults instantiates a new InlineResponse200345Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200345Rules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200345Rules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200345Rules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200345Rules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200345Rules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200345Rules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200345Rules) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200345Rules) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200345Rules) GetUplink() InlineResponse200344Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200345Rules) GetUplinkOk() (*InlineResponse200344Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200345Rules) SetUplink(v InlineResponse200344Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200345Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetCap

`func (o *InlineResponse200345Rules) GetCap() InlineResponse200345Cap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *InlineResponse200345Rules) GetCapOk() (*InlineResponse200345Cap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *InlineResponse200345Rules) SetCap(v InlineResponse200345Cap)`

SetCap sets Cap field to given value.

### HasCap

`func (o *InlineResponse200345Rules) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200345Rules) GetActions() []InlineResponse200344Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200345Rules) GetActionsOk() (*[]InlineResponse200344Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200345Rules) SetActions(v []InlineResponse200344Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InlineResponse200345Rules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


