# InlineResponse200331Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | The unique ID of the rule. | [optional] 
**Slot** | Pointer to **string** | The SIM slot that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;sim3&#39;. | [optional] 
**Uplink** | Pointer to [**InlineResponse200331Uplink**](InlineResponse200331Uplink.md) |  | [optional] 
**Cap** | Pointer to [**InlineResponse200331Cap**](InlineResponse200331Cap.md) |  | [optional] 
**Actions** | Pointer to [**[]InlineResponse200331Actions**](InlineResponse200331Actions.md) | The defined actions that will occur when the profile rule is triggered. At least one action and no more than two actions may be defined for a rule. | [optional] 

## Methods

### NewInlineResponse200331Rules

`func NewInlineResponse200331Rules() *InlineResponse200331Rules`

NewInlineResponse200331Rules instantiates a new InlineResponse200331Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200331RulesWithDefaults

`func NewInlineResponse200331RulesWithDefaults() *InlineResponse200331Rules`

NewInlineResponse200331RulesWithDefaults instantiates a new InlineResponse200331Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200331Rules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200331Rules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200331Rules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200331Rules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200331Rules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200331Rules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200331Rules) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200331Rules) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200331Rules) GetUplink() InlineResponse200331Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200331Rules) GetUplinkOk() (*InlineResponse200331Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200331Rules) SetUplink(v InlineResponse200331Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200331Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetCap

`func (o *InlineResponse200331Rules) GetCap() InlineResponse200331Cap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *InlineResponse200331Rules) GetCapOk() (*InlineResponse200331Cap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *InlineResponse200331Rules) SetCap(v InlineResponse200331Cap)`

SetCap sets Cap field to given value.

### HasCap

`func (o *InlineResponse200331Rules) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200331Rules) GetActions() []InlineResponse200331Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200331Rules) GetActionsOk() (*[]InlineResponse200331Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200331Rules) SetActions(v []InlineResponse200331Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InlineResponse200331Rules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


