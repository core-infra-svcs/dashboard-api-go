# InlineResponse200201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficShapingEnabled** | Pointer to **bool** | Whether traffic shaping rules are applied to clients on your SSID. | [optional] 
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200201Rules**](InlineResponse200201Rules.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewInlineResponse200201

`func NewInlineResponse200201() *InlineResponse200201`

NewInlineResponse200201 instantiates a new InlineResponse200201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200201WithDefaults

`func NewInlineResponse200201WithDefaults() *InlineResponse200201`

NewInlineResponse200201WithDefaults instantiates a new InlineResponse200201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficShapingEnabled

`func (o *InlineResponse200201) GetTrafficShapingEnabled() bool`

GetTrafficShapingEnabled returns the TrafficShapingEnabled field if non-nil, zero value otherwise.

### GetTrafficShapingEnabledOk

`func (o *InlineResponse200201) GetTrafficShapingEnabledOk() (*bool, bool)`

GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingEnabled

`func (o *InlineResponse200201) SetTrafficShapingEnabled(v bool)`

SetTrafficShapingEnabled sets TrafficShapingEnabled field to given value.

### HasTrafficShapingEnabled

`func (o *InlineResponse200201) HasTrafficShapingEnabled() bool`

HasTrafficShapingEnabled returns a boolean if a field has been set.

### GetDefaultRulesEnabled

`func (o *InlineResponse200201) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *InlineResponse200201) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *InlineResponse200201) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *InlineResponse200201) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200201) GetRules() []InlineResponse200201Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200201) GetRulesOk() (*[]InlineResponse200201Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200201) SetRules(v []InlineResponse200201Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200201) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


