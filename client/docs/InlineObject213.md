# InlineObject213

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the dashboard administrator. This attribute can not be updated. | 
**Name** | **string** | The name of the dashboard administrator | 
**OrgAccess** | **string** | The privilege of the dashboard administrator on the organization. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;enterprise&#39; or &#39;none&#39; | 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdAdminsTags1**](OrganizationsOrganizationIdAdminsTags1.md) | The list of tags that the dashboard administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdAdminsNetworks1**](OrganizationsOrganizationIdAdminsNetworks1.md) | The list of networks that the dashboard administrator has privileges on | [optional] 
**AuthenticationMethod** | Pointer to **string** | No longer used as of Cisco SecureX end-of-life. Can be one of &#39;Email&#39;. The default is Email authentication. | [optional] 

## Methods

### NewInlineObject213

`func NewInlineObject213(email string, name string, orgAccess string, ) *InlineObject213`

NewInlineObject213 instantiates a new InlineObject213 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject213WithDefaults

`func NewInlineObject213WithDefaults() *InlineObject213`

NewInlineObject213WithDefaults instantiates a new InlineObject213 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineObject213) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject213) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject213) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *InlineObject213) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject213) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject213) SetName(v string)`

SetName sets Name field to given value.


### GetOrgAccess

`func (o *InlineObject213) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineObject213) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineObject213) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.


### GetTags

`func (o *InlineObject213) GetTags() []OrganizationsOrganizationIdAdminsTags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject213) GetTagsOk() (*[]OrganizationsOrganizationIdAdminsTags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject213) SetTags(v []OrganizationsOrganizationIdAdminsTags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject213) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineObject213) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineObject213) GetNetworksOk() (*[]OrganizationsOrganizationIdAdminsNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineObject213) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineObject213) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *InlineObject213) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *InlineObject213) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *InlineObject213) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *InlineObject213) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


