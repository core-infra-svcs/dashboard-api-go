# OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** | A rule identifier of the format meraki:intrusion/snort/GID/&lt;gid&gt;/SID/&lt;sid&gt;. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents | 
**Message** | Pointer to **string** | Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET | [optional] 

## Methods

### NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules

`func NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules(ruleId string, ) *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules`

NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules instantiates a new OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRulesWithDefaults

`func NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRulesWithDefaults() *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules`

NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRulesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetMessage

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


