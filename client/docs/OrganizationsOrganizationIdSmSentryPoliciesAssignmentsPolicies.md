# OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** | The Sentry Policy Id, if updating an existing Sentry Policy | [optional] 
**SmNetworkId** | **string** | The Id of the Systems Manager Network | 
**Scope** | **string** | The scope of the Sentry Policy | 
**Tags** | **[]string** | The tags for the Sentry Policy | 
**GroupPolicyId** | **string** | The Group Policy Id | 

## Methods

### NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies

`func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies(smNetworkId string, scope string, tags []string, groupPolicyId string, ) *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies`

NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPoliciesWithDefaults

`func NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPoliciesWithDefaults() *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies`

NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsPoliciesWithDefaults instantiates a new OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetSmNetworkId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetSmNetworkId() string`

GetSmNetworkId returns the SmNetworkId field if non-nil, zero value otherwise.

### GetSmNetworkIdOk

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetSmNetworkIdOk() (*string, bool)`

GetSmNetworkIdOk returns a tuple with the SmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNetworkId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetSmNetworkId(v string)`

SetSmNetworkId sets SmNetworkId field to given value.


### GetScope

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTags

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetGroupPolicyId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *OrganizationsOrganizationIdSmSentryPoliciesAssignmentsPolicies) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


