# InlineObject155

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroup** | Pointer to [**OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup.md) |  | [optional] 
**DestinationGroup** | Pointer to [**OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup.md) |  | [optional] 
**Acls** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls**](OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls.md) | An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy  | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL  | [optional] 

## Methods

### NewInlineObject155

`func NewInlineObject155() *InlineObject155`

NewInlineObject155 instantiates a new InlineObject155 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject155WithDefaults

`func NewInlineObject155WithDefaults() *InlineObject155`

NewInlineObject155WithDefaults instantiates a new InlineObject155 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroup

`func (o *InlineObject155) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *InlineObject155) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *InlineObject155) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *InlineObject155) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *InlineObject155) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *InlineObject155) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *InlineObject155) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *InlineObject155) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetAcls

`func (o *InlineObject155) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *InlineObject155) GetAclsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *InlineObject155) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *InlineObject155) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *InlineObject155) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *InlineObject155) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *InlineObject155) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *InlineObject155) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

