# InlineResponse200416Statuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to **string** | Overall status of network group | [optional] 
**ByProductType** | Pointer to [**[]InlineResponse200416StatusesByProductType**](InlineResponse200416StatusesByProductType.md) | List of status counts by product type | [optional] 

## Methods

### NewInlineResponse200416Statuses

`func NewInlineResponse200416Statuses() *InlineResponse200416Statuses`

NewInlineResponse200416Statuses instantiates a new InlineResponse200416Statuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200416StatusesWithDefaults

`func NewInlineResponse200416StatusesWithDefaults() *InlineResponse200416Statuses`

NewInlineResponse200416StatusesWithDefaults instantiates a new InlineResponse200416Statuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse200416Statuses) GetOverall() string`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200416Statuses) GetOverallOk() (*string, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200416Statuses) SetOverall(v string)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200416Statuses) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByProductType

`func (o *InlineResponse200416Statuses) GetByProductType() []InlineResponse200416StatusesByProductType`

GetByProductType returns the ByProductType field if non-nil, zero value otherwise.

### GetByProductTypeOk

`func (o *InlineResponse200416Statuses) GetByProductTypeOk() (*[]InlineResponse200416StatusesByProductType, bool)`

GetByProductTypeOk returns a tuple with the ByProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProductType

`func (o *InlineResponse200416Statuses) SetByProductType(v []InlineResponse200416StatusesByProductType)`

SetByProductType sets ByProductType field to given value.

### HasByProductType

`func (o *InlineResponse200416Statuses) HasByProductType() bool`

HasByProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


