# InlineObject292

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the SAML administrator | 
**OrgAccess** | **string** | The privilege of the SAML administrator on the organization. Can be one of &#39;none&#39;, &#39;read-only&#39;, &#39;full&#39; or &#39;enterprise&#39; | 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesTags1**](OrganizationsOrganizationIdSamlRolesTags1.md) | The list of tags that the SAML administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesNetworks1**](OrganizationsOrganizationIdSamlRolesNetworks1.md) | The list of networks that the SAML administrator has privileges on | [optional] 

## Methods

### NewInlineObject292

`func NewInlineObject292(role string, orgAccess string, ) *InlineObject292`

NewInlineObject292 instantiates a new InlineObject292 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject292WithDefaults

`func NewInlineObject292WithDefaults() *InlineObject292`

NewInlineObject292WithDefaults instantiates a new InlineObject292 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *InlineObject292) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InlineObject292) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InlineObject292) SetRole(v string)`

SetRole sets Role field to given value.


### GetOrgAccess

`func (o *InlineObject292) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineObject292) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineObject292) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.


### GetTags

`func (o *InlineObject292) GetTags() []OrganizationsOrganizationIdSamlRolesTags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject292) GetTagsOk() (*[]OrganizationsOrganizationIdSamlRolesTags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject292) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject292) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineObject292) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineObject292) GetNetworksOk() (*[]OrganizationsOrganizationIdSamlRolesNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineObject292) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineObject292) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


