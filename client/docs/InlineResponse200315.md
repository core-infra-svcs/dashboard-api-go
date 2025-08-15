# InlineResponse200315

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200315Items**](InlineResponse200315Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200315Meta**](InlineResponse200315Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200315

`func NewInlineResponse200315(items []InlineResponse200315Items, ) *InlineResponse200315`

NewInlineResponse200315 instantiates a new InlineResponse200315 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200315WithDefaults

`func NewInlineResponse200315WithDefaults() *InlineResponse200315`

NewInlineResponse200315WithDefaults instantiates a new InlineResponse200315 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200315) GetItems() []InlineResponse200315Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200315) GetItemsOk() (*[]InlineResponse200315Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200315) SetItems(v []InlineResponse200315Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200315) GetMeta() InlineResponse200315Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200315) GetMetaOk() (*InlineResponse200315Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200315) SetMeta(v InlineResponse200315Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200315) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


