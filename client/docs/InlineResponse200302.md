# InlineResponse200302

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InlineResponse200302Items**](InlineResponse200302Items.md) | Organization Alert counts by type | 
**Meta** | [**InlineResponse200301Meta**](InlineResponse200301Meta.md) |  | 

## Methods

### NewInlineResponse200302

`func NewInlineResponse200302(items []InlineResponse200302Items, meta InlineResponse200301Meta, ) *InlineResponse200302`

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


### GetMeta

`func (o *InlineResponse200302) GetMeta() InlineResponse200301Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200302) GetMetaOk() (*InlineResponse200301Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200302) SetMeta(v InlineResponse200301Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


