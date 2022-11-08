# InlineObject137

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]NetworksNetworkIdSwitchStpStpBridgePriority**](NetworksNetworkIdSwitchStpStpBridgePriority.md) | STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewInlineObject137

`func NewInlineObject137() *InlineObject137`

NewInlineObject137 instantiates a new InlineObject137 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject137WithDefaults

`func NewInlineObject137WithDefaults() *InlineObject137`

NewInlineObject137WithDefaults instantiates a new InlineObject137 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *InlineObject137) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject137) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject137) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject137) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *InlineObject137) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *InlineObject137) GetStpBridgePriorityOk() (*[]NetworksNetworkIdSwitchStpStpBridgePriority, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *InlineObject137) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *InlineObject137) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


