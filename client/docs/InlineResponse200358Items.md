# InlineResponse200358Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to [**InlineResponse200358Authority**](InlineResponse200358Authority.md) |  | [optional] 
**Feature** | Pointer to [**InlineResponse200358Feature**](InlineResponse200358Feature.md) |  | [optional] 
**Status** | Pointer to **string** | Feature certificate authority lifecycle state: untrusted (new), trusted (in use), revoked, or pending_revoked (revocation in progress). | [optional] 
**Contents** | Pointer to **string** | PEM encoded certificate authority contents with newlines as \\n; null if not generated | [optional] 
**Chain** | Pointer to **string** | PEM bundle of the complete certification path, ordered from this certificate authority up to the self signed root, with newlines as \\n. Only returned when includeChain is true. If a complete path is unavailable for any certificate authority in the response, the request fails rather than returning a partial path. | [optional] 

## Methods

### NewInlineResponse200358Items

`func NewInlineResponse200358Items() *InlineResponse200358Items`

NewInlineResponse200358Items instantiates a new InlineResponse200358Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358ItemsWithDefaults

`func NewInlineResponse200358ItemsWithDefaults() *InlineResponse200358Items`

NewInlineResponse200358ItemsWithDefaults instantiates a new InlineResponse200358Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *InlineResponse200358Items) GetAuthority() InlineResponse200358Authority`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *InlineResponse200358Items) GetAuthorityOk() (*InlineResponse200358Authority, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *InlineResponse200358Items) SetAuthority(v InlineResponse200358Authority)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *InlineResponse200358Items) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetFeature

`func (o *InlineResponse200358Items) GetFeature() InlineResponse200358Feature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *InlineResponse200358Items) GetFeatureOk() (*InlineResponse200358Feature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *InlineResponse200358Items) SetFeature(v InlineResponse200358Feature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *InlineResponse200358Items) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200358Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200358Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200358Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200358Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200358Items) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200358Items) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200358Items) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200358Items) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetChain

`func (o *InlineResponse200358Items) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *InlineResponse200358Items) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *InlineResponse200358Items) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *InlineResponse200358Items) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


