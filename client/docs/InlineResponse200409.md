# InlineResponse200409

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200409Items**](InlineResponse200409Items.md) | List of Appliance VLAN Assignments | [optional] 
**Meta** | Pointer to [**InlineResponse200277Meta**](InlineResponse200277Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200409

`func NewInlineResponse200409() *InlineResponse200409`

NewInlineResponse200409 instantiates a new InlineResponse200409 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200409WithDefaults

`func NewInlineResponse200409WithDefaults() *InlineResponse200409`

NewInlineResponse200409WithDefaults instantiates a new InlineResponse200409 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200409) GetItems() []InlineResponse200409Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200409) GetItemsOk() (*[]InlineResponse200409Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200409) SetItems(v []InlineResponse200409Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200409) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200409) GetMeta() InlineResponse200277Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200409) GetMetaOk() (*InlineResponse200277Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200409) SetMeta(v InlineResponse200277Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200409) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


