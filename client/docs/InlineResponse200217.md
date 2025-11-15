# InlineResponse200217

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

### NewInlineResponse200217

`func NewInlineResponse200217() *InlineResponse200217`

NewInlineResponse200217 instantiates a new InlineResponse200217 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200217WithDefaults

`func NewInlineResponse200217WithDefaults() *InlineResponse200217`

NewInlineResponse200217WithDefaults instantiates a new InlineResponse200217 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse200217) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse200217) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse200217) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse200217) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200217) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200217) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200217) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200217) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200217) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200217) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200217) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200217) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200217) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200217) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200217) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200217) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200217) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200217) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200217) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200217) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetIsDefaultGroup

`func (o *InlineResponse200217) GetIsDefaultGroup() bool`

GetIsDefaultGroup returns the IsDefaultGroup field if non-nil, zero value otherwise.

### GetIsDefaultGroupOk

`func (o *InlineResponse200217) GetIsDefaultGroupOk() (*bool, bool)`

GetIsDefaultGroupOk returns a tuple with the IsDefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultGroup

`func (o *InlineResponse200217) SetIsDefaultGroup(v bool)`

SetIsDefaultGroup sets IsDefaultGroup field to given value.

### HasIsDefaultGroup

`func (o *InlineResponse200217) HasIsDefaultGroup() bool`

HasIsDefaultGroup returns a boolean if a field has been set.

### GetRequiredIpMappings

`func (o *InlineResponse200217) GetRequiredIpMappings() []string`

GetRequiredIpMappings returns the RequiredIpMappings field if non-nil, zero value otherwise.

### GetRequiredIpMappingsOk

`func (o *InlineResponse200217) GetRequiredIpMappingsOk() (*[]string, bool)`

GetRequiredIpMappingsOk returns a tuple with the RequiredIpMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredIpMappings

`func (o *InlineResponse200217) SetRequiredIpMappings(v []string)`

SetRequiredIpMappings sets RequiredIpMappings field to given value.

### HasRequiredIpMappings

`func (o *InlineResponse200217) HasRequiredIpMappings() bool`

HasRequiredIpMappings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200217) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200217) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200217) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200217) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200217) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200217) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200217) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200217) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


