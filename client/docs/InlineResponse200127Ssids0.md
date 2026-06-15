# InlineResponse200127Ssids0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthorized** | Pointer to **bool** | Whether the client is currently authorized on this SSID | [optional] 
**AuthorizedAt** | Pointer to **NullableString** | When the client was authorized on this SSID in ISO 8601 format. This field will be null if the client is not currently authorized. | [optional] 
**ExpiresAt** | Pointer to **NullableString** | When the client&#39;s authorization on this SSID expires in ISO 8601 format. This field will be null if the client is not currently authorized or if there is no expiration time for the authorization. | [optional] 

## Methods

### NewInlineResponse200127Ssids0

`func NewInlineResponse200127Ssids0() *InlineResponse200127Ssids0`

NewInlineResponse200127Ssids0 instantiates a new InlineResponse200127Ssids0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200127Ssids0WithDefaults

`func NewInlineResponse200127Ssids0WithDefaults() *InlineResponse200127Ssids0`

NewInlineResponse200127Ssids0WithDefaults instantiates a new InlineResponse200127Ssids0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthorized

`func (o *InlineResponse200127Ssids0) GetIsAuthorized() bool`

GetIsAuthorized returns the IsAuthorized field if non-nil, zero value otherwise.

### GetIsAuthorizedOk

`func (o *InlineResponse200127Ssids0) GetIsAuthorizedOk() (*bool, bool)`

GetIsAuthorizedOk returns a tuple with the IsAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthorized

`func (o *InlineResponse200127Ssids0) SetIsAuthorized(v bool)`

SetIsAuthorized sets IsAuthorized field to given value.

### HasIsAuthorized

`func (o *InlineResponse200127Ssids0) HasIsAuthorized() bool`

HasIsAuthorized returns a boolean if a field has been set.

### GetAuthorizedAt

`func (o *InlineResponse200127Ssids0) GetAuthorizedAt() string`

GetAuthorizedAt returns the AuthorizedAt field if non-nil, zero value otherwise.

### GetAuthorizedAtOk

`func (o *InlineResponse200127Ssids0) GetAuthorizedAtOk() (*string, bool)`

GetAuthorizedAtOk returns a tuple with the AuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAt

`func (o *InlineResponse200127Ssids0) SetAuthorizedAt(v string)`

SetAuthorizedAt sets AuthorizedAt field to given value.

### HasAuthorizedAt

`func (o *InlineResponse200127Ssids0) HasAuthorizedAt() bool`

HasAuthorizedAt returns a boolean if a field has been set.

### SetAuthorizedAtNil

`func (o *InlineResponse200127Ssids0) SetAuthorizedAtNil(b bool)`

 SetAuthorizedAtNil sets the value for AuthorizedAt to be an explicit nil

### UnsetAuthorizedAt
`func (o *InlineResponse200127Ssids0) UnsetAuthorizedAt()`

UnsetAuthorizedAt ensures that no value is present for AuthorizedAt, not even an explicit nil
### GetExpiresAt

`func (o *InlineResponse200127Ssids0) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse200127Ssids0) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse200127Ssids0) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse200127Ssids0) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *InlineResponse200127Ssids0) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *InlineResponse200127Ssids0) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


