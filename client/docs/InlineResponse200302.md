# InlineResponse200302

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200302Items**](InlineResponse200302Items.md) | The top-level propery containing all status data. | [optional] 
**Meta** | Pointer to [**InlineResponse200302Meta**](InlineResponse200302Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200302

`func NewInlineResponse200302() *InlineResponse200302`

NewInlineResponse200302 instantiates a new InlineResponse200302 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302WithDefaults

`func NewInlineResponse200302WithDefaults() *InlineResponse200302`

NewInlineResponse200302WithDefaults instantiates a new InlineResponse200302 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200302) GetItems() []InlineResponse200302Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200302) GetItemsOk() (*[]InlineResponse200302Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200302) SetItems(v []InlineResponse200302Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200302) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200302) GetMeta() InlineResponse200302Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200302) GetMetaOk() (*InlineResponse200302Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200302) SetMeta(v InlineResponse200302Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200302) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


