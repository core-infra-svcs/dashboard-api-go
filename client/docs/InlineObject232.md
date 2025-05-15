# InlineObject232

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules**](OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules.md) | An ordered array of the firewall rules (not including the default rule) | [optional] 
**SyslogDefaultRule** | Pointer to **bool** | Log the special default rule (boolean value - enable only if you&#39;ve configured a syslog server) (optional) | [optional] 

## Methods

### NewInlineObject232

`func NewInlineObject232() *InlineObject232`

NewInlineObject232 instantiates a new InlineObject232 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject232WithDefaults

`func NewInlineObject232WithDefaults() *InlineObject232`

NewInlineObject232WithDefaults instantiates a new InlineObject232 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject232) GetRules() []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject232) GetRulesOk() (*[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject232) SetRules(v []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject232) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSyslogDefaultRule

`func (o *InlineObject232) GetSyslogDefaultRule() bool`

GetSyslogDefaultRule returns the SyslogDefaultRule field if non-nil, zero value otherwise.

### GetSyslogDefaultRuleOk

`func (o *InlineObject232) GetSyslogDefaultRuleOk() (*bool, bool)`

GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDefaultRule

`func (o *InlineObject232) SetSyslogDefaultRule(v bool)`

SetSyslogDefaultRule sets SyslogDefaultRule field to given value.

### HasSyslogDefaultRule

`func (o *InlineObject232) HasSyslogDefaultRule() bool`

HasSyslogDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


