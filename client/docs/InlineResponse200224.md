# InlineResponse200224

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200224Items**](InlineResponse200224Items.md) | Organization Alert counts by type | 
**Meta** | [**InlineResponse200223Meta**](InlineResponse200223Meta.md) |  | 

## Methods

### NewInlineResponse200224

`func NewInlineResponse200224(items []InlineResponse200224Items, meta InlineResponse200223Meta, ) *InlineResponse200224`

NewInlineResponse200224 instantiates a new InlineResponse200224 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200224WithDefaults

`func NewInlineResponse200224WithDefaults() *InlineResponse200224`

NewInlineResponse200224WithDefaults instantiates a new InlineResponse200224 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200224) GetItems() []InlineResponse200224Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200224) GetItemsOk() (*[]InlineResponse200224Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200224) SetItems(v []InlineResponse200224Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200224) GetMeta() InlineResponse200223Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200224) GetMetaOk() (*InlineResponse200223Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200224) SetMeta(v InlineResponse200223Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


