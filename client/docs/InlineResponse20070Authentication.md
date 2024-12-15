# InlineResponse20070Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to configure MD5 authentication between BGP peers | [optional] 

## Methods

### NewInlineResponse20070Authentication

`func NewInlineResponse20070Authentication() *InlineResponse20070Authentication`

NewInlineResponse20070Authentication instantiates a new InlineResponse20070Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20070AuthenticationWithDefaults

`func NewInlineResponse20070AuthenticationWithDefaults() *InlineResponse20070Authentication`

NewInlineResponse20070AuthenticationWithDefaults instantiates a new InlineResponse20070Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InlineResponse20070Authentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineResponse20070Authentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineResponse20070Authentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineResponse20070Authentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


