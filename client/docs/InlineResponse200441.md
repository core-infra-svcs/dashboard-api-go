# InlineResponse200441

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200441Items**](InlineResponse200441Items.md) | List of available network enforcement targets according to the provided filters. Maximum 100 items per page. | [optional] 
**Meta** | Pointer to [**InlineResponse200286Meta**](InlineResponse200286Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200441

`func NewInlineResponse200441() *InlineResponse200441`

NewInlineResponse200441 instantiates a new InlineResponse200441 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200441WithDefaults

`func NewInlineResponse200441WithDefaults() *InlineResponse200441`

NewInlineResponse200441WithDefaults instantiates a new InlineResponse200441 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200441) GetItems() []InlineResponse200441Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200441) GetItemsOk() (*[]InlineResponse200441Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200441) SetItems(v []InlineResponse200441Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200441) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200441) GetMeta() InlineResponse200286Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200441) GetMetaOk() (*InlineResponse200286Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200441) SetMeta(v InlineResponse200286Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200441) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


