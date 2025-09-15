# InlineObject217

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroup** | [**OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1**](OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1.md) |  | 
**DestinationGroup** | [**OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1**](OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1.md) |  | 
**Acls** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1**](OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1.md) | An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: []) | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL (default: \&quot;default\&quot;) | [optional] 

## Methods

### NewInlineObject217

`func NewInlineObject217(sourceGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1, destinationGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1, ) *InlineObject217`

NewInlineObject217 instantiates a new InlineObject217 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject217WithDefaults

`func NewInlineObject217WithDefaults() *InlineObject217`

NewInlineObject217WithDefaults instantiates a new InlineObject217 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroup

`func (o *InlineObject217) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *InlineObject217) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *InlineObject217) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup1)`

SetSourceGroup sets SourceGroup field to given value.


### GetDestinationGroup

`func (o *InlineObject217) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *InlineObject217) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *InlineObject217) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup1)`

SetDestinationGroup sets DestinationGroup field to given value.


### GetAcls

`func (o *InlineObject217) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *InlineObject217) GetAclsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *InlineObject217) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls1)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *InlineObject217) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *InlineObject217) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *InlineObject217) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *InlineObject217) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *InlineObject217) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


