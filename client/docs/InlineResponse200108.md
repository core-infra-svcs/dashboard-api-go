# InlineResponse200108

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Meraki auth user id | [optional] 
**Email** | Pointer to **string** | Email address of the user | [optional] 
**Name** | Pointer to **string** | Name of the user | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time of the user | [optional] 
**AccountType** | Pointer to **string** | Authorization type for user. | [optional] 
**IsAdmin** | Pointer to **bool** | Whether or not the user is a Dashboard administrator | [optional] 
**Authorizations** | Pointer to [**[]NetworksNetworkIdMerakiAuthUsersAuthorizations**](NetworksNetworkIdMerakiAuthUsersAuthorizations.md) | User authorization info | [optional] 

## Methods

### NewInlineResponse200108

`func NewInlineResponse200108() *InlineResponse200108`

NewInlineResponse200108 instantiates a new InlineResponse200108 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108WithDefaults

`func NewInlineResponse200108WithDefaults() *InlineResponse200108`

NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200108) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200108) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200108) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200108) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200108) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200108) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200108) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200108) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200108) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200108) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200108) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200108) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200108) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200108) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200108) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200108) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccountType

`func (o *InlineResponse200108) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InlineResponse200108) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InlineResponse200108) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InlineResponse200108) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetIsAdmin

`func (o *InlineResponse200108) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *InlineResponse200108) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *InlineResponse200108) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *InlineResponse200108) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAuthorizations

`func (o *InlineResponse200108) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *InlineResponse200108) GetAuthorizationsOk() (*[]NetworksNetworkIdMerakiAuthUsersAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *InlineResponse200108) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *InlineResponse200108) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


