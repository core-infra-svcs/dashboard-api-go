# InlineResponse200520

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200520Items**](InlineResponse200520Items.md) | Paginated list of scanning api receivers by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200286Meta**](InlineResponse200286Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200520

`func NewInlineResponse200520() *InlineResponse200520`

NewInlineResponse200520 instantiates a new InlineResponse200520 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200520WithDefaults

`func NewInlineResponse200520WithDefaults() *InlineResponse200520`

NewInlineResponse200520WithDefaults instantiates a new InlineResponse200520 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200520) GetItems() []InlineResponse200520Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200520) GetItemsOk() (*[]InlineResponse200520Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200520) SetItems(v []InlineResponse200520Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200520) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200520) GetMeta() InlineResponse200286Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200520) GetMetaOk() (*InlineResponse200286Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200520) SetMeta(v InlineResponse200286Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200520) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


