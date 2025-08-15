# InlineObject236

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules**](OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules.md) | An ordered array of the firewall rules (not including the default rule) | [optional] 
**SyslogDefaultRule** | Pointer to **bool** | Log the special default rule (boolean value - enable only if you&#39;ve configured a syslog server) (optional) | [optional] 

## Methods

### NewInlineObject236

`func NewInlineObject236() *InlineObject236`

NewInlineObject236 instantiates a new InlineObject236 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject236WithDefaults

`func NewInlineObject236WithDefaults() *InlineObject236`

NewInlineObject236WithDefaults instantiates a new InlineObject236 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject236) GetRules() []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject236) GetRulesOk() (*[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject236) SetRules(v []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject236) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSyslogDefaultRule

`func (o *InlineObject236) GetSyslogDefaultRule() bool`

GetSyslogDefaultRule returns the SyslogDefaultRule field if non-nil, zero value otherwise.

### GetSyslogDefaultRuleOk

`func (o *InlineObject236) GetSyslogDefaultRuleOk() (*bool, bool)`

GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDefaultRule

`func (o *InlineObject236) SetSyslogDefaultRule(v bool)`

SetSyslogDefaultRule sets SyslogDefaultRule field to given value.

### HasSyslogDefaultRule

`func (o *InlineObject236) HasSyslogDefaultRule() bool`

HasSyslogDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


