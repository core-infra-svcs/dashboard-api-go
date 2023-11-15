# InlineResponse200162

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID associated with the SAML role | [optional] 
**Role** | Pointer to **string** | The role of the SAML administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the SAML administrator on the organization | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesNetworks**](OrganizationsOrganizationIdSamlRolesNetworks.md) | The list of networks that the SAML administrator has privileges on | [optional] 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdSamlRolesTags**](OrganizationsOrganizationIdSamlRolesTags.md) | The list of tags that the SAML administrator has privleges on | [optional] 

## Methods

### NewInlineResponse200162

`func NewInlineResponse200162() *InlineResponse200162`

NewInlineResponse200162 instantiates a new InlineResponse200162 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200162WithDefaults

`func NewInlineResponse200162WithDefaults() *InlineResponse200162`

NewInlineResponse200162WithDefaults instantiates a new InlineResponse200162 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200162) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200162) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200162) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200162) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *InlineResponse200162) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InlineResponse200162) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InlineResponse200162) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InlineResponse200162) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrgAccess

`func (o *InlineResponse200162) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineResponse200162) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineResponse200162) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *InlineResponse200162) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse200162) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse200162) GetNetworksOk() (*[]OrganizationsOrganizationIdSamlRolesNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse200162) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse200162) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200162) GetTags() []OrganizationsOrganizationIdSamlRolesTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200162) GetTagsOk() (*[]OrganizationsOrganizationIdSamlRolesTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200162) SetTags(v []OrganizationsOrganizationIdSamlRolesTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200162) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

