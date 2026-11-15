# InlineResponse200456

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200456Items**](InlineResponse200456Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200456Meta**](InlineResponse200456Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200456

`func NewInlineResponse200456(items []InlineResponse200456Items, ) *InlineResponse200456`

NewInlineResponse200456 instantiates a new InlineResponse200456 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200456WithDefaults

`func NewInlineResponse200456WithDefaults() *InlineResponse200456`

NewInlineResponse200456WithDefaults instantiates a new InlineResponse200456 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200456) GetItems() []InlineResponse200456Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200456) GetItemsOk() (*[]InlineResponse200456Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200456) SetItems(v []InlineResponse200456Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200456) GetMeta() InlineResponse200456Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200456) GetMetaOk() (*InlineResponse200456Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200456) SetMeta(v InlineResponse200456Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200456) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


