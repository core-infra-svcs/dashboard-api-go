# InlineResponse200328

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200328Items**](InlineResponse200328Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200328Meta**](InlineResponse200328Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200328

`func NewInlineResponse200328(items []InlineResponse200328Items, ) *InlineResponse200328`

NewInlineResponse200328 instantiates a new InlineResponse200328 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200328WithDefaults

`func NewInlineResponse200328WithDefaults() *InlineResponse200328`

NewInlineResponse200328WithDefaults instantiates a new InlineResponse200328 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200328) GetItems() []InlineResponse200328Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200328) GetItemsOk() (*[]InlineResponse200328Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200328) SetItems(v []InlineResponse200328Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200328) GetMeta() InlineResponse200328Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200328) GetMetaOk() (*InlineResponse200328Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200328) SetMeta(v InlineResponse200328Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200328) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


