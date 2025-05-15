# InlineObject112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. Only allowed If the user is not Dashboard administrator. | [optional] 
**Password** | Pointer to **string** | The password for this user account. Only allowed If the user is not Dashboard administrator. | [optional] 
**EmailPasswordToUser** | Pointer to **bool** | Whether or not Meraki should email the password to user. Default is false. | [optional] 
**Authorizations** | Pointer to [**[]NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations**](NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations.md) | Authorization zones and expiration dates for the user. | [optional] 

## Methods

### NewInlineObject112

`func NewInlineObject112() *InlineObject112`

NewInlineObject112 instantiates a new InlineObject112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject112WithDefaults

`func NewInlineObject112WithDefaults() *InlineObject112`

NewInlineObject112WithDefaults instantiates a new InlineObject112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject112) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject112) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject112) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject112) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject112) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject112) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject112) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject112) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmailPasswordToUser

`func (o *InlineObject112) GetEmailPasswordToUser() bool`

GetEmailPasswordToUser returns the EmailPasswordToUser field if non-nil, zero value otherwise.

### GetEmailPasswordToUserOk

`func (o *InlineObject112) GetEmailPasswordToUserOk() (*bool, bool)`

GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPasswordToUser

`func (o *InlineObject112) SetEmailPasswordToUser(v bool)`

SetEmailPasswordToUser sets EmailPasswordToUser field to given value.

### HasEmailPasswordToUser

`func (o *InlineObject112) HasEmailPasswordToUser() bool`

HasEmailPasswordToUser returns a boolean if a field has been set.

### GetAuthorizations

`func (o *InlineObject112) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *InlineObject112) GetAuthorizationsOk() (*[]NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *InlineObject112) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *InlineObject112) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


