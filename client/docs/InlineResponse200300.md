# InlineResponse200300

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200300Items**](InlineResponse200300Items.md) | Organization Alert counts by type | 
**Meta** | [**InlineResponse200299Meta**](InlineResponse200299Meta.md) |  | 

## Methods

### NewInlineResponse200300

`func NewInlineResponse200300(items []InlineResponse200300Items, meta InlineResponse200299Meta, ) *InlineResponse200300`

NewInlineResponse200300 instantiates a new InlineResponse200300 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300WithDefaults

`func NewInlineResponse200300WithDefaults() *InlineResponse200300`

NewInlineResponse200300WithDefaults instantiates a new InlineResponse200300 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200300) GetItems() []InlineResponse200300Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200300) GetItemsOk() (*[]InlineResponse200300Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200300) SetItems(v []InlineResponse200300Items)`

SetItems sets Items field to given value.


### GetMeta

`func (o *InlineResponse200300) GetMeta() InlineResponse200299Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200300) GetMetaOk() (*InlineResponse200299Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200300) SetMeta(v InlineResponse200299Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


