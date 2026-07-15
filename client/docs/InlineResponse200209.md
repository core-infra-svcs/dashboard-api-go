# InlineResponse200209

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]InlineResponse200209StpBridgePriority**](InlineResponse200209StpBridgePriority.md) | STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewInlineResponse200209

`func NewInlineResponse200209() *InlineResponse200209`

NewInlineResponse200209 instantiates a new InlineResponse200209 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200209WithDefaults

`func NewInlineResponse200209WithDefaults() *InlineResponse200209`

NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *InlineResponse200209) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineResponse200209) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineResponse200209) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineResponse200209) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *InlineResponse200209) GetStpBridgePriority() []InlineResponse200209StpBridgePriority`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *InlineResponse200209) GetStpBridgePriorityOk() (*[]InlineResponse200209StpBridgePriority, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *InlineResponse200209) SetStpBridgePriority(v []InlineResponse200209StpBridgePriority)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *InlineResponse200209) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


