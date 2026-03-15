# InlineResponse200413

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200413Items**](InlineResponse200413Items.md) | Wireless LAN controller layer 3 interfaces historical status | [optional] 
**Meta** | Pointer to [**InlineResponse200240Meta**](InlineResponse200240Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200413

`func NewInlineResponse200413() *InlineResponse200413`

NewInlineResponse200413 instantiates a new InlineResponse200413 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200413WithDefaults

`func NewInlineResponse200413WithDefaults() *InlineResponse200413`

NewInlineResponse200413WithDefaults instantiates a new InlineResponse200413 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200413) GetItems() []InlineResponse200413Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200413) GetItemsOk() (*[]InlineResponse200413Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200413) SetItems(v []InlineResponse200413Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200413) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200413) GetMeta() InlineResponse200240Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200413) GetMetaOk() (*InlineResponse200240Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200413) SetMeta(v InlineResponse200240Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200413) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


