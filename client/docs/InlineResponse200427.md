# InlineResponse200427

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200427Items**](InlineResponse200427Items.md) | The organization&#39;s RSSI information between sensor-gateway pairs. | 
**Meta** | Pointer to [**InlineResponse200427Meta**](InlineResponse200427Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200427

`func NewInlineResponse200427(items []InlineResponse200427Items, ) *InlineResponse200427`

NewInlineResponse200427 instantiates a new InlineResponse200427 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200427WithDefaults

`func NewInlineResponse200427WithDefaults() *InlineResponse200427`

NewInlineResponse200427WithDefaults instantiates a new InlineResponse200427 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200427) GetItems() []InlineResponse200427Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200427) GetItemsOk() (*[]InlineResponse200427Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200427) SetItems(v []InlineResponse200427Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200427) GetMeta() InlineResponse200427Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200427) GetMetaOk() (*InlineResponse200427Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200427) SetMeta(v InlineResponse200427Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200427) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


