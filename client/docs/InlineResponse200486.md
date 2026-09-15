# InlineResponse200486

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200486Items**](InlineResponse200486Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200279Meta**](InlineResponse200279Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200486

`func NewInlineResponse200486() *InlineResponse200486`

NewInlineResponse200486 instantiates a new InlineResponse200486 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200486WithDefaults

`func NewInlineResponse200486WithDefaults() *InlineResponse200486`

NewInlineResponse200486WithDefaults instantiates a new InlineResponse200486 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200486) GetItems() []InlineResponse200486Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200486) GetItemsOk() (*[]InlineResponse200486Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200486) SetItems(v []InlineResponse200486Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200486) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200486) GetMeta() InlineResponse200279Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200486) GetMetaOk() (*InlineResponse200279Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200486) SetMeta(v InlineResponse200279Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200486) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


