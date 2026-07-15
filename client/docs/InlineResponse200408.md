# InlineResponse200408

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200408Items**](InlineResponse200408Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200408Meta**](InlineResponse200408Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200408

`func NewInlineResponse200408(items []InlineResponse200408Items, ) *InlineResponse200408`

NewInlineResponse200408 instantiates a new InlineResponse200408 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200408WithDefaults

`func NewInlineResponse200408WithDefaults() *InlineResponse200408`

NewInlineResponse200408WithDefaults instantiates a new InlineResponse200408 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200408) GetItems() []InlineResponse200408Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200408) GetItemsOk() (*[]InlineResponse200408Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200408) SetItems(v []InlineResponse200408Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200408) GetMeta() InlineResponse200408Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200408) GetMetaOk() (*InlineResponse200408Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200408) SetMeta(v InlineResponse200408Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200408) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


