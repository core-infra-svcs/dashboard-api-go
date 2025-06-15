# InlineObject205

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficShapingEnabled** | Pointer to **bool** | Whether traffic shaping rules are applied to clients on your SSID. | [optional] 
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200207Rules**](InlineResponse200207Rules.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewInlineObject205

`func NewInlineObject205() *InlineObject205`

NewInlineObject205 instantiates a new InlineObject205 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject205WithDefaults

`func NewInlineObject205WithDefaults() *InlineObject205`

NewInlineObject205WithDefaults instantiates a new InlineObject205 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficShapingEnabled

`func (o *InlineObject205) GetTrafficShapingEnabled() bool`

GetTrafficShapingEnabled returns the TrafficShapingEnabled field if non-nil, zero value otherwise.

### GetTrafficShapingEnabledOk

`func (o *InlineObject205) GetTrafficShapingEnabledOk() (*bool, bool)`

GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingEnabled

`func (o *InlineObject205) SetTrafficShapingEnabled(v bool)`

SetTrafficShapingEnabled sets TrafficShapingEnabled field to given value.

### HasTrafficShapingEnabled

`func (o *InlineObject205) HasTrafficShapingEnabled() bool`

HasTrafficShapingEnabled returns a boolean if a field has been set.

### GetDefaultRulesEnabled

`func (o *InlineObject205) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *InlineObject205) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *InlineObject205) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *InlineObject205) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject205) GetRules() []InlineResponse200207Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject205) GetRulesOk() (*[]InlineResponse200207Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject205) SetRules(v []InlineResponse200207Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject205) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


