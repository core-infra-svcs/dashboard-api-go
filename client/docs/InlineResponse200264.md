# InlineResponse200264

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Admin&#39;s ID | [optional] 
**Name** | Pointer to **string** | Admin&#39;s username | [optional] 
**Email** | Pointer to **string** | Admin&#39;s email address | [optional] 
**OrgAccess** | Pointer to **string** | Admin&#39;s level of access to the organization | [optional] 
**AccountStatus** | Pointer to **string** | Status of the admin&#39;s account | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** | Indicates whether two-factor authentication is enabled | [optional] 
**HasApiKey** | Pointer to **bool** | Indicates whether the admin has an API key | [optional] 
**LastActive** | Pointer to **time.Time** | Time when the admin was last active | [optional] 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdAdminsTags**](OrganizationsOrganizationIdAdminsTags.md) | Admin tag information | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdAdminsNetworks**](OrganizationsOrganizationIdAdminsNetworks.md) | Admin network access information | [optional] 
**AuthenticationMethod** | Pointer to **string** | Admin&#39;s authentication method | [optional] 

## Methods

### NewInlineResponse200264

`func NewInlineResponse200264() *InlineResponse200264`

NewInlineResponse200264 instantiates a new InlineResponse200264 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200264WithDefaults

`func NewInlineResponse200264WithDefaults() *InlineResponse200264`

NewInlineResponse200264WithDefaults instantiates a new InlineResponse200264 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200264) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200264) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200264) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200264) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200264) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200264) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200264) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200264) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200264) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200264) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200264) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200264) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgAccess

`func (o *InlineResponse200264) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineResponse200264) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineResponse200264) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *InlineResponse200264) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetAccountStatus

`func (o *InlineResponse200264) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *InlineResponse200264) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *InlineResponse200264) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *InlineResponse200264) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *InlineResponse200264) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *InlineResponse200264) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *InlineResponse200264) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *InlineResponse200264) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetHasApiKey

`func (o *InlineResponse200264) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *InlineResponse200264) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *InlineResponse200264) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.

### HasHasApiKey

`func (o *InlineResponse200264) HasHasApiKey() bool`

HasHasApiKey returns a boolean if a field has been set.

### GetLastActive

`func (o *InlineResponse200264) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *InlineResponse200264) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *InlineResponse200264) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *InlineResponse200264) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200264) GetTags() []OrganizationsOrganizationIdAdminsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200264) GetTagsOk() (*[]OrganizationsOrganizationIdAdminsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200264) SetTags(v []OrganizationsOrganizationIdAdminsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200264) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse200264) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse200264) GetNetworksOk() (*[]OrganizationsOrganizationIdAdminsNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse200264) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse200264) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *InlineResponse200264) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *InlineResponse200264) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *InlineResponse200264) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *InlineResponse200264) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


