# InlineResponse200212

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The ID of the adaptive policy group | [optional] 
**Name** | Pointer to **string** | The name of the adaptive policy group | [optional] 
**Sgt** | Pointer to **int32** | The security group tag for the adaptive policy group | [optional] 
**Description** | Pointer to **string** | The description for the adaptive policy group | [optional] 
**PolicyObjects** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects**](OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects.md) | The policy objects for the adaptive policy group | [optional] 
**IsDefaultGroup** | Pointer to **bool** | Whether the adaptive policy group is the default group | [optional] 
**RequiredIpMappings** | Pointer to **[]string** | List of required IP mappings for the adaptive policy group | [optional] 
**CreatedAt** | Pointer to **string** | Created at timestamp for the adaptive policy group | [optional] 
**UpdatedAt** | Pointer to **string** | Updated at timestamp for the adaptive policy group | [optional] 

## Methods

### NewInlineResponse200212

`func NewInlineResponse200212() *InlineResponse200212`

NewInlineResponse200212 instantiates a new InlineResponse200212 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200212WithDefaults

`func NewInlineResponse200212WithDefaults() *InlineResponse200212`

NewInlineResponse200212WithDefaults instantiates a new InlineResponse200212 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse200212) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse200212) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse200212) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse200212) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200212) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200212) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200212) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200212) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200212) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200212) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200212) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200212) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200212) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200212) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200212) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200212) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200212) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200212) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200212) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200212) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetIsDefaultGroup

`func (o *InlineResponse200212) GetIsDefaultGroup() bool`

GetIsDefaultGroup returns the IsDefaultGroup field if non-nil, zero value otherwise.

### GetIsDefaultGroupOk

`func (o *InlineResponse200212) GetIsDefaultGroupOk() (*bool, bool)`

GetIsDefaultGroupOk returns a tuple with the IsDefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultGroup

`func (o *InlineResponse200212) SetIsDefaultGroup(v bool)`

SetIsDefaultGroup sets IsDefaultGroup field to given value.

### HasIsDefaultGroup

`func (o *InlineResponse200212) HasIsDefaultGroup() bool`

HasIsDefaultGroup returns a boolean if a field has been set.

### GetRequiredIpMappings

`func (o *InlineResponse200212) GetRequiredIpMappings() []string`

GetRequiredIpMappings returns the RequiredIpMappings field if non-nil, zero value otherwise.

### GetRequiredIpMappingsOk

`func (o *InlineResponse200212) GetRequiredIpMappingsOk() (*[]string, bool)`

GetRequiredIpMappingsOk returns a tuple with the RequiredIpMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredIpMappings

`func (o *InlineResponse200212) SetRequiredIpMappings(v []string)`

SetRequiredIpMappings sets RequiredIpMappings field to given value.

### HasRequiredIpMappings

`func (o *InlineResponse200212) HasRequiredIpMappings() bool`

HasRequiredIpMappings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200212) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200212) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200212) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200212) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200212) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200212) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200212) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200212) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


