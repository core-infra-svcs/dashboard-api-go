# InlineResponse2025Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to [**InlineResponse2025ResultAuthority**](InlineResponse2025ResultAuthority.md) |  | [optional] 
**DeletedAt** | Pointer to **string** | ISO8601 timestamp when the CA was deleted (delete operations only) | [optional] 

## Methods

### NewInlineResponse2025Result

`func NewInlineResponse2025Result() *InlineResponse2025Result`

NewInlineResponse2025Result instantiates a new InlineResponse2025Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2025ResultWithDefaults

`func NewInlineResponse2025ResultWithDefaults() *InlineResponse2025Result`

NewInlineResponse2025ResultWithDefaults instantiates a new InlineResponse2025Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *InlineResponse2025Result) GetAuthority() InlineResponse2025ResultAuthority`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *InlineResponse2025Result) GetAuthorityOk() (*InlineResponse2025ResultAuthority, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *InlineResponse2025Result) SetAuthority(v InlineResponse2025ResultAuthority)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *InlineResponse2025Result) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDeletedAt

`func (o *InlineResponse2025Result) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InlineResponse2025Result) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InlineResponse2025Result) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InlineResponse2025Result) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


