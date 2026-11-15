# InlineResponse200305CountsRulesByAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **int32** | Child rules configured with alert/detect | [optional] 
**Block** | Pointer to **int32** | Child rules configured to block/drop | [optional] 
**Other** | Pointer to **int32** | Child rules configured with other actions | [optional] 

## Methods

### NewInlineResponse200305CountsRulesByAction

`func NewInlineResponse200305CountsRulesByAction() *InlineResponse200305CountsRulesByAction`

NewInlineResponse200305CountsRulesByAction instantiates a new InlineResponse200305CountsRulesByAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305CountsRulesByActionWithDefaults

`func NewInlineResponse200305CountsRulesByActionWithDefaults() *InlineResponse200305CountsRulesByAction`

NewInlineResponse200305CountsRulesByActionWithDefaults instantiates a new InlineResponse200305CountsRulesByAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *InlineResponse200305CountsRulesByAction) GetAlert() int32`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *InlineResponse200305CountsRulesByAction) GetAlertOk() (*int32, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *InlineResponse200305CountsRulesByAction) SetAlert(v int32)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *InlineResponse200305CountsRulesByAction) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetBlock

`func (o *InlineResponse200305CountsRulesByAction) GetBlock() int32`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *InlineResponse200305CountsRulesByAction) GetBlockOk() (*int32, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *InlineResponse200305CountsRulesByAction) SetBlock(v int32)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *InlineResponse200305CountsRulesByAction) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetOther

`func (o *InlineResponse200305CountsRulesByAction) GetOther() int32`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *InlineResponse200305CountsRulesByAction) GetOtherOk() (*int32, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *InlineResponse200305CountsRulesByAction) SetOther(v int32)`

SetOther sets Other field to given value.

### HasOther

`func (o *InlineResponse200305CountsRulesByAction) HasOther() bool`

HasOther returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


