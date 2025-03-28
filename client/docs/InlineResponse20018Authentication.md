# InlineResponse20018Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | APN auth type. | [optional] [default to "none"]
**Username** | Pointer to **string** | APN username, if type is set. | [optional] 
**Password** | Pointer to **string** | APN password, if type is set (if APN password is not supplied, the password is left unchanged). | [optional] 

## Methods

### NewInlineResponse20018Authentication

`func NewInlineResponse20018Authentication() *InlineResponse20018Authentication`

NewInlineResponse20018Authentication instantiates a new InlineResponse20018Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20018AuthenticationWithDefaults

`func NewInlineResponse20018AuthenticationWithDefaults() *InlineResponse20018Authentication`

NewInlineResponse20018AuthenticationWithDefaults instantiates a new InlineResponse20018Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20018Authentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20018Authentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20018Authentication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20018Authentication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse20018Authentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse20018Authentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse20018Authentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse20018Authentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *InlineResponse20018Authentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineResponse20018Authentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineResponse20018Authentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineResponse20018Authentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


