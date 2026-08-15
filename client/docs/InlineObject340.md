# InlineObject340

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role of the SAML administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the SAML administrator on the organization. Can be one of &#39;none&#39;, &#39;read-only&#39;, &#39;full&#39; or &#39;enterprise&#39; or a custom role in the format custom-role:ID:NAME. | [optional] 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesTags1**](OrganizationsOrganizationIdSamlRolesTags1.md) | The list of tags that the SAML administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesNetworks1**](OrganizationsOrganizationIdSamlRolesNetworks1.md) | The list of networks that the SAML administrator has privileges on | [optional] 

## Methods

### NewInlineObject340

`func NewInlineObject340() *InlineObject340`

NewInlineObject340 instantiates a new InlineObject340 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject340WithDefaults

`func NewInlineObject340WithDefaults() *InlineObject340`

NewInlineObject340WithDefaults instantiates a new InlineObject340 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *InlineObject340) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InlineObject340) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InlineObject340) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InlineObject340) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrgAccess

`func (o *InlineObject340) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineObject340) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineObject340) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *InlineObject340) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject340) GetTags() []OrganizationsOrganizationIdSamlRolesTags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject340) GetTagsOk() (*[]OrganizationsOrganizationIdSamlRolesTags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject340) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject340) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineObject340) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineObject340) GetNetworksOk() (*[]OrganizationsOrganizationIdSamlRolesNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineObject340) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineObject340) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


