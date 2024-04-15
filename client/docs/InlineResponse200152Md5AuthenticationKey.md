# InlineResponse200152Md5AuthenticationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | MD5 authentication key index. Key index must be between 1 to 255 | [optional] 
**Passphrase** | Pointer to **string** | MD5 authentication passphrase | [optional] 

## Methods

### NewInlineResponse200152Md5AuthenticationKey

`func NewInlineResponse200152Md5AuthenticationKey() *InlineResponse200152Md5AuthenticationKey`

NewInlineResponse200152Md5AuthenticationKey instantiates a new InlineResponse200152Md5AuthenticationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200152Md5AuthenticationKeyWithDefaults

`func NewInlineResponse200152Md5AuthenticationKeyWithDefaults() *InlineResponse200152Md5AuthenticationKey`

NewInlineResponse200152Md5AuthenticationKeyWithDefaults instantiates a new InlineResponse200152Md5AuthenticationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200152Md5AuthenticationKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200152Md5AuthenticationKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200152Md5AuthenticationKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200152Md5AuthenticationKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassphrase

`func (o *InlineResponse200152Md5AuthenticationKey) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineResponse200152Md5AuthenticationKey) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineResponse200152Md5AuthenticationKey) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineResponse200152Md5AuthenticationKey) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


