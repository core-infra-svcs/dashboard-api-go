# InlineObject140

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchPorts** | Pointer to [**[]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts1**](NetworksNetworkIdSwitchLinkAggregationsSwitchPorts1.md) | Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported. | [optional] 
**SwitchProfilePorts** | Pointer to [**[]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts**](NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts.md) | Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported. | [optional] 

## Methods

### NewInlineObject140

`func NewInlineObject140() *InlineObject140`

NewInlineObject140 instantiates a new InlineObject140 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject140WithDefaults

`func NewInlineObject140WithDefaults() *InlineObject140`

NewInlineObject140WithDefaults instantiates a new InlineObject140 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchPorts

`func (o *InlineObject140) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts1`

GetSwitchPorts returns the SwitchPorts field if non-nil, zero value otherwise.

### GetSwitchPortsOk

`func (o *InlineObject140) GetSwitchPortsOk() (*[]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts1, bool)`

GetSwitchPortsOk returns a tuple with the SwitchPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPorts

`func (o *InlineObject140) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts1)`

SetSwitchPorts sets SwitchPorts field to given value.

### HasSwitchPorts

`func (o *InlineObject140) HasSwitchPorts() bool`

HasSwitchPorts returns a boolean if a field has been set.

### GetSwitchProfilePorts

`func (o *InlineObject140) GetSwitchProfilePorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts`

GetSwitchProfilePorts returns the SwitchProfilePorts field if non-nil, zero value otherwise.

### GetSwitchProfilePortsOk

`func (o *InlineObject140) GetSwitchProfilePortsOk() (*[]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts, bool)`

GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilePorts

`func (o *InlineObject140) SetSwitchProfilePorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts)`

SetSwitchProfilePorts sets SwitchProfilePorts field to given value.

### HasSwitchProfilePorts

`func (o *InlineObject140) HasSwitchProfilePorts() bool`

HasSwitchProfilePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


