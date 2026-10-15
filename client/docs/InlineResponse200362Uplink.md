# InlineResponse200362Uplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | The uplink priority for this rule. One of: 1 or 2, where 1 is the highest priority. | [optional] 
**IsPreferred** | Pointer to **bool** | Whether this uplink is the preferred one to use once all rules have been exhausted. | [optional] 

## Methods

### NewInlineResponse200362Uplink

`func NewInlineResponse200362Uplink() *InlineResponse200362Uplink`

NewInlineResponse200362Uplink instantiates a new InlineResponse200362Uplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200362UplinkWithDefaults

`func NewInlineResponse200362UplinkWithDefaults() *InlineResponse200362Uplink`

NewInlineResponse200362UplinkWithDefaults instantiates a new InlineResponse200362Uplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *InlineResponse200362Uplink) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200362Uplink) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200362Uplink) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200362Uplink) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIsPreferred

`func (o *InlineResponse200362Uplink) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *InlineResponse200362Uplink) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *InlineResponse200362Uplink) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *InlineResponse200362Uplink) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


