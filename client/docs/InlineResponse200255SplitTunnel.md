# InlineResponse200255SplitTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, VPN split tunnel is enabled. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200255SplitTunnelRules**](InlineResponse200255SplitTunnelRules.md) | List of VPN split tunnel rules. | [optional] 

## Methods

### NewInlineResponse200255SplitTunnel

`func NewInlineResponse200255SplitTunnel() *InlineResponse200255SplitTunnel`

NewInlineResponse200255SplitTunnel instantiates a new InlineResponse200255SplitTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200255SplitTunnelWithDefaults

`func NewInlineResponse200255SplitTunnelWithDefaults() *InlineResponse200255SplitTunnel`

NewInlineResponse200255SplitTunnelWithDefaults instantiates a new InlineResponse200255SplitTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200255SplitTunnel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200255SplitTunnel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200255SplitTunnel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200255SplitTunnel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200255SplitTunnel) GetRules() []InlineResponse200255SplitTunnelRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200255SplitTunnel) GetRulesOk() (*[]InlineResponse200255SplitTunnelRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200255SplitTunnel) SetRules(v []InlineResponse200255SplitTunnelRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200255SplitTunnel) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


