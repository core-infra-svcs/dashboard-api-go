# InlineResponse200503

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200503Items**](InlineResponse200503Items.md) | Wireless LAN controller layer 2 interfaces historical status | [optional] 
**Meta** | Pointer to [**InlineResponse200279Meta**](InlineResponse200279Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200503

`func NewInlineResponse200503() *InlineResponse200503`

NewInlineResponse200503 instantiates a new InlineResponse200503 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200503WithDefaults

`func NewInlineResponse200503WithDefaults() *InlineResponse200503`

NewInlineResponse200503WithDefaults instantiates a new InlineResponse200503 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200503) GetItems() []InlineResponse200503Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200503) GetItemsOk() (*[]InlineResponse200503Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200503) SetItems(v []InlineResponse200503Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200503) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200503) GetMeta() InlineResponse200279Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200503) GetMetaOk() (*InlineResponse200279Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200503) SetMeta(v InlineResponse200279Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200503) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


