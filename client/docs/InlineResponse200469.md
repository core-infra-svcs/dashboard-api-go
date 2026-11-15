# InlineResponse200469

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageItems** | Pointer to **int32** | Number of items in the page | [optional] 
**Contents** | Pointer to [**[]InlineResponse200469Contents**](InlineResponse200469Contents.md) | Array of daily metrics | [optional] 

## Methods

### NewInlineResponse200469

`func NewInlineResponse200469() *InlineResponse200469`

NewInlineResponse200469 instantiates a new InlineResponse200469 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200469WithDefaults

`func NewInlineResponse200469WithDefaults() *InlineResponse200469`

NewInlineResponse200469WithDefaults instantiates a new InlineResponse200469 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageItems

`func (o *InlineResponse200469) GetPageItems() int32`

GetPageItems returns the PageItems field if non-nil, zero value otherwise.

### GetPageItemsOk

`func (o *InlineResponse200469) GetPageItemsOk() (*int32, bool)`

GetPageItemsOk returns a tuple with the PageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageItems

`func (o *InlineResponse200469) SetPageItems(v int32)`

SetPageItems sets PageItems field to given value.

### HasPageItems

`func (o *InlineResponse200469) HasPageItems() bool`

HasPageItems returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200469) GetContents() []InlineResponse200469Contents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200469) GetContentsOk() (*[]InlineResponse200469Contents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200469) SetContents(v []InlineResponse200469Contents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200469) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


