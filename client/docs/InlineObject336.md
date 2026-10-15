# InlineObject336

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the firewall rule | [optional] 
**RulesetId** | Pointer to **string** | Firewall ruleset ID to associate the rule with | [optional] 
**Policy** | Pointer to **string** | Rule policy - allow or deny traffic | [optional] 
**Enabled** | Pointer to **bool** | Whether the rule is enabled | [optional] 
**Priority** | Pointer to **int32** | Rule priority (lower numbers &#x3D; higher priority) | [optional] 
**Description** | Pointer to **string** | Description of the firewall rule | [optional] 
**Sources** | Pointer to [**OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources.md) |  | [optional] 
**Destinations** | Pointer to [**OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations.md) |  | [optional] 

## Methods

### NewInlineObject336

`func NewInlineObject336() *InlineObject336`

NewInlineObject336 instantiates a new InlineObject336 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject336WithDefaults

`func NewInlineObject336WithDefaults() *InlineObject336`

NewInlineObject336WithDefaults instantiates a new InlineObject336 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject336) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject336) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject336) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject336) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRulesetId

`func (o *InlineObject336) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineObject336) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineObject336) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *InlineObject336) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineObject336) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineObject336) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineObject336) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineObject336) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject336) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject336) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject336) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject336) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineObject336) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineObject336) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineObject336) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineObject336) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject336) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject336) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject336) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject336) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *InlineObject336) GetSources() OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InlineObject336) GetSourcesOk() (*OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InlineObject336) SetSources(v OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *InlineObject336) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *InlineObject336) GetDestinations() OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InlineObject336) GetDestinationsOk() (*OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InlineObject336) SetDestinations(v OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *InlineObject336) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


