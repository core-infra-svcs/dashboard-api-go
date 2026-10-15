# InlineResponse200460

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** | Current page index | [optional] 
**PageItems** | Pointer to **int32** | Number of items in the page | [optional] 
**TotalItems** | Pointer to **int32** | Total number of items | [optional] 
**HasNext** | Pointer to **bool** | Is there more data to load | [optional] 
**HasPrevious** | Pointer to **bool** | Is there previous data to load | [optional] 
**Contents** | Pointer to [**[]InlineResponse200460Contents**](InlineResponse200460Contents.md) | Metrics content | [optional] 

## Methods

### NewInlineResponse200460

`func NewInlineResponse200460() *InlineResponse200460`

NewInlineResponse200460 instantiates a new InlineResponse200460 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200460WithDefaults

`func NewInlineResponse200460WithDefaults() *InlineResponse200460`

NewInlineResponse200460WithDefaults instantiates a new InlineResponse200460 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *InlineResponse200460) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *InlineResponse200460) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *InlineResponse200460) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *InlineResponse200460) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageItems

`func (o *InlineResponse200460) GetPageItems() int32`

GetPageItems returns the PageItems field if non-nil, zero value otherwise.

### GetPageItemsOk

`func (o *InlineResponse200460) GetPageItemsOk() (*int32, bool)`

GetPageItemsOk returns a tuple with the PageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageItems

`func (o *InlineResponse200460) SetPageItems(v int32)`

SetPageItems sets PageItems field to given value.

### HasPageItems

`func (o *InlineResponse200460) HasPageItems() bool`

HasPageItems returns a boolean if a field has been set.

### GetTotalItems

`func (o *InlineResponse200460) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *InlineResponse200460) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *InlineResponse200460) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *InlineResponse200460) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetHasNext

`func (o *InlineResponse200460) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *InlineResponse200460) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *InlineResponse200460) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *InlineResponse200460) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrevious

`func (o *InlineResponse200460) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *InlineResponse200460) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *InlineResponse200460) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *InlineResponse200460) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200460) GetContents() []InlineResponse200460Contents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200460) GetContentsOk() (*[]InlineResponse200460Contents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200460) SetContents(v []InlineResponse200460Contents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200460) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


