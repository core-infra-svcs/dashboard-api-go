# InlineResponse200301

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200300Items**](InlineResponse200300Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200219Meta**](InlineResponse200219Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200301

`func NewInlineResponse200301() *InlineResponse200301`

NewInlineResponse200301 instantiates a new InlineResponse200301 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200301WithDefaults

`func NewInlineResponse200301WithDefaults() *InlineResponse200301`

NewInlineResponse200301WithDefaults instantiates a new InlineResponse200301 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200301) GetItems() []InlineResponse200300Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200301) GetItemsOk() (*[]InlineResponse200300Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200301) SetItems(v []InlineResponse200300Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200301) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200301) GetMeta() InlineResponse200219Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200301) GetMetaOk() (*InlineResponse200219Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200301) SetMeta(v InlineResponse200219Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200301) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


