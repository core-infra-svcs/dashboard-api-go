# InlineResponse20050

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse20049Rules**](InlineResponse20049Rules.md) | An ordered array of the firewall rules (not including the default rule) | [optional] 
**SyslogDefaultRule** | Pointer to **bool** | Log the special default rule (boolean value - enable only if you&#39;ve configured a syslog server) (optional) | [optional] 

## Methods

### NewInlineResponse20050

`func NewInlineResponse20050() *InlineResponse20050`

NewInlineResponse20050 instantiates a new InlineResponse20050 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20050WithDefaults

`func NewInlineResponse20050WithDefaults() *InlineResponse20050`

NewInlineResponse20050WithDefaults instantiates a new InlineResponse20050 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse20050) GetRules() []InlineResponse20049Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse20050) GetRulesOk() (*[]InlineResponse20049Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse20050) SetRules(v []InlineResponse20049Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse20050) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSyslogDefaultRule

`func (o *InlineResponse20050) GetSyslogDefaultRule() bool`

GetSyslogDefaultRule returns the SyslogDefaultRule field if non-nil, zero value otherwise.

### GetSyslogDefaultRuleOk

`func (o *InlineResponse20050) GetSyslogDefaultRuleOk() (*bool, bool)`

GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDefaultRule

`func (o *InlineResponse20050) SetSyslogDefaultRule(v bool)`

SetSyslogDefaultRule sets SyslogDefaultRule field to given value.

### HasSyslogDefaultRule

`func (o *InlineResponse20050) HasSyslogDefaultRule() bool`

HasSyslogDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


