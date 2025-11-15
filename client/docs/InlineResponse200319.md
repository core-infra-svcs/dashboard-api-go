# InlineResponse200319

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200319Items**](InlineResponse200319Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200319Meta**](InlineResponse200319Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200319

`func NewInlineResponse200319(items []InlineResponse200319Items, ) *InlineResponse200319`

NewInlineResponse200319 instantiates a new InlineResponse200319 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200319WithDefaults

`func NewInlineResponse200319WithDefaults() *InlineResponse200319`

NewInlineResponse200319WithDefaults instantiates a new InlineResponse200319 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200319) GetItems() []InlineResponse200319Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200319) GetItemsOk() (*[]InlineResponse200319Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200319) SetItems(v []InlineResponse200319Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200319) GetMeta() InlineResponse200319Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200319) GetMetaOk() (*InlineResponse200319Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200319) SetMeta(v InlineResponse200319Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200319) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


