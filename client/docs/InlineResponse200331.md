# InlineResponse200331

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200331Items**](InlineResponse200331Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200331Meta**](InlineResponse200331Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200331

`func NewInlineResponse200331(items []InlineResponse200331Items, ) *InlineResponse200331`

NewInlineResponse200331 instantiates a new InlineResponse200331 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200331WithDefaults

`func NewInlineResponse200331WithDefaults() *InlineResponse200331`

NewInlineResponse200331WithDefaults instantiates a new InlineResponse200331 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200331) GetItems() []InlineResponse200331Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200331) GetItemsOk() (*[]InlineResponse200331Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200331) SetItems(v []InlineResponse200331Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200331) GetMeta() InlineResponse200331Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200331) GetMetaOk() (*InlineResponse200331Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200331) SetMeta(v InlineResponse200331Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200331) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


