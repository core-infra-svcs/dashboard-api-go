# InlineObject70

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingRulesRules**](NetworksNetworkIdApplianceTrafficShapingRulesRules.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewInlineObject70

`func NewInlineObject70() *InlineObject70`

NewInlineObject70 instantiates a new InlineObject70 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject70WithDefaults

`func NewInlineObject70WithDefaults() *InlineObject70`

NewInlineObject70WithDefaults instantiates a new InlineObject70 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRulesEnabled

`func (o *InlineObject70) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *InlineObject70) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *InlineObject70) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *InlineObject70) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject70) GetRules() []NetworksNetworkIdApplianceTrafficShapingRulesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject70) GetRulesOk() (*[]NetworksNetworkIdApplianceTrafficShapingRulesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject70) SetRules(v []NetworksNetworkIdApplianceTrafficShapingRulesRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject70) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


