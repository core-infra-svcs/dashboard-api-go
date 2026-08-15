# InlineResponse200423

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200423Items**](InlineResponse200423Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200423Meta**](InlineResponse200423Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200423

`func NewInlineResponse200423(items []InlineResponse200423Items, ) *InlineResponse200423`

NewInlineResponse200423 instantiates a new InlineResponse200423 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200423WithDefaults

`func NewInlineResponse200423WithDefaults() *InlineResponse200423`

NewInlineResponse200423WithDefaults instantiates a new InlineResponse200423 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200423) GetItems() []InlineResponse200423Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200423) GetItemsOk() (*[]InlineResponse200423Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200423) SetItems(v []InlineResponse200423Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200423) GetMeta() InlineResponse200423Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200423) GetMetaOk() (*InlineResponse200423Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200423) SetMeta(v InlineResponse200423Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200423) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


