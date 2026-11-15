# InlineResponse200113Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to configure MD5 authentication between BGP peers | [optional] 

## Methods

### NewInlineResponse200113Authentication

`func NewInlineResponse200113Authentication() *InlineResponse200113Authentication`

NewInlineResponse200113Authentication instantiates a new InlineResponse200113Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200113AuthenticationWithDefaults

`func NewInlineResponse200113AuthenticationWithDefaults() *InlineResponse200113Authentication`

NewInlineResponse200113AuthenticationWithDefaults instantiates a new InlineResponse200113Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InlineResponse200113Authentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineResponse200113Authentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineResponse200113Authentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineResponse200113Authentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


