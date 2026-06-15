# InlineResponse200405

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200405Items**](InlineResponse200405Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200405Meta**](InlineResponse200405Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200405

`func NewInlineResponse200405(items []InlineResponse200405Items, ) *InlineResponse200405`

NewInlineResponse200405 instantiates a new InlineResponse200405 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200405WithDefaults

`func NewInlineResponse200405WithDefaults() *InlineResponse200405`

NewInlineResponse200405WithDefaults instantiates a new InlineResponse200405 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200405) GetItems() []InlineResponse200405Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200405) GetItemsOk() (*[]InlineResponse200405Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200405) SetItems(v []InlineResponse200405Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200405) GetMeta() InlineResponse200405Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200405) GetMetaOk() (*InlineResponse200405Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200405) SetMeta(v InlineResponse200405Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200405) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


