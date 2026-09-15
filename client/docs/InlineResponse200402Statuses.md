# InlineResponse200402Statuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to **string** | Overall status of network group | [optional] 
**ByProductType** | Pointer to [**[]InlineResponse200402StatusesByProductType**](InlineResponse200402StatusesByProductType.md) | List of status counts by product type | [optional] 

## Methods

### NewInlineResponse200402Statuses

`func NewInlineResponse200402Statuses() *InlineResponse200402Statuses`

NewInlineResponse200402Statuses instantiates a new InlineResponse200402Statuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200402StatusesWithDefaults

`func NewInlineResponse200402StatusesWithDefaults() *InlineResponse200402Statuses`

NewInlineResponse200402StatusesWithDefaults instantiates a new InlineResponse200402Statuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse200402Statuses) GetOverall() string`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200402Statuses) GetOverallOk() (*string, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200402Statuses) SetOverall(v string)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200402Statuses) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByProductType

`func (o *InlineResponse200402Statuses) GetByProductType() []InlineResponse200402StatusesByProductType`

GetByProductType returns the ByProductType field if non-nil, zero value otherwise.

### GetByProductTypeOk

`func (o *InlineResponse200402Statuses) GetByProductTypeOk() (*[]InlineResponse200402StatusesByProductType, bool)`

GetByProductTypeOk returns a tuple with the ByProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProductType

`func (o *InlineResponse200402Statuses) SetByProductType(v []InlineResponse200402StatusesByProductType)`

SetByProductType sets ByProductType field to given value.

### HasByProductType

`func (o *InlineResponse200402Statuses) HasByProductType() bool`

HasByProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


