# InlineResponse200208

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptivePolicyId** | Pointer to **string** | The ID for the adaptive policy | [optional] 
**SourceGroup** | Pointer to [**OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup.md) |  | [optional] 
**DestinationGroup** | Pointer to [**OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup.md) |  | [optional] 
**Acls** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls**](OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls.md) | The access control lists for the adaptive policy | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL | [optional] 
**CreatedAt** | Pointer to **string** | The created at timestamp for the adaptive policy | [optional] 
**UpdatedAt** | Pointer to **string** | The updated at timestamp for the adaptive policy | [optional] 

## Methods

### NewInlineResponse200208

`func NewInlineResponse200208() *InlineResponse200208`

NewInlineResponse200208 instantiates a new InlineResponse200208 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200208WithDefaults

`func NewInlineResponse200208WithDefaults() *InlineResponse200208`

NewInlineResponse200208WithDefaults instantiates a new InlineResponse200208 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptivePolicyId

`func (o *InlineResponse200208) GetAdaptivePolicyId() string`

GetAdaptivePolicyId returns the AdaptivePolicyId field if non-nil, zero value otherwise.

### GetAdaptivePolicyIdOk

`func (o *InlineResponse200208) GetAdaptivePolicyIdOk() (*string, bool)`

GetAdaptivePolicyIdOk returns a tuple with the AdaptivePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyId

`func (o *InlineResponse200208) SetAdaptivePolicyId(v string)`

SetAdaptivePolicyId sets AdaptivePolicyId field to given value.

### HasAdaptivePolicyId

`func (o *InlineResponse200208) HasAdaptivePolicyId() bool`

HasAdaptivePolicyId returns a boolean if a field has been set.

### GetSourceGroup

`func (o *InlineResponse200208) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *InlineResponse200208) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *InlineResponse200208) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *InlineResponse200208) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *InlineResponse200208) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *InlineResponse200208) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *InlineResponse200208) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *InlineResponse200208) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetAcls

`func (o *InlineResponse200208) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *InlineResponse200208) GetAclsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *InlineResponse200208) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *InlineResponse200208) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *InlineResponse200208) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *InlineResponse200208) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *InlineResponse200208) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *InlineResponse200208) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200208) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200208) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200208) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200208) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200208) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200208) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200208) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200208) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


