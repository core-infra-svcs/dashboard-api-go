# InlineResponse200287Policies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** | The Id of the Sentry Policy | [optional] 
**NetworkId** | Pointer to **string** | The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group. | [optional] 
**SmNetworkId** | Pointer to **string** | The Id of the Systems Manager Network the Sentry Policy is assigned to | [optional] 
**Tags** | Pointer to **[]string** | The tags of the Sentry Policy | [optional] 
**Scope** | Pointer to **string** | The scope of the Sentry Policy | [optional] 
**GroupNumber** | Pointer to **string** | The number of the Group Policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The Id of the Group Policy. This is associated with the network specified by the networkId. | [optional] 
**Priority** | Pointer to **string** | The priority of the Sentry Policy | [optional] 
**CreatedAt** | Pointer to **string** | The creation time of the Sentry Policy | [optional] 
**LastUpdatedAt** | Pointer to **string** | The last update time of the Sentry Policy | [optional] 

## Methods

### NewInlineResponse200287Policies

`func NewInlineResponse200287Policies() *InlineResponse200287Policies`

NewInlineResponse200287Policies instantiates a new InlineResponse200287Policies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200287PoliciesWithDefaults

`func NewInlineResponse200287PoliciesWithDefaults() *InlineResponse200287Policies`

NewInlineResponse200287PoliciesWithDefaults instantiates a new InlineResponse200287Policies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *InlineResponse200287Policies) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *InlineResponse200287Policies) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *InlineResponse200287Policies) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *InlineResponse200287Policies) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200287Policies) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200287Policies) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200287Policies) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200287Policies) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSmNetworkId

`func (o *InlineResponse200287Policies) GetSmNetworkId() string`

GetSmNetworkId returns the SmNetworkId field if non-nil, zero value otherwise.

### GetSmNetworkIdOk

`func (o *InlineResponse200287Policies) GetSmNetworkIdOk() (*string, bool)`

GetSmNetworkIdOk returns a tuple with the SmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNetworkId

`func (o *InlineResponse200287Policies) SetSmNetworkId(v string)`

SetSmNetworkId sets SmNetworkId field to given value.

### HasSmNetworkId

`func (o *InlineResponse200287Policies) HasSmNetworkId() bool`

HasSmNetworkId returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200287Policies) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200287Policies) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200287Policies) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200287Policies) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetScope

`func (o *InlineResponse200287Policies) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse200287Policies) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse200287Policies) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse200287Policies) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetGroupNumber

`func (o *InlineResponse200287Policies) GetGroupNumber() string`

GetGroupNumber returns the GroupNumber field if non-nil, zero value otherwise.

### GetGroupNumberOk

`func (o *InlineResponse200287Policies) GetGroupNumberOk() (*string, bool)`

GetGroupNumberOk returns a tuple with the GroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumber

`func (o *InlineResponse200287Policies) SetGroupNumber(v string)`

SetGroupNumber sets GroupNumber field to given value.

### HasGroupNumber

`func (o *InlineResponse200287Policies) HasGroupNumber() bool`

HasGroupNumber returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200287Policies) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200287Policies) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200287Policies) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200287Policies) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200287Policies) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200287Policies) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200287Policies) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200287Policies) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200287Policies) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200287Policies) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200287Policies) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200287Policies) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200287Policies) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200287Policies) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200287Policies) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200287Policies) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


