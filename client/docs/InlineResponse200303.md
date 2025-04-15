# InlineResponse200303

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200303Items**](InlineResponse200303Items.md) | Array of Limited Access Roles | [optional] 
**Meta** | Pointer to [**InlineResponse200220Meta**](InlineResponse200220Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200303

`func NewInlineResponse200303() *InlineResponse200303`

NewInlineResponse200303 instantiates a new InlineResponse200303 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200303WithDefaults

`func NewInlineResponse200303WithDefaults() *InlineResponse200303`

NewInlineResponse200303WithDefaults instantiates a new InlineResponse200303 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200303) GetItems() []InlineResponse200303Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200303) GetItemsOk() (*[]InlineResponse200303Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200303) SetItems(v []InlineResponse200303Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200303) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200303) GetMeta() InlineResponse200220Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200303) GetMetaOk() (*InlineResponse200220Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200303) SetMeta(v InlineResponse200220Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200303) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


