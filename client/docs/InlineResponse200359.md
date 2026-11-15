# InlineResponse200359

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to [**InlineResponse200358Authority**](InlineResponse200358Authority.md) |  | [optional] 
**Feature** | Pointer to [**InlineResponse200358Feature**](InlineResponse200358Feature.md) |  | [optional] 
**Status** | Pointer to **string** | Feature certificate authority lifecycle state: untrusted (new), trusted (in use), revoked, or pending_revoked (revocation in progress). | [optional] 
**Contents** | Pointer to **string** | PEM encoded certificate authority contents with newlines as \\n; null if not generated | [optional] 

## Methods

### NewInlineResponse200359

`func NewInlineResponse200359() *InlineResponse200359`

NewInlineResponse200359 instantiates a new InlineResponse200359 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200359WithDefaults

`func NewInlineResponse200359WithDefaults() *InlineResponse200359`

NewInlineResponse200359WithDefaults instantiates a new InlineResponse200359 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *InlineResponse200359) GetAuthority() InlineResponse200358Authority`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *InlineResponse200359) GetAuthorityOk() (*InlineResponse200358Authority, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *InlineResponse200359) SetAuthority(v InlineResponse200358Authority)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *InlineResponse200359) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetFeature

`func (o *InlineResponse200359) GetFeature() InlineResponse200358Feature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *InlineResponse200359) GetFeatureOk() (*InlineResponse200358Feature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *InlineResponse200359) SetFeature(v InlineResponse200358Feature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *InlineResponse200359) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200359) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200359) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200359) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200359) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200359) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200359) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200359) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200359) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


