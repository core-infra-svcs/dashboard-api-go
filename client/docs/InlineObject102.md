# InlineObject102

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. Only allowed If the user is not Dashboard administrator. | [optional] 
**Password** | Pointer to **string** | The password for this user account. Only allowed If the user is not Dashboard administrator. | [optional] 
**EmailPasswordToUser** | Pointer to **bool** | Whether or not Meraki should email the password to user. Default is false. | [optional] 
**Authorizations** | Pointer to [**[]NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations**](NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations.md) | Authorization zones and expiration dates for the user. | [optional] 

## Methods

### NewInlineObject102

`func NewInlineObject102() *InlineObject102`

NewInlineObject102 instantiates a new InlineObject102 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject102WithDefaults

`func NewInlineObject102WithDefaults() *InlineObject102`

NewInlineObject102WithDefaults instantiates a new InlineObject102 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject102) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject102) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject102) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject102) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject102) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject102) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject102) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject102) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmailPasswordToUser

`func (o *InlineObject102) GetEmailPasswordToUser() bool`

GetEmailPasswordToUser returns the EmailPasswordToUser field if non-nil, zero value otherwise.

### GetEmailPasswordToUserOk

`func (o *InlineObject102) GetEmailPasswordToUserOk() (*bool, bool)`

GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPasswordToUser

`func (o *InlineObject102) SetEmailPasswordToUser(v bool)`

SetEmailPasswordToUser sets EmailPasswordToUser field to given value.

### HasEmailPasswordToUser

`func (o *InlineObject102) HasEmailPasswordToUser() bool`

HasEmailPasswordToUser returns a boolean if a field has been set.

### GetAuthorizations

`func (o *InlineObject102) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *InlineObject102) GetAuthorizationsOk() (*[]NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *InlineObject102) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *InlineObject102) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


