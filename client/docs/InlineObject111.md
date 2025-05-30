# InlineObject111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the user | 
**Name** | Pointer to **string** | Name of the user. Only required If the user is not a Dashboard administrator. | [optional] 
**Password** | Pointer to **string** | The password for this user account. Only required If the user is not a Dashboard administrator. | [optional] 
**AccountType** | Pointer to **string** | Authorization type for user. Can be &#39;Guest&#39; or &#39;802.1X&#39; for wireless networks, or &#39;Client VPN&#39; for MX networks. Defaults to &#39;802.1X&#39;. | [optional] [default to "802.1X"]
**EmailPasswordToUser** | Pointer to **bool** | Whether or not Meraki should email the password to user. Default is false. | [optional] 
**IsAdmin** | Pointer to **bool** | Whether or not the user is a Dashboard administrator. | [optional] 
**Authorizations** | [**[]NetworksNetworkIdMerakiAuthUsersAuthorizations1**](NetworksNetworkIdMerakiAuthUsersAuthorizations1.md) | Authorization zones and expiration dates for the user. | 

## Methods

### NewInlineObject111

`func NewInlineObject111(email string, authorizations []NetworksNetworkIdMerakiAuthUsersAuthorizations1, ) *InlineObject111`

NewInlineObject111 instantiates a new InlineObject111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject111WithDefaults

`func NewInlineObject111WithDefaults() *InlineObject111`

NewInlineObject111WithDefaults instantiates a new InlineObject111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineObject111) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject111) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject111) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *InlineObject111) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject111) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject111) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject111) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject111) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject111) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject111) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject111) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccountType

`func (o *InlineObject111) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InlineObject111) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InlineObject111) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InlineObject111) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetEmailPasswordToUser

`func (o *InlineObject111) GetEmailPasswordToUser() bool`

GetEmailPasswordToUser returns the EmailPasswordToUser field if non-nil, zero value otherwise.

### GetEmailPasswordToUserOk

`func (o *InlineObject111) GetEmailPasswordToUserOk() (*bool, bool)`

GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPasswordToUser

`func (o *InlineObject111) SetEmailPasswordToUser(v bool)`

SetEmailPasswordToUser sets EmailPasswordToUser field to given value.

### HasEmailPasswordToUser

`func (o *InlineObject111) HasEmailPasswordToUser() bool`

HasEmailPasswordToUser returns a boolean if a field has been set.

### GetIsAdmin

`func (o *InlineObject111) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *InlineObject111) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *InlineObject111) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *InlineObject111) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAuthorizations

`func (o *InlineObject111) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersAuthorizations1`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *InlineObject111) GetAuthorizationsOk() (*[]NetworksNetworkIdMerakiAuthUsersAuthorizations1, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *InlineObject111) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersAuthorizations1)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


