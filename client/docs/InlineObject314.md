# InlineObject314

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the firewall rule | 
**RulesetId** | **string** | Firewall ruleset ID to associate the rule with | 
**Policy** | **string** | Rule policy - allow or deny traffic | 
**Enabled** | Pointer to **bool** | Whether the rule is enabled | [optional] 
**Priority** | Pointer to **int32** | Rule priority (lower numbers &#x3D; higher priority) | [optional] 
**Description** | Pointer to **string** | Description of the firewall rule | [optional] 
**Sources** | [**OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources.md) |  | 
**Destinations** | [**OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations.md) |  | 

## Methods

### NewInlineObject314

`func NewInlineObject314(name string, rulesetId string, policy string, sources OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources, destinations OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations, ) *InlineObject314`

NewInlineObject314 instantiates a new InlineObject314 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject314WithDefaults

`func NewInlineObject314WithDefaults() *InlineObject314`

NewInlineObject314WithDefaults instantiates a new InlineObject314 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject314) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject314) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject314) SetName(v string)`

SetName sets Name field to given value.


### GetRulesetId

`func (o *InlineObject314) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineObject314) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineObject314) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.


### GetPolicy

`func (o *InlineObject314) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineObject314) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineObject314) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetEnabled

`func (o *InlineObject314) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject314) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject314) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject314) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineObject314) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineObject314) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineObject314) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineObject314) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject314) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject314) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject314) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject314) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *InlineObject314) GetSources() OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InlineObject314) GetSourcesOk() (*OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InlineObject314) SetSources(v OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources)`

SetSources sets Sources field to given value.


### GetDestinations

`func (o *InlineObject314) GetDestinations() OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InlineObject314) GetDestinationsOk() (*OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InlineObject314) SetDestinations(v OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations)`

SetDestinations sets Destinations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


