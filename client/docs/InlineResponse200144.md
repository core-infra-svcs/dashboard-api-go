# InlineResponse200144

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Meraki managed Id of the user record. | [optional] 
**Email** | Pointer to **string** | User email. | [optional] 
**FullName** | Pointer to **string** | User full name. | [optional] 
**Username** | Pointer to **string** | The users username. | [optional] 
**HasPassword** | Pointer to **bool** | A boolean denoting if the user has a password associated with the record. | [optional] 
**Tags** | Pointer to **string** | The set of tags the user is scoped to. | [optional] 
**AdGroups** | Pointer to **[]string** | Active Directory Groups the user belongs to. | [optional] 
**AzureAdGroups** | Pointer to **[]string** | Azure Active Directory Groups the user belongs to. | [optional] 
**SamlGroups** | Pointer to **[]string** | SAML Groups the user belongs to. | [optional] 
**AsmGroups** | Pointer to **[]string** | Apple School Manager Groups the user belongs to. | [optional] 
**IsExternal** | Pointer to **bool** | Whether the user was created using an external integration, or via the Meraki Dashboard. | [optional] 
**DisplayName** | Pointer to **string** | The user display name. | [optional] 
**HasIdentityCertificate** | Pointer to **bool** | A boolean indicating if the user has an associated identity certificate.. | [optional] 
**UserThumbnail** | Pointer to **string** | The url where the users thumbnail is hosted. | [optional] 

## Methods

### NewInlineResponse200144

`func NewInlineResponse200144() *InlineResponse200144`

NewInlineResponse200144 instantiates a new InlineResponse200144 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200144WithDefaults

`func NewInlineResponse200144WithDefaults() *InlineResponse200144`

NewInlineResponse200144WithDefaults instantiates a new InlineResponse200144 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200144) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200144) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200144) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200144) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200144) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200144) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200144) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200144) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *InlineResponse200144) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InlineResponse200144) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InlineResponse200144) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InlineResponse200144) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse200144) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse200144) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse200144) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse200144) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetHasPassword

`func (o *InlineResponse200144) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *InlineResponse200144) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *InlineResponse200144) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *InlineResponse200144) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200144) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200144) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200144) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200144) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAdGroups

`func (o *InlineResponse200144) GetAdGroups() []string`

GetAdGroups returns the AdGroups field if non-nil, zero value otherwise.

### GetAdGroupsOk

`func (o *InlineResponse200144) GetAdGroupsOk() (*[]string, bool)`

GetAdGroupsOk returns a tuple with the AdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdGroups

`func (o *InlineResponse200144) SetAdGroups(v []string)`

SetAdGroups sets AdGroups field to given value.

### HasAdGroups

`func (o *InlineResponse200144) HasAdGroups() bool`

HasAdGroups returns a boolean if a field has been set.

### GetAzureAdGroups

`func (o *InlineResponse200144) GetAzureAdGroups() []string`

GetAzureAdGroups returns the AzureAdGroups field if non-nil, zero value otherwise.

### GetAzureAdGroupsOk

`func (o *InlineResponse200144) GetAzureAdGroupsOk() (*[]string, bool)`

GetAzureAdGroupsOk returns a tuple with the AzureAdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdGroups

`func (o *InlineResponse200144) SetAzureAdGroups(v []string)`

SetAzureAdGroups sets AzureAdGroups field to given value.

### HasAzureAdGroups

`func (o *InlineResponse200144) HasAzureAdGroups() bool`

HasAzureAdGroups returns a boolean if a field has been set.

### GetSamlGroups

`func (o *InlineResponse200144) GetSamlGroups() []string`

GetSamlGroups returns the SamlGroups field if non-nil, zero value otherwise.

### GetSamlGroupsOk

`func (o *InlineResponse200144) GetSamlGroupsOk() (*[]string, bool)`

GetSamlGroupsOk returns a tuple with the SamlGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroups

`func (o *InlineResponse200144) SetSamlGroups(v []string)`

SetSamlGroups sets SamlGroups field to given value.

### HasSamlGroups

`func (o *InlineResponse200144) HasSamlGroups() bool`

HasSamlGroups returns a boolean if a field has been set.

### GetAsmGroups

`func (o *InlineResponse200144) GetAsmGroups() []string`

GetAsmGroups returns the AsmGroups field if non-nil, zero value otherwise.

### GetAsmGroupsOk

`func (o *InlineResponse200144) GetAsmGroupsOk() (*[]string, bool)`

GetAsmGroupsOk returns a tuple with the AsmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmGroups

`func (o *InlineResponse200144) SetAsmGroups(v []string)`

SetAsmGroups sets AsmGroups field to given value.

### HasAsmGroups

`func (o *InlineResponse200144) HasAsmGroups() bool`

HasAsmGroups returns a boolean if a field has been set.

### GetIsExternal

`func (o *InlineResponse200144) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *InlineResponse200144) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *InlineResponse200144) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *InlineResponse200144) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineResponse200144) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineResponse200144) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineResponse200144) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineResponse200144) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHasIdentityCertificate

`func (o *InlineResponse200144) GetHasIdentityCertificate() bool`

GetHasIdentityCertificate returns the HasIdentityCertificate field if non-nil, zero value otherwise.

### GetHasIdentityCertificateOk

`func (o *InlineResponse200144) GetHasIdentityCertificateOk() (*bool, bool)`

GetHasIdentityCertificateOk returns a tuple with the HasIdentityCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIdentityCertificate

`func (o *InlineResponse200144) SetHasIdentityCertificate(v bool)`

SetHasIdentityCertificate sets HasIdentityCertificate field to given value.

### HasHasIdentityCertificate

`func (o *InlineResponse200144) HasHasIdentityCertificate() bool`

HasHasIdentityCertificate returns a boolean if a field has been set.

### GetUserThumbnail

`func (o *InlineResponse200144) GetUserThumbnail() string`

GetUserThumbnail returns the UserThumbnail field if non-nil, zero value otherwise.

### GetUserThumbnailOk

`func (o *InlineResponse200144) GetUserThumbnailOk() (*string, bool)`

GetUserThumbnailOk returns a tuple with the UserThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserThumbnail

`func (o *InlineResponse200144) SetUserThumbnail(v string)`

SetUserThumbnail sets UserThumbnail field to given value.

### HasUserThumbnail

`func (o *InlineResponse200144) HasUserThumbnail() bool`

HasUserThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


