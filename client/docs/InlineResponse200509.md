# InlineResponse200509

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200509Items**](InlineResponse200509Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200283Meta**](InlineResponse200283Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200509

`func NewInlineResponse200509() *InlineResponse200509`

NewInlineResponse200509 instantiates a new InlineResponse200509 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200509WithDefaults

`func NewInlineResponse200509WithDefaults() *InlineResponse200509`

NewInlineResponse200509WithDefaults instantiates a new InlineResponse200509 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200509) GetItems() []InlineResponse200509Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200509) GetItemsOk() (*[]InlineResponse200509Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200509) SetItems(v []InlineResponse200509Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200509) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200509) GetMeta() InlineResponse200283Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200509) GetMetaOk() (*InlineResponse200283Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200509) SetMeta(v InlineResponse200283Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200509) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


