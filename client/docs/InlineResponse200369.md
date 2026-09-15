# InlineResponse200369

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200133**](InlineResponse200133.md) | Outer container to hold the different syslog server configurations per network | [optional] 
**Meta** | Pointer to [**InlineResponse200369Meta**](InlineResponse200369Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200369

`func NewInlineResponse200369() *InlineResponse200369`

NewInlineResponse200369 instantiates a new InlineResponse200369 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200369WithDefaults

`func NewInlineResponse200369WithDefaults() *InlineResponse200369`

NewInlineResponse200369WithDefaults instantiates a new InlineResponse200369 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200369) GetItems() []InlineResponse200133`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200369) GetItemsOk() (*[]InlineResponse200133, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200369) SetItems(v []InlineResponse200133)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200369) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200369) GetMeta() InlineResponse200369Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200369) GetMetaOk() (*InlineResponse200369Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200369) SetMeta(v InlineResponse200369Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200369) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


